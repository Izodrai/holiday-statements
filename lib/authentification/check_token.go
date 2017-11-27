package authentification

import (
	"../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

func Check_token_and_json(c *gin.Context, json *tools.Request, need_admin bool) bool {

	if c.BindJSON(json) == nil {

		if json.Data != nil {
			switch t := json.Data.(type) {
			case string:
				if !utf8.Valid([]byte(t)) {
					c.JSON(http.StatusForbidden, gin.H{
						"code": http.StatusForbidden,
						"msg":  "invalid format you not transmit utf8",
					})
					return false
				}
			}
		}

		if u, ok := tools.Connected_users[json.User_id]; ok {
			if json.Token == u.Token {

				u.Update_activity()

				if !need_admin {
					return true
				}

				if _, ok_admin := tools.Admins[u.Login]; ok_admin {
					return true
				}

				c.JSON(http.StatusForbidden, gin.H{
					"code": http.StatusForbidden,
					"msg":  "invalid rights",
				})
				return false
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"code": http.StatusForbidden,
			"msg":  "invalid token (maybe you are not connected)",
		})
		return false
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "invalid parameters",
	})
	return false
}
