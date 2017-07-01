package applescript_cli

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/meonlol/syncitunes/tree"
)

// TODO: Add interface direct to itunes library? https://github.com/josephw/titl
// https://metacpan.org/pod/distribution/Mac-iTunes/doc/file_format.pod

type Adapter struct {
}

// return id of (make new folder playlist at playlist id 1 with properties {name:\"root\"})
func (a Adapter) NewFolder(name string, parentID int) (int, error) {
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

func (a Adapter) NewPlaylist(name string, parentID int) (int, error) {
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

func (a Adapter) GetPlaylistIDByName(name string) (int, error) {
	out, err := a.runAppleScriptForItunes(`return id in playlist "` + name + `"`)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a Adapter) GetParentIDForPlaylist(id int) (int, error) {
	strID := strconv.Itoa(id)
	out, err := a.runAppleScriptForItunes(`return id of parent in playlist id ` + strID)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a Adapter) AddFileToPlaylist(filePath string, playlistID int) (int, error) {
	strID := strconv.Itoa(playlistID)
	out, err := a.runAppleScriptForItunes(`return id of (add POSIX file "` + filePath + `" to user playlist id ` + strID + `)`)
	intID, _ := strconv.Atoi(out)
	return intID, err
}

func (a Adapter) DeletePlaylistByID(id int) error {
	strID := strconv.Itoa(id)
	_, err := a.runAppleScriptForItunes(`delete playlist id ` + strID)
	return err
}

func (a Adapter) UpdateTreeWithExisting(tree *tree.Node) {
}

func (a Adapter) GetLibrary() (string, error) {
	out, err := a.runAppleScriptForItunes(`
	set resultList to {}
	repeat with currentPlaylist in (get every playlist)
		set playlistName to name of currentPlaylist
		set playlistID to id of currentPlaylist
		set parentID to -1
		try
			set parentID to id of parent in currentPlaylist
		end try
		set trackLocations to {}
		if class of currentPlaylist is user playlist then
			#Don't get tracks for folders
			try
				set trackLocations to {id, location} of every track in currentPlaylist
			end try
		end if
		set isSmart to false
		if class of currentPlaylist is not folder playlist then
			try
				set isSmart to smart of currentPlaylist
			end try
		end if
		if not isSmart then
			copy {playlistName, playlistID, parentID, trackLocations} to end of resultList
		end if
	end repeat
	return resultList
	`)
	return out, err
}

func (a Adapter) runAppleScriptForItunes(commandLines ...string) (string, error) {
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

func (a Adapter) runAppleScript(commandLines ...string) ([]byte, error) {
	params := []string{"-s", "s"}
	for _, v := range commandLines {
		params = append(params, "-e", v)
	}

	cmd := exec.Command("osascript", params...)
	return cmd.CombinedOutput()
}
