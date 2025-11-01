/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {

    start:=&ListNode{Next:head}
    current:=start

    check:=make(map[int]bool,len(nums))
    for _,v:= range nums{
        check[v]=true
    }
    for current!=nil{
        if current.Next!=nil&&check[current.Next.Val] {
        current.Next=current.Next.Next}else{
        current=current.Next}
    }
    return start.Next

}