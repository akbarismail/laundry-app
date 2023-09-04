package delivery

import (
	"clean-code/config"
	"clean-code/delivery/api"
	"clean-code/repository"
	"clean-code/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uomUseCase usecase.UomUseCase
	engine     *gin.Engine
}

func (s *Server) initControllers() {
	api.NewUomController(s.uomUseCase, s.engine)
}

func (s *Server) Run() {
	if err := s.engine.Run(); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDBConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}

	db := con.Conn()

	uomRepository := repository.NewUomRepository(db)

	uomUseCase := usecase.NewUomUseCase(uomRepository)

	engine := gin.Default()
	server := Server{
		uomUseCase: uomUseCase,
		engine:     engine,
	}
	server.initControllers()

	return &server
}
