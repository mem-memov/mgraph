package mgraph

import (
	"github.com/mem-memov/mbook"
)

type Graph struct {
	storage *mbook.Book
}

// NewGraph creates new graph instance.
//
// Its increment parameter specifies the number of entries for which new space is provided each time it runs out.
func NewGraph(increment uint) (*Graph, error) {
	storage, err := mbook.NewBook(increment * 6)
	if err != nil {
		return nil, err
	}

	g := &Graph {
		storage: storage,
	}

	return g, nil
}

func (g *Graph) Create() (*Vertex, error) {
	return nil, nil
}

func (g *Graph) Read(identifier uint) (*Vertex, error) {
	return nil, nil
}

