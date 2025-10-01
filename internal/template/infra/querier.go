package infra_template

const QuerierPgxInfraTemplate = `
package storage_infra

import 	(
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Querier interface {
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
		Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	}

	pgxQuerier struct {
		pgx *pgxpool.Pool
	}
)

func NewQuerier(pool *pgxpool.Pool) Querier {
	return &pgxQuerier{pgx: pool}
}

func (p *pgxQuerier) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.pgx.QueryRow(ctx, sql, args...)
}
func (p *pgxQuerier) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return p.pgx.Exec(ctx, sql, args...)
}

func (p *pgxQuerier) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.pgx.Query(ctx, sql, args...)
}
`

const QuerierMariaDBInfraTemplate = `
package storage_infra

import "database/sql"

type (
	Querier interface {
			QueryRow(ctx context.Context, query string, args ...any) *sql.Row
			Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
			Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	}

	sqlQuerier struct {
			db *sql.DB
	}
)

func NewQuerier(db *sql.DB) Querier {
	return &sqlQuerier{db: db}
}

func (q *sqlQuerier) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return q.db.QueryRowContext(ctx, query, args...)
}

func (q *sqlQuerier) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return q.db.ExecContext(ctx, query, args...)
}

func (q *sqlQuerier) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return q.db.QueryContext(ctx, query, args...)
}
`

const QuerierMongoDBInfraTemplate = `
package storage_infra

import (
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Querier interface {
		Collection(name string) *mongo.Collection
		FindOne(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
		Find(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
		InsertOne(ctx context.Context, collection string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		InsertMany(ctx context.Context, collection string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
		UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		DeleteOne(ctx context.Context, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		DeleteMany(ctx context.Context, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		Disconnect(ctx context.Context) error
	}

	mongoQuerier struct {
		client   *mongo.Client
		database *mongo.Database
	}
)

func NewQuerier(client *mongo.Client) Querier {
	dn := "{{.DatabaseName}}"
	return &mongoQuerier{
		client:   client,
		database: client.Database(dn),
	}
}

func (m *mongoQuerier) Collection(name string) *mongo.Collection {
	return m.database.Collection(name)
}

func (m *mongoQuerier) FindOne(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.database.Collection(collection).FindOne(ctx, filter, opts...)
}

func (m *mongoQuerier) Find(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.database.Collection(collection).Find(ctx, filter, opts...)
}

func (m *mongoQuerier) InsertOne(ctx context.Context, collection string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.database.Collection(collection).InsertOne(ctx, document, opts...)
}

func (m *mongoQuerier) InsertMany(ctx context.Context, collection string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return m.database.Collection(collection).InsertMany(ctx, documents, opts...)
}

func (m *mongoQuerier) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).UpdateOne(ctx, filter, update, opts...)
}

func (m *mongoQuerier) UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).UpdateMany(ctx, filter, update, opts...)
}

func (m *mongoQuerier) DeleteOne(ctx context.Context, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.database.Collection(collection).DeleteOne(ctx, filter, opts...)
}

func (m *mongoQuerier) DeleteMany(ctx context.Context, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.database.Collection(collection).DeleteMany(ctx, filter, opts...)
}

func (m *mongoQuerier) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
`

const QuerierNeo4jInfraTemplate = `
package storage_infra

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type (
	Querier interface {
		ExecuteWrite(ctx context.Context, query string, params map[string]interface{}) (neo4j.ResultWithContext, error)
		ExecuteRead(ctx context.Context, query string, params map[string]interface{}) (neo4j.ResultWithContext, error)
		Close(ctx context.Context) error
	}

	neo4jQuerier struct {
		driver neo4j.DriverWithContext
	}
)

func NewQuerier(driver neo4j.DriverWithContext) Querier {
	return &neo4jQuerier{driver: driver}
}

func (n *neo4jQuerier) ExecuteWrite(ctx context.Context, query string, params map[string]interface{}) (neo4j.ResultWithContext, error) {
	session := n.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	return session.Run(ctx, query, params)
}

func (n *neo4jQuerier) ExecuteRead(ctx context.Context, query string, params map[string]interface{}) (neo4j.ResultWithContext, error) {
	session := n.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	return session.Run(ctx, query, params)
}

func (n *neo4jQuerier) Close(ctx context.Context) error {
	return n.driver.Close(ctx)
}
`
