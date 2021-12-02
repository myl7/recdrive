package recdrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	uploadQuery = map[string]string{
		"storage":   "moss",
		"disk_type": "cloud",
	}
)

func (drive *Drive) StartUploadToID(id string, filename string, filesize int64) (StartUploadInfo, error) {
	path := fmt.Sprintf("/file/%s", id)
	s := drive.opt.ApiEndpoint + path
	s, err := appendQuery(s, uploadQuery)
	if err != nil {
		return StartUploadInfo{}, err
	}

	s, err = appendQuery(s, map[string]string{
		"file_name": filename,
		"byte":      strconv.FormatInt(filesize, 10),
	})
	if err != nil {
		return StartUploadInfo{}, err
	}

	req, err := http.NewRequest("GET", s, nil)
	req.Header.Set(authTokenField, drive.opt.AuthToken)
	res, err := drive.opt.HttpClient.Do(req)
	if err != nil {
		return StartUploadInfo{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return StartUploadInfo{}, err
	}

	body = removeBom(body)

	var status ResStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return StartUploadInfo{}, err
	}

	e := NewErrFromStatus(status, path)
	if e != nil {
		return StartUploadInfo{}, *e
	}

	var startUploadRes StartUploadRes
	err = json.Unmarshal(body, &startUploadRes)
	if err != nil {
		return StartUploadInfo{}, err
	}

	return startUploadRes.Entity, nil
}

type StartUploadInfo struct {
	UploadParams [][]struct {
		Key         string `json:"key"`
		RequestType string `json:"request_type"`
		Value       string `json:"value"`
	} `json:"upload_params"`
	UploadChunkSize string `json:"upload_chunk_size"`
	UploadToken     string `json:"upload_token"`
}

type StartUploadRes struct {
	Entity StartUploadInfo `json:"entity"`
}

func (drive *Drive) DoUpload(file io.Reader, info StartUploadInfo) (string, error) {
	chunkSize, err := strconv.Atoi(info.UploadChunkSize)
	if err != nil {
		return "", err
	}

	buf := make([]byte, chunkSize)
	for i := range info.UploadParams {
		params := info.UploadParams[i]
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}

		req, err := http.NewRequest(params[2].Value, params[1].Value, bytes.NewReader(buf[:n]))
		req.Header.Set(authTokenField, drive.opt.AuthToken)
		req.Header.Set("Content-Length", strconv.Itoa(n))
		res, err := drive.opt.HttpClient.Do(req)
		if err != nil {
			return "", err
		}

		if res.StatusCode != 200 {
			return "", ErrReqFailed{
				StatusCode: res.StatusCode,
				Message:    "",
				Path:       "",
			}
		}
	}

	path := "/file/complete"
	s := drive.opt.ApiEndpoint + path
	body, err := json.Marshal(map[string]string{"upload_token": info.UploadToken})
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

	var doUploadRes DoUploadRes
	err = json.Unmarshal(body, &doUploadRes)
	if err != nil {
		return "", err
	}

	return doUploadRes.Entity.ID, nil
}

type DoUploadRes struct {
	Entity struct {
		ID string `json:"number"`
	} `json:"entity"`
}

func (drive *Drive) Upload(path string, file io.Reader, filename string, filesize int64) (string, error) {
	id, err := drive.QueryID(path)
	if err != nil {
		return "", err
	}

	info, err := drive.StartUploadToID(id, filename, filesize)
	if err != nil {
		return "", err
	}

	id, err = drive.DoUpload(file, info)
	if err != nil {
		return "", err
	}

	return id, nil
}
