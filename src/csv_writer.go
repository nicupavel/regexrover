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
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Writer struct {
	cacheMaxSize int
	file         *os.File
	writer       *csv.Writer
	cache        map[string]string
}

func NewWriter(cacheMaxSize int) (*Writer, error) {
	timeStr := time.Now().Format(time.RFC3339)
	fileName := fmt.Sprintf("found_matches_%s.csv", strings.ReplaceAll(timeStr, ":", "_"))

	outputFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}

	writer := csv.NewWriter(outputFile)
	cache := make(map[string]string)

	return &Writer{
			cacheMaxSize: cacheMaxSize,
			file:         outputFile,
			writer:       writer,
			cache:        cache},
		nil
}

func (w *Writer) WriteWithCache(key string, value string, forceWrite bool) error {
	if len(key) > 0 && len(value) > 0 {
		w.cache[key] = value
	}

	if forceWrite || len(w.cache) > w.cacheMaxSize {
		records := make([][]string, 1)
		for k, v := range w.cache {
			r := []string{k, v}
			records = append(records, r)
		}
		err := w.writer.WriteAll(records)
		if err != nil {
			log.Print("Error: ", err)
			return err
		}
		clear(w.cache)
		records = nil
	}
	return nil
}

func (w *Writer) WriteAll(records [][]string) error {
	err := w.writer.WriteAll(records)
	return err
}

func (w *Writer) Close() error {
	return w.file.Close()
}
