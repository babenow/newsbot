package storage

import (
	"context"
	"time"

	"github.com/babenow/newsbot/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

type dbSource struct {
	ID        int64     `db:"id"`
	Name      string    `db:"source_name"`
	FeedURL   string    `db:"feed_url"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var sources []dbSource
	var mSources []model.Source

	if err := s.db.SelectContext(ctx, &sources, `SELECT * FROM sources`); err != nil {
		return nil, err
	}
	for _, source := range sources {
		m := model.Source(source)
		mSources = append(mSources, m)
	}
	return mSources, nil
}

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var source dbSource
	if err := s.db.GetContext(ctx, &source, "SELECT * FROM sources WHERE id = $1", id); err != nil {
		return nil, err
	}
	return (*model.Source)(&source), nil

}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	var id int64

	row := s.db.QueryRowContext(ctx,
		"INSERT INTO sources (source_name,feed_url,created_at)VALUES($1, $2, $3) RETURNING id",
		source.Name,
		source.FeedURL,
		source.CreatedAt,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := s.db.ExecContext(ctx,
		"DELETE FROM sources WHERE id = $1",
		id,
	); err != nil {
		return err
	}

	return nil
}
