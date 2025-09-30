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

func HTTPServerParser(fwk, db string) (string, error) {
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
	case "mongodb":
		t.DBInstanceType = "*mongo.Client"
		t.DBCloseConnection = "db.Disconnect(context.TODO())"
		t.DBImport = `"go.mongodb.org/mongo-driver/mongo"`
	}
	return ParseTemplate(main_template.HTTPServerTemplate, t)
}
