// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"testing"

	cobbler "github.com/cobbler/cobblerclient"
)

// Test_ExtractGroupFlags_Parent exercises extractGroupFlags' --parent branch directly
// and does not need a live cobbler server.
func Test_ExtractGroupFlags_Parent(t *testing.T) {
	cmd := newDistroGroupAddCmd()
	err := cmd.Flags().Set("parent", "some-uid")
	cobbler.FailOnError(t, err)

	_, _, _, _, parent, parentSet, err := extractGroupFlags(cmd)
	cobbler.FailOnError(t, err)

	if !parentSet {
		t.Fatal("extractGroupFlags should report parentSet=true after --parent=some-uid")
	}
	if parent != "some-uid" {
		t.Fatalf("extractGroupFlags should return parent=\"some-uid\", got %q", parent)
	}
}
