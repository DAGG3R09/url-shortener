package shorteners

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encodeToBase62(t *testing.T) {
	var k int64 = 123

	t.Run("encode key", func(t *testing.T) {
		url := encodeToBase62(k)
		assert.Equal(t, "1Z", url)
	})
}

func Test_decodeToBase62(t *testing.T) {
	urls := []string{"100000", "ZZZZZZ"}
	expected := []int64{916132832, 56800235583}

	for i := range urls {
		t.Run("decode key", func(t *testing.T) {
			key, err := decodeBase62String(urls[i])

			assert.NoError(t, err)
			assert.Equal(t, key, expected[i])

		})
	}

}

func Test_EncodeDecode(t *testing.T) {
	b62 := NewBase62Shortener()
	key := int64(916132832)
	url := "100000"

	t.Run("shorten", func(t *testing.T) {
		url := b62.Shorten(0)
		assert.Equal(t, "0", url)
	})

	t.Run("shorten", func(t *testing.T) {
		url := b62.Shorten(key)
		assert.Equal(t, "100000", url)
	})

	t.Run("un-shorten", func(t *testing.T) {
		decodedKey, err := b62.Unshorten(url)
		assert.NoError(t, err)
		assert.Equal(t, key, decodedKey)
	})

	t.Run("un-shorten: invalid characters", func(t *testing.T) {
		decodedKey, err := b62.Unshorten("5a1s8d7asd/1231231231///")
		assert.Error(t, err)
		assert.Empty(t, decodedKey)
	})

}
