package main

import (
	"github.com/Unknwon/goconfig"
)

const (
	CONFIG_FILE_PATH                 = "./conf/controller.conf"
	DEFAULT_TCP_BIND                 = ":10050"
	DEFAULT_MYSQL_ADDR               = "127.0.0.1:3306"
	DEFAULT_MYSQL_USER               = "root"
	DEFAULT_MYSQL_DBNAME             = "owl"
	DEFAULT_MYSQL_PASSWORD           = ""
	DEFAULT_MAX_CONN                 = 20
	DEFAULT_MAX_IDLE_CONN            = 5
	DEFAULT_LOG_FILE                 = "./logs/controller.log"
	DEFAULT_LOG_EXPIRE_DAYS          = 7
	DEFAULT_LOG_LEVEL                = 3
	DEFAULT_MAX_PACKET_SIZE          = 4096
	DEFAULT_LOAD_STRATEGIES_INTERVAL = 300 //seconds
	DEFAULT_INSPECTOR_INTERVAL       = 6   //minutes
	DEFAULT_TASK_POOL_SIZE           = 4096
	DEFAULT_RESULT_POOL_SIZE         = 4096
	DEFAULT_HTTP_SERVER              = ":10051"
	DEFAULT_TASK_SIZE                = 10
	DEFAULT_WORKER_COUNT             = 5
	DEFAULT_SEND_MAIL_SCRIPT         = "./scripts/send_mail.py"
	DEFAULT_SEND_SMS_SCRIPT          = "./scripts/send_sms.py"
	DEFAULT_SEND_WECHAT_SCRIPT       = "./scripts/send_wechat.py"
)

var GlobalConfig *Config

type Config struct {
	//MYSQL CONFIG
	MYSQL_ADDR          string //mysql ip地址
	MYSQL_USER          string //mysql 登陆用户名
	MYSQL_PASSWORD      string //mysql 登陆密码
	MYSQL_DBNAME        string //mysql 数据库名称
	MYSQL_MAX_IDLE_CONN int    //mysql 最大空闲连接数
	MYSQL_MAX_CONN      int    //mysql 最大连接数
	//SERVER CONFIG
	TCP_BIND string //tcp监听地址和端口
	//LOG CONFIG
	LOG_FILE        string //日志保存目录
	LOG_LEVEL       int    //日志级别
	LOG_EXPIRE_DAYS int    //日志保留天数

	MAX_PACKET_SIZE int //tcp 消息报最大size

	LOAD_STRATEGIES_INTERVAL int //获取策略时间间隔 单位秒

	INSPECTOR_INTERVAL int //inspector 查询tsdb的时间间隔 单位分钟

	TASK_POOL_SIZE int //任务池的缓冲大小

	RESULT_POOL_SIZE int //结果池的缓冲大小

	HTTP_SERVER string //Http服务的地址

	TASK_SIZE int //单次获取任务数

	WORKER_COUNT int //处理结果池的协程数

	SEND_MAIL_SCRIPT string

	SEND_SMS_SCRIPT string

	SEND_WECHAT_SCRIPT string
}

func InitGlobalConfig() error {
	cfg, err := goconfig.LoadConfigFile(CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	GlobalConfig = &Config{
		TCP_BIND:                 cfg.MustValue(goconfig.DEFAULT_SECTION, "tcp_bind", DEFAULT_TCP_BIND),
		MYSQL_ADDR:               cfg.MustValue(goconfig.DEFAULT_SECTION, "mysql_addr", DEFAULT_MYSQL_ADDR),
		MYSQL_USER:               cfg.MustValue(goconfig.DEFAULT_SECTION, "mysql_user", DEFAULT_MYSQL_USER),
		MYSQL_DBNAME:             cfg.MustValue(goconfig.DEFAULT_SECTION, "mysql_dbname", DEFAULT_MYSQL_DBNAME),
		MYSQL_PASSWORD:           cfg.MustValue(goconfig.DEFAULT_SECTION, "mysql_password", DEFAULT_MYSQL_PASSWORD),
		MYSQL_MAX_CONN:           cfg.MustInt(goconfig.DEFAULT_SECTION, "mysql_max_conn", DEFAULT_MAX_CONN),
		MYSQL_MAX_IDLE_CONN:      cfg.MustInt(goconfig.DEFAULT_SECTION, "mysql_max_idle_conn", DEFAULT_MAX_IDLE_CONN),
		LOG_FILE:                 cfg.MustValue(goconfig.DEFAULT_SECTION, "log_file", "./logs/controller.log", DEFAULT_LOG_FILE),
		LOG_EXPIRE_DAYS:          cfg.MustInt(goconfig.DEFAULT_SECTION, "log_expire_days", DEFAULT_LOG_EXPIRE_DAYS),
		LOG_LEVEL:                cfg.MustInt(goconfig.DEFAULT_SECTION, "log_level", DEFAULT_LOG_LEVEL),
		MAX_PACKET_SIZE:          cfg.MustInt(goconfig.DEFAULT_SECTION, "max_packt_size", DEFAULT_MAX_PACKET_SIZE),
		LOAD_STRATEGIES_INTERVAL: cfg.MustInt(goconfig.DEFAULT_SECTION, "load_strategies_interval", DEFAULT_LOAD_STRATEGIES_INTERVAL),
		INSPECTOR_INTERVAL:       cfg.MustInt(goconfig.DEFAULT_SECTION, "inspector_interval", DEFAULT_INSPECTOR_INTERVAL),
		TASK_POOL_SIZE:           cfg.MustInt(goconfig.DEFAULT_SECTION, "task_pool_size", DEFAULT_TASK_POOL_SIZE),
		RESULT_POOL_SIZE:         cfg.MustInt(goconfig.DEFAULT_SECTION, "result_pool_size", DEFAULT_RESULT_POOL_SIZE),
		HTTP_SERVER:              cfg.MustValue(goconfig.DEFAULT_SECTION, "http_server", DEFAULT_HTTP_SERVER),
		TASK_SIZE:                cfg.MustInt(goconfig.DEFAULT_SECTION, "task_size", DEFAULT_TASK_SIZE),
		WORKER_COUNT:             cfg.MustInt(goconfig.DEFAULT_SECTION, "worker_count", DEFAULT_WORKER_COUNT),
		SEND_MAIL_SCRIPT:         cfg.MustValue(goconfig.DEFAULT_SECTION, "send_mail_script", DEFAULT_SEND_MAIL_SCRIPT),
		SEND_SMS_SCRIPT:          cfg.MustValue(goconfig.DEFAULT_SECTION, "send_sms_script", DEFAULT_SEND_SMS_SCRIPT),
		SEND_WECHAT_SCRIPT:       cfg.MustValue(goconfig.DEFAULT_SECTION, "send_wechat_script", DEFAULT_SEND_WECHAT_SCRIPT),
	}
	return nil
}
