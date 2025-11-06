package pkg

import (
	"bytes"
	"text/template"

	main_template "github.com/daffadon/fndn/internal/template/main"
	"github.com/daffadon/fndn/internal/types"
)

func ParseTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("tmpl").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func HTTPServerParser(fwk, db, mq, cache string) (string, error) {
	var t types.HTTPServerParse
	switch fwk {
	case "gin":
		t.FrameworkImport = `"github.com/gin-gonic/gin"`
		t.FrameworkRouter = "*gin.Engine"
		t.RouterHandler = "r"
	case "chi":
		t.FrameworkImport = `"github.com/go-chi/chi/v5"`
		t.FrameworkRouter = "*chi.Mux"
		t.RouterHandler = "r"
	case "echo":
		t.FrameworkImport = `"github.com/labstack/echo/v4"`
		t.FrameworkRouter = "*echo.Echo"
		t.RouterHandler = "r"
	case "fiber":
		t.FrameworkImport = `"github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/middleware/adaptor"
		`
		t.FrameworkRouter = "*fiber.App"
		t.RouterHandler = "adaptor.FiberApp(r)"
	case "gorilla/mux":
		t.FrameworkImport = `"github.com/gorilla/mux"`
		t.FrameworkRouter = "*mux.Router"
		t.RouterHandler = "router.WarpWithCorsAndLogger(r)"
	}

	switch db {
	case "postgresql":
		t.DBInstanceType = "*pgxpool.Pool"
		t.DBCloseConnection = "db.Close()"
		t.DBImport = `"github.com/jackc/pgx/v5/pgxpool"`
	case "mariadb":
		t.DBInstanceType = "*sql.DB"
		t.DBCloseConnection = "db.Close()"
		t.DBImport = `"database/sql"`
	case "clickhouse":
		t.DBInstanceType = "clickhouse.Conn"
		t.DBCloseConnection = "db.Close()"
		t.DBImport = `"github.com/ClickHouse/clickhouse-go/v2"`
	case "mongodb", "ferretdb":
		t.DBInstanceType = "*mongo.Client"
		t.DBCloseConnection = "db.Disconnect(ctx)"
		t.DBImport = `"go.mongodb.org/mongo-driver/mongo"`
	case "neo4j":
		t.DBInstanceType = "neo4j.DriverWithContext"
		t.DBCloseConnection = "db.Close(ctx)"
		t.DBImport = `"github.com/neo4j/neo4j-go-driver/v5/neo4j"`
	}

	switch mq {
	case "nats":
		t.MQImport = `"github.com/nats-io/nats.go"`
		t.MQInstance = "mq *nats.Conn,"
		t.MQCloseConn = `
			defer func() {
				if err := mq.Drain(); err != nil {
					logger.Error().Err(err).Msg("Failed to close mq client connection")
				}
			}()`
	case "rabbitmq":
		t.MQImport = `amqp "github.com/rabbitmq/amqp091-go"`
		t.MQInstance = "mq *amqp.Connection,"
		t.MQCloseConn = `
			defer func() {
				if err := mq.Close(); err != nil {
					logger.Error().Err(err).Msg("Failed to close mq client connection")
				}
			}()`
	case "kafka":
		t.MQImport = `"github.com/segmentio/kafka-go"`
		t.MQInstance = "mq *kafka.Conn,"
		t.MQCloseConn = `
			defer func() {
				if err := mq.Close(); err != nil {
					logger.Error().Err(err).Msg("Failed to close mq client connection")
				}
			}()`
	}

	switch cache {
	case "redis":
		t.CacheImport = `"github.com/redis/go-redis/v9"`
		t.CacheInstanceType = "*redis.Client"
		t.CacheCloseConn = `defer func() {
				if err := cache.Close(); err != nil {
					logger.Error().Err(err).Msg("Failed to close cache connection")
				}
			}()`
	case "valkey":
		t.CacheImport = `"github.com/valkey-io/valkey-go"`
		t.CacheInstanceType = "valkey.Client"
		t.CacheCloseConn = `defer cache.Close()`
	}

	return ParseTemplate(main_template.HTTPServerTemplate, t)
}
