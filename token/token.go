package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authCustomClaims struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(user_id string, user_name string) (string, error) {

	// token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	// if err != nil {
	// 	return "", err
	// }

	token_lifespan := 1

	// claims := jwt.MapClaims{}
	// claims["authorized"] = true
	claims := &authCustomClaims{
		user_id,
		user_name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
			Issuer:    "Sandhu-Sahil",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// []byte(os.Getenv("API_SECRET"))
	return token.SignedString([]byte("Sandhu-Sahil"))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (string, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("Sandhu-Sahil"), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := fmt.Sprintf("%.0f", claims["user_id"])
		return uid, nil
	}
	return "", nil
}
