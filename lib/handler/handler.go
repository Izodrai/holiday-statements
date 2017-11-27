package handler

import (
	"../authentification"
	"../workers/events"
	"../workers/users"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	a := router.Group("/authentification")
	{
		a.POST("/login", authentification.Login)
		a.POST("/logout", authentification.Logout)
	}

	ev := router.Group("/events")
	{
		ev.POST("/get/all/all", events.Get_all)
		ev.POST("/get/all/active", events.Get_all_active)
		ev.POST("/get/all/deactivate", events.Get_all_deactivate)
		ev.POST("/get/all/archive", events.Get_all_archive)

		ev.POST("/add", events.Add)
		ev.POST("/get/one/:event_id", events.Get)
		ev.POST("/update/:event_id", events.Update)
		ev.POST("/deactivate/:event_id", events.Deactivate)

		ev.POST("/participant/add/:event_id", events.Add_participant)
		ev.POST("/participant/del/:event_id", events.Del_participant)

		ev.POST("/spending/add/:event_id", events.Add_spending)
		ev.POST("/spending/del/:event_id", events.Del_spending)
		ev.POST("/spending/update/:event_id", events.Update_spending)
	}

	us := router.Group("/users")
	{
		us.POST("/search", users.Search)
		us.POST("/get/connected", users.Get_connected)

		us.POST("/add", users.Add)

		us.POST("/get/all", users.Get_all)
		us.POST("/friends/get/all", users.Friends_get_all)
		us.POST("/friends/add", users.Friends_add)
		// us.POST("/friends/delete", users.Friends_delete)

		us.POST("/update/password", users.Update_password)
		us.POST("/update/email", users.Update_email)
	}
}
