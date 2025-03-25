package commonParams

import (
	"context"
	"database/sql"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/timandy/routine"
	"log"
	"sky-take-out/common/properties"
)

var Thread = routine.NewThreadLocal[map[string]interface{}]()
var Db *sql.DB
var RedisDb *redis.Client
var ServerPort string
var ServerHost string
var Ctx = context.Background()
var DatabaseUser string
var DatabasePassword string
var DatabaseName string
var DatabaseHost string
var DatabasePort string
var JwtProperties properties.JwtProperties

func init() {
	var cfg, err = ini.Load("resources/app.conf")
	if err != nil {
		log.Fatalf("Fail to parse 'resources/app.conf': %v", err)
	}

	ServerPort = cfg.Section("server").Key("port").MustString("8080")
	ServerHost = cfg.Section("server").Key("host").MustString("127.0.0.1")
	DatabaseUser = cfg.Section("database").Key("user").MustString("")
	DatabasePassword = cfg.Section("database").Key("password").MustString("")
	DatabaseHost = cfg.Section("database").Key("host").MustString("127.0.0.1")
	DatabasePort = cfg.Section("database").Key("port").MustString("3306")
	DatabaseName = cfg.Section("database").Key("name").MustString("")
	JwtProperties.AdminSecretKey = cfg.Section("jwt").Key("admin-secret-key").MustString("")
	JwtProperties.AdminTtl = cfg.Section("jwt").Key("admin-ttl").MustInt64(0)
	JwtProperties.AdminTokenName = cfg.Section("jwt").Key("admin-token-name").MustString("")

	dsn := DatabaseUser + ":" + DatabasePassword + "@tcp(" + DatabaseHost + ":" + DatabasePort + ")/" + DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     cfg.Section("redis").Key("host").MustString("127.0.0.1") + ":" + cfg.Section("redis").Key("port").MustString("6379"),
		Password: "",
		DB:       0,
	})
}
