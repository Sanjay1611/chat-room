package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sanjay/roam/pkg/versionedController/v1"
	"github.com/sanjay/roam/pkg/versionedController/v1/chatRooms"
	"github.com/sanjay/roam/pkg/versionedController/v1/messages"
	"github.com/sanjay/roam/pkg/versionedController/v1/users"
)

type Server struct {
	Router *gin.Engine
}

var (
	controllers = []v1.Controller{
		users.New(),
		messages.New(),
		chatRooms.New(),
	}
)

func InitializeServer() (*Server, error) {

	router := newRouter()

	s := &Server{
		Router: router,
	}

	s.registerControllers()
	return s, nil
}

// newRouter will initialize all routes
func newRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	router := gin.Default()

	return router
}

func (s *Server) registerControllers() {
	for _, controller := range controllers {
		controller.Register(s.Router)
	}
}

// Run the server on the given port
func (s *Server) Run(port string) {
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, s.Router))
}
