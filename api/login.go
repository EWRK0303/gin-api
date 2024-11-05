package api

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("AFaGfgddjtyrjty46$xds")

func Login(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	if err := c.Bind(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}

	tokenString, _ := GenerateToken(loginInfo.Username)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": gin.H{"id": 5, "token": tokenString}})
}

func GenerateToken(username string) (string, error) {
	// 设置 JWT 过期时间
	expirationTime := time.Now().Add(30 * time.Second) // 有效期为 24 小时

	// 创建 JWT 标头
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 JWT 载荷（Payload）
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = expirationTime.Unix()

	// 签署 JWT
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
