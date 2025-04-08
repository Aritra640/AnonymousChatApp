package controller

import (
	ws "github.com/Aritra640/AnonymousChatApp/server/Handlers/WS"
	"github.com/labstack/echo/v4"
)

//Admins can join their group with out needing permissions

func GroupRoutes(e *echo.Echo) {
  e.POST("/group_create")
  e.Any("/groupWS" , ws.WSgroup)
  e.DELETE("/group_delete")
  e.PUT("/group_create_admin") //turn a regular user to admin

  e.POST("/group_join")
  e.PUT("/group_exit")
}
