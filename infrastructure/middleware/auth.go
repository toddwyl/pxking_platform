package middleware

import (
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/common"
	"github.com/go-eagle/eagle/infrastructure/common/errcode"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth authorize user
func Auth(paths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ignore some path
		// eg: register, login, logout
		if len(paths) > 0 {
			path := c.Request.URL.Path
			pathsStr := strings.Join(paths, "|")
			reg := regexp.MustCompile("(" + pathsStr + ")")
			if reg.MatchString(path) {
				return
			}
		}

		// Parse the json web token.
		ctx, err := common.ParseRequest(c, global.Conf.JwtSecret)
		if err != nil {
			common.NewResponse().Error(c, errcode.ErrInvalidToken)
			c.Abort()
			return
		}

		// set uid to context
		c.Set("uid", ctx.UserID)

		c.Next()
	}
}
