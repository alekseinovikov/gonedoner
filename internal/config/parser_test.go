package config

import (
	"reflect"
	"testing"
)

func TestParseFileByPath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    Config
	}{
		{
			"Test Configuration Parsing",
			args{filePath: "testdata/config1.yml"},
			nil,
			Config{
				Version: "0.0.1",
				Services: map[string]Service{
					"ping": {
						Command: "ping",
						Args:    []string{"google.com"},
						Env: map[string]string{
							"ENV":   "production",
							"DEBUG": "false",
						},
					},
					"pong": {
						Command: "pong",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseFileByPath(tt.args.filePath)
			if !reflect.DeepEqual(got, tt.wantErr) {
				t.Errorf("ParseFileByPath() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want) {
				t.Errorf("ParseFileByPath() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
