package alidrive

import (
	"time"
)

// RespHandle response bean methods
type RespHandle interface {
	IsAvailable() bool   // check available
	GetCode() string     // get err code
	GetMessage() string  // get err message
	SetCode(code string) // set err code
}

// RespError common response bean
type RespError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (resp *RespError) IsAvailable() bool {
	return resp.Code == ""
}

func (resp *RespError) GetCode() string {
	return resp.Code
}

func (resp *RespError) GetMessage() string {
	return resp.Message
}

func (resp *RespError) SetCode(code string) {
	resp.Code = code
}

// UserInfo user_info response bean
type UserInfo struct {
	RespError
	DomainId       string                 `json:"domain_id"`
	UserId         string                 `json:"user_id"`
	Avatar         string                 `json:"avatar"`
	CreatedAt      int64                  `json:"created_at"`
	UpdatedAt      int64                  `json:"updated_at"`
	Email          string                 `json:"email"`
	NickName       string                 `json:"nick_name"`
	Phone          string                 `json:"phone"`
	Role           string                 `json:"role"`
	Status         string                 `json:"status"`
	UserName       string                 `json:"user_name"`
	Description    string                 `json:"description"`
	DefaultDriveId string                 `json:"default_drive_id"`
	UserData       map[string]interface{} `json:"user_data"`
}

// Files folder files response bean
type Files struct {
	RespError
	Items      []File `json:"items"`
	NextMarker string `json:"next_marker"`
	Readme     string `json:"readme"` // Deprecated
	Paths      []Path `json:"paths"`
}

// Path path bean
type Path struct {
	Name   string `json:"name"`
	FileId string `json:"file_id"`
}

/** 秒传
{
	"name":"mikuclub.mp4",
	"content_hash":"C733AC50D1F964C0398D0E403F3A30C37EFC2ADD",
	"size":1141068377,
	"content_type":"video/mp4"
}
*/

// File file response bean
type File struct {
	RespError
	DriveId       string     `json:"drive_id"`
	CreatedAt     *time.Time `json:"created_at"`
	DomainId      string     `json:"domain_id"`
	EncryptMode   string     `json:"encrypt_mode"`
	FileExtension string     `json:"file_extension"`
	FileId        string     `json:"file_id"`
	Hidden        bool       `json:"hidden"`
	Name          string     `json:"name"`
	ParentFileId  string     `json:"parent_file_id"`
	Starred       bool       `json:"starred"`
	Status        string     `json:"status"`
	Type          string     `json:"type"`
	UpdatedAt     *time.Time `json:"updated_at"`
	// 文件多出部分
	Category           string                 `json:"category"`
	ContentHash        string                 `json:"content_hash"`
	ContentHashName    string                 `json:"content_hash_name"`
	ContentType        string                 `json:"content_type"`
	Crc64Hash          string                 `json:"crc_64_hash"`
	DownloadUrl        string                 `json:"download_url"`
	PunishFlag         int64                  `json:"punish_flag"`
	Size               int64                  `json:"size"`
	Thumbnail          string                 `json:"thumbnail"`
	Url                string                 `json:"url"`
	ImageMediaMetadata map[string]interface{} `json:"image_media_metadata"`

	Paths []Path `json:"paths"`
}

type DownloadResp struct {
	RespError
	Expiration string `json:"expiration"`
	Method     string `json:"method"`
	Size       int64  `json:"size"`
	Url        string `json:"url"`
	//RateLimit struct{
	//	PartSize int `json:"part_size"`
	//	PartSpeed int `json:"part_speed"`
	//} `json:"rate_limit"`//rate limit
}

// TokenLoginResp token_login response bean
type TokenLoginResp struct {
	RespError
	Goto string `json:"goto"`
}

// TokenResp token response bean
type TokenResp struct {
	RespError
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`

	UserInfo

	DefaultSboxDriveId string        `json:"default_sbox_drive_id"`
	ExpireTime         *time.Time    `json:"expire_time"`
	State              string        `json:"state"`
	ExistLink          []interface{} `json:"exist_link"`
	NeedLink           bool          `json:"need_link"`
	PinSetup           bool          `json:"pin_setup"`
	IsFirstLogin       bool          `json:"is_first_login"`
	NeedRpVerify       bool          `json:"need_rp_verify"`
	DeviceId           string        `json:"device_id"`
}

// OfficePreviewUrlResp office_preview_url response bean
type OfficePreviewUrlResp struct {
	RespError
	PreviewUrl  string `json:"preview_url"`
	AccessToken string `json:"access_token"`
}

type VideoPreviewUrlResp struct {
	RespError
	TemplateList []struct {
		TemplateId string `json:"template_id"`
		Status     string `json:"status"`
		Url        string `json:"url"`
	} `json:"template_list"`
}

type VideoPreviewPlayInfoResp struct {
	RespError
	VideoPreviewPlayInfo struct {
		LiveTranscodingTaskList []struct {
			TemplateId string `json:"template_id"`
			Status     string `json:"status"`
			Url        string `json:"url"`
			Stage      string `json:"stage"`
		} `json:"live_transcoding_task_list"`
	} `json:"video_preview_play_info"`
}
