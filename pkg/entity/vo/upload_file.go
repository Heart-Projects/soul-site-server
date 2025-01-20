package vo

type UploadedFile struct {
	Key  string `json:"key"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}
