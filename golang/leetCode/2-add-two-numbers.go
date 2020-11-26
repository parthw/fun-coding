// Problem - https://leetcode.com/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	l3 := &ListNode{}
	currentl3 := l3
	for {
		sum := 0
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			sum = l2.Val + carry
		}
		if l2 == nil {
			sum = l1.Val + carry
		}
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + carry
		}
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		//Insert(l3, sum)
		currentl3.Next = &ListNode{Val: sum}
		currentl3 = currentl3.Next
	}
	if carry != 0 {
		//Insert(l3, carry)
		currentl3.Next = &ListNode{Val: carry}
		currentl3 = currentl3.Next
	}

	return l3.Next
}

func Insert(l *ListNode, value int) {
	if l.Next == nil {
		l.Next = &ListNode{Val: value}
		return
	}
	Insert(l.Next, value)
	return
}