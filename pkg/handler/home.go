package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}

func Login(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	defer c.Request().Body.Close()
	fmt.Println(user)

	if err != nil {
		log.Println("Failed Processing")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	fmt.Println(user)
	if user.Username == "admin" && user.Password == "1234" {
		cookie :=&http.Cookie{}

		cookie.Name = "sessionId"
		cookie.Value = "login"
		cookie.Expires = time.Now().Add(48*time.Hour)
		c.SetCookie(cookie)

		token, err := GenerateToken()
		if err != nil {
			log.Println("Error creating JWT token")
			return echo.NewHTTPError(http.StatusInternalServerError, "Error with creating Token")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in",
			"token": token,
		})
	}

	return nil
}

func GenerateToken() (string, error) {
	var hmacSecret []byte
	claims := JwtClaims{"Phuoc", jwt.StandardClaims{
		Id: "phuoc",
		ExpiresAt: time.Now().Add(48*time.Hour).Unix(),
	}}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	hmacSecret = []byte("9pwmXFcerw")
	// Sign and get the complete encoded token as a string using the secret
	token, err := rawToken.SignedString(hmacSecret)
	if err != nil {
		fmt.Println("Token is invalid", err)
		return "Error", nil
	}
	fmt.Println(token)
	return token, nil
}
