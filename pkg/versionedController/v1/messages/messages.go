package messages

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjay/roam/pkg/models"
	"github.com/sanjay/roam/pkg/store"
	v1 "github.com/sanjay/roam/pkg/versionedController/v1"
)

//MessageController is the API handler for "/messages" route
type MessageController struct {
	v1.GenericController
	routePath string
}

var _ v1.Controller = &MessageController{}

// New creates a new controller for Message API
func New() *MessageController {
	return &MessageController{
		routePath: v1.MessageRoute,
	}
}

// Get responses with the list of available messages
func (controller *MessageController) Get(c *gin.Context) {

	var messages []models.Message
	for _, message := range store.Messages {
		messages = append(messages, message)
	}
	controller.ResponseJSON(c, http.StatusOK, messages)
}

func (controller *MessageController) Post(c *gin.Context) {
	req := postRequestModel{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		controller.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, found := store.Users[req.UserID]
	if !found {
		controller.ResponseError(c, http.StatusConflict, "User does not exist")
		return
	}

	if user.ChatRoomID == "" {
		controller.ResponseError(c, http.StatusConflict, "User has not joined any chat yet. Please enter a chat room first")
		return
	}

	message := models.NewMessage().SetContent(req.Content).SetUserID(req.UserID)
	store.AddMessage(message)
	controller.ResponseJSON(c, http.StatusOK, map[string]interface{}{
		"message": "Message sent",
	})
}

func (controller *MessageController) Register(router *gin.Engine) {
	v1.RegisterController(router, controller, controller.routePath)
}
