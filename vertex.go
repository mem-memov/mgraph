package mgraph

type Vertex struct {
	identifier uint
}

func (v *Vertex) Each(f func (vertex *Vertex)) {
	f(v)
}

func (v *Vertex) Connect(vertex *Vertex) error {
	return nil
}

func (v *Vertex) Disconnect(vertex *Vertex) error {
	return nil
}

func (v *Vertex) Delete() {

}