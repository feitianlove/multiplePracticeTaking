package raft

import (
	"fmt"
	"sync"
	"time"
)

/*
	raft raft leader选举相关的逻辑
*/

//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//投票请求参数
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	Term         int // 候选人的任期号
	CandidateId  int // 候选人的ID
	LastLogIndex int //候选人的最后日志索引
	LastLogTerm  int //候选人的最后日志任期号
}

//
// example RequestVote RPC reply structure.
// field names must start with capital letters!
// 响应参数
type RequestVoteReply struct {
	// Your data here (2A).
	CurrentTerm int  //当前任期号
	VoteGranted bool // 是否获取到了该节点的投票
}

//
// example RequestVote RPC handler.
// 接收到投票请求进行处理
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	fmt.Printf("jasdkfjaksdjfaksdjfaksdjfalkRequestVote: %+v\n", *args)
	//当前节点的最后一次日志索引获取
	lastIndex := len(rf.Logs) - 1
	//当前节点的最后一次日志任期获取
	lastTerm := rf.Logs[lastIndex].Term

	if args.Term < rf.CurrentTerm {
		reply.CurrentTerm = rf.CurrentTerm
		reply.VoteGranted = false
	} else {
		if args.Term > rf.CurrentTerm {
			// 将当前raft节点的状态变成follower
			rf.CurrentTerm = args.Term
			rf.IsLeader = false
			rf.VotedFor = -1
		}
		// 如果当前节点votedFor为-1或者candidateID
		if rf.VotedFor == -1 || rf.VotedFor == args.CandidateId {
			// 并且候选人的日志至少和自己一样新，那么就投票给他
			if args.LastLogTerm == lastTerm && args.LastLogIndex >= lastIndex || args.LastLogTerm > lastTerm {
				rf.ResetTimer <- struct{}{}
				rf.IsLeader = false
				rf.VotedFor = args.CandidateId
				reply.VoteGranted = true
			}
		}

	}
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

// 启动raft进程
func (rf *Raft) ElectionDaemon() {
	fmt.Println("ElectionDaemon")
	for {
		select {
		//接收到请求之后的处理
		case <-rf.ResetTimer:
			fmt.Println(1)
			if !rf.ElectionTimer.Stop() {
				// 发送超时 不让自己在变成Candidate
				<-rf.ElectionTimer.C
			}
			// 重置选举超时
			rf.ElectionTimer.Reset(rf.ElectionTimeOut)
		case <-rf.ElectionTimer.C:
			fmt.Println(2)
			// 超时，follower没有在指定时间内接收到leader的信息，自己变成candidate，向其他节点发起投票请求
			go rf.CanvassVotes()
			//重置选举超时
			rf.ElectionTimer.Reset(rf.ElectionTimeOut)
		}

	}
}

//填充请求参数
func (rf *Raft) fillVoteArgs(args *RequestVoteArgs) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	//任期号+1
	rf.CurrentTerm += 1
	rf.VotedFor = rf.me // 投票给自己
	args.Term = rf.CurrentTerm
	args.CandidateId = rf.me
	args.LastLogIndex = len(rf.Logs) - 1
	args.LastLogTerm = rf.Logs[args.LastLogIndex].Term
}

//发起选举请求
func (rf *Raft) CanvassVotes() {
	var voteArgs RequestVoteArgs
	rf.fillVoteArgs(&voteArgs)

	//获取节点数量，正式发起投票请求

	peers := len(rf.peers)
	//设置缓存channel，大小为peers，保存结构
	replyCh := make(chan RequestVoteReply, peers)
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < peers; i++ {
		if i == rf.me {
			// 不让自己在变成Candidate
			//rf.ResetTimer <- struct{}{}
		} else {
			wg.Add(1)
			count++
			//发起投票
			go func(n int) {
				defer func() {
					wg.Done()
				}()

				var reply RequestVoteReply
				//投票rpc 请求结果
				doneCh := make(chan bool, 1)
				go func() {
					ok := rf.sendRequestVote(n, &voteArgs, &reply)
					doneCh <- ok
					fmt.Println(ok, n, reply)
				}()
				select {
				case ok := <-doneCh:
					fmt.Printf("reply%+v\n", reply)
					if !ok {
						return
					}
					//响应的投票结果传入reply channel中
					replyCh <- reply
				}
				fmt.Println("ok", reply)
			}(i)

		}
	}
	// 另启一个协程关闭通道
	// TODO 这里插一个面试题： 从一个关闭的通道中读写数据：写数据会panic，如果有数据则会返回数据，如果没有数据则会返回nil值
	go func() {
		wg.Wait()
		close(replyCh)
		fmt.Println("end")
	}()
	//遍历缓存通道，获取每一个响应中的结果
	var votes = 1 //统计票数结果，自己会给自己投一票，初始值为1
	for reply := range replyCh {
		// 得到了当前返回值的票
		fmt.Println("---------------reply", reply)
		if reply.VoteGranted == true {
			votes++
			if votes > peers/2 {
				rf.mu.Lock()
				rf.IsLeader = true
				rf.mu.Unlock()
				//  重置相关状态
				rf.resetOnElection()
				// 发起心跳机制，防止追随者变成候选人
				go rf.heartbeatDaemon()
				// 当选leader之后发起日志复制操作
				go rf.logEntryAgreeDaemon()
				return
			}
		} else if reply.CurrentTerm > voteArgs.Term {
			// 当follower的任期号(reply.CurrentTerm)大于当前candidate的任期号(voteArgs.Term)
			// 改变状态，重新回到follower 状态
			rf.mu.Lock()
			rf.IsLeader = false
			rf.VotedFor = -1
			rf.CurrentTerm = reply.CurrentTerm
			rf.mu.Unlock()
			rf.ResetTimer <- struct{}{}
			return
		}
	}
	//统计所有的票数，决定是否能当成leader
}

//启动心跳进程  rf.ResetTimer <- struct{}{}
func (rf *Raft) heartbeatDaemon() {
	for {
		if _, isLeader := rf.GetState(); isLeader {
			// 只要是leader 就可以不断重置选举超时
			fmt.Println("CanvassVotes", 3)
			rf.ResetTimer <- struct{}{}
		} else {
			break
		}
		//设置心跳间隔
		time.Sleep(rf.HeartbeatInterval)
	}
}

//当⼀个领导⼈刚获得权⼒的时候，他初始化所有的 nextIndex 值为⾃⼰的最后⼀条⽇志的 index 加 1
func (rf *Raft) resetOnElection() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	//节点数量
	count := len(rf.peers)
	//日志长度恰好是最后一个日志长度+1
	length := len(rf.Logs)
	for i := 0; i < count; i++ {

		rf.NextIndex[i] = length
		//对于每一个服务器，已经复制给他的最高日志最高索引值
		rf.MatchIndex[i] = 0
		if i == rf.me {
			// leader日志 复制给自己
			rf.MatchIndex[i] = length - 1

		}
	}
}
