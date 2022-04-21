/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
     result := ListNode{Val:0}
    var current *ListNode = &result
    for l1 != nil || l2 != nil || carry == 1  {
        sum := carry
        carry = 0
        if l1 != nil{
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        if sum > 9 {
            carry = 1
            sum %= 10
        }
        current.Val = sum
        
        if l1 != nil || l2 != nil || carry == 1{
        current.Next = &ListNode{}
        current = current.Next
        }
    }
    return &result
}