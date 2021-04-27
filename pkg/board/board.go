package board

import (
	tf "github.com/djthorpe/go-tensorflow/pkg/tensorflow"
	graphviz "github.com/goccy/go-graphviz"
	cgraph "github.com/goccy/go-graphviz/cgraph"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Board struct {
	*graphviz.Graphviz
	*cgraph.Graph
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewBoardWithModel(model *tf.Model) (*Board, error) {
	return NewBoardWithGraph(model.Graph())
}

func NewBoardWithGraph(graph *tf.Graph) (*Board, error) {
	this := new(Board)
	g := graphviz.New()
	if graph, err := g.Graph(); err != nil {
		return nil, err
	} else {
		this.Graphviz = g
		this.Graph = graph
	}

	// Add nodes
	nodes := make(map[string]*cgraph.Node)
	for _, op := range graph.Ops() {
		key := op.Name()
		if node, err := this.Graph.CreateNode(key); err != nil {
			this.Graphviz.Close()
			return nil, err
		} else {
			node.SetComment(op.Type())
			nodes[key] = node
		}
	}

	// Add edges
	for _, op := range graph.Ops() {
		src := nodes[op.Name()]
		for i := 0; i < op.NumOutputs(); i++ {
			output := op.Output(i)
			for _, consumer := range output.Consumers() {
				key := consumer.Operation().Name()
				if dest, exists := nodes[key]; exists {
					this.Graph.CreateEdge("edge", src, dest)
				}
			}
		}
	}

	// TODO: Need to add output nodes

	return this, nil
}

func (this *Board) Dispose() error {
	return this.Graphviz.Close()
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Board) WritePNG(path string) error {
	return this.Graphviz.RenderFilename(this.Graph, graphviz.PNG, path)
}
