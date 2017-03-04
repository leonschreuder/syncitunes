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
	var output string
	if parentID != "" {
		output, _ = a.runAppleScriptForItunes(`return id of (make new user playlist at playlist id ` + parentID + ` with properties {name:"` + name + `"})`)
	} else {
		output, _ = a.runAppleScriptForItunes(`return id of (make new user playlist with properties {name:"` + name + `"})`)
	}
	return output
}

func (a applescriptInterface) getPlaylistIDByName(name string) (string, error) {
	return a.runAppleScriptForItunes(`return id in playlist "` + name + `"`)
}

func (a applescriptInterface) getParentIDForPlaylist(id string) (string, error) {
	return a.runAppleScriptForItunes(`return id of parent in playlist id ` + id)
}

func (a applescriptInterface) addFileToPlaylist(filePath, playlistID string) (string, error) {
	return a.runAppleScriptForItunes(`return id of (add POSIX file "` + filePath + `" to user playlist id ` + playlistID + `)`)
}

func (a applescriptInterface) deletePlaylistByID(id string) error {
	_, err := a.runAppleScriptForItunes(`delete playlist id ` + id)
	return err
}

func (a applescriptInterface) runAppleScriptForItunes(commandLines ...string) (string, error) {
	iTunesLines := []string{`tell application "iTunes"`}
	iTunesLines = append(iTunesLines, commandLines...)
	iTunesLines = append(iTunesLines, `end tell`)

	requestOutput, err := a.runAppleScript(iTunesLines...)
	stdOut := strings.TrimSpace(string(requestOutput))
	if err != nil {
		return "", errors.New(stdOut)
	}
	return stdOut, err
}

func (a applescriptInterface) runAppleScript(commandLines ...string) ([]byte, error) {
	params := []string{"-s", "s"}
	for _, v := range commandLines {
		params = append(params, "-e", v)
	}

	cmd := exec.Command("osascript", params...)
	return cmd.CombinedOutput()
}
