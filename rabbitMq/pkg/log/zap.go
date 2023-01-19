package log

import (
	"github.com/jonny-mark/package-demo/rabbitMq/pkg/utils"
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

const (
	// json输出
	WriterJson = "json"
)

const (
	// RotateTimeDaily 按天切割
	RotateTimeDaily = "daily"
	// RotateTimeHourly 按小时切割
	RotateTimeHourly = "hourly"
)

// zapLogger logger struct
type zapLogger struct {
	sugarLogger *zap.SugaredLogger
}

// For mapping config logger to system logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// newZapLogger new zap log
func newZapLogger(cfg *Config) Logger {
	encoderCfg := getEncoderConfig(cfg)

	encoder := getEncoder(cfg, encoderCfg)

	var cores []zapcore.Core
	var options []zap.Option

	// init option
	option := zap.Fields(
		zap.String("ip", utils.GetLocalIP()),
		zap.String("app_id", cfg.Name),
		zap.String("instance_id", utils.GetHostname()),
	)
	options = append(options, option)

	// info
	cores = append(cores, getInfoCore(encoder, cfg))

	// warning
	cores = append(cores, getWarnCore(encoder, cfg))

	// error
	cores = append(cores, getErrorCore(encoder, cfg))

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	if cfg.Stacktrace {
		options = append(options, zap.AddCaller(), zap.AddStacktrace(getLoggerLevel(cfg)))
	}

	// 构造日志
	return &zapLogger{zap.New(combinedCore, options...).Sugar()}
}

func getEncoderConfig(cfg *Config) zapcore.EncoderConfig {
	var encoderCfg zapcore.EncoderConfig

	if cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	return encoderCfg
}

func getEncoder(cfg *Config, encoderCfg zapcore.EncoderConfig) zapcore.Encoder {
	var encoder zapcore.Encoder
	if cfg.Format == WriterJson {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}
	return encoder
}

func getLoggerLevel(cfg *Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func getInfoCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	infoWrite := getLogWriterWithTime(cfg, cfg.LoggerInfoFile)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
}

func getWarnCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	warnWrite := getLogWriterWithTime(cfg, cfg.LoggerWarnFile)
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(warnWrite), warnLevel)
}

func getErrorCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	errorFilename := cfg.LoggerErrorFile
	errorWrite := getLogWriterWithTime(cfg, errorFilename)
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), errorLevel)
}

// getLogWriterWithTime 按时间(小时)进行切割
func getLogWriterWithTime(cfg *Config, filename string) io.Writer {
	var logFullPath string
	var MaxAge time.Duration
	rotationPolicy := cfg.LogRollingPolicy
	if cfg.MaxAge > 0 {
		MaxAge = time.Duration(cfg.MaxAge) * 24 * time.Hour
	} else {
		MaxAge = 7 * 24 * time.Hour
	}
	//backupCount := cfg.LogBackupCount
	// 默认
	var rotateDuration time.Duration
	if rotationPolicy == RotateTimeHourly {
		rotateDuration = time.Hour
		logFullPath = cfg.Director + "/" + "%Y-%m-%d-%H" + "/" + filename
	} else if rotationPolicy == RotateTimeDaily {
		rotateDuration = time.Hour * 24
		logFullPath = cfg.Director + "/" + "%Y-%m-%d" + "/" + filename
	}
	hook, err := rotatelogs.New(
		logFullPath, // 时间格式使用shell的date时间格式
		//rotatelogs.WithLinkName(logFullPath), // 生成软链，指向最新日志文件
		//rotatelogs.WithRotationCount(backupCount),   // 文件最大保存份数
		rotatelogs.WithMaxAge(MaxAge),               // 文件最大保存时间（天）
		rotatelogs.WithRotationTime(rotateDuration), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// Info logger
func (l *zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Warn logger
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// Error logger
func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Panic logger
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Fatal logger
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Infof logger
func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

// Warnf logger
func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

// Errorf logger
func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

// Fatalf logger
func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

// Panicf logger
func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	type Fields zapcore.Field

	newLogger := l.sugarLogger.With(f...)
	return &zapLogger{newLogger}
}
