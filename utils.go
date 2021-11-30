package recdrive

import "net/url"

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
