package logger

import "github.com/rs/zerolog/log"

func Info(info string) {
	log.Info().Msgf(info)
}

func Error(err error) {
	log.Error().Msgf(err.Error())
}
