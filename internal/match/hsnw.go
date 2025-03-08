package match

func search_users(h *Hierarchy, user *User, n_users int) {

}

func insert_user(h *Hierarchy, user *User) {

	new_vertex := create_vertex(user)

	//initial user
	if h.ep == "" {
		for _, layer := range h.layers {
			layer.AddVertex(new_vertex)
		}
		h.ep = user.user_id
		return
	}

	layer_i := h.CalcLayerI()

	ep := h.ep

	//go's troll ass way of making a set
	nearest_neighbors := make(map[string]*struct{})

	for i := 0; i < layer_i; i++ {
		search_users(h, user)
	}

}
