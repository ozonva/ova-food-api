package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

type logger struct {
	Filename string
	zerolog.Logger
}

var GlobalLogger *logger
var once sync.Once

func InitLogger(logfile string) {
	once.Do(func() {
		filename := append([]rune(time.Now().Format("2006.01.02_")), []rune(logfile)...)
		f, err := os.OpenFile(string(filename), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		GlobalLogger = &logger{
			logfile,
			zerolog.New(f).With().Logger(),
		}
		GlobalLogger.Logger.Info().Msg("Logger init ok")
	})
}
