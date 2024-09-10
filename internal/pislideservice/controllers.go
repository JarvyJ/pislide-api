package pislideservice

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPiSlideshowService(api *huma.API) {

	huma.Register(*api, huma.Operation{
		OperationID: "get-pislide-service-status",
		Method:      http.MethodGet,
		Path:        "/api/v1/pislide-service",
		Summary:     "Get PiSlide service status",
		Description: "Show the current PiSlide service status as per systemctl",
		Tags:        []string{"PiSlide-Service"},
	}, func(ctx context.Context, input *struct{}) (*PiSlideStatusOutput, error) {
		resp := &PiSlideStatusOutput{}
		output, statusCode, err := GetPiSlideServiceStatus()
		if err != nil {
			return nil, huma.Error500InternalServerError("Error while executing systemctl")
		}
		resp.Body.CommandOutput = output
		resp.Body.StatusCode = statusCode
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "change-pislide-service",
		Method:      http.MethodPut,
		Path:        "/api/v1/pislide-service",
		Summary:     "Start/Stop/Restart PiSlide service",
		Description: "Start/Stop/Restart the PiSlide service per systemctl",
		Tags:        []string{"PiSlide-Service"},
	}, func(ctx context.Context, input *PiSlideServiceInput) (*PiSlideStatusOutput, error) {
		resp := &PiSlideStatusOutput{}
		output, statusCode, err := "", 0, errors.New("") // TODO: how to intialize an empty error?

		switch input.Action {
		case "start":
			output, statusCode, err = StartPiSlideService()
		case "stop":
			output, statusCode, err = StopPiSlideService()
		case "restart":
			output, statusCode, err = RestartPiSlideService()
		}

		if err != nil {
			return nil, huma.Error500InternalServerError("Error while executing systemctl")
		}
		resp.Body.CommandOutput = output
		resp.Body.StatusCode = statusCode
		return resp, nil
	})
}
