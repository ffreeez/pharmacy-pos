package sysloghandler

import (
	"strconv"

	syslogmodel "pharmacy-pos/pkg/db/models/syslog"
	syslogservice "pharmacy-pos/pkg/service/syslog"
	"pharmacy-pos/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SyslogHandler 处理系统日志相关的 HTTP 请求
type SyslogHandler struct {
	SyslogService *syslogservice.SyslogService
}

// NewSyslogHandler 创建一个新的 SyslogHandler 实例
func NewSyslogHandler(db *gorm.DB) *SyslogHandler {
	return &SyslogHandler{
		SyslogService: syslogservice.NewSyslogService(db),
	}
}

// CreateSyslog 创建新的系统日志条目
func (sh *SyslogHandler) CreateSyslog(c *gin.Context) {
	var log syslogmodel.Syslog
	if err := c.ShouldBindJSON(&log); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := sh.SyslogService.CreateSyslog(&log)
	if err != nil {
		response.InternalServerError(c, "Failed to create syslog")
		return
	}

	response.Created(c, log, "success")
}

// GetAllSyslogs 获取所有系统日志条目
func (sh *SyslogHandler) GetAllSyslogs(c *gin.Context) {
	logs, err := sh.SyslogService.GetAllSyslogs()
	if err != nil {
		response.InternalServerError(c, "Failed to get all syslogs")
		return
	}

	response.OK(c, logs, "success")
}

// DeleteSyslogByID 根据ID删除系统日志条目
func (sh *SyslogHandler) DeleteSyslogByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid syslog ID")
		return
	}

	syslogID := uint(id)
	err = sh.SyslogService.DeleteSyslogByID(syslogID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete syslog")
		return
	}

	response.OK(c, gin.H{"message": "Syslog deleted successfully"}, "success")
}
