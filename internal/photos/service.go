package photos

import (
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/JarvyJ/pislide-api/internal/slideshow"
	"github.com/JarvyJ/pislide-api/internal/util"
)

var photoDir = filepath.Join(util.GetSlideDir(), "slideshows")

func GetAllFolders() []string {
	folders := []string{}
	entries, _ := os.ReadDir(photoDir)
	for _, element := range entries {
		if element.IsDir() {
			folders = append(folders, element.Name())
		}
	}
	return folders
}

func GetAllFilesInFolder(foldername string) []string {
	allFiles := []string{}
	folderToTraverse := filepath.Join(photoDir, foldername)

	strippedPhotoDirLen := len(filepath.Join(photoDir, foldername) + string(os.PathSeparator))

	visit := func(path string, di fs.DirEntry, err error) error {
		// weird, but eliminates the root dirs
		if len(path) > strippedPhotoDirLen {
			fileName := path[strippedPhotoDirLen:]
			allFiles = append(allFiles, fileName)
		}
		return nil
	}
	filepath.WalkDir(folderToTraverse, visit)
	return allFiles
}

func MakeFolder(foldername string) {
	newFolder := filepath.Join(photoDir, foldername)
	// TODO: find out if i care if the folder already exists or not
	os.Mkdir(newFolder, 0755)
}

func RenameFolder(oldFoldername string, newFoldername string) {
	oldFolder := filepath.Join(photoDir, oldFoldername)
	newFolder := filepath.Join(photoDir, newFoldername)
	err := os.Rename(oldFolder, newFolder)
	if err == nil {
		if oldFoldername == slideshow.GetActiveSlideshow() {
			slideshow.SetActiveSlideshow(newFoldername)
		}
	}
}

func GetAllCaptions(foldername string) map[string]string {
	allFiles := GetAllFilesInFolder(foldername)
	fileSet := make(map[string]bool)
	potentialCaptionFiles := []string{}
	for _, file := range allFiles {
		if strings.HasSuffix(file, ".txt") {
			potentialCaptionFiles = append(potentialCaptionFiles, file)
		} else {
			fileSet[file] = true
		}
	}

	captions := map[string]string{}
	for _, file := range potentialCaptionFiles {
		associatedFile := strings.TrimSuffix(file, ".txt")
		if fileSet[associatedFile] == true {
			// TODO: probably some error handling
			caption, _ := os.ReadFile(filepath.Join(photoDir, foldername, file))
			captions[associatedFile] = strings.TrimSpace(string(caption))
		}
	}
	return captions
}

func SetCaptions(foldername string, captions map[string]string) error {
	for filename, _ := range captions {
		associatedFile := filepath.Join(photoDir, foldername, filename)
		// what a dumb way to check if a file exists...
		_, err := os.Stat(associatedFile)
		if err != nil {
			return errors.New("Can only set caption for a photo that exists")
		}
	}
	for filename, caption := range captions {
		associatedFile := filepath.Join(photoDir, foldername, filename)
		captionsFile := associatedFile + ".txt"
		os.WriteFile(captionsFile, []byte(caption), 0644)
	}
	return nil
}

func DeleteCaptions(foldername string, fileCaptionsToDelete []string) {
	for _, caption := range fileCaptionsToDelete {
		os.Remove(filepath.Join(photoDir, foldername, caption+".txt"))
	}
}

func UploadFiles(foldername string, fileHeaderMap map[string][]*multipart.FileHeader) error {
	for _, fileData := range fileHeaderMap {
		for _, fileHeader := range fileData {
			file, err := fileHeader.Open()
			defer file.Close()
			if err != nil {
				return errors.New("Something went wrong opening the uploaded file")
			}

			filePath := filepath.Join(photoDir, foldername, fileHeader.Filename)
			out, err := os.Create(filePath)
			defer out.Close()
			if err != nil {
				return errors.New("Unable to start writing file to directory")
			}

			_, err = io.Copy(out, file)
			if err != nil {
				return errors.New("Unable to write uploaded file")
			}
		}
	}
	return nil
}
