package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" 
)

// DB struct
type DB struct {
	*sql.DB
}

var db *gorp.DbMap
var rawConn *sql.Conn

// Init
func Init() {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	fmt.Println(dbinfo)
	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	rawConn, err = db.Db.Conn(context.Background())
	if err != nil {
		log.Fatalf("Failed to get raw connection: %v", err)
	}

}

// ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

// GetDB returns the gorp DbMap
func GetDB() *gorp.DbMap {
	return db
}

// GetRawConn returns the raw SQL connection
func GetRawConn() *sql.Conn {
	return rawConn
}

// RedisClient ...
var RedisClient *_redis.Client

// InitRedis initializes Redis connection
func InitRedis(selectDB ...int) {
	dbNum := 0
	if len(selectDB) > 0 {
		dbNum = selectDB[0]
	}

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})
}

// GetRedis returns Redis client
func GetRedis() *_redis.Client {
	return RedisClient
}
