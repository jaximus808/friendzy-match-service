package match

type User struct {
	user_id      string
	trait_vector []float64
}

func create_user(id string, trait_vec []float64) *User {
	return &User{
		user_id:      id,
		trait_vector: trait_vec,
	}
}
