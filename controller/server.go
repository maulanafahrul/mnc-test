package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/maulanafahrul/mnc-test/manager"
	"github.com/maulanafahrul/mnc-test/middleware"
)

type Server interface {
	Run()
}

type server struct {
	serviceManager manager.ServiceManager

	srv *gin.Engine
}

func (s *server) Run() {
	// setup sessions
	store := cookie.NewStore([]byte("secret"))
	s.srv.Use(sessions.Sessions("mysession", store))
	s.initController()
	s.srv.Use(middleware.LoggerMiddleware())
	s.srv.Run()

}

func (s *server) initController() {
	NewUserConroller(s.srv, s.serviceManager.GetUserService())
	NewLoginController(s.srv, s.serviceManager.GetLoginService())
}

func NewServer() Server {
	repo := manager.NewRepoManager()
	service := manager.NewServiceManager(repo)

	srv := gin.Default()
	return &server{
		serviceManager: service,
		srv:            srv,
	}
}
