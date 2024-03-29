package linked_list

// ListNoder provides listNode interface
type ListNoder interface {
	HasCycle(head *listNode) bool
	NextNode(list *listNode) *listNode
}

// ListNode is definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

// HasCycle determine if list has a cycle in it
// Return true if has, else return false
func (l *listNode) HasCycle(head *listNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// NextNode sets next node for listNode
func (l *listNode) SetNextNode(list *listNode) *listNode {
	l.Next = list
	return l.Next
}

// NewListNode returns new listNode
func NewListNode(val int) *listNode {
	return &listNode{val, nil}
}
