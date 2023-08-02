package controller

import (
	"TaskApp/model"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func LoginController(c echo.Context) error {
	var doLogin model.DoLogin
	if err := c.Bind(&doLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	// Lakukan validasi data, misalnya pastikan email dan password tidak kosong
	if doLogin.Email != "nayeon535@gmail.com" || doLogin.Password != "atulcantik" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email and password are required",
		})
	}

	// Lakukan proses sign-in, misalnya dengan memeriksa email dan password di database
	// Jika sign-in berhasil, berikan response sesuai kebutuhan aplikasi
	// Jika sign-in gagal, berikan response sesuai kebutuhan aplikasi

	// Contoh response jika sign-in berhasil
	claims := &jwtCustomClaims{
		Name:  "Shooofia",
		Id:    1,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success login",
		Data: echo.Map{
			"token": t,
		},
	})
}
