package logger

type AttrKeys struct {
	App          AppAttrKeys
	HTTP         HTTPAttrKeys
	IP           string
	PID          string
	ResponseTime string
	Tags         string
}

type AppAttrKeys struct {
	Name    string
	Version string
}

type HTTPAttrKeys struct {
	Body     string
	BodySize string
	Headers  string
	Method   string
	Path     string
	Status   string
}

var AttrKey = AttrKeys{
	App: AppAttrKeys{
		Name:    "name",
		Version: "version",
	},
	HTTP: HTTPAttrKeys{
		Body:     "body",
		BodySize: "body_size",
		Headers:  "headers",
		Method:   "method",
		Path:     "path",
		Status:   "status",
	},
	IP:           "ip",
	PID:          "pid",
	ResponseTime: "response_time",
	Tags:         "tags",
}
