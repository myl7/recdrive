package recdrive

import (
	"bytes"
	"net/url"
	"path/filepath"
	"strings"
)

func appendQuery(s string, query map[string]string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func cleanPath(path string) string {
	path = filepath.ToSlash(path)
	path = filepath.Join("/", path)
	path = filepath.Clean(path)
	if path == "/" {
		return path
	} else {
		return strings.TrimSuffix(path, "/")
	}
}

// path should be cleaned first
func splitPath(path string) []string {
	ps := strings.Split(path, "/")
	return ps[1:]
}

func removeBom(b []byte) []byte {
	return bytes.TrimPrefix(b, []byte{0xef, 0xbb, 0xbf})
}

func Filename(item ListItem) string {
	if item.FileExt != "" {
		return item.Name + "." + item.FileExt
	} else {
		return item.Name
	}
}
