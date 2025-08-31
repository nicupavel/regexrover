package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	rr "github.com/nicupavel/regexrover/src"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx            context.Context
	crawler        *rr.Crawler
	outputFileName string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetOutputFileName() string {
	return a.outputFileName
}

func (a *App) GetConfig() rr.Config {
	config, err := rr.ReadEnvConfig()
	if err != nil {
		log.Print("Error loading .env file")
	}
	rr.PrintConfig(config)
	return config
}

func (a *App) EmitErrorEvent(format string, optionalData ...interface{}) {
	runtime.EventsEmit(a.ctx, "ErrorEvent", fmt.Sprintf(format, optionalData...))
}

func (a *App) EmitStatusEvent(format string, optionalData ...interface{}) {
	runtime.EventsEmit(a.ctx, "StatusEvent", fmt.Sprintf(format, optionalData...))
}

func (a *App) EmitStateEvent(format string, optionalData ...interface{}) {
	runtime.EventsEmit(a.ctx, "StateEvent", optionalData...)
}

func (a *App) RunCrawler(config rr.Config, urlOrKeywordList []string, mode int) {
	log.Printf("%+v\n", config)
	log.Print(urlOrKeywordList)
	log.Print(mode)

	rr.SetLogger(a.EmitStatusEvent, rr.EmptyLogger, a.EmitStateEvent, a.EmitErrorEvent)
	writer, err := rr.NewWriter(config.OutputDriver, config.MatchOutputChunks)
	if err != nil {
		a.EmitErrorEvent("Error cannot open file to write results %s", err)
	}

	a.outputFileName = writer.GetFileName()

	var crawlUrls []string

	if mode == 1 {
		// Google Search Mode
		a.EmitStatusEvent("Searching keywords on Google")
		crawlUrls = rr.SearchAllKeywords(urlOrKeywordList, config)
	} else {
		// List of Urls Mode
		a.EmitStatusEvent("URL list crawl mode")
		crawlUrls = urlOrKeywordList
	}

	a.crawler, err = rr.NewCrawler(config, writer)

	if err != nil {
		a.EmitErrorEvent("Error cannot initiliza crawler %s", err)
	}

	totalUrls, _ := a.crawler.QueueAdd(crawlUrls)

	if totalUrls > 0 {
		a.EmitStateEvent("%d start URLs to crawl", totalUrls)
	} else {
		a.EmitErrorEvent("Error no URLs to crawl")
	}

	// Handle keyboard interrupt
	syscallChannel := make(chan os.Signal, syscall.SIGINT)
	signal.Notify(syscallChannel, os.Interrupt)
	go func() {
		for sig := range syscallChannel {
			log.Printf("Got signal %v finishing data write.", sig)
			// Stop queue and recursive crawl
			a.crawler.Stop()
			// Write remaining cache
			writer.WriteWithCache("", "", true)
		}
	}()

	err = a.crawler.Run()
	if err != nil {
		a.EmitErrorEvent("Error cannot run crawler %s", err)
	}

	a.EmitStatusEvent("Finished crawling")

	// Write remaining cache
	writer.WriteWithCache("", "", true)
	defer writer.Close()
}

func (a *App) StopCrawler() {
	a.crawler.Stop()
}

func (a *App) GetCrawlerStats() rr.CrawlerStats {
	return a.crawler.GetStats()
}
