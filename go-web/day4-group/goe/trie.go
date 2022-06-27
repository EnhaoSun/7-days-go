package goe

import "strings"

type node struct {
	// 例如/p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true (模糊匹配)
}

func (n *node) insert(pattern string, parts []string, height int) {
	// parts已经遍历完了，插入结束
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 寻找第一个成功匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// 如果匹配上了某个节点，或者当前节点为模糊匹配
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 找出所有匹配成功的节点, 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 使用n.pattern == ""来判断路由规则是否匹配成功
		// 拥有pattern值的，都是叶子节点
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
