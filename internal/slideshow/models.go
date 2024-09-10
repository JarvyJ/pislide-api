package slideshow

type SlideshowsInput struct {
	Body struct {
		ActiveSlideshow string `json:"active_slideshow" example:"world-wonders" doc:"The active slideshow"`
	}
}

type BaseSlideShowOutput struct {
	Duration           float32 `json:"duration" example:"3.2" doc:"Slide duration in seconds"`
	Sort               string  `json:"sort" default:"filename" enum:"filename,natural,random"`
	TransitionDuration float32 `json:"transition_duration" example:"3.2" doc:"Duration of fade transition in seconds"`
	Display            string  `json:"display" default:"none" enum:"filename,caption,none"`
}

type SlideshowsOutput struct {
	Body struct {
		ActiveSlideshow string                         `json:"active_slideshow" example:"world-wonders" doc:"The active slideshow"`
		Slideshows      map[string]BaseSlideShowOutput `json:"slideshows"`
	}
}

type SpecificSlideshowInput struct {
	FolderName string `path:"foldername" maxLength:"30" example:"world-wonders" doc:"Folder the images reside in"`
	// some of these are "omitempty", and we don't want that when returning, hence the duplication in BaseSlideShowOutput
	Body struct {
		Duration           float32 `json:"duration" example:"3.2" doc:"Slideshow delay in seconds"`
		Sort               string  `json:"sort,omitempty" default:"filename" enum:"filename,natural,random"`
		TransitionDuration float32 `json:"transition_duration,omitempty" example:"3.2" doc:"Duration of fade transition in seconds"`
		Display            string  `json:"display,omitempty" default:"none" enum:"filename,caption,none"`
	}
}

type SpecificSlideshowOutput struct {
	Body struct {
		BaseSlideShowOutput
		Command string `json:"command" doc:"Actual rayimg command that will run the slideshow"`
	}
}
