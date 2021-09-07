package alidrive

// ListReq list request bean
type ListReq struct {
	DriveId               string `json:"drive_id"`
	Fields                string `json:"fields"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageUrlProcess       string `json:"image_url_process"`
	Limit                 int    `json:"limit"`
	Marker                string `json:"marker"`
	OrderBy               string `json:"order_by"`
	OrderDirection        string `json:"order_direction"`
	ParentFileId          string `json:"parent_file_id"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

// GetReq get request bean
type GetReq struct {
	DriveId               string `json:"drive_id"`
	FileId                string `json:"file_id"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

// DownloadReq download request bean
type DownloadReq struct {
	DriveId   string `json:"drive_id"`
	FileId    string `json:"file_id"`
	ExpireSec int    `json:"expire_sec"`
	FileName  string `json:"file_name"`
}

// SearchReq search request bean
type SearchReq struct {
	DriveId               string `json:"drive_id"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageUrlProcess       string `json:"image_url_process"`
	Limit                 int    `json:"limit"`
	Marker                string `json:"marker"`
	OrderBy               string `json:"order_by"` //"type ASC,updated_at DESC"

	Query string `json:"query"` // "name match '测试文件'"

	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

// TokenLoginReq token_login request bean
type TokenLoginReq struct {
	Token string `json:"token"`
}

// GetTokenReq get_token request bean
type GetTokenReq struct {
	Code string `json:"code"`
}

// RefreshTokenReq refresh_token request bean
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

// OfficePreviewUrlReq office_preview_url request bean
type OfficePreviewUrlReq struct {
	AccessToken string `json:"access_token"`
	DriveId     string `json:"drive_id"`
	FileId      string `json:"file_id"`
}

// VideoPreviewUrlReq video preview url request bean
type VideoPreviewUrlReq struct {
	DriveId   string `json:"drive_id"`
	FileId    string `json:"file_id"`
	ExpireSec int    `json:"expire_sec"`
}

// VideoPreviewPlayInfoReq video preview play info req
type VideoPreviewPlayInfoReq struct {
	Category   string `json:"category"`
	DriveId    string `json:"drive_id"`
	FileId     string `json:"file_id"`
	TemplateId string `json:"template_id"`
}
