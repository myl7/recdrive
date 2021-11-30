package recdrive

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

var (
	queryDefault = map[string]string{
		"disk_type": "cloud",
		"is_rec":    "false",
		"category":  "all",
	}
)

type ResStatus struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
