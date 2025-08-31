package regexrover

import (
	"reflect"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFileLines(t *testing.T) {
	type args struct {
		lines []string
		path  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AppendFileLines(tt.args.lines, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("AppendFileLines() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
