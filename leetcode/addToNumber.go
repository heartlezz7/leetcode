package leetcode

// #################### add two number ####################
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var current = &ListNode{}
	var head = current
	if l1.Next == nil && l2.Next == nil {

		sum := l1.Val + l2.Val
		if sum > 9 {

			current = &ListNode{Val: sum % 10, Next: nil}
			head = current
			current.Next = current
			s := &ListNode{Val: 1, Next: nil}
			current.Next = s

		} else {
			return &ListNode{Val: sum, Next: nil}

		}
		return head
	}

	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return head.Next
}
