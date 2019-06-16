// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/LGUG2Z/call-me/handlers"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/LGUG2Z/call-me/restapi/operations"

	"github.com/LGUG2Z/call-me/models"
)

//go:generate swagger generate server --target ../../call-me --name CallMe --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.CallMeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CallMeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	api.GetMaybeHandler = operations.GetMaybeHandlerFunc(handlers.GetMaybe)
	api.PostMaybeHandler = operations.PostMaybeHandlerFunc(handlers.PostMaybe)
	api.DeleteMaybeHandler = operations.DeleteMaybeHandlerFunc(handlers.DeleteMaybe)

	ApiKey := os.Getenv("API_KEY")
	if len(ApiKey) == 0 {
		log.Fatal("API_KEY must be set")
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-KEY" header is set
	api.APIKeyHeaderAuth = func(token string) (*models.Principal, error) {
		if token == ApiKey {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.DeleteMaybeHandler == nil {
		api.DeleteMaybeHandler = operations.DeleteMaybeHandlerFunc(func(params operations.DeleteMaybeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation .DeleteMaybe has not yet been implemented")
		})
	}
	if api.GetMaybeHandler == nil {
		api.GetMaybeHandler = operations.GetMaybeHandlerFunc(func(params operations.GetMaybeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation .GetMaybe has not yet been implemented")
		})
	}
	if api.PostMaybeHandler == nil {
		api.PostMaybeHandler = operations.PostMaybeHandlerFunc(func(params operations.PostMaybeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation .PostMaybe has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
