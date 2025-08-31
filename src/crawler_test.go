package regexrover

import (
	"reflect"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	type args struct {
		config Config
		writer Writer
	}
	tests := []struct {
		name    string
		args    args
		want    *Crawler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCrawler(tt.args.config, tt.args.writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCrawler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCrawler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrawler_QueueAdd(t *testing.T) {
	type args struct {
		crawlUrls []string
	}
	tests := []struct {
		name    string
		crawler *Crawler
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.crawler.QueueAdd(tt.args.crawlUrls)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crawler.QueueAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Crawler.QueueAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrawler_Stop(t *testing.T) {
	tests := []struct {
		name    string
		crawler *Crawler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.crawler.Stop()
		})
	}
}

func TestCrawler_Run(t *testing.T) {
	tests := []struct {
		name    string
		crawler *Crawler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.crawler.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Crawler.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
