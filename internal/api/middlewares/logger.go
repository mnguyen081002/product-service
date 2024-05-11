package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (e *GinMiddleware) Logger(c *gin.Context) {

	start := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	context := *c
	context.Writer = blw
	body, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewReader(body))
	c.Next()
	path := context.Request.URL.Path
	raw := context.Request.URL.RawQuery

	end := time.Now()
	latency := end.Sub(start)

	clientIP := context.ClientIP()
	method := context.Request.Method
	statusCode := context.Writer.Status()

	l := e.logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel))
	logger := l.Info
	if statusCode >= 400 && statusCode < 500 {
		logger = l.Warn
	}
	if statusCode >= 500 {
		logger = l.Error
	}

	logger("Request",
		zap.String("Path", path),
		zap.String("Raw", raw),
		zap.String("Body", string(body)),
		zap.String("ClientIP", clientIP),
		zap.String("Method", method),
		zap.Int("StatusCode", statusCode),
		zap.Duration("Latency", latency),
		zap.String("Response", blw.body.String()),
	)

}
