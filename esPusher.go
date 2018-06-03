package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Actor model
// type Actor struct {
// 	id             int
// 	uuid           string
// 	login          string
// 	avatar_url     string
// 	created_at     string
// 	updated_at     string
// 	events_count   int
// 	last_event_at  string
// 	longest_streak int
// }

func printRows(wg *sync.WaitGroup, id int, uuid string, login string) {
	defer wg.Done()
	fmt.Println(id, uuid, login)
}

func main() {
	var (
		id    int
		uuid  string
		login string
	)
	var wg sync.WaitGroup
	db, err := sql.Open("mysql", "root:password@/git_events")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// var queryString = "SELECT id, uuid, login, avatar_url, events_count, last_event_at, longest_streak, created_at, updated_at FROM actors LIMIT 100"
	var queryString = "SELECT id, uuid, login FROM actors LIMIT 100000"
	rows, err := db.Query(queryString)
	defer rows.Close()
	fmt.Println(time.Now())
	for rows.Next() {
		err := rows.Scan(&id, &uuid, &login)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(id, uuid, login)
		wg.Add(1)
		go printRows(&wg, id, uuid, login)
	}
	time.Now()
	fmt.Println(time.Now())
	// time.Sleep(1 * time.Second)
}

// func main() {
// 	db, err := sql.Open("mysql", "root:password@/git_events")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	var v struct {
// 		Data []Actor
// 	}
// 	var queryString = "SELECT id, uuid, login, avatar_url, events_count, last_event_at, longest_streak, created_at, updated_at FROM actors LIMIT 100"
// 	rows, err := db.Query(queryString)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var a Actor
// 		if err := rows.Scan(&a.id, &a.uuid, &a.login, &a.avatar_url, &a.created_at, &a.updated_at, &a.events_count, &a.last_event_at, &a.longest_streak); err != nil {
// 			v.Data = append(v.Data, a)
// 		}
// 	}
// 	fmt.Println(v)
// }
