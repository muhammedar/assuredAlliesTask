package essence

//Statistics ...
type Statistics struct {
	AvgRequestsHandlingTime float64 `json:"avgRequestsHandlingTime"`
	RequestHandlingCount    int     `json:"requestHandlingCount"`
	WordCount               int     `json:"WordCount"`
	sumOfRequestTimes       float64
}

//NewStatistics inits a new statistics holder object with zero values
func NewStatistics() *Statistics {
	return &Statistics{
		AvgRequestsHandlingTime: 0,
		RequestHandlingCount:    0,
		WordCount:               0,
	}
}

//Update updates the statistics object
func (s *Statistics) Update(wordCount int, reqTime float64) {
	s.RequestHandlingCount++
	s.sumOfRequestTimes += reqTime
	if s.RequestHandlingCount != 0 {
		s.AvgRequestsHandlingTime = s.sumOfRequestTimes / float64(s.RequestHandlingCount)
	}
}
