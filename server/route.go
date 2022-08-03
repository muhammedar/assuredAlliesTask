package server

func (s *Server) setRoutes() {

	s.e.GET("/dictionary", s.getPossibleMatches)
	s.e.GET("/statistics", s.getStatistics)
	s.e.POST("/update_dictionary", s.updateDictionary)

}
