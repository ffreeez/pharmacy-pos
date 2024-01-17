package syslogrepo

import (
	syslogmodel "pharmacy-pos/pkg/db/models/syslog"
	logger "pharmacy-pos/pkg/util/logger"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// CreateSysLog 创建新的系统日志记录
func CreateSysLog(db *gorm.DB, sysLog *syslogmodel.Syslog) error {
	result := db.Create(sysLog)
	if result.Error != nil {
		logs.Errorf("创建新系统日志记录失败, error: %v", result.Error)
		return result.Error
	}
	logs.Infof("创建新系统日志记录成功")
	return nil
}

// GetAllSysLogs 获取所有系统日志的信息
func GetAllSysLogs(db *gorm.DB) ([]syslogmodel.Syslog, error) {
	var sysLogs []syslogmodel.Syslog
	result := db.Find(&sysLogs)
	if result.Error != nil {
		logs.Errorf("获取所有系统日志信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有系统日志信息成功")
	return sysLogs, nil
}

// DeleteSysLogByID 根据ID删除系统日志记录
func DeleteSysLogByID(db *gorm.DB, id uint) error {
	result := db.Delete(&syslogmodel.Syslog{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除系统日志记录失败, ID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除系统日志记录成功, ID: %d", id)
	return nil
}
