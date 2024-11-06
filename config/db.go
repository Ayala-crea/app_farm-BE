package config

import (
	"gobizdevelop/helper/atdb"
	"os"
	"strconv"
)

var MongoString string = os.Getenv("MONGODB_URI")
var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "gobizdev",
}
var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)

var (
	PostgresHost     = os.Getenv("POSTGRES_HOST")
	PostgresPort, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	PostgresUser     = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresDBName   = os.Getenv("POSTGRES_DBNAME")
	PostgresSSLMode  = os.Getenv("POSTGRES_SSLMODE")
)

// var PostgresConn *sql.DB
// var ErrorPostgresConn error

var PostgresConn, ErrorPostgresConn = atdb.PostgresConnect(
	PostgresHost,
	PostgresPort,
	PostgresUser,
	PostgresPassword,
	PostgresDBName,
	PostgresSSLMode,
)

// // MongoDB connection status
// if ErrorMongoconn != nil {
// 	log.Fatalf("Failed to connect to MongoDB: %v", ErrorMongoconn)
// } else {
// 	log.Println("Successfully connected to MongoDB.")
// }

// // PostgreSQL connection status
// if ErrorPostgresConn != nil {
// 	log.Fatalf("Failed to connect to PostgreSQL: %v", ErrorPostgresConn)
// } else {
// 	log.Println("Successfully connected to PostgreSQL.")
// }
