package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	temp1 := head
	lengthList := 0

	for temp1 != nil {
		temp1 = temp1.Next
		lengthList += 1
	}

	if n == lengthList {
		return head.Next
	}


	if n == 1{
		temp1 = head
		var temp2 *ListNode
		for temp1.Next != nil{
			temp2 = temp1
			temp1 = temp1.Next
		}

		temp2.Next = nil
		return head

	}

	temp1 = head
	
	for i := 0; i < lengthList-n-1; i++ {
		temp1 = temp1.Next
	}

	temp2 := temp1.Next
	temp1.Next = temp2.Next

	return head
}

func main() {

}
