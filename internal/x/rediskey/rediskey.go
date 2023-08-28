package rediskey

func K(suffix string) string {
	return "roguestats|" + suffix
}

func ResetToken(userID string) string {
	return K("reset_token|" + userID)
}
