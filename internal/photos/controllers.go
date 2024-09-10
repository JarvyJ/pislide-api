package photos

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPhotos(api *huma.API) {
	huma.Register(*api, huma.Operation{
		OperationID: "get-all-photo-folders",
		Method:      http.MethodGet,
		Path:        "/api/v1/photos",
		Summary:     "Get list of photo folders",
		Description: "Get list of folders that may contain photos",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *struct{}) (*foldersOutput, error) {
		resp := &foldersOutput{}
		folders := GetAllFolders()
		resp.Body.Folders = folders
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "create-a-folder",
		Method:      http.MethodPost,
		Path:        "/api/v1/photos",
		Summary:     "Create a folder",
		Description: "Create a folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *folderRenameInput) (*foldersOutput, error) {
		resp := &foldersOutput{}
		MakeFolder(input.Body.NewFoldername)
		resp.Body.Folders = GetAllFolders()
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-files-in-folders",
		Method:      http.MethodGet,
		Path:        "/api/v1/photos/{foldername}",
		Summary:     "Get all files in a folder",
		Description: "Get list of all files in a folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *folderInput) (*folderFilesOutput, error) {
		resp := &folderFilesOutput{}
		resp.Body.Files = GetAllFilesInFolder(input.Foldername)
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "rename-a-folder",
		Method:      http.MethodPut,
		Path:        "/api/v1/photos/{foldername}",
		Summary:     "Rename a folder",
		Description: "Rename a folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *folderRenameInput) (*folderFilesOutput, error) {
		resp := &folderFilesOutput{}
		RenameFolder(input.Foldername, input.Body.NewFoldername)
		resp.Body.Files = GetAllFilesInFolder(input.Body.NewFoldername)
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "get-captions-from-folder",
		Method:      http.MethodGet,
		Path:        "/api/v1/photos/{foldername}/captions",
		Summary:     "Get all captions in a folder",
		Description: "Go through and find all valid captions files associated with images",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *folderInput) (*captionsOutput, error) {
		resp := &captionsOutput{}
		captions := GetAllCaptions(input.Foldername)
		resp.Body.Captions = captions
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "set-captions-in-folder",
		Method:      http.MethodPut,
		Path:        "/api/v1/photos/{foldername}/captions",
		Summary:     "Set captions in a folder",
		Description: "Upsert captions into the folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *captionsSetInput) (*captionsOutput, error) {
		resp := &captionsOutput{}
		err := SetCaptions(input.Foldername, input.Body.Captions)
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error())
		}
		captions := GetAllCaptions(input.Foldername)
		resp.Body.Captions = captions
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "delete-captions-in-folder",
		Method:      http.MethodDelete,
		Path:        "/api/v1/photos/{foldername}/captions",
		Summary:     "Delete captions in a folder",
		Description: "Delete specified captions from a folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *captionsDeleteInput) (*captionsOutput, error) {
		resp := &captionsOutput{}
		DeleteCaptions(input.Foldername, input.Body.CaptionsToDelete)
		captions := GetAllCaptions(input.Foldername)
		resp.Body.Captions = captions
		return resp, nil
	})

	huma.Register(*api, huma.Operation{
		OperationID: "upload-files",
		Method:      http.MethodPost,
		Path:        "/api/v1/photos/{foldername}",
		Summary:     "Upload a file to a folder",
		Tags:        []string{"Photos"},
	}, func(ctx context.Context, input *fileUploadToFolder) (*folderFilesOutput, error) {
		// TODO: check foldername dir exists

		err := UploadFiles(input.Foldername, input.RawBody.File)
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error())
		}

		output := &folderFilesOutput{}
		output.Body.Files = GetAllFilesInFolder(input.Foldername)
		return output, nil
	})
}
