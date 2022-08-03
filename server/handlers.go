package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (s *Server) getPossibleMatches(c echo.Context) error {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	prefix := c.QueryParam("prefix")
	if prefix == "" {
		return fmt.Errorf("server: failed to read prefix")
	}

	matches := s.services.Dictionary.GetPossibleMatches(prefix)

	s.services.Statistics.Update(s.services.Dictionary.WordCount(), s.calculateRequestTime(float64(startTime)))

	return c.JSON(http.StatusOK, matches)
}

func (s *Server) getStatistics(c echo.Context) error {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	stats := s.services.Statistics

	s.services.Statistics.Update(s.services.Dictionary.WordCount(), s.calculateRequestTime(float64(startTime)))

	return c.JSON(http.StatusOK, stats)
}

func (s *Server) updateDictionary(c echo.Context) error {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	var req []string

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	s.services.Dictionary.Update(req)
	s.services.Statistics.Update(s.services.Dictionary.WordCount(), s.calculateRequestTime(float64(startTime)))

	return c.JSON(http.StatusOK, "server: dictionary was updated successfully")

}
