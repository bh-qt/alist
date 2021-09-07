package controllers

import (
	"github.com/bh-qt/alist/alidrive"
	"github.com/bh-qt/alist/conf"
	"github.com/bh-qt/alist/server/models"
	"github.com/bh-qt/alist/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

type DownReq struct {
	Password string `form:"pw"`
}

// Down handle download request
func Down(c *gin.Context) {
	if !conf.Conf.Server.Download {
		c.JSON(200,MetaResponse(403,"not allowed download and preview"))
		return
	}
	filePath := c.Param("path")[1:]
	var down DownReq
	if err := c.ShouldBindQuery(&down); err != nil {
		c.JSON(200, MetaResponse(400, "Bad Request."))
		return
	}
	log.Debugf("down:%s", filePath)
	dir, name := filepath.Split(filePath)
	fileModel, err := models.GetFileByDirAndName(dir, name)
	if err != nil {
		if fileModel == nil {
			c.JSON(200, MetaResponse(404, "Path not found."))
			return
		}
		c.JSON(200, MetaResponse(500, err.Error()))
		return
	}
	if fileModel.Password != "" && down.Password != utils.Get16MD5Encode(fileModel.Password) {
		if down.Password == "" {
			c.JSON(200, MetaResponse(401, "need password."))
		} else {
			c.JSON(200, MetaResponse(401, "wrong password."))
		}
		return
	}
	if fileModel.Type == "folder" {
		c.JSON(200, MetaResponse(406, "无法下载目录."))
		return
	}
	drive := utils.GetDriveByName(strings.Split(filePath, "/")[0])
	if drive == nil {
		c.JSON(200, MetaResponse(500, "找不到drive."))
		return
	}
	file, err := alidrive.GetDownLoadUrl(fileModel.FileId, drive)
	if err != nil {
		c.JSON(200, MetaResponse(500, err.Error()))
		return
	}
	c.Redirect(302, file.Url)
	return
}
