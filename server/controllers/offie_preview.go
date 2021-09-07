package controllers

import (
	"github.com/bh-qt/alist/alidrive"
	"github.com/bh-qt/alist/conf"
	"github.com/bh-qt/alist/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type OfficePreviewReq struct {
	FileId string `json:"file_id" binding:"required"`
}

// handle office_preview request
func OfficePreview(c *gin.Context) {
	if !conf.Conf.Server.Download {
		c.JSON(200,MetaResponse(403,"not allowed download and preview"))
		return
	}
	drive := utils.GetDriveByName(c.Param("drive"))
	if drive == nil {
		c.JSON(200, MetaResponse(400, "drive isn't exist."))
		return
	}
	var req OfficePreviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, MetaResponse(400, "Bad Request:"+err.Error()))
		return
	}
	log.Debugf("preview_req:%+v", req)
	preview, err := alidrive.GetOfficePreviewUrl(req.FileId, drive)
	if err != nil {
		c.JSON(200, MetaResponse(500, err.Error()))
		return
	}
	c.JSON(200, DataResponse(preview))
}

