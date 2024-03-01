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
	"fmt"
	"log"
	"regexp"
	"sync/atomic"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type CrawlerStats struct {
	scraped  uint32
	requests uint32
	pending  int64
	matches  uint32
}

type Crawler struct {
	c           *colly.Collector
	q           *queue.Queue
	keepRunning bool
	matchRegexp *regexp.Regexp
	stats       CrawlerStats
}

func NewCrawler(config Config, writer Writer) (*Crawler, error) {
	crawler := Crawler{}
	var err error

	// Create crawl queue
	crawler.q, err = queue.New(
		config.CrawlThreads,
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)

	if err != nil {
		return &crawler, err
	}

	crawler.keepRunning = true
	crawler.matchRegexp = regexp.MustCompile(config.CrawlMatchRegex)
	crawler.stats = CrawlerStats{
		scraped:  0,
		requests: 0,
		pending:  0,
		matches:  0,
	}

	crawler.c = colly.NewCollector(
		colly.MaxDepth(config.CrawlDepth),
		colly.URLFilters(
			regexp.MustCompile(config.CrawlAllowedUrlsRegex),
			regexp.MustCompile(`^https?://google\.com/.*`),
		),

		colly.DisallowedDomains(config.CrawlIgnoreDomains...),
	)
	if len(config.CrawlCacheDir) > 0 {
		crawler.c.CacheDir = config.CrawlCacheDir
	}
	crawler.c.UserAgent = config.CrawlUserAgent

	// On every a element which has href attribute visit again
	// On every defined crawl tag match regex and save
	/*
		crawler.c.OnHTML("html", func(e *colly.HTMLElement) {
			e.ForEach("a", func(_ int, el *colly.HTMLElement) {
				if crawler.keepRunning {
					link := el.Attr("href")
					crawler.c.Visit(el.Request.AbsoluteURL(link))
					log.Print("Visiting ", el.Request.AbsoluteURL(link))
				}
			})
			e.ForEach(config.CrawlTag, func(_ int, el *colly.HTMLElement) {
				regexpMatches := crawler.matchRegexp.FindAllString(e.Text, -1)
				if len(regexpMatches) > 0 {
					for _, v := range regexpMatches {
						log.Printf(" * Found match: %s from %q\n", v, e.Request.URL)
						writer.WriteWithCache(v, e.Request.URL.String(), false)
					}
				}
			})
		})
	*/

	crawler.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if crawler.keepRunning {
			link := e.Attr("href")
			e.Request.Visit(link)
			if config.CrawlLog {
				log.Print("Visiting ", e.Request.AbsoluteURL(link))
			}
		}
	})

	crawler.c.OnHTML(config.CrawlTag, func(e *colly.HTMLElement) {
		regexpMatches := crawler.matchRegexp.FindAllString(e.Text, -1)
		if len(regexpMatches) > 0 {
			for _, v := range regexpMatches {
				atomic.AddUint32(&crawler.stats.matches, 1)
				writer.WriteWithCache(v, e.Request.URL.String(), false)
				if config.CrawlLog {
					log.Printf(" * Found match: %s from %q\n", v, e.Request.URL)
				}
			}
		}
	})

	crawler.c.OnRequest(func(r *colly.Request) {
		atomic.AddUint32(&crawler.stats.requests, 1)
		atomic.AddInt64(&crawler.stats.pending, 1)
	})

	crawler.c.OnScraped(func(res *colly.Response) {
		atomic.AddUint32(&crawler.stats.scraped, 1)
		atomic.AddInt64(&crawler.stats.pending, -1)
		if !config.CrawlLog {
			crawler.printStats()
		}
	})

	return &crawler, nil
}

func (crawler *Crawler) QueueAdd(crawlUrls []string) (int, error) {
	for _, link := range crawlUrls {
		if err := crawler.q.AddURL(link); err != nil {
			log.Printf("failed to add URL: %s\n%v", link, err)
		} else {
			crawler.stats.pending++
		}
	}

	return crawler.q.Size()
}

func (crawler *Crawler) Stop() {
	crawler.keepRunning = false
	crawler.q.Stop()
}

func (crawler *Crawler) Run() error {
	return crawler.q.Run(crawler.c)
}

func (crawler *Crawler) GetStats() CrawlerStats {
	return crawler.stats
}

func (crawler *Crawler) printStats() {
	fmt.Printf("\033[2K\r%+v", crawler.stats)
}
