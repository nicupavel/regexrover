// Copyright 2024 Nicu Pavel <npavel@linuxconsulting.ro>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package regexrover

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type SQLiteWriter struct {
	cacheMaxSize int
	tableName    string
	mutex        *sync.Mutex
	db           *sql.DB
	cache        map[string]string
}

func (w *SQLiteWriter) Init(cacheMaxSize int) error {
	timeStr := time.Now().Format(time.RFC3339)
	dbName := fmt.Sprintf("found_matches_%s.sqlite", strings.ReplaceAll(timeStr, ":", "_"))
	tableName := "crawl_results"

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		errorLog("SQLiteWriter Init: ", err)
		return err
	}

	createTableSQL := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				match TEXT,
				url TEXT
			)`, tableName)

	_, err = db.Exec(createTableSQL)

	if err != nil {
		errorLog("Error creating table: ", err)
		return err
	}

	w.cacheMaxSize = cacheMaxSize
	w.tableName = tableName
	w.mutex = &sync.Mutex{}
	w.db = db
	w.cache = make(map[string]string)

	return nil
}

func (w *SQLiteWriter) WriteWithCache(key string, value string, forceWrite bool) error {
	if len(key) > 0 && len(value) > 0 {
		w.cache[key] = value
	}

	if forceWrite || len(w.cache) > w.cacheMaxSize {
		records := make([][]string, 0)
		w.mutex.Lock()
		for k, v := range w.cache {
			r := []string{k, v}
			records = append(records, r)
		}
		clear(w.cache)
		w.mutex.Unlock()

		err := w.WriteAll(records)
		if err != nil {
			errorLog("Error: ", err)
			return err
		}
	}
	return nil
}

func (w *SQLiteWriter) WriteAll(records [][]string) error {
	w.mutex.Lock()
	err := w.insertRows(records)
	w.mutex.Unlock()
	return err
}

func (w *SQLiteWriter) Close() error {
	return w.db.Close()
}

func (w *SQLiteWriter) insertRows(records [][]string) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			errorLog("Error inserting rows: ", err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	insertRowSQL := fmt.Sprintf("INSERT INTO %s (match, url) VALUES (?, ?)", w.tableName)
	stmt, err := tx.Prepare(insertRowSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range records {
		_, err := stmt.Exec(d[0], d[1])
		if err != nil {
			return err
		}
	}

	return nil
}
