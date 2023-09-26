package myserver

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type logFields struct {
	RequestID  string
	Method     string
	RemoteIP   string
	Host       string
	Path       string
	StatusCode int
	Latency    float64
}

func (lf *logFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		// Str("request_id", lf.RequestID).
		// Str("method", lf.Method).
		// Str("remote_ip", lf.RemoteIP).
		// Str("host", lf.Host).
		Str("path", fmt.Sprintf("[%s] %s", lf.Method, lf.Path)).
		Int("status_code", lf.StatusCode).
		Float64("latency", lf.Latency)
}

// Middleware requestid + logger + recover for request traceability
func ZMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		start := time.Now()
		var myerr error

		log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel)
		rid := strings.ReplaceAll(uuid.New().String(), "-", "")

		fields := &logFields{
			RequestID: rid,
			RemoteIP:  c.IP(),
			Method:    c.Method(),
			Host:      c.Hostname(),
			Path:      c.Path(),
		}

		defer func() {
			rvr := recover()
			fields.StatusCode = c.Context().Response.StatusCode()
			fields.Latency = time.Since(start).Seconds()

			if myerr != nil {
				fields.StatusCode = http.StatusInternalServerError
				c.Status(http.StatusInternalServerError)
				c.JSON(map[string]interface{}{
					"message": "internal server error",
					"result":  nil,
				})
			}
			switch {
			case rvr != nil:
				log.Error().EmbedObject(fields).Msg("panic recover")
			case fields.StatusCode >= 500:
				log.Error().EmbedObject(fields).Msg(myerr.Error())
			case fields.StatusCode >= 400:
				log.Info().EmbedObject(fields).Msg("client error")
			case fields.StatusCode >= 300:
				log.Warn().EmbedObject(fields).Msg("redirect")
			case fields.StatusCode >= 200:
				log.Info().EmbedObject(fields).Msg("success")
			case fields.StatusCode >= 100:
				log.Info().EmbedObject(fields).Msg("informative")
			default:
				log.Warn().EmbedObject(fields).Msg("unknown status")
			}
		}()

		myerr = c.Next()
		return nil
	}
}
