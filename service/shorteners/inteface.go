package shorteners

// type ShortnerIface interface {
// 	Shorten(string) string
// }

type ShortenerIface interface {
	Shorten(key int64) string
	Unshorten(shortURL string) (int64, error)
}
