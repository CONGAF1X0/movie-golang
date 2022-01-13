package global

import (
	"TicketSales/pkg/logger"
	"TicketSales/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	Logger          *logger.Logger
)
