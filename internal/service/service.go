package service

// Add takes two integers and returns their sum.
func Add(a, b int) int {
	return a + b
}

// Greet returns a greeting message for the provided name.
func Greet(name string) string {
	if name == "" {
		return "Hello, Guest!"
	}
	return "Hello, " + name + "!"
}
