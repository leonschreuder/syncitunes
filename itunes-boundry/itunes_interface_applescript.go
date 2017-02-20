package itunes

import (
	"errors"
	"os/exec"
	"strings"
)

// TODO: Add interface direct to itunes library? https://github.com/josephw/titl
// https://metacpan.org/pod/distribution/Mac-iTunes/doc/file_format.pod

type applescriptInterface struct {
}

func (a applescriptInterface) newFolder(name string) string {
	output, _ := a.runAppleScriptForItunes(`make new folder playlist with properties {name:"` + name + `"}`)
	words := strings.Split(string(output), " ") //returns: folder playlist id 71268 of source id 66 of application "iTunes"
	return words[3]
}

func (a applescriptInterface) newPlaylist(name, parentID string) string {
	var output []byte
	if parentID != "" {
		output, _ = a.runAppleScriptForItunes(`return id of (make new user playlist at playlist id ` + parentID + ` with properties {name:"` + name + `"})`)
	} else {
		output, _ = a.runAppleScriptForItunes(`return id of (make new user playlist with properties {name:"` + name + `"})`)
	}
	return strings.TrimSpace(string(output))
}

func (a applescriptInterface) getPlaylistIDByName(name string) (string, error) {
	requestOutput, err := a.runAppleScriptForItunes(
		`return id in playlist "` + name + `"`,
	)
	stdOut := strings.TrimSpace(string(requestOutput))
	if err != nil {
		return "", errors.New(stdOut)
	}
	return stdOut, err
}

func (a applescriptInterface) getParentIDForPlaylist(id string) (string, error) {
	requestOutput, err := a.runAppleScriptForItunes(
		`return id of parent in playlist id ` + id,
	)
	stdOut := strings.TrimSpace(string(requestOutput))
	if err != nil {
		return "", errors.New(stdOut)
	}
	return stdOut, err
}

func (a applescriptInterface) addFileToPlaylist(filePath, playlistID string) {
	a.runAppleScriptForItunes(
		`add POSIX file "` + filePath + `" to user playlist id ` + playlistID,
	)
}

func (a applescriptInterface) deletePlaylistByID(id string) {
	a.runAppleScriptForItunes(`delete playlist id ` + id)
}

func (a applescriptInterface) runAppleScriptForItunes(commandLines ...string) ([]byte, error) {
	iTunesLines := []string{`tell application "iTunes"`}
	iTunesLines = append(iTunesLines, commandLines...)
	iTunesLines = append(iTunesLines, `end tell`)
	return a.runAppleScript(iTunesLines...)
}

func (a applescriptInterface) runAppleScript(commandLines ...string) ([]byte, error) {
	params := []string{"-s", "s"}
	for _, v := range commandLines {
		params = append(params, "-e", v)
	}

	cmd := exec.Command("osascript", params...)
	return cmd.CombinedOutput()
}
