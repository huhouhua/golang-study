package authorization

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	JWTExpireMinute = 720
)

var APITokenSecret = []byte("lqyy8azdsbnjrry6pzkznn52z3v2g734")

type AuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(payload jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GeneralJwtToken(userName string, id string) (string, error) {
	claims := AuthClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JWTExpireMinute * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(APITokenSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckPassword(password, id, hash string) bool {
	t := sha1.New()
	_, err := io.WriteString(t, password+id)
	if err != nil {
		return false
	}
	if fmt.Sprintf("%x", t.Sum(nil)) == hash {
		return true
	}
	return false
}

func VerifyToken(tokenString string, secret []byte) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("parser failed")
		}

		return secret, nil
	})

	if err != nil {
		log.Println("VerifyToken", err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, errors.New("verify token failed")
}

func GeneralSession() string {
	return fmt.Sprintf("seesion-%s-%s", xid.New(), xid.New())
}
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var userName string
		//jwt
		payload, err := VerifyToken(token, APITokenSecret)
		if err != nil {
			log.Println("AuthRequired", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if payload == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		currentUser, ok := (*payload)["username"]
		if currentUser == "" || !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userName = currentUser.(string)
		c.Header("user", userName)
		c.Set("user", userName)
		c.Next()
		return
	}
}
