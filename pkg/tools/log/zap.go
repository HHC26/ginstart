package log

import (
	"ginstart/global"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.SugaredLogger {
	log := global.Conf.Zap

	fileName := log.Path
	level := getLoggerLevel(log.Level)

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	if log.MaxAge < 0 {
		log.MaxAge = 0
	}

	// 定义日志切割配置
	hook := &lumberjack.Logger{
		Filename:   fileName,       // 日志文件的位置
		MaxSize:    log.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: log.MaxBackups, // 保留旧文件的最大个数
		Compress:   log.Compress,   // 是否压缩 disabled by default
		MaxAge:     log.MaxAge,     // 单位: 天
	}

	// 判断是否控制台输出日志
	var syncer zapcore.WriteSyncer
	if log.Level == "debug" {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	} else {
		syncer = zapcore.AddSync(hook)
	}
	// 定义zap配置信息
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,          // 自定义时间格式
		EncodeLevel:    customLevelEncoder,         // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	// 判断是否json格式输出
	if log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, syncer, level)

	logger := zap.New(core)

	// 判断是否显示代码行号
	if log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger.Sugar()
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}
