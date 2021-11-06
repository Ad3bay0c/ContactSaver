package middleware

import (
	"github.com/Ad3bay0c/ContactSaver/server/responses"
	"github.com/Ad3bay0c/ContactSaver/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := strings.TrimSpace(c.Request.Header.Get("Authorization"))
		if header == ""{
			responses.JSON(c, http.StatusUnauthorized, "", "No Token, Authorization denied", nil)
			c.Abort()
			return
		}
		jwtToken, err := services.VerifyToken(header)

		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"message":"", "error": "Invalid Token", "code": 401, "data": nil})
			responses.JSON(c, http.StatusUnauthorized, "", []string{"Invalid Token"}, nil)
			c.Abort()
			return
		}
		token := jwtToken.Claims.(jwt.MapClaims)
		userId := token["userId"].(string)
		expire := token["expiresAt"].(float64)
		if expire < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Expired", "message": ""})
			c.Abort()
			return
		}
		c.Set("userID", userId)
		c.Next()
	}
}
