//+build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"path"
)

func main() {
	fmt.Println("Running build script for the Riff trigger")

	var cmd = exec.Command("")

	// appdir is the directory where the app is stored. For example if you app is called
	// lambda this would be <path>/lambda/src/lambda
	appDir := os.Args[1]

	appName := path.Dir(appDir)

	// Clean up
	fmt.Println("Cleaning up previous executables")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "del", "/q", appName+".so")
	} else {
		cmd = exec.Command("rm", "-f", appName+".so")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = appDir

	err := cmd.Run()
	if err != nil {
		fmt.Printf(err.Error())
	}

	// Build an executable for Linux
	fmt.Println("Building a new handler file")
	cmd = exec.Command("go", "build", "-buildmode=plugin", "-o", appName+".so")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = appDir
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOPATH=%s", filepath.Join(appDir, "..", "..")), "GOOS=linux")

	err = cmd.Run()
	if err != nil {
		fmt.Printf(err.Error())
	}
}


