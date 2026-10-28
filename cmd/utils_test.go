package cmd

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
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

// Test_printDumpVars_defaultCaseDoesNotGarbleUnhandledKinds guards the same %s-in-a-default-case
// anti-pattern that printField (cmd/root.go) was already fixed for: %s on a value whose kind isn't
// explicitly handled and doesn't implement Stringer renders as literal "%!s(uint=42)" garbage
// instead of its content. The switch's explicit cases cover every kind the XML-RPC client's generic
// interface{} decoder can actually produce today (bool, int64, float64, string, time.Time -
// which implements Stringer -, []interface{}, map[string]interface{}), so this isn't currently
// reachable from a real dump_vars response - but %v is a strictly safer default for whatever kind
// a future Cobbler API field or decoder change adds, with no change in behavior for the one kind
// (string) that reaches it today.
func Test_printDumpVars_defaultCaseDoesNotGarbleUnhandledKinds(t *testing.T) {
	cmd := &cobra.Command{}
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)

	printDumpVars(cmd, map[string]interface{}{"priority": uint(42)})

	got := buf.String()
	if strings.Contains(got, "%!s(") {
		t.Fatalf("printDumpVars garbled an unhandled kind: %q", got)
	}
}
