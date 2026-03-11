package logger_test

import (
	"github.com/munhdalai/go_helper/logger"
	"go.uber.org/zap"
)

func ExampleInitLogger() {
	// Development горимд logger эхлүүлэх
	err := logger.InitLogger(logger.DebugLevel, "/tmp/app.log", false)
	if err != nil {
		panic(err)
	}

	// Production горимд logger эхлүүлэх (stacktrace-гүй)
	err = logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)
	if err != nil {
		panic(err)
	}
}

func ExampleInfo() {
	logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)

	logger.Info("Хэрэглэгч нэвтэрлээ",
		zap.String("username", "bat"),
		zap.Int("user_id", 123),
	)
}

func ExampleDebug() {
	logger.InitLogger(logger.DebugLevel, "/tmp/app.log", false)

	logger.Debug("Query ажиллалаа",
		zap.String("sql", "SELECT * FROM users"),
		zap.Duration("duration", 0),
	)
}

func ExampleError() {
	logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)

	logger.Error("Database холболт амжилтгүй",
		zap.String("host", "localhost"),
		zap.Int("port", 5432),
	)
}

func ExampleWarn() {
	logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)

	logger.Warn("Хүсэлтийн хугацаа удаашралтай",
		zap.String("endpoint", "/api/users"),
		zap.Float64("duration_ms", 2500.5),
	)
}

func ExampleNewSugar() {
	logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)

	sugar := logger.NewSugar("user-service")
	sugar.Infow("Хэрэглэгч үүсгэлээ",
		"name", "Бат",
		"email", "bat@example.com",
	)
}

func ExampleGetLogger() {
	logger.InitLogger(logger.InfoLevel, "/tmp/app.log", true)

	l := logger.GetLogger()
	l.Info("Global logger ашиглаж байна")
}
