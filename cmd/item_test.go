package cmd

import (
	"io"
	"net/http"
	"strings"
	"testing"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// cannedHTTPClient is a minimal cobbler.HTTPClient implementation that always
// answers with a fixed, pre-baked XML-RPC response body regardless of what
// was requested. It lets resolveUID's server-response handling be exercised
// (find_items faults, and malformed find_items results) without requiring a
// live Cobbler server or a mock framework: cobblerclient.NewClient accepts
// any type satisfying the small HTTPClient interface.
type cannedHTTPClient struct {
	body string
}

func (c *cannedHTTPClient) Post(string, string, io.Reader) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(c.body)),
	}, nil
}

func newCannedClient(body string) cobbler.Client {
	client := cobbler.NewClient(&cannedHTTPClient{body: body}, cobbler.ClientConfig{
		URL:      "http://resolveuid-canned-test.invalid/",
		Username: "cobbler",
		Password: "cobbler",
	})
	client.Token = "canned-token"
	return client
}

// findItemsFaultResponse is a well-formed XML-RPC fault, as find_items would
// yield e.g. on an invalid session token.
const findItemsFaultResponse = `<?xml version="1.0"?>
<methodResponse>
<fault>
<value><struct>
<member><name>faultCode</name><value><int>1</int></value></member>
<member><name>faultString</name><value><string>simulated find_items failure</string></value></member>
</struct></value>
</fault>
</methodResponse>`

// findItemsNonStructResponse simulates find_items(..., expand=true) returning
// a result entry that isn't a struct, which is unexpected but not something
// resolveUID can rule out just from the type signature of FindItems.
const findItemsNonStructResponse = `<?xml version="1.0"?>
<methodResponse>
<params><param><value><array><data>
<value><string>not-a-struct-result</string></value>
</data></array></value></param></params>
</methodResponse>`

// findItemsNoUIDResponse simulates find_items(..., expand=true) returning a
// single matching struct that (unexpectedly) has no "uid" member.
const findItemsNoUIDResponse = `<?xml version="1.0"?>
<methodResponse>
<params><param><value><array><data>
<value><struct>
<member><name>name</name><value><string>some-distro</string></value></member>
</struct></value>
</data></array></value></param></params>
</methodResponse>`

// TestResolveUID_UIDGivenReturnsImmediately exercises the fast path of
// resolveUID: when --uid is supplied it must be returned as-is without any
// server round-trip, so this is safe to run without a live Cobbler server
// (the client is never touched).
func TestResolveUID_UIDGivenReturnsImmediately(t *testing.T) {
	got, err := resolveUID(nil, "distro", "", "some-uid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "some-uid" {
		t.Fatalf("expected %q, got %q", "some-uid", got)
	}
}

// TestResolveUID_UIDTakesPrecedenceOverName confirms that when both --name and
// --uid are supplied, --uid wins and no name resolution (server round-trip)
// happens.
func TestResolveUID_UIDTakesPrecedenceOverName(t *testing.T) {
	got, err := resolveUID(nil, "distro", "some-name", "some-uid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "some-uid" {
		t.Fatalf("expected %q, got %q", "some-uid", got)
	}
}

// TestResolveUID_NoIdentifierGiven confirms that omitting both --name and
// --uid produces a clear error without contacting the server.
func TestResolveUID_NoIdentifierGiven(t *testing.T) {
	_, err := resolveUID(nil, "distro", "", "")
	if err == nil {
		t.Fatal("expected an error when neither --name nor --uid is given")
	}
	if !strings.Contains(err.Error(), "--name") || !strings.Contains(err.Error(), "--uid") {
		t.Fatalf("expected error to mention --name/--uid, got: %v", err)
	}
}

// TestResolveUID_NotFound exercises the zero-match path against a live
// server: resolving a name that doesn't exist must produce a clear
// "no <what> found" error.
func TestResolveUID_NotFound(t *testing.T) {
	setupClient(t)
	_, err := resolveUID(&Client, "distro", "this-distro-does-not-exist-resolveuid-test", "")
	if err == nil {
		t.Fatal("expected an error for a nonexistent name")
	}
	if !strings.Contains(err.Error(), "no distro found") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// TestResolveUID_Ambiguous exercises the multiple-match path against a live
// server. NetworkInterface names are only unique per-system (not globally),
// so creating two systems each with an interface of the same name is the
// simplest way to produce a genuine ambiguous-name collision.
func TestResolveUID_Ambiguous(t *testing.T) {
	setupClient(t)

	// Each cleanup is registered via t.Cleanup() immediately after its resource is
	// successfully created (not batched at the end), and uses t.Errorf rather than
	// cobbler.FailOnError/t.Fatal: a failure partway through setup or in one cleanup
	// step must not skip the others and leak state on the live server.
	systemAName := "test-resolveuid-ambiguous-sys-a"
	systemBName := "test-resolveuid-ambiguous-sys-b"
	systemA, err := createSystem(Client, systemAName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(systemA.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemAName, err)
		}
	})
	systemB, err := createSystem(Client, systemBName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(systemB.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemBName, err)
		}
	})

	ifaceName := "eth-resolveuid-ambiguous"
	ifaceA := cobbler.NewNetworkInterface()
	ifaceA.Name = ifaceName
	ifaceA.SystemUid = systemA.Uid
	_, err = Client.CreateNetworkInterface(systemA.Uid, ifaceA)
	cobbler.FailOnError(t, err)
	// Recursive system deletion above already removes this interface; no separate
	// cleanup needed here (and DeleteNetworkInterface after the system is gone
	// would just fail on an already-deleted uid).
	ifaceB := cobbler.NewNetworkInterface()
	ifaceB.Name = ifaceName
	ifaceB.SystemUid = systemB.Uid
	_, err = Client.CreateNetworkInterface(systemB.Uid, ifaceB)
	cobbler.FailOnError(t, err)

	_, err = resolveUID(&Client, "network_interface", ifaceName, "")
	if err == nil {
		t.Fatal("expected an error for an ambiguous name")
	}
	if !strings.Contains(err.Error(), "multiple") || !strings.Contains(err.Error(), "--uid") {
		t.Fatalf("expected error to mention the name is ambiguous and to use --uid, got: %v", err)
	}
}

