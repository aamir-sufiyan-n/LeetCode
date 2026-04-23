/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    current:=head
    for current!=nil && current.Next!=nil{
        val:=gcd(current.Val,current.Next.Val)
        node:=&ListNode{Val:val}
        node.Next=current.Next
        current.Next=node
        current=node.Next
    }
    return head

}
func gcd(a,b int)int{
    for b!=0{
        a,b=b,a%b
    }
    return a 
}