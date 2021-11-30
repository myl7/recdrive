package recdrive

import (
	"net/url"
	"path/filepath"
	"strings"
)

func appendDefaultQuery(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range queryDefault {
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

func splitPath(path string) []string {
	path = cleanPath(path)
	ps := strings.Split(path, "/")
	return ps[1:]
}
