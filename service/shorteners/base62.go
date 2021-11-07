package shorteners

import (
	"strings"

	"github.com/ansel1/merry/v2"
)

var _ ShortenerIface = (*Base62)(nil)

type Base62 struct {
}

func NewBase62Shortener() *Base62 {
	s := Base62{}

	return &s
}

func (b *Base62) Shorten(key int64) string {
	return encodeToBase62(key)
}

func (b *Base62) Unshorten(shortURL string) (int64, error) {
	return decodeBase62String(shortURL)
}

func encodeToBase62(i int64) string {
	if i == 0 {
		return string(Base62Charset[0])
	}

	base62Encode := []byte{}
	rem := int64(0)

	for i > 0 {
		rem, i = i%Base62CharsetSize, i/Base62CharsetSize

		base62Encode = append(base62Encode, Base62Charset[rem])
	}

	return reverse(base62Encode)
}

func decodeBase62String(shortURL string) (int64, error) {
	var (
		power int64 = 1
		sum   int64
		index int
	)

	for i := len(shortURL) - 1; i >= 0; i-- {
		index = strings.IndexByte(Base62Charset, shortURL[i])

		if index == -1 {
			return 0, merry.Errorf("Invalid character in URL %c", shortURL[i])
		}

		sum += power * int64(index)
		power *= Base62CharsetSize
	}

	return sum, nil
}

func reverse(a []byte) string {
	s, e := 0, len(a)-1

	for s < e {
		a[s], a[e] = a[e], a[s]

		s++
		e--
	}

	return string(a)
}

var (
	Base62CharsetSize = int64(len(Base62Charset))
)

const (
	Base62Charset string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
