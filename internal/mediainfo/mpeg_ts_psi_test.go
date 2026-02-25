package mediainfo

import "testing"

func TestParsePATRejectsTooShortSectionLength(t *testing.T) {
	payload := make([]byte, 9)

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("parsePAT panicked: %v", recovered)
		}
	}()

	programs, consumed := parsePAT(payload)
	if programs != nil {
		t.Fatalf("expected nil programs, got %v", programs)
	}
	if consumed != 0 {
		t.Fatalf("expected consumed=0, got %d", consumed)
	}
}

func TestParseSDTRejectsTooShortSectionLength(t *testing.T) {
	payload := make([]byte, 12)
	payload[1] = 0x42

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("parseSDT panicked: %v", recovered)
		}
	}()

	name, provider, serviceType := parseSDT(payload, 0)
	if name != "" || provider != "" || serviceType != "" {
		t.Fatalf("expected empty SDT result, got name=%q provider=%q serviceType=%q", name, provider, serviceType)
	}
}
