package config

import (
	"fmt"
	"gobizdevelop/helper/atdb"
	"log"
	"os"
)

// MongoDB configuration
var MongoString string = os.Getenv("MONGOSTRING")
var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "gobizdev",
}
var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)

// PostgreSQL configuration
var PostgresString string = os.Getenv("POSTGRESSTRING")

var PostgresConn, ErrorPostgresConn = atdb.PostgresConnect(PostgresString)

func init() {
	// MongoDB connection status
	if ErrorMongoconn != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", ErrorMongoconn)
	} else {
		fmt.Println("Successfully connected to MongoDB!")
	}

	// PostgreSQL connection status
	if ErrorPostgresConn != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", ErrorPostgresConn)
	} else {
		fmt.Println("Successfully connected to PostgreSQL!")
	}
}
