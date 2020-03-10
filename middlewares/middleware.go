package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/control-money/utilitys"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

type CheckAPIKey struct {
	ApiKey string
}

func (client *CheckAPIKey) Check(context *gin.Context) {
	apiKey := context.Request.Header.Get("api-key")
	if apiKey == client.ApiKey {
		context.Next()
	} else {
		context.AbortWithStatus(401)
	}
}

func SetUserID(context *gin.Context)  {
	err := utilitys.SetUserID(context)
	if err != nil {
		utilitys.Response(context, nil, 401, "parse userID error: " + err.Error())
		context.Abort()
	}
	context.Next()
}