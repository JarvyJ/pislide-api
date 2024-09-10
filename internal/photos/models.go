package photos

import "mime/multipart"

type foldersOutput struct {
	Body struct {
		Folders []string `json:"folders" doc:"List of all the folders that may contain photos"`
	}
}

type folderFilesOutput struct {
	Body struct {
		Files []string `json:"files" doc:"List of all the files (recursively) in the folder"`
	}
}

type folderInput struct {
	Foldername string `path:"foldername"`
}

type folderRenameInput struct {
	Foldername string `path:"foldername"`
	Body       struct {
		NewFoldername string `json:"new_foldername" doc:"The new foldername to change to`
	}
}

type captionsOutput struct {
	Body struct {
		Captions map[string]string `json:"captions" doc:"Map of captions associated with files"`
	}
}

type captionsSetInput struct {
	Foldername string `path:"foldername"`
	Body       struct {
		Captions map[string]string `json:"captions" doc:"Map of captions associated with files. Will upsert and not delete existing if not specified"`
	}
}

type captionsDeleteInput struct {
	Foldername string `path:"foldername"`
	Body       struct {
		CaptionsToDelete []string `json:"captions_to_delete" doc:"List of captions to be deleted"`
	}
}

type fileUploadToFolder struct {
	Foldername string `path:"foldername"`
	RawBody    multipart.Form
}
