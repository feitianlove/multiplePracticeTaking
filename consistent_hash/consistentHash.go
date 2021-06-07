package consistent_hash

import "C"
import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Node struct {
	IPAddr   string
	replicas int // 本数目相当于权重，如果为0则为剔除该节点，如果大于0，则表示虚拟节点的个数
}

type Hash func([]byte) uint32

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Consistent struct {
	HashFn    Hash
	Nodes     UInt32Slice       // 已排序的节点哈希切片
	HashMap   map[uint32]string // 节点哈希和KEY的map，键是哈希值，值是节点Key
	Resources map[uint32]bool   // 是否健康 在不需要迁移数据的系统中不需要
	Weight    map[string]int    //节点变化需要重新均衡所有节点
	sync.RWMutex
}

func NewConsistentHash(fn Hash) *Consistent {
	m := Consistent{
		HashFn:    fn,
		HashMap:   make(map[uint32]string),
		Resources: make(map[uint32]bool),
		Weight:    make(map[string]int),
	}
	if m.HashFn == nil {
		m.HashFn = crc32.ChecksumIEEE
	}
	return &m
}
func (c *Consistent) IsEmpty() bool {
	return len(c.Nodes) == 0
}

func (c *Consistent) AddNode(nodes []Node) {
	c.Lock()
	defer c.Unlock()
	for _, node := range nodes {
		for i := 0; i < node.replicas; i++ {
			hash := c.HashFn([]byte(strconv.Itoa(i) + node.IPAddr))
			c.Nodes = append(c.Nodes, hash)
			c.HashMap[hash] = node.IPAddr
			c.Resources[hash] = true
			c.Weight[node.IPAddr] = node.replicas
		}
	}
	sort.Sort(c.Nodes)
}

func (c *Consistent) GetNode(key string) string {
	if c.IsEmpty() {
		return ""
	}
	hash := c.HashFn([]byte(key))
	idx := sort.Search(len(c.Nodes), func(i int) bool {
		return c.Nodes[i] >= hash
	})
	if idx >= len(c.Nodes) {
		idx = 0
	}
	return c.HashMap[c.Nodes[idx]]
}
func (c *Consistent) RemoveNode(node string) error {
	c.Lock()
	defer c.Unlock()
	_, ok := c.Weight[node]
	if !ok {
		return errors.New(fmt.Sprintf("%s不存在系统中\n", node))
	}
	// 删除当前节点重新计算hash值
	delete(c.Weight, node)
	c.BalanceNode()
	return nil
}

func (c *Consistent) BalanceNode() {
	// 清空
	c.HashMap = make(map[uint32]string)
	c.Resources = make(map[uint32]bool)
	c.Nodes = UInt32Slice{}
	for ip, replica := range c.Weight {
		for i := 0; i < replica; i++ {
			hash := c.HashFn([]byte(strconv.Itoa(i) + ip))
			c.Nodes = append(c.Nodes, hash)
			c.HashMap[hash] = ip
			c.Resources[hash] = true
		}
	}
}

func (c *Consistent) ChangeWeight(node Node) error {
	c.Lock()
	defer c.Unlock()
	_, ok := c.Weight[node.IPAddr]
	if !ok {
		return errors.New(fmt.Sprintf("%s不存在系统中\n", node.IPAddr))
	}
	// 删除当前节点重新计算hash值
	c.Weight[node.IPAddr] = node.replicas
	c.BalanceNode()
	return nil
}
