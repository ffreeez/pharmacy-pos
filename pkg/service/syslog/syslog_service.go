package syslogservice

import (
	syslogmodel "pharmacy-pos/pkg/db/models/syslog"
	syslogrepo "pharmacy-pos/pkg/db/repository/syslog"

	"gorm.io/gorm"
)

// SyslogService 提供系统日志相关的服务
type SyslogService struct {
	DB *gorm.DB
}

// NewSyslogService 创建一个新的 SyslogService 实例
func NewSyslogService(db *gorm.DB) *SyslogService {
	return &SyslogService{DB: db}
}

// CreateSyslog 创建新的系统日志条目
func (ss *SyslogService) CreateSyslog(log *syslogmodel.Syslog) error {
	return syslogrepo.CreateSysLog(ss.DB, log)
}

// GetAllSyslogs 获取所有系统日志条目
func (ss *SyslogService) GetAllSyslogs() ([]syslogmodel.Syslog, error) {
	return syslogrepo.GetAllSysLogs(ss.DB)
}

// DeleteSyslogByID 根据ID删除系统日志条目
func (ss *SyslogService) DeleteSyslogByID(id uint) error {
	return syslogrepo.DeleteSysLogByID(ss.DB, id)
}
