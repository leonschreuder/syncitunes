package main

import (
	"errors"
	"os/exec"
	"strings"
)

func newPlaylist(name string) string {
	output, _ := runAppleScriptForItunes(`make new user playlist with properties {name:"` + name + `"}`)
	words := strings.Split(string(output), " ") //gets: "user playlist id 57494 of source id 66 of application "iTunes""
	return words[3]
}

func getPlaylistIDByName(name string) (string, error) {
	requestOutput, err := runAppleScriptForItunes(
		`try`,
		`    return id in playlist "`+name+`"`,
		`on error number -1728`,
		`    return`,
		`end try`,
	)
	if err != nil {
		return string(requestOutput), err
	} else if string(requestOutput) == "" {
		return string(requestOutput), errors.New("no playlist found")
	} else {
		return strings.TrimSpace(string(requestOutput)), err
	}
}

func deletePlaylistByID(id string) {
	runAppleScriptForItunes(`delete user playlist id ` + id)
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
