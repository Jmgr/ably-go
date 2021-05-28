package proto_test

import (
	"bytes"
	"testing"

	"github.com/ably/ably-go/ably/internal/ablyutil"
	"github.com/ably/ably-go/ably/proto"
)

// TestProtocolMessageEncodeZeroSerials tests that zero-valued serials are
// explicitly encoded into msgpack (as required by the realtime API)
func TestProtocolMessageEncodeZeroSerials(t *testing.T) {
	msg := proto.ProtocolMessage{
		ID:               "test",
		MsgSerial:        0,
		ConnectionSerial: 0,
	}
	encoded, err := ablyutil.MarshalMsgpack(msg)
	if err != nil {
		t.Fatal(err)
	}
	// expect a 3-element map with both the serial fields set to zero
	expected := []byte("\x83\xB0connectionSerial\x00\xA2id\xA4test\xA9msgSerial\x00")
	if !bytes.Equal(encoded, expected) {
		t.Fatalf("unexpected msgpack encoding\nexpected: %x\nactual:   %x", expected, encoded)
	}
}

func TestIfFlagIsSet(t *testing.T) {
	flags := proto.FlagAttachResume
	flags.Set(proto.FlagPresence)
	flags.Set(proto.FlagPublish)
	flags.Set(proto.FlagSubscribe)
	flags.Set(proto.FlagPresenceSubscribe)

	if expected, actual := proto.FlagPresence, flags&proto.FlagPresence; expected != actual {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
	if expected, actual := proto.FlagPublish, flags&proto.FlagPublish; expected != actual {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
	if expected, actual := proto.FlagSubscribe, flags&proto.FlagSubscribe; expected != actual {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
	if expected, actual := proto.FlagPresenceSubscribe, flags&proto.FlagPresenceSubscribe; expected != actual {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
	if expected, actual := proto.FlagAttachResume, flags&proto.FlagAttachResume; expected != actual {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
	if expected, actual := proto.FlagHasBacklog, flags&proto.FlagAttachResume; expected == actual {
		t.Fatalf("Shouldn't contain flag %v", expected)
	}
}

func TestIfHasFlg(t *testing.T) {
	flags := proto.FlagAttachResume | proto.FlagPresence | proto.FlagPublish
	if !flags.Has(proto.FlagAttachResume) {
		t.Fatalf("Should contain flag %v", proto.FlagAttachResume)
	}
	if !flags.Has(proto.FlagPresence) {
		t.Fatalf("Should contain flag %v", proto.FlagPresence)
	}
	if !flags.Has(proto.FlagPublish) {
		t.Fatalf("Should contain flag %v", proto.FlagPublish)
	}
	if flags.Has(proto.FlagHasBacklog) {
		t.Fatalf("Shouldn't contain flag %v", proto.FlagHasBacklog)
	}
}
