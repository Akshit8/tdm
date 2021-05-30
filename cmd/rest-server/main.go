package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/postgresql"
	"github.com/Akshit8/tdm/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatalln("couldn't connect to db: %w", err)
	}
	defer db.Close()

	repo := postgresql.NewTask(db)
	svc := service.Newtask(repo)

	task, err := svc.Create(context.Background(), "new task", internal.PriorityHigh, internal.Dates{})

	fmt.Printf("NEW task %#v, err %s\n", task, err)

	err = svc.Update(context.Background(),
		task.ID,
		"changed task",
		internal.PriorityHigh,
		internal.Dates{
			Due: time.Now().Add(2 * time.Hour),
		},
		false,
	)
	if err != nil {
		log.Fatalln("coulndn't update task", err)
	}

	updatedTask, err := svc.Task(context.Background(), task.ID)
	if err != nil {
		log.Fatalln("couldn't find task", err)
	}

	fmt.Printf("UPDATED task %#v, err %s\n", updatedTask, err)

}

func newDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return db, nil
}
