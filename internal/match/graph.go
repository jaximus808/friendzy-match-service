package match

import (
	"math"
	"math/rand"
)

type Hierarchy struct {
	layers  []*Graph
	maxk    int
	efLimit int
	mL      float64
	ep      string
}

type Graph struct {
	vertices map[string]*Vertex
	layer    int
}

type Vertex struct {
	neighbors map[string]*Vertex
	user      *User
}

func create_hieracry(layers_n int, maxK int, efLimit int, mL float64) *Hierarchy {
	h := Hierarchy{
		layers:  make([]*Graph, layers_n),
		maxk:    maxK,
		efLimit: efLimit,
		mL:      mL,
		ep:      "",
	}

	for i, _ := range h.layers {
		h.layers[i] = create_graph(i)
	}

	return &h
}

func create_graph(layer int) *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
		layer:    layer,
	}

}

func create_vertex(user *User) *Vertex {
	return &Vertex{
		neighbors: make(map[string]*Vertex),
		user:      user,
	}
}

func (v *Vertex) AddNeighbor(v_new *Vertex) {
	v.neighbors[v_new.user.user_id] = v_new
}

func (g *Graph) AddVertex(v_new *Vertex) {
	g.vertices[v_new.user.user_id] = v_new
}

func (h *Hierarchy) CalcLayerI() int {
	layer_i := math.Floor(-1 * math.Log(rand.Float64()*h.mL))
	layer_i = math.Min(layer_i, float64(len(h.layers)-1))
	return int(layer_i)
}
