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

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	rr "github.com/nicupavel/regexrover/src"
)

func main() {

	config, err := rr.ReadEnvConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rr.PrintConfig(config)

	writer, err := rr.NewWriter(config.OutputDriver, config.MatchOutputChunks)
	if err != nil {
		log.Fatal("Cannot open output file")
	}

	var crawlUrls []string

	// Should Google Search be used to retrieve urls for crawling or get them from a file
	crawlUrls, err = rr.ReadFileLines(config.CrawlUrlsFile)

	if err != nil {
		log.Printf("Cannot read list of Urls from %s using Google Search to build url list.", config.CrawlUrlsFile)

		keywords, ferr := rr.ReadFileLines(config.KeywordsFile)
		if ferr != nil {
			log.Printf("Cannot load keywords from %s. Aborting Google Search.", config.KeywordsFile)
		}
		crawlUrls = rr.SearchAllKeywords(keywords, config)
	}

	crawler, cerr := rr.NewCrawler(config, writer)

	if cerr != nil {
		log.Fatal(cerr)
	}

	totalUrls, _ := crawler.QueueAdd(crawlUrls)

	if totalUrls > 0 {
		log.Printf("%d start URLs to crawl", totalUrls)
	} else {
		log.Fatal("No urls to crawl !")
	}

	// Handle keyboard interrupt
	syscallChannel := make(chan os.Signal, syscall.SIGINT)
	signal.Notify(syscallChannel, os.Interrupt)
	go func() {
		for sig := range syscallChannel {
			log.Printf("Got signal %v finishing data write.", sig)
			// Stop queue and recursive crawl
			crawler.Stop()
			// Write remaining cache
			writer.WriteWithCache("", "", true)
		}
	}()

	errC := crawler.Run()
	if errC != nil {
		log.Printf("failed to run: %s", errC)
	}

	log.Print("Finished running.")

	// Write remaining cache
	writer.WriteWithCache("", "", true)
	defer writer.Close()
}