// TestResolveUID_FindItemsError exercises the branch where the find_items
// RPC call itself fails (e.g. an invalid/expired token server-side): the
// error must be propagated as-is. This is done against a canned HTTP
// transport rather than a live server since provoking a genuine server-side
// find_items failure on demand isn't practical.
func TestResolveUID_FindItemsError(t *testing.T) {
	client := newCannedClient(findItemsFaultResponse)
	_, err := resolveUID(&client, "distro", "some-name", "")
	if err == nil {
		t.Fatal("expected an error when find_items fails")
	}
	if !strings.Contains(err.Error(), "simulated find_items failure") {
		t.Fatalf("expected the find_items error to be propagated, got: %v", err)
	}
}

// TestResolveUID_UnexpectedResultType exercises the defensive branch that
// guards against a single find_items match that isn't a struct (map). This
// shouldn't happen against a real Cobbler server, but resolveUID must still
// fail clearly rather than panicking on the failed type assertion.
func TestResolveUID_UnexpectedResultType(t *testing.T) {
	client := newCannedClient(findItemsNonStructResponse)
	_, err := resolveUID(&client, "distro", "some-name", "")
	if err == nil {
		t.Fatal("expected an error for a non-struct find_items result")
	}
	if !strings.Contains(err.Error(), "unexpected result type") {
		t.Fatalf("expected an 'unexpected result type' error, got: %v", err)
	}
}

// TestResolveUID_NoUIDInResult exercises the defensive branch that guards
// against a single find_items match whose struct has no "uid" member. As
// with TestResolveUID_UnexpectedResultType this shouldn't happen against a
// real Cobbler server, but resolveUID must still fail clearly.
func TestResolveUID_NoUIDInResult(t *testing.T) {
	client := newCannedClient(findItemsNoUIDResponse)
	_, err := resolveUID(&client, "distro", "some-name", "")
	if err == nil {
		t.Fatal("expected an error for a find_items result without a uid")
	}
	if !strings.Contains(err.Error(), "no uid found in result") {
		t.Fatalf("expected a 'no uid found in result' error, got: %v", err)
	}
}

// newRemoveItemFlags builds a *cobra.Command carrying exactly the flags
// listed, letting each of RemoveItemRecursive's `cmd.Flags().Get*` error
// branches be exercised individually and without any server interaction:
// pflag itself errors when a lookup misses, and that happens before
// RemoveItemRecursive ever touches the (unset, zero-value) global Client.
func newRemoveItemFlags(withName, withUID, withRecursive bool) *cobra.Command {
	cmd := &cobra.Command{}
	if withName {
		cmd.Flags().String("name", "", "")
	}
	if withUID {
		cmd.Flags().String("uid", "", "")
	}
	if withRecursive {
		cmd.Flags().Bool("recursive", false, "")
	}
	return cmd
}

// TestRemoveItemRecursive_MissingNameFlag exercises the error branch taken
// when the calling command forgot to register --name: cmd.Flags().GetString
// fails before any server interaction happens.
func TestRemoveItemRecursive_MissingNameFlag(t *testing.T) {
	err := RemoveItemRecursive(newRemoveItemFlags(false, false, false), nil, "distro")
	if err == nil {
		t.Fatal("expected an error when the --name flag isn't registered")
	}
}

// TestRemoveItemRecursive_MissingUIDFlag exercises the error branch taken
// when the calling command forgot to register --uid.
func TestRemoveItemRecursive_MissingUIDFlag(t *testing.T) {
	err := RemoveItemRecursive(newRemoveItemFlags(true, false, false), nil, "distro")
	if err == nil {
		t.Fatal("expected an error when the --uid flag isn't registered")
	}
}

// TestRemoveItemRecursive_MissingRecursiveFlag exercises the error branch
// taken when the calling command forgot to register --recursive.
func TestRemoveItemRecursive_MissingRecursiveFlag(t *testing.T) {
	err := RemoveItemRecursive(newRemoveItemFlags(true, true, false), nil, "distro")
	if err == nil {
		t.Fatal("expected an error when the --recursive flag isn't registered")
	}
}

// TestRemoveItemRecursive_NoIdentifier exercises the branch where
// resolveUID's own error (neither --name nor --uid supplied) is propagated
// by RemoveItemRecursive. All three flags are registered but left at their
// zero value, so this never reaches the (unset) global Client either.
func TestRemoveItemRecursive_NoIdentifier(t *testing.T) {
	err := RemoveItemRecursive(newRemoveItemFlags(true, true, true), nil, "distro")
	if err == nil {
		t.Fatal("expected an error when neither --name nor --uid is set")
	}
	if !strings.Contains(err.Error(), "--name") || !strings.Contains(err.Error(), "--uid") {
		t.Fatalf("expected error to mention --name/--uid, got: %v", err)
	}
}
