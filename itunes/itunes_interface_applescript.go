package itunes

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

// TODO: Add interface direct to itunes library? https://github.com/josephw/titl
// https://metacpan.org/pod/distribution/Mac-iTunes/doc/file_format.pod

type ApplescriptInterface struct {
}

// return id of (make new folder playlist at playlist id 1 with properties {name:\"root\"})
func (a ApplescriptInterface) NewFolder(name string, parentID int) (int, error) {
	var output string
	var err error
	if parentID > 0 {
		strID := strconv.Itoa(parentID)
		output, err = a.runAppleScriptForItunes(`return id of (make new folder playlist at playlist id ` + strID + ` with properties {name:"` + name + `"})`)
	} else {
		output, err = a.runAppleScriptForItunes(`return id of (make new folder playlist with properties {name:"` + name + `"})`)
	}
	// output, err := a.runAppleScriptForItunes(`make new folder playlist with properties {name:"` + name + `"}`)
	if err != nil {
		return -1, err
	}
	// words := strings.Split(string(output), " ") //returns: folder playlist id 71268 of source id 66 of application "iTunes"
	// intID, err := strconv.Atoi(words[3])
	intID, err := strconv.Atoi(output)
	return intID, err
}

func (a ApplescriptInterface) NewPlaylist(name string, parentID int) (int, error) {
	var output string
	var err error
	if parentID > 0 {
		strID := strconv.Itoa(parentID)
		output, err = a.runAppleScriptForItunes(`return id of (make new user playlist at playlist id ` + strID + ` with properties {name:"` + name + `"})`)
	} else {
		output, err = a.runAppleScriptForItunes(`return id of (make new user playlist with properties {name:"` + name + `"})`)
	}
	if err != nil {
		return -1, err
	}
	intID, err := strconv.Atoi(output)
	return intID, err
}

func (a ApplescriptInterface) GetPlaylistIDByName(name string) (int, error) {
	out, err := a.runAppleScriptForItunes(`return id in playlist "` + name + `"`)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a ApplescriptInterface) GetParentIDForPlaylist(id int) (int, error) {
	strID := strconv.Itoa(id)
	out, err := a.runAppleScriptForItunes(`return id of parent in playlist id ` + strID)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a ApplescriptInterface) AddFileToPlaylist(filePath string, playlistID int) (int, error) {
	strID := strconv.Itoa(playlistID)
	out, err := a.runAppleScriptForItunes(`return id of (add POSIX file "` + filePath + `" to user playlist id ` + strID + `)`)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a ApplescriptInterface) DeletePlaylistByID(id int) error {
	strID := strconv.Itoa(id)
	_, err := a.runAppleScriptForItunes(`delete playlist id ` + strID)
	return err
}

func (a ApplescriptInterface) runAppleScriptForItunes(commandLines ...string) (string, error) {
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

func (a ApplescriptInterface) runAppleScript(commandLines ...string) ([]byte, error) {
	params := []string{"-s", "s"}
	for _, v := range commandLines {
		params = append(params, "-e", v)
	}

	cmd := exec.Command("osascript", params...)
	return cmd.CombinedOutput()
}
