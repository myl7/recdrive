package recdrive

import (
	"path/filepath"
	"strings"
)

type Drive struct {
	opt Options
}

func NewDrive(options Options) *Drive {
	return &Drive{
		opt: options.Build(),
	}
}

const (
	authTokenField     = "x-auth-token"
	apiEndpointDefault = "https://recapi.ustc.edu.cn/api/v2"
)

func (drive *Drive) QueryID(path string) (string, error) {
	if path == "/" {
		return "0", nil
	}

	base := filepath.Base(path)

	dirId := "0"
	dir := filepath.Dir(path)
	if dir != "/" {
		ps := splitPath(path)
		for i := range ps {
			items, err := drive.ListByID(dirId)
			if err != nil {
				return "", err
			}

			found := false
			for j := range items {
				if items[j].Name == ps[i] {
					if items[j].Type != "folder" {
						return "", ErrNotFolder{Path: strings.Join(ps[:i+1], "/")}
					}

					dirId = items[j].ID
					found = true
					break
				}
			}
			if !found {
				return "", ErrNotFound{Path: strings.Join(ps[:i+1], "/")}
			}
		}
	}

	items, err := drive.ListByID(dirId)
	if err != nil {
		return "", err
	}

	for i := range items {
		if items[i].Name == base {
			return items[i].ID, nil
		}
	}
	return "", ErrNotFound{Path: path}
}

type ResStatus struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
