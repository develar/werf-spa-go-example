package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "go.deanishe.net/env"
  "go.uber.org/zap"
)

func main() {
  serverPort := env.Get("SERVER_PORT", "8080")
  var logger *zap.Logger
  if env.GetBool("USE_DEV_LOGGING") {
    logger, _ = zap.NewProduction()
  } else {
    logger, _ = zap.NewDevelopment(zap.WithCaller(false))
  }
  defer logger.Sync()

  http.HandleFunc("/health-check", func(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(200)
  })
  http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
    _, _ = writer.Write([]byte("Hello World!"))
  })

  server := &http.Server{Addr: ":" + serverPort, Handler: nil}
  logger.Info("started", zap.String("address", server.Addr))

  go func() {
    err := server.ListenAndServe()
    if err != nil {
      log.Fatal(err)
    }
  }()

  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
  <-stop

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  err := server.Shutdown(ctx)
  if err != nil {
    log.Fatal(err)
  }
}