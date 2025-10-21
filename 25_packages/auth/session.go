package auth

// private
func extractSession() string {
	return "loggedin"
}

// public
func GetSession() string {
	return extractSession()
}
