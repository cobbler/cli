package cmd

import (
	"bytes"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"reflect"
	"strings"
	"testing"
	"time"
)

// Test_convertMapStringToMapInterface confirms the conversion is a plain, literal
// string->interface{} copy with no implicit JSON-sniffing: a flag value that happens to look like
// JSON (e.g. an IPv6-literal kernel option such as "ip=[::1]") must never be silently reinterpreted
// as structured data. Fields that genuinely need list/nested-object values have a dedicated "-json"
// flag (e.g. --autoinstall-meta-json) instead.
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
			name: "a value that looks like a json list is kept as a raw string",
			args: map[string]string{"agama_modules": `["root", "user"]`},
			want: map[string]interface{}{"agama_modules": `["root", "user"]`},
		},
		{
			name: "a value that looks like a json object is kept as a raw string",
			args: map[string]string{"cloud_init_set_hostname": `{"hostname": "example-host"}`},
			want: map[string]interface{}{"cloud_init_set_hostname": `{"hostname": "example-host"}`},
		},
		{
			name: "an IPv6-literal kernel option is not mistaken for malformed json",
			args: map[string]string{"ip": `[::1]`},
			want: map[string]interface{}{"ip": `[::1]`},
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

// Test_mergeJSONObjectFlag covers the shared helper behind every "-json" companion flag
// (--autoinstall-meta-json, --kernel-options-json, --kernel-options-post-json) that lets a value
// carry a list/nested-object the flat key=value stringToString flag syntax can't represent.
func Test_mergeJSONObjectFlag(t *testing.T) {
	t.Run("flag not set leaves base untouched", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		base := map[string]interface{}{"foo": "bar"}

		cobbler.FailOnError(t, mergeJSONObjectFlag(cmd, "extra-json", base))

		if !reflect.DeepEqual(base, map[string]interface{}{"foo": "bar"}) {
			t.Errorf("expected base to be untouched, got %v", base)
		}
	})

	t.Run("flag set merges and overrides", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		cobbler.FailOnError(t, cmd.Flags().Set("extra-json", `{"foo": "overridden", "baz": [1, 2]}`))
		base := map[string]interface{}{"foo": "bar"}

		cobbler.FailOnError(t, mergeJSONObjectFlag(cmd, "extra-json", base))

		if base["foo"] != "overridden" {
			t.Errorf("expected \"foo\" to be overridden, got %v", base["foo"])
		}
		baz, ok := base["baz"].([]interface{})
		if !ok || len(baz) != 2 {
			t.Errorf("expected \"baz\" to carry a decoded 2-element list, got %v", base["baz"])
		}
	})

	t.Run("malformed json returns an error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		cobbler.FailOnError(t, cmd.Flags().Set("extra-json", `{not valid json`))
		base := map[string]interface{}{}

		if err := mergeJSONObjectFlag(cmd, "extra-json", base); err == nil {
			t.Fatal("expected an error for malformed JSON, got nil")
		}
	})
}

// Test_mergeJSONStringMapFlag is mergeJSONObjectFlag's flat map[string]string counterpart, used by
// --template-files-json.
func Test_mergeJSONStringMapFlag(t *testing.T) {
	t.Run("flag not set leaves base untouched", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		base := map[string]string{"motd": "snippets/motd"}

		cobbler.FailOnError(t, mergeJSONStringMapFlag(cmd, "extra-json", base))

		if !reflect.DeepEqual(base, map[string]string{"motd": "snippets/motd"}) {
			t.Errorf("expected base to be untouched, got %v", base)
		}
	})

	t.Run("flag set merges and overrides", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		cobbler.FailOnError(t, cmd.Flags().Set("extra-json", `{"motd": "snippets/motd-v2", "banner": "snippets/banner"}`))
		base := map[string]string{"motd": "snippets/motd"}

		cobbler.FailOnError(t, mergeJSONStringMapFlag(cmd, "extra-json", base))

		want := map[string]string{"motd": "snippets/motd-v2", "banner": "snippets/banner"}
		if !reflect.DeepEqual(base, want) {
			t.Errorf("got = %v, want %v", base, want)
		}
	})

	t.Run("malformed json returns an error", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("extra-json", "", "")
		cobbler.FailOnError(t, cmd.Flags().Set("extra-json", `{not valid json`))
		base := map[string]string{}

		if err := mergeJSONStringMapFlag(cmd, "extra-json", base); err == nil {
			t.Fatal("expected an error for malformed JSON, got nil")
		}
	})
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
