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
	"os"
	"strings"
	"sync"
	"time"
)

type CSVWriter struct {
	cacheMaxSize int
	file         *os.File
	mutex        *sync.Mutex
	writer       *csv.Writer
	cache        map[string]string
}

func (w *CSVWriter) Init(cacheMaxSize int) error {
	timeStr := time.Now().Format(time.RFC3339)
	fileName := fmt.Sprintf("found_matches_%s.csv", strings.ReplaceAll(timeStr, ":", "_"))

	outputFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		errorLog("CSVWriter Init: ", err)
		return err
	}

	w.cacheMaxSize = cacheMaxSize
	w.file = outputFile
	w.mutex = &sync.Mutex{}
	w.writer = csv.NewWriter(outputFile)
	w.cache = make(map[string]string)

	return nil
}

func (w *CSVWriter) WriteWithCache(key string, value string, forceWrite bool) error {
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
			errorLog("WriteWithCache: ", err)
			return err
		}

	}
	return nil
}

func (w *CSVWriter) WriteAll(records [][]string) error {
	w.mutex.Lock()
	err := w.writer.WriteAll(records)
	w.mutex.Unlock()
	return err
}

func (w *CSVWriter) GetFileName() string {
	return w.file.Name()
}

func (w *CSVWriter) Close() error {
	return w.file.Close()
}
