package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/url"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var file_path string
	args := os.Args
	if len(args) == 1 {
		panic("Argument required, share the path of chrome history file")
	}
	file_path = args[1]
	// copy the file
	file_path = copyFile(file_path)
	db, err := sql.Open("sqlite3", file_path)
	if err != nil {
		panic(err)
	}
	// to find the schema of the table
	//rows, err := db.Query(`SELECT sql FROM sqlite_master where type="table" AND name="urls"`)

	// aggregate the results for different urls
	rows, err := db.Query("Select url, visit_count from urls where url like '%linkedin%' order by visit_count")
	// club the same base urls

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	urlsCount := map[string]int{}
	for rows.Next() {
		var web_url, count string
		if err := rows.Scan(&web_url, &count); err != nil {
			panic(err)
		}
		fmt.Println(web_url)
		base, err := url.Parse(web_url)
		if err != nil {
			panic(err)

		}

		urlsCount[base.Hostname()]++
	}
	for key, value := range urlsCount {
		fmt.Println(key, "+++", value)
	}
}

func copyFile(file string) string {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	file_name := "xyzzz"
	destination, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, f)
	if err != nil {
		panic(err)
	}
	return file_name
}
