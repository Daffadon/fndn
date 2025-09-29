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
