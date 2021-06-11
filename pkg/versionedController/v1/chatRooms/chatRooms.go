package chatRooms

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjay/roam/pkg/models"
	"github.com/sanjay/roam/pkg/store"
	v1 "github.com/sanjay/roam/pkg/versionedController/v1"
)

//ChatRoomController is the API handler for "/chat_rooms" route
type ChatRoomController struct {
	v1.GenericController
	routePath string
}

var _ v1.Controller = &ChatRoomController{}

// New creates a new controller for ChatRoom API
func New() *ChatRoomController {
	return &ChatRoomController{
		routePath: v1.ChatRoomRoute,
	}
}

// Get responses with the list of available chat rooms
func (controller *ChatRoomController) Get(c *gin.Context) {

	var chatRooms []models.ChatRoom
	for _, chatRoom := range store.ChatRooms {
		chatRooms = append(chatRooms, chatRoom)
	}
	controller.ResponseJSON(c, http.StatusOK, chatRooms)
}

// Post adds a new chat room to the list
func (controller *ChatRoomController) Post(c *gin.Context) {
	req := postRequestModel{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		controller.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	chatRoom := models.NewChatRoom().SetName(req.Name)

	store.AddChatRoom(chatRoom)
	controller.ResponseJSON(c, http.StatusOK, map[string]interface{}{
		"message": "Chat Room created successfully",
	})
}

func (controller *ChatRoomController) Register(router *gin.Engine) {
	v1.RegisterController(router, controller, controller.routePath)
}
