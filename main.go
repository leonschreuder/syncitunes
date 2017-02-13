package main

import (
	"errors"
	"os/exec"
	"strings"
)

func newFolder(name string) string {
	output, _ := runAppleScriptForItunes(`make new folder playlist with properties {name:"` + name + `"}`)
	words := strings.Split(string(output), " ") //returns: folder playlist id 71268 of source id 66 of application "iTunes"
	return words[3]
}

func newPlaylist(name, parentID string) string {
	var output []byte
	if parentID != "" {
		output, _ = runAppleScriptForItunes(`return id of (make new user playlist at playlist id ` + parentID + ` with properties {name:"` + name + `"})`)
	} else {
		output, _ = runAppleScriptForItunes(`return id of (make new user playlist with properties {name:"` + name + `"})`)
	}
	return strings.TrimSpace(string(output))
}

func getPlaylistIDByName(name string) (string, error) {
	requestOutput, err := runAppleScriptForItunes(
		`return id in playlist "` + name + `"`,
	)
	stdOut := strings.TrimSpace(string(requestOutput))
	if err != nil {
		return "", errors.New(stdOut)
	}
	return stdOut, err
}

func getParentIDForPlaylist(id string) (string, error) {
	requestOutput, err := runAppleScriptForItunes(
		`return id of parent in playlist id ` + id,
	)
	stdOut := strings.TrimSpace(string(requestOutput))
	if err != nil {
		return "", errors.New(stdOut)
	}
	return stdOut, err
}

func deletePlaylistByID(id string) {
	runAppleScriptForItunes(`delete playlist id ` + id)
}

func runAppleScriptForItunes(commandLines ...string) ([]byte, error) {
	iTunesLines := []string{`tell application "iTunes"`}
	iTunesLines = append(iTunesLines, commandLines...)
	iTunesLines = append(iTunesLines, `end tell`)
	return runAppleScript(iTunesLines...)
}

func runAppleScript(commandLines ...string) ([]byte, error) {
	params := []string{"-s", "s"}
	for _, v := range commandLines {
		params = append(params, "-e", v)
	}

	cmd := exec.Command("osascript", params...)
	return cmd.CombinedOutput()
}
