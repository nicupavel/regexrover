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

type Writer interface {
	WriterInit
	WriteWithCache(key string, value string, forceWrite bool) error
	WriteAll(records [][]string) error
	GetFileName() string
	Close() error
}

type WriterInit interface {
	Init(cacheMaxSize int) error
}

func NewWriter(writerType string, cacheMaxSize int) (Writer, error) {
	var writer Writer
	var err error

	switch writerType {
	case "csv":
		writer = &CSVWriter{}
		err = writer.Init(cacheMaxSize)
	case "sqlite":
		writer = &SQLiteWriter{}
		err = writer.Init(cacheMaxSize)
	default:
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return writer, nil
}
