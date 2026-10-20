package cmd

import (
	"reflect"
	"testing"
	"time"
)

func Test_convertMapStringToMapInterface(t *testing.T) {
	tests := []struct {
		name string
		args map[string]string
		want map[string]interface{}
	}{
		{
			name: "flat string value is left unchanged",
			args: map[string]string{"foo": "bar"},
			want: map[string]interface{}{"foo": "bar"},
		},
		{
			name: "json list value is decoded",
			args: map[string]string{"agama_modules": `["root", "user"]`},
			want: map[string]interface{}{"agama_modules": []interface{}{"root", "user"}},
		},
		{
			name: "json object value is decoded",
			args: map[string]string{"cloud_init_set_hostname": `{"hostname": "example-host"}`},
			want: map[string]interface{}{
				"cloud_init_set_hostname": map[string]interface{}{"hostname": "example-host"},
			},
		},
		{
			name: "leading whitespace is tolerated",
			args: map[string]string{"agama_modules": `  ["root"]`},
			want: map[string]interface{}{"agama_modules": []interface{}{"root"}},
		},
		{
			name: "malformed json-looking value falls back to the raw string",
			args: map[string]string{"agama_modules": `[root`},
			want: map[string]interface{}{"agama_modules": `[root`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertMapStringToMapInterface(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertMapStringToMapInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_covertFloatToTime(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "float",
			args:    args{t: 1725887974.610569},
			want:    time.Date(2024, time.September, 9, 13, 19, 34, 610569, time.UTC),
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := covertFloatToUtcTime(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("covertFloatToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covertFloatToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
