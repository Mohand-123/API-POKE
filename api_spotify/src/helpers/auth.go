package helpers

func CheckAuth(token string) bool {
	// Exemple simple
	return token == "secret"
}
