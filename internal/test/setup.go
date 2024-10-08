package test

import (
	"database/sql"
	"net/http/httptest"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/connection"
	_ "github.com/lib/pq"
)

type Persistence struct {
	Repo *persistence.Queries
	DB   *sql.DB
}

func TestWithServerAndDB() (Server *httptest.Server, database *Persistence, close func()) {
	db, closeDB := connection.OpenPostgresConnection()
	repo := persistence.New(db)
	database = &Persistence{
		Repo: repo,
		DB:   db,
	}
	repository := domain.NewRepository(repo)
	srv := applications.Initialize(&repository)

	Server = httptest.NewServer(srv)
	close = func() {
		Server.Close()
		closeDB()
	}
	return
}
