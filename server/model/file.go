package model

// FileInfo represents the minimal version of Mattermost's FileInfo model.
// We only use yhe fields defined below in Focalboard.
// This model exists to be used in both plugin and personal-server SKU, hence
// not reusing Mattermost's FileInfo type.
type FileInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
	CreateAt  int64  `json:"createAt"`
	DeleteAt  int64  `json:"deleteAt"`
	Archived  bool   `json:"archived"`
}
