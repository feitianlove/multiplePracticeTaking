package raft

/*
	raft日志复制
*/

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

//日志结构
type LogEntry struct {
	Term    int         //任期号
	Command interface{} //client发送的命令
}

//日志复制请求结构
type AppendEntriesArgs struct {
	Term int // leader期
	//在raft中，有可能会出现直接连上follower，此时follower需要给client发送leaderID，方便follower重定向
	LeaderId     int        //leader ID  so follower can redirect clients 节点日志请求重定向到leader上
	PrevLogIndex int        // prevLogIndex 紧邻新⽇志条⽬之前的那个⽇志条⽬的索引
	PrevLogTerm  int        //prevLogTerm 紧邻新⽇志条⽬之前的那个⽇志条⽬的任期
	Entries      []LogEntry //准备提交的日志条目
	LeaderCommit int        //leader已经提交的日志索引
}

//日志复制响应结构
type AppendEntriesReply struct {
	CurrentTerm int  //leader任期，主要用于leader更新自己
	Success     bool //结果为真 如果跟随者所含有的条⽬和prevLogIndex以及prevLogTerm匹配上了
	//自定义冲突日志相关变量
	ConflictTerm int //冲突日志的任期编号
	FirstIndex   int //存储第一个冲突编号的日志索引
}

//唤醒一致性检查
func (rf *Raft) wakeupConsistencyCheck() {
	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			rf.NewEntryCond[i].Broadcast()
		}
	}
}

//应用日志进程
func (rf *Raft) applyEntryDaemon() {
	//日志提交完成之后将日志进行应用，然后进行返回
	for {
		var logs []LogEntry
		rf.mu.Lock()
		//判断 如果节点最后的应用日志索引与已提交的日志索引相等，说明所有的日志已经被应用
		for rf.LastApplied == rf.CommitIndex {
			rf.CommitCond.Wait()
			select {
			//检测是否有中断
			case <-rf.shutDown:
				rf.mu.Unlock()
				close(rf.applyCh)
				return
			default:

			}
			//获取最后的应用索引，与最新提交的应用索引
			last, cur := rf.LastApplied, rf.CommitIndex
			//还有一部分已经提交的日志没有被应用
			if last < cur {

				rf.LastApplied = rf.CommitIndex
				//找到已经提交但未更新的这一部分日志
				logs = make([]LogEntry, cur-last)
				copy(logs, rf.Logs[last+2:cur+1])
			}
			rf.mu.Unlock()
			//对还没有应用的日志，进行应用
			for i := 0; i < cur-last; i++ {
				reply := ApplyMsg{
					Index:   last + i,
					Command: logs[i].Command,
				}
				//传回响应
				rf.applyCh <- reply
			}
		}
	}
}

//启动日志复制进程
func (rf *Raft) logEntryAgreeDaemon() {
	//遍历节点，向其他每个节点发起日志复制操作
	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			go rf.consistencyCheckDaemon(i)
		}
	}
}

//发起日志复制操作
func (rf *Raft) consistencyCheckDaemon(n int) {
	for {
		rf.mu.Lock()
		// 每个节点都在等待client将命令提交给leader
		rf.NewEntryCond[n].Wait()
		select {
		case <-rf.shutDown:
			rf.mu.Unlock()
			return
		default:

		}
		//判断节点角色，只有leader才能发起日志 复制
		if rf.IsLeader {
			var args AppendEntriesArgs
			args.Term = rf.CurrentTerm
			args.LeaderId = rf.me
			args.LeaderCommit = rf.CommitIndex
			args.PrevLogIndex = rf.NextIndex[n] - 1
			args.PrevLogTerm = rf.Logs[args.PrevLogIndex].Term
			//判断是否有新的日志
			//leader 的日志长度大于leader所知道的follower n的日志长度
			if rf.NextIndex[n] < len(rf.Logs) {
				args.Entries = append(args.Entries, rf.Logs[rf.NextIndex[n]:]...)
			} else {
				args.Entries = nil
			}
			rf.mu.Unlock()
			//新建响应通道，在发起日志复制请求之后，存储结果
			replyCh := make(chan AppendEntriesReply, 1)
			go func() {
				var reply AppendEntriesReply
				//发起日志复制请求
				if rf.sendAppendEntries(n, &args, &reply) {
					replyCh <- reply
				}
			}()
			//获取响应
			select {
			case reply := <-replyCh:
				rf.mu.Lock()
				if reply.Success {
					//响应成功
					rf.MatchIndex[n] = reply.FirstIndex
					rf.NextIndex[n] = rf.MatchIndex[n] + 1
					//提交日志（更新已经提交的日志索引）
					rf.updateCommitIndex()
				} else {
					//判断当前传回来的term与leader自己的term谁大
					if reply.CurrentTerm > args.Term {
						//更新当前leader任期号
						rf.CurrentTerm = reply.CurrentTerm
						rf.VotedFor = -1
					}
					// leader是接收的每个节点返回的响应，在进行当前节点的时候，角色状态可能已经修改过
					if rf.IsLeader {
						rf.IsLeader = false
						//变更角色状态一致性检查
						rf.wakeupConsistencyCheck()
					}
					rf.mu.Unlock()
					rf.ResetTimer <- struct{}{}
					return
				}
				//解决冲突相关问题
				//know 当前leader是否能查到冲突
				//lastIndex 代表节点中最后一个包含冲突任期编号的日志索引
				know, lastIndex := false, 0
				if reply.ConflictTerm != 0 {
					//查找最后一个包含冲突的日志索引
					for i := len(rf.Logs) - 1; i > 0; i-- {
						//找到产生冲突的任期编号
						if rf.Logs[i].Term == reply.ConflictTerm {
							know = true
							lastIndex = i
							break
						}
					}
					//如果找到了冲突编号
					if know {
						//判断当前获取到的冲突编号索引与响应中的冲突日志编号索引大小

						if lastIndex > reply.FirstIndex {
							//满足当前的条件下，说明在最后一个产生之前已经有了一个冲突编号，只允许存在第一个
							lastIndex = reply.FirstIndex
						}
						rf.NextIndex[n] = lastIndex
					} else {
						rf.NextIndex[n] = reply.FirstIndex
					}
				} else {
					//如果响应中没有返回冲突的任期编号
					rf.NextIndex[n] = reply.FirstIndex
				}
				// 1<= rf.nextIndex<=len(raft.Logs)
				rf.NextIndex[n] = min(max(rf.NextIndex[n], 1), len(rf.Logs))
			}
			rf.mu.Unlock()
		} else {
			//当前节点不是leader，直接返回
			rf.mu.Unlock()
			return
		}
	}
}

