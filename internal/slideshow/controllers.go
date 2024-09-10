package slideshow

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterSlideshow(api *huma.API) {
	// Register GET /greeting/{name}
	huma.Register(*api, huma.Operation{
		OperationID: "create-slideshow",
		Method:      http.MethodPost,
		Path:        "/api/v1/slideshows/{foldername}",
		Summary:     "Create a slideshow",
		Description: "Create a slideshow by setting all the required options.",
		Tags:        []string{"Slideshows"},
	}, func(ctx context.Context, input *SpecificSlideshowInput) (*SpecificSlideshowOutput, error) {
		resp := &SpecificSlideshowOutput{}
		resp.Body.Duration = input.Body.Duration
		resp.Body.Sort = input.Body.Sort
		resp.Body.Display = input.Body.Display
		resp.Body.TransitionDuration = input.Body.TransitionDuration

		command := CreateSlideshow(input)

		resp.Body.Command = command
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "update-specific-slideshow",
		Method:      http.MethodPut,
		Path:        "/api/v1/slideshows/{foldername}",
		Summary:     "Update an existing slideshow",
		Description: "Update a slideshow by setting all the required options.",
		Tags:        []string{"Slideshows"},
	}, func(ctx context.Context, input *SpecificSlideshowInput) (*SpecificSlideshowOutput, error) {
		resp := &SpecificSlideshowOutput{}
		resp.Body.Duration = input.Body.Duration
		resp.Body.Sort = input.Body.Sort
		resp.Body.Display = input.Body.Display
		resp.Body.TransitionDuration = input.Body.TransitionDuration

		command, err := UpdateExistingSlideshow(input)
		if err != nil {
			return nil, huma.Error404NotFound(err.Error())
		}

		resp.Body.Command = command
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "set-active-slideshow",
		Method:      http.MethodPut,
		Path:        "/api/v1/slideshows",
		Summary:     "Set active slideshow",
		Description: "Set the active slideshow",
		Tags:        []string{"Slideshows"},
	}, func(ctx context.Context, input *SlideshowsInput) (*SlideshowsOutput, error) {
		resp := &SlideshowsOutput{}
		SetActiveSlideshow(input.Body.ActiveSlideshow)

		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "get-slideshow",
		Method:      http.MethodGet,
		Path:        "/api/v1/slideshows/{foldername}",
		Summary:     "Get specific slideshow",
		Description: "Get a specific slideshow's settings",
		Tags:        []string{"Slideshows"},
	}, func(ctx context.Context, input *struct {
		FolderName string `path:"foldername" example:"world-wonders" doc:"slideshow folder to get settings for"`
	}) (*SpecificSlideshowOutput, error) {
		slideshow_settings, feh_command, err := GetSpecificSlideshow(input.FolderName)
		if err != nil {
			return nil, huma.Error404NotFound(err.Error())
		}
		resp := &SpecificSlideshowOutput{}
		resp.Body.BaseSlideShowOutput = *slideshow_settings
		resp.Body.Command = feh_command
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-slideshows",
		Method:      http.MethodGet,
		Path:        "/api/v1/slideshows",
		Summary:     "Get all slideshows",
		Description: "Get the active slideshow and all slideshow settings",
		Tags:        []string{"Slideshows"},
	}, func(ctx context.Context, input *struct{}) (*SlideshowsOutput, error) {
		resp := &SlideshowsOutput{}
		resp.Body.Slideshows = GetAllSlideshows()
		resp.Body.ActiveSlideshow = GetActiveSlideshow()
		return resp, nil
	})
}
