package elasticsearch

import (
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

// GetConnection create elasticSearch objec connection
func GetConnection() (*elastic.Client, error) {
	// Instantiate a client instance of the elastic library
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(os.Getenv("ELASTIC_HOST")),
		elastic.SetHealthcheckInterval(5*time.Second), // quit trying after 5 seconds
	)

	if err != nil {
		// (Bad Request): Failed to parse content to map if mapping bad
		log.Println("elastic.NewClient() ERROR: ", err)
		log.Fatalf("quiting connection..")

		return nil, err
	}

	// Print client information
	log.Println("client:", client)
	//log.Println("client TYPE:", reflect.TypeOf(client))

	return client, nil
}
