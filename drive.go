package recdrive

type Drive struct {
}

func NewDrive(options Options) *Drive {
	return nil
}

const (
	authTokenField     = "x-auth-token"
	apiEndpointDefault = "https://recapi.ustc.edu.cn/api/v2"
)

var (
	queryDefault = map[string]string{
		"disk_type": "cloud",
		"is_rec":    "false",
		"category":  "all",
	}
)
