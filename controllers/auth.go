package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/EkyGalih/dukcapil-web/config"
	"github.com/EkyGalih/dukcapil-web/utils"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in, omitempty"`
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/auth/login.html", gin.H{
		"title": "Login",
	})
}

func LoginSubmit(c *gin.Context) {
	var in LoginPayload
	if err := c.ShouldBind(&in); err != nil {
		c.HTML(http.StatusBadRequest, "pages/auth/login.html", gin.H{
			"title": "Login",
			"error": "Input tidak valid",
		})
		return
	}

	// call Flask/FastAPI login endpoint
	var out LoginResponse
	status, err := utils.AuthorizedJSON(http.MethodPost, config.API_BASE_URL+config.API_LOGIN_PATH, "", in, &out)
	if err != nil || status >= 400 || out.AccessToken == "" {
		c.HTML(http.StatusUnauthorized, "pages/auth/login.html", gin.H{
			"title": "Login",
			"error": "Login gagal",
		})
		return
	}

	// set HttpOnly cookie with token
	maxAge := config.COOKIE_MAX_AGE_S
	// if backend tells us TTL via ExpiresIn, prefer that
	if out.ExpiresIn > 0 && out.ExpiresIn < int64(maxAge) {
		maxAge = int(out.ExpiresIn)
	}

	c.SetSameSite(http.SameSite(config.COOKIE_SAMESITE))
	c.SetCookie(config.COOKIE_NAME, out.AccessToken, maxAge, "/", config.COOKIE_DOMAIN, config.COOKIE_SECURE, config.COOKIE_HTTP_ONLY)

	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	// clear cookie
	c.SetSameSite(http.SameSite(config.COOKIE_SAMESITE))
	c.SetCookie(config.COOKIE_NAME, "", -1, "/", config.COOKIE_DOMAIN, config.COOKIE_SECURE, config.COOKIE_HTTP_ONLY)
	c.Redirect(http.StatusFound, "/login")
}
