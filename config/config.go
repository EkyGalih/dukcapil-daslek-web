package config

const (
	API_BASE_URL = "http://127.0.0.1:8000"
	API_LOGIN_PATH = "/auth/login"
)

const (
	COOKIE_NAME = "access_token"
	COOKIE_MAX_AGE_S = 60 * 60 * 24 // 1 day
	COOKIE_SECURE = false 	// true if https
	COOKIE_HTTP_ONLY = true // false if https
	COOKIE_SAMESITE = 2 // 0: default, 1: lax, 2: strict, 3: none
	COOKIE_DOMAIN = "" // e.g ".example.com" if needed
)