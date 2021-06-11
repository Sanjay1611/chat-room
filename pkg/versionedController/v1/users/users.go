package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjay/roam/pkg/models"
	"github.com/sanjay/roam/pkg/store"
	v1 "github.com/sanjay/roam/pkg/versionedController/v1"
)

//UserController is the API handler for "/user" route
type UserController struct {
	v1.GenericController
	routePath string
}

var _ v1.Controller = &UserController{}

// New creates a new controller for User API
func New() *UserController {
	return &UserController{
		routePath: v1.UserRoute,
	}
}

// Get responses with the list of available users
func (controller *UserController) Get(c *gin.Context) {

	var users []models.User
	for _, user := range store.Users {
		users = append(users, user)
	}
	controller.ResponseJSON(c, http.StatusOK, users)
}

// Create user
func (controller *UserController) Post(c *gin.Context) {
	req := postRequestModel{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		controller.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.NewUser().SetName(req.Name).SetAge(req.Age)
	store.AddUser(user)
	controller.ResponseJSON(c, http.StatusOK, map[string]interface{}{
		"message": "User created successfully",
	})
}

// Enter chat room
func (controller *UserController) Put(c *gin.Context) {
	req := putRequestModel{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		controller.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	_, found := store.ChatRooms[req.ChatRoomID]
	if !found {
		controller.ResponseError(c, http.StatusConflict, "Chat room  does not exist")
		return
	}

	user, found := store.Users[req.UserID]
	if !found {
		controller.ResponseError(c, http.StatusConflict, "User does not exist")
		return
	}

	user.ChatRoomID = req.ChatRoomID
	store.Users[req.UserID] = user
	controller.ResponseJSON(c, http.StatusOK, map[string]interface{}{
		"message": "User added to chat room successfully",
	})
}

func (controller *UserController) Register(router *gin.Engine) {
	v1.RegisterController(router, controller, controller.routePath)
}
