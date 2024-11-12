package cmd

import (
	"bytes"
	cobbler "github.com/cobbler/cobblerclient"
	"testing"
)

func setupClient(t *testing.T) {
	cfgFile = "../testing/.cobbler.yaml"
	initConfig()
	err := generateCobblerClient()
	cobbler.FailOnError(t, err)
}

func FailOnNonEmptyStream(t *testing.T, buffer *bytes.Buffer) {
	if buffer.Available() > 0 {
		t.Fatal("stream wasn't empty!")
	}
}
