/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root==nil{ return true }
    return isMirror(root.Left,root.Right)

}
func isMirror(r1,r2 *TreeNode)bool{
    if r1==nil && r2==nil{return true}
    if r1==nil || r2==nil{return false}
    return( r1.Val==r2.Val )&& isMirror(r1.Left,r2.Right) &&isMirror(r2.Left,r1.Right)
}
