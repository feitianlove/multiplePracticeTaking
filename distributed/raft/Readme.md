# raft协议的简单实现
## 实现raft的选举
- 1、更具raft paper添加相应raft所包含的状态属性
- 2、添加相应的自定义属性
- 3、查找入口Make函数: 对raft结构各个成员进行初始化
- 4、初始化完成之后准备选举
    
    - 填充请求参数和响应参数（论文有）
    - 遍历节点，向每个节点发送投票请求，接收请求结果与响应
    - 关闭计数器与通道
    - 遍历缓存通道，获取每一个响应中的结果
    - 统计票数，当票数超过一半，当选leader
    - 当选leader重置相关状态、发起心跳
    - 如果没有当选为leader，需要考虑当前任期的是不是最新的，通过响应的任期号进行更新
- 5、 其他follower处理接收到的投票请求 RequestVote
    - 如果候选人的任期没有自己的大，就返回false
    - 如果 votedFor 为-1或者为 candidateId，并且候选⼈的⽇志⾄少和⾃⼰⼀样新，那么就投票给他
## 实现raft的日志复制
- 1、当raft当选leader之后就可以发起一个日志复制的操作
- 2、修改raft结构
    - 添加日志提交和日志复制的一致性检查
    - 添加一个applyCh接收日志响应
    - 添加shutdown作为中断标志
- 3、实现日志复制
    - 实现重置函数resetOnElection
    - 填充start函数： 修改默认值、将命令以日志的方式传递给leader、更新自己的nextIndex与matchIndex、唤醒一致性检查
    - 实现一致性检查函数
- 4、实现向每个节点发起日志复制操作
- 5、首先实现日志请求与响应结构
- 6、发起日志rpc请求的日志复制send函数
- 7、接收follower的响应
- 8、如果响应成功，更新相关数据
- 9、更新响应的索引
- 10、如果响应失败
    - 判断任期号
    - 然后变更leader的任期号和角色状态
    - 根据冲突编号 进行属性修改
-11、判断已经提交的日志索引，与当前最后一个被应用的日志索引大小