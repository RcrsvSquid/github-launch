package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const SSHExample = "git@github.com:Rcrsvsquid/github-launch.git"

func main() {
	remoteName := parseArgs()
	remotePath := fmt.Sprintf("remote.%s.url", remoteName)
	gitRemoteUrlCmd := exec.Command("git", "config", "--local", remotePath)
	out, err := gitRemoteUrlCmd.Output()

	if err != nil {
		var waitStatus syscall.WaitStatus
		if exitErr, ok := err.(*exec.ExitError); ok {
			waitStatus = exitErr.Sys().(syscall.WaitStatus)
			statusCode := waitStatus.ExitStatus()
			handleGitRemoteErr(statusCode)
		}
		log.Fatal(err)
	}

	var remoteUrl string = strings.TrimSpace(string(out))
	if !(len(remoteUrl) > 0) {
		handleGitRemoteErr(1)
	}

	if !strings.HasPrefix(remoteUrl, "http") {
		remoteUrl = sshToHttp(remoteUrl)
	}

	fmt.Println(remoteUrl)
	browserCmd := exec.Command("open", remoteUrl)
	err = browserCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() (remoteName string) {
	switch true {
	case !(len(os.Args) > 1):
		remoteName = "origin"
	case os.Args[1] == "-h" || os.Args[1] == "--help":
		fmt.Println("usage: gh-launch [remoteName:{origin}]")
		fmt.Println("Launches github homepage from a git directory")
		os.Exit(0)
	default:
		remoteName = os.Args[1]
	}

	return
}

func handleGitRemoteErr(statusCode int) {
	switch statusCode {
	case 128:
		log.Println("Not a git repository")
	case 1:
		log.Println("Remote url doesn't exist")
	default:
		log.Println("Unknown error occurred")
	}

	os.Exit(statusCode)
}

func sshToHttp(sshStr string) string {
	splitAtSymbol := strings.Split(sshStr, "@")
	splitUserRepo := strings.Split(splitAtSymbol[1], ":")
	return fmt.Sprintf("https://%s/%s", splitUserRepo[0], splitUserRepo[1])
}
