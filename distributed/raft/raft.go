package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	"math/rand"
	"sync"
	"time"
)
import "github.com/feitianlove/multiplePracticeTaking/labrpc"

// import "bytes"
// import "encoding/gob"

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make().
//
type ApplyMsg struct {
	Index       int
	Command     interface{}
	UseSnapshot bool   // ignore for lab2; only used in lab3
	Snapshot    []byte // ignore for lab2; only used in lab3
}

//
// A Go object implementing a single Raft peer.
// 需要定义一个结构来保存有关每个日志条目的信息， 论文中的表格
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	// me 是自己在当前整个网络中的索引号
	me int // this peer's index into peers[]

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	// 论文中：Persistent state on all servers 需要持久化的状态变量
	CurrentTerm  int          //当前节点号
	VotedFor     int          //投票给谁
	Logs         []LogEntry   //日志结构
	CommitCond   *sync.Cond   //主要用于提交日志索引条目同志
	NewEntryCond []*sync.Cond //唤醒每个节点一致性检查

	CommitIndex int   //已知已提交的最⾼的⽇志条⽬的索
	LastApplied int   //已经被应⽤到状态机的最⾼的⽇志条⽬的索引
	NextIndex   []int //对于每⼀台服务器，发送到该服务器的下⼀个⽇志条⽬的索引（初始值为领导者最后的⽇志条⽬的索引+1）
	MatchIndex  []int //对于每⼀台服务器，已知的已经复制到该服务器的最⾼⽇志条⽬的索引（初始值为0，单调递增）
	//自定义属性
	IsLeader          bool             //查看当前节点是不是leader
	ElectionTimer     *time.Timer      //选举超时实例
	ElectionTimeOut   time.Duration    // 选举超时（选中leader重置超时）
	ResetTimer        chan interface{} //重置选举超时
	HeartbeatInterval time.Duration    // 心跳超时
	applyCh           chan ApplyMsg    //日志提交成功之后，向服务发起的应用消息
	shutDown          chan struct{}    //中断标志
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	var term int
	var isleader bool
	// Your code here (2A).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	term = rf.CurrentTerm
	isleader = rf.IsLeader
	return term, isleader
}

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := -1
	term := 0
	isLeader := false //假定当前不是leader
	select {
	case <-rf.shutDown:
		return index, term, isLeader
	default:

	}

	// Your code here (2B).

	rf.mu.Lock()
	defer rf.mu.Unlock()
	// 判断当前节点是否是leader，如果是leader，从client中添加日志
	if rf.IsLeader {
		log := LogEntry{
			Term:    rf.CurrentTerm,
			Command: command,
		}
		rf.Logs = append(rf.Logs, log)
		index = len(rf.Logs) - 1
		term = rf.CurrentTerm
		isLeader = true
		rf.NextIndex[rf.me] = index + 1
		rf.MatchIndex[rf.me] = index
		// client日志添加到leader成功，进行一致性检查
		go rf.consistencyCheckDaemon()
	}

	return index, term, isLeader
}

//
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
// 创建一个raft节点服务
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me
	rf.applyCh = applyCh
	// Your initialization code here (2A, 2B, 2C).
	// 初始化状态所有都是follower
	rf.IsLeader = false
	rf.VotedFor = -1
	//初始化重置超时选举
	rf.ResetTimer = make(chan interface{})
	// 选举的超时时间150-300随机时间
	rf.ElectionTimeOut = time.Millisecond * (150 + time.Duration(rand.Int63()%150))
	//超时器
	rf.ElectionTimer = time.NewTimer(rf.ElectionTimeOut)
	//心跳超时时间 1s
	rf.HeartbeatInterval = time.Second

	rf.NextIndex = make([]int, len(peers))
	rf.MatchIndex = make([]int, len(peers))
	rf.Logs = make([]LogEntry, 1)
	rf.Logs[0] = LogEntry{
		Term:    0,
		Command: nil,
	}
	//日子复制初始化
	rf.shutDown = make(chan struct{})
	// commit ch
	rf.CommitCond = sync.NewCond(&rf.mu)
	rf.NewEntryCond = make([]*sync.Cond, len(peers))
	for i := 0; i < len(peers); i++ {
		rf.NewEntryCond[i] = sync.NewCond(&rf.mu)
	}
	for i := 0; i < len(peers); i++ {
		//初始化为leader最后的索引值+1，也就是leader的日志长度+1

		rf.NextIndex[i] = len(rf.Logs) + 1
	}
	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())
	// 完成初始化之后启动选举进程
	rf.ElectionDaemon()
	// 日志应用进程
	return rf
}
