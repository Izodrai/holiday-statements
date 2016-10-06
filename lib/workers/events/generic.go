package events

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get_all(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Get_all",
	})
}

func Get_all_active(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Get_all_active",
	})
}

func Get_all_deactivate(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Get_all_deactivate",
	})
}

func Get_all_archive(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Get_all_archive",
	})
}

func Add(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Add",
	})
}

func Get(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Get",
	})
}

func Update(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Update",
	})
}

func Deactivate(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Deactivate",
	})
}

func Add_participant(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Add_participant",
	})
}

func Del_participant(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Del_participant",
	})
}

func Add_spending(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Add_spending",
	})
}

func Del_spending(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Del_spending",
	})
}

func Update_spending(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "events.Update_spending",
	})
}
