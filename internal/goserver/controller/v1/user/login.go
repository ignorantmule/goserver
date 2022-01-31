package user

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/pkg/auth"
	"github.com/marmotedu/goserver/pkg/core"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
	"github.com/marmotedu/goserver/pkg/token"
)

// LoginResponse defines the response fields for `/login`.
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginRequest defines the request fields for `/login`.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login login goserver and return a jwt token.
// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json
// "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post].
func (u *UserController) Login(c *gin.Context) {
	// Binding the data with the user struct.
	var r LoginRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	// Get the user information by the login username.
	user, err := u.srv.Users().Get(c, r.Username, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)

		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(user.Password, r.Password); err != nil {
		core.WriteResponse(c, errno.ErrPasswordIncorrect, nil)

		return
	}

	// Sign the json web token.
	t, err := token.Sign(r.Username)
	if err != nil {
		core.WriteResponse(c, errno.ErrToken, nil)

		return
	}

	core.WriteResponse(c, nil, LoginResponse{Token: t})
}
