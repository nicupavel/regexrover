package regexrover

import (
	"reflect"
	"testing"
)

func TestGoogleSearch(t *testing.T) {
	type args struct {
		query      string
		apiKey     string
		cx         string
		startIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    SearchResults
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GoogleSearch(tt.args.query, tt.args.apiKey, tt.args.cx, tt.args.startIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoogleSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchAllKeywords(t *testing.T) {
	type args struct {
		keywords []string
		config   Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchAllKeywords(tt.args.keywords, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchAllKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}
