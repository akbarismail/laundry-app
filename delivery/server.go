package delivery

import (
	"clean-code/config"
	"clean-code/delivery/api"
	"clean-code/delivery/middleware"
	"clean-code/manager"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) initMiddleWares() {
	s.engine.Use(middleware.RequestLogMiddleWare(s.log))
}

func (s *Server) initControllers() {
	routerGroup := s.engine.Group("/api/v1")

	api.NewUomController(s.ucManager.UomUC(), routerGroup).Route()
	api.NewCustomerController(s.ucManager.CustomerUC(), routerGroup).Route()
	api.NewEmployeeController(s.ucManager.EmployeeUC(), routerGroup).Route()
	api.NewAuthController(s.ucManager.AuthUC(), s.ucManager.UserUC(), routerGroup).Route()
}

func (s *Server) Run() {
	s.initMiddleWares()
	s.initControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}

	rm := manager.NewRepoManager(infraManager)
	ucm := manager.NewUseCaseManager(rm)

	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)
	engine := gin.Default()
	log := logrus.New()

	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
		log:       log,
	}
}
