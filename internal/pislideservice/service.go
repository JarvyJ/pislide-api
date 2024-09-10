package pislideservice

import (
	"errors"
	"os/exec"
)

func getSystemCtlPath() (string, error) {
	path, err := exec.LookPath("systemctl")
	if err != nil {
		return "", errors.New("Not able to find systemctl")
	}
	return path, nil
}

func piSlideSericeAction(action string) (string, int, error) {
	path, err := getSystemCtlPath()
	if err != nil {
		return "", 0, err
	}
	command := exec.Command(path, action, "pislide.service")
	fullOuput, err := command.CombinedOutput()
	statusCode := 0
	if exiterr, ok := err.(*exec.ExitError); ok {
		statusCode = exiterr.ExitCode()
	}
	return string(fullOuput), statusCode, nil
}

func GetPiSlideServiceStatus() (string, int, error) {
	return piSlideSericeAction("status")
}

func StartPiSlideService() (string, int, error) {
	return piSlideSericeAction("start")
}

func StopPiSlideService() (string, int, error) {
	return piSlideSericeAction("stop")
}

func RestartPiSlideService() (string, int, error) {
	return piSlideSericeAction("restart")
}
