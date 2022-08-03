package essence

import (
	"log"
	"os"
)

//Services wraps the bushiness logic services needed in the server
type Services struct {
	//Dictionary represents the dictionary service and its logic
	Dictionary *Dictionary
	//Statistics represents the statistics regarding the systems data and operations
	Statistics *Statistics
}

//NewServices inits the business logic service needed for the main server
func NewServices() (*Services, error) {
	svs := &Services{}

	svs.Statistics = NewStatistics()

	dicLocation := os.Getenv("DICTIONARY")
	log.Println("server: reading dictionary from: ", dicLocation)

	d, err := NewDictionary(dicLocation)
	if err != nil {
		return nil, err
	}

	svs.Dictionary = d
	svs.Statistics.WordCount = len(d.Words)

	return svs, nil
}
