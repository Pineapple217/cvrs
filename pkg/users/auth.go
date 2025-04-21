package users

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const SECRET = "adsjfkaweijrfsdjfkla"
const DUR = 3600 * 24

type JwtClaims struct {
	Username string `json:"usn"`
	UserId   int    `json:"uid"`
	IsAdmin  bool   `json:"adm"`
	jwt.RegisteredClaims
}

func CreateJWT(user *ent.User) (string, error) {
	expiration := time.Second * time.Duration(DUR)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Username: user.Username,
		UserId:   user.ID,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cvrs",
		},
	})

	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func Auth(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				c.Set("isAuth", false)
				return next(c)
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &JwtClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				slog.Debug("failed to parse jwt", "err", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}
			if !token.Valid {
				slog.Debug("Invalid jwt token")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}
			c.Set("isAuth", true)
			c.Set("claims", *claims)
			return next(c)
		}
	}
}

func IsAuth(c echo.Context) (bool, JwtClaims) {
	isAuth, ok := c.Get("isAuth").(bool)
	if !ok || !isAuth {
		return false, JwtClaims{}
	}
	claims, ok := c.Get("claims").(JwtClaims)
	if !ok { // ok should never be false
		slog.Warn("failed to cast claims")
	}
	return true, claims
}

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		IsAuth, _ := IsAuth(c)
		if !IsAuth {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		IsAuth, claims := IsAuth(c)
		if !IsAuth {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if !claims.IsAdmin {
			return echo.NewHTTPError(http.StatusForbidden)
		}
		return next(c)
	}
}
