package itunes

type Interface interface {
	NewFolder(name string, id int) (int, error)
	NewPlaylist(name string, parentID int) (int, error)
	GetPlaylistIDByName(name string) (int, error)
	//GetPlaylistIDByNameInParent(name string) (int, error)
	GetParentIDForPlaylist(id int) (int, error)
	AddFileToPlaylist(filePath string, playlistID int) (int, error)
	DeletePlaylistByID(id int) error
	// UpdateTreeWithExisting(tree *tree.Node)
}

type ItemType int

const (
	File ItemType = iota
	Dir
	Playlist
)
