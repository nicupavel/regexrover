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
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey                string
	Cx                    string
	OutputDriver          string
	KeywordsFile          string
	CrawlIgnoreDomains    []string
	CrawlAllowedUrlsRegex string
	CrawlUserAgent        string
	CrawlTag              string
	CrawlMatchRegex       string
	CrawlUrlsFile         string
	CrawlCacheDir         string
	CrawlDepth            int
	CrawlThreads          int
	MatchOutputChunks     int
	MaxSearchResults      int
}

func ReadEnvConfig() (Config, error) {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		return config, err
	}

	config.ApiKey = os.Getenv("GOOGLE_SEARCH_API_KEY")
	config.Cx = os.Getenv("GOOGLE_SEARCH_ID")
	config.KeywordsFile = os.Getenv("KEYWORDS_FILE")
	config.CrawlUrlsFile = os.Getenv("CRAWL_URLS_FILE")
	config.CrawlCacheDir = os.Getenv("CRAWL_CACHE_DIR")
	config.CrawlAllowedUrlsRegex = os.Getenv("CRAWL_ALLOWED_URLS_REGEX")
	config.CrawlUserAgent = os.Getenv("CRAWL_USER_AGENT")
	config.CrawlTag = os.Getenv("CRAWL_TAG")
	config.CrawlMatchRegex = os.Getenv("CRAWL_MATCH_REGEX")
	config.MaxSearchResults = 99 // Max results that Google Custom Search will give

	_crawlDepth := os.Getenv("CRAWL_DEPTH")
	_crawlThreads := os.Getenv("CRAWL_THREADS")
	_crawlIgnoreDomains := os.Getenv("CRAWL_IGNORE_DOMAINS")
	_matchOutputChunks := os.Getenv("MATCH_OUTPUT_CHUNKS")
	_outputDriver := os.Getenv("OUTPUT_DRIVER")

	config.CrawlThreads, err = strconv.Atoi(_crawlThreads)
	if err != nil {
		config.CrawlThreads = 10
	}

	config.CrawlDepth, err = strconv.Atoi(_crawlDepth)
	if err != nil {
		config.CrawlDepth = 1
	}

	config.CrawlIgnoreDomains = strings.Split(strings.TrimSpace(_crawlIgnoreDomains), ",")

	config.MatchOutputChunks, err = strconv.Atoi(_matchOutputChunks)
	if err != nil {
		config.MatchOutputChunks = 5
	}

	if strings.EqualFold(_outputDriver, "sqlite") || strings.EqualFold(_outputDriver, "csv") {
		config.OutputDriver = strings.ToLower(_outputDriver)
	} else {
		config.OutputDriver = "csv"
	}

	return config, nil
}

func PrintConfig(config Config) {
	log.Printf("%+v\n", config)
}
