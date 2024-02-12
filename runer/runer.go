package runer

import (
	"launchingprogramlist/config"
	"os/exec"
	"runtime"
	"time"
)

func RunApp(app config.App) error {
	osName := runtime.GOOS

	day := config.IntJson(time.Now().Weekday())
	dayExist := false
	for _, number := range app.Days {
		if number == day {
			dayExist = true
			break
		}
	}
	if !dayExist {
		return nil
	}
	switch osName {
	case "windows":
		return runWindows(app.Path)
	case "linux":
		return runLinux(app.Path)
	case "android":
		return runAndroid(app.Path)
	}

	return nil
}

func runWindows(path config.PathJson) error {
	return runCommon(path)
}

func runLinux(path config.PathJson) error {
	return runCommon(path)
}

func runAndroid(path config.PathJson) error {
	return runCommon(path)
}

func runCommon(path config.PathJson) error {
	cmd := exec.Command(string(path))
	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
