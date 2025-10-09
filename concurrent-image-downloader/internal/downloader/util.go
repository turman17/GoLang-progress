package downloader

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ensureDir(p string) error {
	return os.MkdirAll(p, 0o755)
}

func fileNameFromURL(raw string) string {
	u, err := url.Parse(raw)
	if err == nil {
		base := path.Base(u.Path)
		if base != "" && base != "/" && base != "." {
			return base
		}
	}
	sum := sha1.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])[:16] + ".bin"
}

func writeFile(dst string, r io.Reader) (int64, error) {
	if err := ensureDir(filepath.Dir(dst)); err != nil {
		return 0, err
	}
	f, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return io.Copy(f, r)
}

func normalizeURL(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("empty url")
	}
	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", errors.New("invalid url")
	}
	return u.String(), nil
}