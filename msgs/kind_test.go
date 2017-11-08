package msgs_test

import (
	"testing"

	"github.com/telyn/midi/msgs"
)

func TestKindOf(t *testing.T) {
	tests := []struct {
		BottomByte byte
		TopByte    byte
		Kind       msgs.Kind
	}{
		{0x00, 0x7F, msgs.DataByte},
		{0x80, 0x8F, msgs.NoteOff},
		{0x90, 0x9F, msgs.NoteOn},
		{0xA0, 0xAF, msgs.KeyPressure},
		{0xB0, 0xBF, msgs.ControlChange},
		{0xC0, 0xCF, msgs.ProgramChange},
		{0xD0, 0xDF, msgs.ChannelPressure},
		{0xE0, 0xEF, msgs.PitchBend},
		{0xF0, 0xF0, msgs.SystemExclusive},
		{0xF1, 0xF1, msgs.SystemCommonTimeCode},
		{0xF2, 0xF2, msgs.SystemCommonSongPositionPointer},
		{0xF3, 0xF3, msgs.SystemCommonSongSelect},
		{0xF4, 0xF5, msgs.SystemCommonUndefined},
		{0xF6, 0xF6, msgs.SystemCommonTuneRequest},
		{0xF7, 0xF7, msgs.SystemCommonEOX},
		{0xF8, 0xFF, msgs.RealTime},
	}
	for _, test := range tests {
		if test.TopByte < test.BottomByte {
			t.Fatalf("Test with range 0x%x -> 0x%x is bad", test.BottomByte, test.TopByte)
		}
		for b := test.BottomByte; b < test.TopByte; b++ {
			if test.Kind != msgs.KindOf(b) {
				t.Errorf("%x: expecting %v, got %v", b, test.Kind, msgs.KindOf(b))
			}
			// this is needed cause 0xFF <= 0xFF == true so it'll just loop forever
			if b == 0xFF {
				break
			}
		}
	}

}

func TestChannelOf(t *testing.T) {
	tests := []struct {
		Bottom byte
		Top    byte
		FF     bool
	}{
		{0x00, 0x7F, true},
		{0x80, 0xEF, false},
		{0xF0, 0xFF, true},
	}
	for _, test := range tests {
		for b := test.Bottom; b <= test.Top; b++ {
			expect := b & 0xF
			if test.FF {
				expect = 0xFF
			}
			actual := msgs.ChannelOf(b)
			if expect != actual {
				t.Errorf("%x: Channel should be %x, was %x", b, expect, actual)
			}
			if b == 0xFF {
				break
			}
		}
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		Bottom byte
		Top    byte
		Bytes  int
	}{
		{0x00, 0x7F, -2},
		{0x80, 0xBF, 2},
		{0xC0, 0xDF, 1},
		{0xE0, 0xEF, 2},
		{0xF0, 0xF0, -1},
		{0xF1, 0xF1, 1},
		{0xF2, 0xF2, 2},
		{0xF3, 0xF3, 1},
		{0xF4, 0xF5, 0},
		{0xF6, 0xF6, 0},
		{0xF7, 0xF7, 0},
		{0xF8, 0xFF, 0},
	}
	for _, test := range tests {
		for b := test.Bottom; b <= test.Top; b++ {
			k := msgs.KindOf(b)
			if test.Bytes != k.Bytes() {
				t.Errorf("%x (%v) expected %d, got %d", b, k, test.Bytes, k.Bytes())
			}
			if b == 0xFF {
				break
			}
		}
	}
}
