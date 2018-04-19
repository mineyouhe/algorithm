package main

import "fmt"

//LRUCache,双向链表和hash函数
//初始化时候调用NewLRUCache
//添加数据Put (如果容量已满则需要删除)
//获取数据Get
//查询数据，如果存在，将该数据节点移动到双向链表头部
//如果不存在，通过Put接口将数据插入双向链表中(如果容量未满，新节点插入链表头部，如果已满，把链表中的最后一个节点内容替换为新内容，移动到头部)

type Node struct {
	Key   interface{}
	Value interface{}
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Head, Tail *Node
	Capacity   int
	Map        map[interface{}]*Node
}

func NewLRUCache(capacity int) *LRUCache {
	l := &LRUCache{}
	l.Capacity = capacity
	l.Head = &Node{}
	l.Tail = &Node{}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	l.Head.Prev = nil
	l.Tail.Next = nil
	l.Map = make(map[interface{}]*Node)
	return l
}

//分离节点
func (l *LRUCache) detach(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

//节点插入头部
func (l *LRUCache) attach(n *Node) {
	n.Prev = l.Head
	n.Next = l.Head.Next
	l.Head.Next = n
	n.Next.Prev = n
}

func (l *LRUCache) Put(k interface{}, v interface{}) {
	oldV, ok := l.Map[k]
	if ok {
		l.detach(oldV)
		oldV.Value = v
	} else {
		var n *Node
		if len(l.Map) >= l.Capacity {
			n = l.Tail.Prev
			l.detach(n)
			delete(l.Map, n.Key)
		} else {
			n = new(Node)
		}
		n.Key = k
		n.Value = v
		l.Map[k] = n
		l.attach(n)
	}
}

func (l *LRUCache) Get(k interface{}) (interface{}, bool) {
	v, ok := l.Map[k]
	if ok {
		l.detach(v)
		l.attach(v)
		return v.Value, true
	}
	return -1, false
}

func main() {
	l := NewLRUCache(4)
	l.Put("str", 1)
	l.Put(1, 2)
	l.Put(2, 2)
	l.Put(3, 2)
	l.Put(4, 2)
	i, ok := l.Get(4)
	fmt.Println(i.(int), ok)
	i, ok = l.Get("str")
	fmt.Println(i.(int), ok)
}
