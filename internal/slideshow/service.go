package slideshow

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JarvyJ/pislide-api/internal/util"
)

var slideDir = filepath.Join(util.GetSlideDir(), "slideshows")

func createSlideshowCommand(input *SpecificSlideshowInput) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("rayimg --recursive --duration %f ", input.Body.Duration))

	sb.WriteString("--sort ")
	sb.WriteString(input.Body.Sort)
	sb.WriteString(" ")

	if input.Body.Display != "" {
		sb.WriteString("--display ")
		sb.WriteString(input.Body.Display)
		sb.WriteString(" ")
	}

	if input.Body.TransitionDuration > 0 {
		sb.WriteString(fmt.Sprintf("--transition-duration %f ", input.Body.TransitionDuration))
	}

	sb.WriteString(".")
	return sb.String()
}

func parseSlideShowSettings(path string) (*BaseSlideShowOutput, string) {
	slideshow_settings := &BaseSlideShowOutput{}
	data, _ := os.ReadFile(path)
	dataString := string(data)
	feh_command := ""
	for _, line := range strings.Split(dataString, "\n") {
		if strings.HasPrefix(line, "rayimg ") {
			feh_command = line
			commands := strings.Split(line, " ")
			for i, command := range commands {
				switch command {

				case "--duration":
					slideshowdelay, _ := strconv.ParseFloat(commands[i+1], 32)
					slideshow_settings.Duration = float32(slideshowdelay)

				case "--sort":
					slideshow_settings.Sort = commands[i+1]

				case "--transition-duration":
					slideshowtransition, _ := strconv.ParseFloat(commands[i+1], 32)
					slideshow_settings.TransitionDuration = float32(slideshowtransition)

				case "--display":
					slideshow_settings.Display = commands[i+1]

				default:
					continue
				}
			}
		}
	}
	return slideshow_settings, feh_command
}

func CreateSlideshow(input *SpecificSlideshowInput) string {
	command := createSlideshowCommand(input)

	path := filepath.Join(slideDir, input.FolderName)
	os.MkdirAll(path, 0755)
	os.WriteFile(filepath.Join(path, "run_slideshow.sh"), []byte("#!/bin/bash\n"+command), 0744)
	return command
}

func UpdateExistingSlideshow(input *SpecificSlideshowInput) (string, error) {
	path := filepath.Join(slideDir, input.FolderName)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("The slideshow folder '%s' does not exist", path))
	}

	command := createSlideshowCommand(input)
	os.WriteFile(filepath.Join(path, "run_slideshow.sh"), []byte("#!/bin/bash\n"+command), 0744)

	return command, nil
}

func SetActiveSlideshow(activeSlideshow string) error {
	folderpath := filepath.Join(slideDir, activeSlideshow)
	_, err := os.Stat(folderpath)
	if os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("The slideshow folder '%s' does not exist", folderpath))
	}

	var sb strings.Builder
	sb.WriteString("#!/bin/bash\n")
	sb.WriteString("cd ")
	sb.WriteString(activeSlideshow)
	sb.WriteString("\n")
	sb.WriteString("./run_slideshow.sh")
	command := sb.String()

	os.WriteFile(filepath.Join("slideshows", "activate_slideshow.sh"), []byte(command), 0744)
	return nil
}

func GetSpecificSlideshow(folderName string) (*BaseSlideShowOutput, string, error) {
	slidehshow_file := filepath.Join(slideDir, folderName, "run_slideshow.sh")
	_, err := os.Stat(slidehshow_file)
	if os.IsNotExist(err) {
		return nil, "", errors.New(fmt.Sprintf("The slideshow file '%s' does not exist", slidehshow_file))
	}
	output, command := parseSlideShowSettings(slidehshow_file)
	return output, command, nil
}

func GetActiveSlideshow() string {
	data, _ := os.ReadFile(filepath.Join(slideDir, "activate_slideshow.sh"))
	dataString := string(data)
	for _, line := range strings.Split(dataString, "\n") {
		if strings.HasPrefix(line, "cd ") {
			return strings.Split(line, " ")[1]
		}
	}
	return ""
}

func GetAllSlideshows() map[string]BaseSlideShowOutput {
	output := map[string]BaseSlideShowOutput{}
	matches, _ := filepath.Glob(filepath.Join(slideDir, "*", "run_slideshow.sh"))
	for _, match := range matches {
		settings, _ := parseSlideShowSettings(match)
		filePathSplit := strings.Split(match, "/")
		output[filePathSplit[len(filePathSplit)-2]] = *settings
	}
	return output
}