//发起日志复制的请求
func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}

//更新日志提交索引
func (rf *Raft) updateCommitIndex() {

}

//接收到日志复制请求之后的处理
func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	select {
	case <-rf.shutDown:
		return
	default:
	}
	rf.mu.Lock()
	defer rf.mu.Unlock()
	//判断任期  Reply false if term < currentTerm
	if args.Term < rf.CurrentTerm {
		//将比leader更新的任期，返回
		reply.CurrentTerm = rf.CurrentTerm
		reply.Success = false
		rf.mu.Unlock()
		return
	}
	//如果当前角色还是leader，更改节点角色(强一致性)
	if rf.IsLeader {
		rf.IsLeader = false
		//唤醒一致性检查
		rf.wakeupConsistencyCheck()
	}
	// 	如果不相等，那么出现client直接访问follower的时候重定向到错误的leader
	if rf.VotedFor != args.LeaderId {
		rf.VotedFor = args.LeaderId
	}
	//当前节点的任期标号小于leader,强制覆盖
	if args.Term > rf.CurrentTerm {
		rf.CurrentTerm = args.Term
	}
	// 重置timer
	rf.ResetTimer <- struct{}{}

	// TODO 在发送附加⽇志 RPC 的时候，领导⼈会把新的⽇志条⽬紧接着之前的条⽬（prevIndex）的索引位置和任期(prevTerm)
	//号包含在⾥⾯。如果跟随者在它的⽇志中找不到包含相同索引位置和任期号的条⽬，那么他就会拒绝接收新的⽇志条⽬。
	prevLogIndex, prevLogTerm := 0, 0
	//大于说明在follower中的日志中可以找到args包含的pervLogIndex
	if len(rf.Logs) > args.PrevLogIndex {
		prevLogIndex = args.PrevLogIndex
		//从follower节点的日志中获取到的与args传入的索引相同的任期编号
		prevLogTerm = rf.Logs[prevLogIndex].Term
	}
	//判断日志是否匹配
	if prevLogIndex == args.PrevLogIndex && prevLogTerm == args.PrevLogTerm {
		reply.Success = true
		//截取当前已经知道的最后一个匹配，将不匹配的丢掉
		rf.Logs = rf.Logs[:prevLogIndex+1]
		//追加日志
		rf.Logs = append(rf.Logs, args.Entries...)
		//获取到更新之后的最后一个索引
		var last = len(rf.Logs) - 1
		// 更新一下commitIndex
		//如果领导者的已知已经提交的最⾼的⽇志条⽬的索引leaderCommit ⼤于 接收者的已知已经提交的最⾼的⽇志条⽬的索引commitIndex
		if args.LeaderCommit > rf.CommitIndex {
			rf.CommitIndex = min(args.LeaderCommit, last)
			//commitIndex 更新
			go func() {
				rf.CommitCond.Broadcast()
			}()
		}
		//更新最后一个冲突日志
		reply.ConflictTerm = rf.Logs[last].Term
		reply.FirstIndex = last
	} else {
		reply.Success = false
		//处理冲突任期编号
		var first = 1
		reply.ConflictTerm = prevLogTerm
		//如果说新的ConflictTerm == 0， leader有更多的日志需要复制，或者follower没有日志
		if reply.ConflictTerm == 0 {
			first = len(rf.Logs)
			//将响应的冲突日志条目设置成当前节点最后一一个日志条目
			reply.ConflictTerm = rf.Logs[first-1].Term
		} else {
			// 不为0说明任期编号有冲突
			for i := prevLogIndex - 1; i > 0; i-- {
				if rf.Logs[i].Term != prevLogTerm {
					//遍历查找产生冲突的日志i
					first = i + 1
					break
				}

			}
		}
		//第一个产生冲突的日志
		reply.FirstIndex = first
	}
	//
	//添加日志
	//返回冲突任期编号
}
