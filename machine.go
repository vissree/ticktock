package ticktock

// LockRequest is a request for a lock on a machine
type Lock struct {
	Id          int
	RequesterId int
	Timestamp   int
}

// Ack is an acknowledgement of a lock request
type Ack struct {
	Id          int
	RequesterId int
	AckerId     int
	Timestamp   int
	Explicit    bool
}

// Unlock is a request to unlock an acquired lock
type Unlock struct {
	Id          int
	RequesterId int
}

// Machine is a machine in the system
type Machine struct {
	Id        int
	Name      string
	HasLock   bool
	LockQueue *Heap[Lock]
	AckList   []Ack
}

// AddLock adds a Lock to the Lock queue
func (m *Machine) AddLock(lock Lock) {
	m.LockQueue.Insert(lock)
}

// AddAckadds an Ack to the Ack list
func (m *Machine) AddAck(ack Ack) {
	m.AckList = append(m.AckList, ack)
}

func (m *Machine) EmptyAckList() {
	m.AckList = []Ack{}
}
