package khatru

import (
	"fmt"
	"github.com/mailru/easyjson"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/nbd-wtf/go-nostr"
	"github.com/tidwall/gjson"
	"unsafe"
)

// Copied from a PR to nostr-tools, TODO: remove this once it's merged
// ProbeEnvelope represents a PROBE message.
type ProbeEnvelope struct {
	nostr.Event
}

func (_ ProbeEnvelope) Label() string { return "PROBE" }

func (v *ProbeEnvelope) FromJSON(data string) error {
	r := gjson.Parse(data)
	arr := r.Array()
	switch len(arr) {
	case 2:
		return easyjson.Unmarshal(unsafe.Slice(unsafe.StringData(arr[1].Raw), len(arr[1].Raw)), &v.Event)
	default:
		return fmt.Errorf("failed to decode PROBE envelope")
	}
}

func (v ProbeEnvelope) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{NoEscapeHTML: true}
	w.RawString(`["PROBE",`)
	v.Event.MarshalEasyJSON(&w)
	w.RawString(`]`)
	return w.BuildBytes()
}
