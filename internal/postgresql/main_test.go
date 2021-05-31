package postgresql_test

import (
	"database/sql"
	"log"
	"os"
	"path"
	"testing"

	"github.com/Akshit8/tdm/internal/env"
	"github.com/Akshit8/tdm/internal/postgresql"
	"github.com/Akshit8/tdm/internal/service"
	_ "github.com/lib/pq"
)

var repo service.TaskRepository

func TestMain(m *testing.M) {
	var err error

	err = env.Load(path.Join("../../", ".test.env"))
	if err != nil {
		log.Fatalln("Couldn't Load config", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DB_ADDRESS"))
	if err != nil {
		log.Fatalln("Couldn't open DB", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Couldn't ping DB", err)
	}

	repo = postgresql.NewTask(db)

	os.Exit(m.Run())
}
