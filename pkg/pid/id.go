package pid

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ┌───────────────────────────────────────────────────────────────────────────┐
// │ 63 │ 62 … 20                              │ 19 … 0                        │
// │ Unused (1 bit)   Timestamp (43 bits)      │ Randomness (20 bits)          │
// └───────────────────────────────────────────────────────────────────────────┘

// ID is a 64-bit unsigned identifier: 43 bits timestamp (ms since epoch), 20 bits randomness, 1 bit unused (sign bit zero).
type ID int64

const (
	timestampBits = 43
	randomBits    = 20
	randomBytes   = 3 // 3 bytes = 24 bits, mask to 20 bits
	randomMask    = (1 << randomBits) - 1
	maxAttempts   = 10
)

var (
	mu         sync.Mutex
	lastTs     int64
	usedRandom = make(map[uint32]struct{}, 1<<randomBits)
	prng       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateRandom() uint32 {
	return uint32(prng.Int31()) & randomMask
}

func New() ID {
	ts := int64(time.Now().UnixMilli()) & ((1 << timestampBits) - 1)

	var r uint32
	for range maxAttempts {
		r = generateRandom()

		mu.Lock()
		if ts != lastTs {
			lastTs = ts
			usedRandom = make(map[uint32]struct{})
		}

		if _, exists := usedRandom[r]; !exists {
			usedRandom[r] = struct{}{}
			mu.Unlock()
			val := (ts << randomBits) | int64(r)
			return ID(val)
		}
		mu.Unlock()
	}

	val := (ts << randomBits) | int64(r)
	return ID(val)
}

func (id ID) Time() time.Time {
	ts := int64(id) >> randomBits
	return time.UnixMilli(ts)
}

func (id ID) Random() uint32 {
	return uint32(int64(id) & randomMask)
}

func (id ID) String() string {
	return EncodeBase32(id)
}

func (id ID) Int() int64 {
	return int64(id)
}

// Crockford's Base32 encoding alphabet
var crockfordBase32 = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")

var decodeCrockford = func() map[rune]byte {
	m := make(map[rune]byte)
	for i, c := range crockfordBase32 {
		m[rune(c)] = byte(i)
		m[rune(strings.ToLower(string(c))[0])] = byte(i)
	}
	// support ambiguous characters
	m['i'], m['l'] = 1, 1
	m['o'] = 0
	return m
}()

func EncodeBase32(val ID) string {
	buf := [13]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	i := len(buf)
	for val > 0 {
		i--
		buf[i] = crockfordBase32[val%32]
		val /= 32
	}
	return string(buf[:])
}

func DecodeBase32(s string) (ID, error) {
	var val uint64
	for _, c := range s {
		d, ok := decodeCrockford[c]
		if !ok {
			return 0, fmt.Errorf("invalid base32 character: %c", c)
		}
		val = val*32 + uint64(d)
	}
	return ID(val), nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	u, err := DecodeBase32(s)
	if err != nil {
		return err
	}
	*id = ID(int64(u))
	return nil
}
