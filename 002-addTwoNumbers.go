/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var result ListNode{
        Val: l1.Val + l2.Val
    }
    for l1.Next != nil || l2.Next != nil {
        sum := l1.Val + l2.Val + carry
        carry = 0
        if sum > 10 {
            carry = 1
        }
        
    }
    return nil
}