package recdrive

import "fmt"

type ErrReqFailed struct {
	StatusCode int
	Message    string
	Path       string
}

func (e ErrReqFailed) Error() string {
	return fmt.Sprintf("req path %s status %d reason %s", e.Path, e.StatusCode, e.Message)
}

func NewErrFromStatus(status ResStatus, path string) *ErrReqFailed {
	if status.StatusCode >= 200 && status.StatusCode < 300 {
		return &ErrReqFailed{
			StatusCode: status.StatusCode,
			Message:    status.Message,
			Path:       path,
		}
	} else {
		return nil
	}
}

type ErrNotFolder struct {
	Path string
}

func (e ErrNotFolder) Error() string {
	return fmt.Sprintf("%s is not folder", e.Path)
}
