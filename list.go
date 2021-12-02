package recdrive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	listQuery = map[string]string{
		"disk_type": "cloud",
		"is_rec":    "false",
		"category":  "all",
	}
)

func (drive *Drive) ListByID(id string) ([]ListItem, error) {
	path := fmt.Sprintf("/folder/content/%s", id)
	s := drive.opt.ApiEndpoint + path
	s, err := appendQuery(s, listQuery)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", s, nil)
	req.Header.Set(authTokenField, drive.opt.AuthToken)
	res, err := drive.opt.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body = removeBom(body)

	var status ResStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}

	e := NewErrFromStatus(status, path)
	if e != nil {
		return nil, *e
	}

	var listRes ListRes
	err = json.Unmarshal(body, &listRes)
	if err != nil {
		return nil, err
	}

	return listRes.Entity.Datas, nil
}

func (drive *Drive) List(path string) ([]ListItem, error) {
	id, err := drive.QueryID(path)
	if err != nil {
		return nil, err
	}

	items, err := drive.ListByID(id)
	if err != nil {
		return nil, err
	}

	return items, nil
}

type ListItem struct {
	CreaterUserNumber   string      `json:"creater_user_number"`
	CreaterUserRealName string      `json:"creater_user_real_name"`
	CreaterUserAvatar   string      `json:"creater_user_avatar"`
	ID                  string      `json:"number"`
	ParentID            string      `json:"parent_number"`
	DiskType            string      `json:"disk_type"`
	IsHistory           bool        `json:"is_history"`
	Name                string      `json:"name"`
	Type                string      `json:"type"`
	FileExt             string      `json:"file_ext"`
	FileType            string      `json:"file_type"`
	Bytes               interface{} `json:"bytes"` // string "" for folder or int for file
	Hash                string      `json:"hash"`
	TranscodeStatus     string      `json:"transcode_status"`
	IsStar              bool        `json:"is_star"`
	IsLock              bool        `json:"is_lock"`
	LockReason          string      `json:"lock_reason"`
	ShareCount          int         `json:"share_count"`
	LastUpdateDate      string      `json:"last_update_date"`
	ParentPathNumber    string      `json:"parent_path_number"`
	ReviewStatus        string      `json:"review_status"`
}

type ListRes struct {
	Entity struct {
		Total int        `json:"total"`
		Datas []ListItem `json:"datas"`
	} `json:"entity"`
}
