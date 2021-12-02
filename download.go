package recdrive

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (drive *Drive) DownloadByID(id string) (string, error) {
	path := "/download"
	s := drive.opt.ApiEndpoint + path
	body, err := json.Marshal(map[string]interface{}{
		"files_list":   []string{id},
		"group_number": nil,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", s, bytes.NewReader(body))
	req.Header.Set(authTokenField, drive.opt.AuthToken)
	res, err := drive.opt.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	body = removeBom(body)

	var status ResStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return "", err
	}

	e := NewErrFromStatus(status, path)
	if e != nil {
		return "", *e
	}

	var downloadRes DownloadRes
	err = json.Unmarshal(body, &downloadRes)
	if err != nil {
		return "", err
	}

	return downloadRes.Entity[id], nil
}

func (drive *Drive) Download(path string) (string, error) {
	id, err := drive.QueryID(path)
	if err != nil {
		return "", err
	}

	s, err := drive.DownloadByID(id)
	if err != nil {
		return "", err
	}

	return s, nil
}

type DownloadRes struct {
	Entity map[string]string `json:"entity"`
}
