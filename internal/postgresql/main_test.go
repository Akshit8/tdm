package postgresql_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/postgresql"
	_ "github.com/lib/pq"
)

var repo internal.TaskRepository

func TestMain(m *testing.M) {
	var err error

	db, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/tdm?sslmode=disable")
	if err != nil {
		log.Fatalln("Couldn't open DB ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Couldn't ping DB ", err)
	}

	repo = postgresql.NewTask(db)

	os.Exit(m.Run())
}
