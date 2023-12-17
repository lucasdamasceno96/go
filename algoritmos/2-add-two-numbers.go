package main 
import "fmt"

func addTwoNumbers(l1 *ListNode, l2* ListNode) * ListNode{
	l1Ptr := l1
	l2Ptr := l2

	dummyHead := &ListNode{}
	output := dummyHead
	carry := 0

	for l1Ptr != nil || l1Ptr != nil {
		sum := carry

		if l1Ptr != nil {
			sum += l1Ptr.Val 
			l1Ptr = l1Ptr.Next
		}

		if l2Ptr != nil {
			sum += l2Ptr.Val
			l2Ptr = l2Ptr.Next
		}

		carry = sum / 10
		digit := sum % 10

		output.Next = &ListNode{Val: digit}
		output = output.Next
	}
	if carry > 0 {
		output.Next = &ListNode{Val: carry}
	}
		return dummyHead.Next

} 