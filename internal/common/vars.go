package common

import (
	"github.com/brocaar/lora-app-server/internal/handler"
	"github.com/brocaar/lora-app-server/internal/nsclient"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// DB holds the database connection pool.
var DB *sqlx.DB

// RedisPool holds the Redis connection pool.
var RedisPool *redis.Pool

// NetworkServerPool holds the connection(s) to the network-server(s).
var NetworkServerPool nsclient.Pool

// Handler holds the handler of events.
var Handler handler.Handler

// ApplicationServerID holds the application-server ID (UUID).
var ApplicationServerID = "6d5db27e-4ce2-4b2b-b5d7-91f069397978"

// ApplicationServerServer holds the hostname:IP of the application-server.
var ApplicationServerServer = "localhost:8001"
