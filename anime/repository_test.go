package anime

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository(t *testing.T) {

	t.Run("GetAnimes(limit = 2, offset = 0) returns []Anime{{...}, {...}}", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		if err != nil {
			t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		animes := []Anime{
			{ID: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
			{ID: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		}

		rows := sqlmock.NewRows([]string{"id", "title", "description", "plot"}).
			AddRow(animes[0].ID, animes[0].Title, animes[0].Description, animes[0].Plot).
			AddRow(animes[1].ID, animes[1].Title, animes[1].Description, animes[1].Plot)

		mock.ExpectQuery("^SELECT (.+) FROM animes LIMIT (.+) OFFSET (.+)$").
			WithArgs(2, 0).
			WillReturnRows(rows)

		repo := mySQLRepository{db}

		got, err := repo.GetAnimes(context.Background(), 2, 0)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := animes

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("GetAnimes(limit = 2, offset = 1) returns []Anime{{...}, {...}}", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		if err != nil {
			t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		animes := []Anime{
			{ID: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
		}

		rows := sqlmock.NewRows([]string{"id", "title", "description", "plot"}).
			AddRow(animes[0].ID, animes[0].Title, animes[0].Description, animes[0].Plot)

		mock.ExpectQuery("^SELECT (.+) FROM animes LIMIT (.+) OFFSET (.+)$").
			WithArgs(2, 1).
			WillReturnRows(rows)

		repo := mySQLRepository{db}

		got, err := repo.GetAnimes(context.Background(), 2, 1)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := animes

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("GetAnime(id = 1) returns &Anime{}", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		if err != nil {
			t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		anime := &Anime{
			ID:          1,
			Title:       "Tokyo Ghoul",
			Description: "Testing...",
			Plot:        "Testing...",
		}

		rows := sqlmock.NewRows([]string{"id", "title", "description", "plot"}).
			AddRow(anime.ID, anime.Title, anime.Description, anime.Plot)

		mock.ExpectQuery("^SELECT (.+) FROM animes WHERE id = (.+)$").
			WithArgs(1).
			WillReturnRows(rows)

		repo := mySQLRepository{db}

		got, err := repo.GetAnime(context.Background(), 1)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := anime

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}
