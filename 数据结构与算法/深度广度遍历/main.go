package main

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

func bfs(root *Node) []int {
	// 在这⾥写代码
	var resp []int
	var queue []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		for _, v := range queue[0].Children {
			queue = append(queue, v)
		}
		resp = append(resp, queue[0].Value)
		queue = queue[1:]
	}
	return resp
}

func dfs(root *Node) []int {
	var stack []*Node
	var resp []int
	check := make(map[*Node]struct{})

	stack = append(stack, root)
	check[root] = struct{}{}
	resp = append(resp, root.Value)

	for len(stack) != 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, v := range c.Children {
			if _, ok := check[v]; !ok {
				check[v] = struct{}{}
				resp = append(resp, v.Value)
				stack = append(stack, c)
				stack = append(stack, v)
				break
			}
		}
	}
	return resp
}

func dfs2(root *Node) []int {
	var resp []int
	var f func(node *Node)
	f = func(node *Node) {
		if node == nil {
			return
		}
		resp = append(resp, node.Value)
		for _, v := range node.Children {
			f(v)
		}
	}
	f(root)
	return resp
}

func main() {
	r := &Node{
		Value: 1,
		Children: []*Node{
			{
				Value: 2,
				Children: []*Node{
					{
						Value: 4,
						Children: []*Node{
							{
								Value:    8,
								Children: []*Node{},
							},
						},
					},
					{
						Value:    5,
						Children: []*Node{},
					},
					{
						Value:    6,
						Children: []*Node{},
					},
				},
			},
			{
				Value: 3,
				Children: []*Node{
					{
						Value:    7,
						Children: []*Node{},
					},
				},
			},
		},
	}
	fmt.Println(bfs(r))
	fmt.Println(dfs(r))
	fmt.Println(dfs2(r))
}
