package regexrover

import "testing"

func TestCSVWriter_Init(t *testing.T) {
	type args struct {
		cacheMaxSize int
	}
	tests := []struct {
		name    string
		w       *CSVWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Init(tt.args.cacheMaxSize); (err != nil) != tt.wantErr {
				t.Errorf("CSVWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSVWriter_WriteWithCache(t *testing.T) {
	type args struct {
		key        string
		value      string
		forceWrite bool
	}
	tests := []struct {
		name    string
		w       *CSVWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.WriteWithCache(tt.args.key, tt.args.value, tt.args.forceWrite); (err != nil) != tt.wantErr {
				t.Errorf("CSVWriter.WriteWithCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSVWriter_WriteAll(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		w       *CSVWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.WriteAll(tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("CSVWriter.WriteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSVWriter_Close(t *testing.T) {
	tests := []struct {
		name    string
		w       *CSVWriter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("CSVWriter.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
