package main

import (
	"fmt"
	my "github.com/vissree/ticktock"
)

func main() {
	m1 := my.Machine{Id: 1, Name: "m1", HasLock: false, LockQueue: my.NewHeap(compareLock), AckList: []my.Ack{}}
	m2 := my.Machine{Id: 2, Name: "m2", HasLock: false, LockQueue: my.NewHeap(compareLock), AckList: []my.Ack{}}
	m3 := my.Machine{Id: 3, Name: "m3", HasLock: false, LockQueue: my.NewHeap(compareLock), AckList: []my.Ack{}}
	m4 := my.Machine{Id: 4, Name: "m4", HasLock: false, LockQueue: my.NewHeap(compareLock), AckList: []my.Ack{}}
	m5 := my.Machine{Id: 5, Name: "m5", HasLock: false, LockQueue: my.NewHeap(compareLock), AckList: []my.Ack{}}

	t := my.NewTree(compareMachine)
	t.Insert(m1)
	t.Insert(m5)
	t.Insert(m2)
	t.Insert(m4)
	t.Insert(m3)

	treeTraversal(t.GetRoot())
}

func compareMachine(a my.Machine, b my.Machine) int {
	if a.Id < b.Id {
		return -1
	} else if a.Id > b.Id {
		return 1
	} else {
		return 0
	}
}

func compareLock(a my.Lock, b my.Lock) int {
	if a.Id < b.Id {
		return -1
	} else if a.Id > b.Id {
		return 1
	} else {
		return 0
	}
}

func treeTraversal(n *my.Node[my.Machine]) {
	if n == nil {
		return
	}

	treeTraversal(n.GetLeft())
	fmt.Println(n.GetValue().Name)
	treeTraversal(n.GetRight())
}
