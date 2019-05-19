package anime

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Close()
	Ping() error
	GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error)
	GetAnime(ctx context.Context, id uint64) (*Anime, error)
}

type mySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(url string) (Repository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &mySQLRepository{db}, nil
}

func (r *mySQLRepository) Close() {
	r.db.Close()
}

func (r *mySQLRepository) Ping() error {
	return r.db.Ping()
}

func (r *mySQLRepository) GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT * FROM animes LIMIT ? OFFSET ?",
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var animes []Anime

	for rows.Next() {
		a := &Anime{}
		if err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.Plot); err != nil {
			return nil, err
		}
		animes = append(animes, *a)
	}

	return animes, nil
}

func (r *mySQLRepository) GetAnime(ctx context.Context, id uint64) (*Anime, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM animes WHERE id = ?", id)

	var anime Anime
	if err := row.Scan(&anime.ID, &anime.Title, &anime.Description, &anime.Plot); err != nil {
		return nil, err
	}

	return &anime, nil
}
