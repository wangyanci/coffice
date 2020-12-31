package logs

import (
	"encoding/json"

	"github.com/wangyanci/coffice/logs/config"
	"github.com/wangyanci/coffice/setting"

	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

// LevelEmergency = iota     // 紧急级别0
// LevelAlert                // 报警级别1
// LevelCritical             // 严重错误级别2
// LevelError                // 错误级别3
// LevelWarning              // 警告级别4
// LevelNotice               // 注意级别5
// LevelInformational        // 报告级别6
// LevelDebug                // 除错级别7
func init() {
	logConfig := config.LogConfig{
		ConsoleConfig: config.ConsoleConfig{
			Colorful: true,
		},
		FileConfig: config.FileConfig{
			Filename: setting.LogPath,
			Separate: []string{"emergency", "alert", "critical",
				"error", "warning", "notice", "info", "debug"},
		},
	}

	logs.Async()
	logs.Async(1e3)
	jsonConfig, _ := json.Marshal(logConfig)
	log := logs.NewLogger(setting.LogCache)               // 创建一个日志记录器，参数为缓冲区的大小
	log.SetLogger(setting.LogAdapter, string(jsonConfig)) // 设置日志记录方式：控制台记录
	log.SetLevel(setting.LogLevel)                        // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	log.EnableFuncCallDepth(true)                         // 输出log时能显示输出文件名和行号（非必须）

	Logger = log
}
