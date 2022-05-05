package model

type FileInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
	CreateAt  int    `json:"createAt"`
	DeleteAt  int    `json:"deleteAt"`
}
