package pislideservice

type PiSlideStatusOutput struct {
	Body struct {
		CommandOutput string `json:"command_output" doc:"The command from the corresponding systemctl call"`
		StatusCode    int    `json:"status_code" doc:"The status code of the response"`
	}
}

type PiSlideServiceInput struct {
	Action string `query:"action" enum:"start,stop,restart" required:"true"`
}
