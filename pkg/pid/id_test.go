package pid

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	id := New()
	if id == 0 {
		t.Fatal("expected non-zero ID")
	}
	if id.Time().IsZero() {
		t.Error("expected non-zero time")
	}
	if id.Random() > randomMask {
		t.Errorf("random bits exceed mask: %d > %d", id.Random(), randomMask)
	}
}

func TestEncodeDecodeBase32(t *testing.T) {
	tests := []ID{
		0,
		1,
		42,
		123456,
		987654321,
		1<<63 - 1,
	}
	for _, v := range tests {
		encoded := EncodeBase32(v)
		decoded, err := DecodeBase32(encoded)
		if err != nil {
			t.Errorf("DecodeBase32(%q) failed: %v", encoded, err)
			continue
		}
		if v != decoded {
			t.Errorf("roundtrip failed: got %d, want %d", decoded, v)
		}
	}
}

func TestDecodeAmbiguousCharacters(t *testing.T) {
	orig := EncodeBase32(1234567890)
	ambiguous := strings.NewReplacer("O", "0", "I", "1", "L", "1").Replace(orig)

	decoded, err := DecodeBase32(ambiguous)
	if err != nil {
		t.Fatalf("DecodeBase32 with ambiguous chars failed: %v", err)
	}
	expected, _ := DecodeBase32(orig)
	if decoded != expected {
		t.Errorf("ambiguous decode mismatch: got %d, want %d", decoded, expected)
	}
}

func TestDecodeBase32Invalid(t *testing.T) {
	_, err := DecodeBase32("!@#$")
	if err == nil {
		t.Fatal("expected error for invalid characters")
	}
}

func TestID_String_Decode(t *testing.T) {
	id := New()
	encoded := id.String()
	decoded, err := DecodeBase32(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if id != decoded {
		t.Errorf("mismatch: original %d != decoded %d", id, decoded)
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	id := New()
	data, err := json.Marshal(id)
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	var decoded ID
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if id != decoded {
		t.Errorf("JSON roundtrip mismatch: got %d, want %d", decoded, id)
	}
}

func TestTimeAccuracy(t *testing.T) {
	now := time.Now()
	id := New()
	idTime := id.Time()
	if idTime.Before(now.Add(-time.Second)) || idTime.After(now.Add(time.Second)) {
		t.Errorf("timestamp mismatch: got %v, want ~%v", idTime, now)
	}
}

func TestRandomMask(t *testing.T) {
	id := New()
	if r := id.Random(); r > randomMask {
		t.Errorf("Random value exceeds mask: %d > %d", r, randomMask)
	}
}

func TestBase32SortableByTimestamp(t *testing.T) {
	// Make two IDs that differ only in timestamp:
	ts1 := int64(1_000_000) // ms after epoch
	ts2 := ts1 + 1
	rand := uint32(0)

	// Build raw values:
	v1 := (ts1 << randomBits) | int64(rand)
	v2 := (ts2 << randomBits) | int64(rand)

	s1 := EncodeBase32(ID(v1))
	s2 := EncodeBase32(ID(v2))

	if len(s1) != 13 || len(s2) != 13 {
		t.Fatalf("expected both to be length 13, got %d and %d", len(s1), len(s2))
	}
	if strings.Compare(s1, s2) >= 0 {
		t.Errorf("expected %q < %q lexicographically (older first)", s1, s2)
	}
}

func TestBase32SortableByRandom(t *testing.T) {
	// Same timestamp but different random
	ts := time.Now().UnixMilli() & ((1 << timestampBits) - 1)
	vSmall := (ts << randomBits) | int64(1)
	vBig := (ts << randomBits) | int64(2)

	sSmall := EncodeBase32(ID(vSmall))
	sBig := EncodeBase32(ID(vBig))

	if strings.Compare(sSmall, sBig) >= 0 {
		t.Errorf("expected %q < %q when same timestamp but random is smaller", sSmall, sBig)
	}
}

func FuzzEncodeDecodeBase32Roundtrip(f *testing.F) {
	f.Add(uint64(1))
	f.Add(uint64(123456789))
	f.Fuzz(func(t *testing.T, val ID) {
		s := EncodeBase32(val)
		got, err := DecodeBase32(s)
		if err != nil {
			t.Errorf("DecodeBase32 failed: %v", err)
		}
		if got != val {
			t.Errorf("roundtrip failed: expected %d, got %d", val, got)
		}
	})
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

func BenchmarkEncodeBase32(b *testing.B) {
	val := New()
	for i := 0; i < b.N; i++ {
		_ = EncodeBase32(val)
	}
}

func BenchmarkDecodeBase32(b *testing.B) {
	str := EncodeBase32(New())
	for i := 0; i < b.N; i++ {
		_, _ = DecodeBase32(str)
	}
}
