package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/Pineapple217/cvrs/pkg/ent/user"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	IsAdmin  bool   `json:"isAdmin"`
	Token    string `json:"token"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) Login(c echo.Context) error {
	body := loginRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return err
	}
	if body.Password == "" || body.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	user, err := h.DB.Client.User.Query().
		Where(user.UsernameEQ(body.Username)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusUnauthorized, "failed to authenticate")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "failed to authenticate")
	}

	token, err := users.CreateJWT(user)
	if err != nil {
		return err
	}

	resp := loginResponse{
		IsAdmin:  user.IsAdmin,
		Id:       user.ID,
		Token:    token,
		Username: user.Username,
	}
	return c.JSON(http.StatusOK, resp)
}
