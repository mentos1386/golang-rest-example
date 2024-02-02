package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mentos1386/golang-rest-example/pkg/api"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"

	"github.com/ogen-go/ogen/middleware"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logging(logger *zap.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		logger := logger.With(
			zap.String("operation", req.OperationName),
			zap.String("operationId", req.OperationID),
		)
		logger.Info("Handling request")
		resp, err := next(req)
		if err != nil {
			logger.Error("Fail", zap.Error(err))
		} else {
			var fields []zapcore.Field
			// Some response types may have a status code.
			// ogen provides a getter for it.
			//
			// You can write your own interface to match any response type.
			if tresp, ok := resp.Type.(interface{ GetStatusCode() int }); ok {
				fields = []zapcore.Field{
					zap.Int("status_code", tresp.GetStatusCode()),
				}
			}
			logger.Info("Success", fields...)
		}
		return resp, err
	}
}

func main() {
	service := api.NewApiService()
	logger, _ := zap.NewDevelopment()

	srv, err := openapi.NewServer(service, openapi.WithMiddleware(Logging(logger)))
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%d", service.Config.Port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	logger.Info("Starting server", zap.String("address", address))
	if err := http.ListenAndServe(address, c.Handler(srv)); err != nil {
		log.Fatal(err)
	}
}
