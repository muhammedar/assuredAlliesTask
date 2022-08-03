package server

import (
	"compress/flate"
	"fmt"
	"net/http"
	"time"

	"dictionaryManager.com/essence"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Server the backend server
type Server struct {
	//services pointer to all business logic services
	services *essence.Services
	//echo is the used web framework
	e *echo.Echo
}

//NewServer create a new server
func NewServer() (*Server, error) {
	s := &Server{}

	svs, err := essence.NewServices()
	if err != nil {
		return nil, fmt.Errorf("server: failed to init services: %v", err)
	}

	s.services = svs
	s.e = echo.New()
	return s, nil
}

//Start start the backend server
func (s *Server) Start() error {
	s.e = echo.New()
	s.e.Debug = true
	logconf := middleware.DefaultLoggerConfig
	logconf.Format = "${time_rfc3339_nano} ${id} ${remote_ip} ${host} ${method} ${uri} ${status} ${error} \n"

	s.e.Use(middleware.LoggerWithConfig(logconf))
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.Secure())
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: flate.BestCompression}))

	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete,
			http.MethodOptions},
	}))

	s.setRoutes()

	return s.e.StartTLS(":8080", "server.crt", "server.key")

}

func (s *Server) calculateRequestTime(startTime float64) float64 {
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	return float64(endTime) - startTime
}
