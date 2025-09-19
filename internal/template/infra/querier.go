package infra_template

const QuerierInfraTemplate = `
package storage_infra

import 	"github.com/jackc/pgx/v5"

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
