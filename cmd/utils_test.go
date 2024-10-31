package cmd

import (
	"reflect"
	"testing"
	"time"
)

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
