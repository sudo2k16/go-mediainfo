package mediainfo

import "testing"

func TestOverallBitRateModeForKindPrefersVariableAcrossAudioStreams(t *testing.T) {
	streams := []Stream{
		{
			Kind: StreamAudio,
			Fields: []Field{
				{Name: "Bit rate mode", Value: "Constant"},
			},
		},
		{
			Kind: StreamAudio,
			JSON: map[string]string{
				"BitRate_Mode": "VBR",
			},
		},
	}

	if got := overallBitRateModeForKind(streams, StreamAudio); got != "Variable" {
		t.Fatalf("overallBitRateModeForKind() = %q, want %q", got, "Variable")
	}
}
