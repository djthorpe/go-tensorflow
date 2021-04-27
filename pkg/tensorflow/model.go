package tensorflow

import (
	"io/ioutil"
	"path/filepath"

	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
	core "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/for_core_protos_go_proto"
	proto "google.golang.org/protobuf/proto"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Model struct {
	*tf.SavedModel
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewModelFromFile(path string, tags []string) (*Model, error) {
	this := new(Model)

	if model, err := tf.LoadSavedModel(path, tags, &tf.SessionOptions{}); err != nil {
		return nil, err
	} else {
		this.SavedModel = model
	}

	// Return success
	return this, nil
}

func ModelTags(path string) ([]string, error) {
	bytes, err := ioutil.ReadFile(filepath.Join(path, "saved_model.pb"))
	if err != nil {
		return nil, err
	}

	savedModel := &core.SavedModel{}
	if err := proto.Unmarshal(bytes, savedModel); err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(savedModel.MetaGraphs))
	for _, graph := range savedModel.MetaGraphs {
		tags = append(tags, graph.MetaInfoDef.Tags...)
	}
	return tags, nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *Model) Tags() []string {
	tags := make([]string, 0)
	for k := range this.SavedModel.Signatures {
		tags = append(tags, k)
	}
	return tags
}

func (this *Model) Graph() *Graph {
	return &Graph{this.SavedModel.Graph}
}

func (this *Model) Operation(name string) *Op {
	return (*Op)(this.SavedModel.Graph.Operation(name))
}

func (this *Model) Session() *Session {
	return &Session{this.SavedModel.Session}
}

func (this *Model) Run(in map[Output]*Tensor, fetches []Output, targets []*Op) ([]*Tensor, error) {
	return this.Session().Run(in, fetches, targets)
}
