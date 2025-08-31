package regexrover

import "testing"

func TestSQLiteWriter_Init(t *testing.T) {
	type args struct {
		cacheMaxSize int
	}
	tests := []struct {
		name    string
		w       *SQLiteWriter
		args    args
		wantErr bool
	}{
		{
			name: "Create DB and table",
			w:    &SQLiteWriter{},
			args: args{
				cacheMaxSize: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Init(tt.args.cacheMaxSize); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteWriter_WriteWithCache(t *testing.T) {
	type args struct {
		key        string
		value      string
		forceWrite bool
	}
	tests := []struct {
		name    string
		w       *SQLiteWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.WriteWithCache(tt.args.key, tt.args.value, tt.args.forceWrite); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteWriter.WriteWithCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteWriter_WriteAll(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		w       *SQLiteWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.WriteAll(tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteWriter.WriteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteWriter_Close(t *testing.T) {
	tests := []struct {
		name    string
		w       *SQLiteWriter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteWriter.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
