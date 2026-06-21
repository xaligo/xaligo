package entity

// Document is the root of xaligo DSL.
type Document struct {
	Root *Node
}

type Position struct {
	Offset int
	Line   int
	Column int
}

// Node is a Vue-like tag node.
type Node struct {
	Tag      string
	Attrs    map[string]string
	Children []*Node
	Text     string
	Position Position
	TextRuns []TextRun
}

type TextRun struct {
	Text     string
	Position Position
}

func (n *Node) Attr(key string) string {
	if n == nil || n.Attrs == nil {
		return ""
	}
	return n.Attrs[key]
}
