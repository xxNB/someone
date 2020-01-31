/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type withNo struct{
    *TreeNode
    No int
}
func isCompleteTree(root *TreeNode) bool {
    nodes := []withNo{{root, 1}}
    for _, i  := range nodes{
        node, v := i, i.No
        if node.Left !=nil{
        nodes = append(nodes, withNo{node.Left, 2*v})
        }
        if node.Right !=nil{
        nodes = append(nodes, withNo{node.Right, 2*v+1})
        }
    }
    if nodes[len(nodes)-1].No == len(nodes){
        return true
    }else{
        return false
    }
}