package regexrover

import (
	"reflect"
	"testing"
)

func TestReadEnvConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Load .env",
			want:    Config{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadEnvConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadEnvConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadEnvConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintConfig(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Print config",
			args: args{
				Config{
					MaxSearchResults: 99,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintConfig(tt.args.config)
		})
	}
}
