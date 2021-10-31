package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port   = 8080
	logger *zap.Logger
)

func init() {
	logger = newZapLogger("log-app", zapcore.DebugLevel)
}

func main() {

	if portEnv, ok := os.LookupEnv("FBT_PORT"); ok {
		if portInt, err := strconv.Atoi(portEnv); err == nil {
			port = portInt
		}
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		indexPostHandler(w, r)
		return
	}

	io.Copy(io.Discard, r.Body)
	r.Body.Close()
	w.WriteHeader(http.StatusAccepted)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("read request: %v", err)))
		return
	}

	defer r.Body.Close()
	logger.Info(string(b))
	w.WriteHeader(http.StatusCreated)
}

func newZapLogger(appName string, level zapcore.Level) *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "ts_app"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.Level.SetLevel(level)

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("initialise zap logger: %v", err)
	}
	return logger.With(zap.String("app", appName))
}
