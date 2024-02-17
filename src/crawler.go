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
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Crawler struct {
	c           *colly.Collector
	q           *queue.Queue
	keepRunning bool
	matchRegexp *regexp.Regexp
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

	crawler.matchRegexp = regexp.MustCompile(config.CrawlMatchRegex)
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

	// On every a element which has href attribute call callback
	crawler.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if crawler.keepRunning {
			link := e.Attr("href")
			crawler.c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	crawler.c.OnHTML(config.CrawlTag, func(e *colly.HTMLElement) {
		regexpMatches := crawler.matchRegexp.FindAllString(e.Text, -1)

		if len(regexpMatches) > 0 {
			for _, v := range regexpMatches {
				log.Printf(" * Found match: %s from %q\n", v, e.Request.URL)
				writer.WriteWithCache(v, e.Request.URL.String(), false)
			}
		}
	})

	crawler.c.OnRequest(func(r *colly.Request) {
		//log.Print("Visiting ", r.URL.String())
	})

	return &crawler, nil
}

func (crawler *Crawler) QueueAdd(crawlUrls []string) (int, error) {
	for _, link := range crawlUrls {
		if err := crawler.q.AddURL(link); err != nil {
			log.Printf("failed to add URL: %s\n%v", link, err)
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
