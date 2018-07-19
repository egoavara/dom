package xpath

import "encoding/xml"

type INode interface {
	//
	INodePrev() INode
	INodeNext() INode
	INodeFirst() INode
	INodeLast() INode
	//
	INodeParent() INode
	INodeChildrun() []INode
}

type (
	INodeElement interface {
		INode
		INodeName() xml.Name
	}

	INodeAttribute interface {
		INode
		INodeAttr() (name xml.Name, value string)
	}

	INodeContents interface {
		INode
		INodeContents() string
	}

	INodeProcessInstruction interface {
		INode
		INodeProcessInstruction() (key string, value string)
	}

	INodeComment interface {
		INode
		INodeComment() string
	}

	INodeDirective interface {
		INode
		INodeDirective() string
	}
)
