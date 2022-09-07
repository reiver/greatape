package server

import (
	"config"
	. "contracts"
	"server/authorize"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/html"
)

type httpServer struct {
	framework *fiber.App
	storage   IStorage
	logger    ILogger
	cache     ICache
}

func New() IServer {
	framework := fiber.
		New(fiber.Config{
			DisableStartupMessage: true,
			Views:                 html.New("./views", ".html"),
			BodyLimit:             config.BodyLimit(),
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}

				ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
				return ctx.Status(code).SendString(err.Error())
			},
		})

	framework.
		Use(
			cors.New(),
			recover.New(),
			helmet.New(),
			// csrf.New(),
			limiter.New(limiter.Config{
				Max:               20,
				Expiration:        30 * time.Second,
				LimiterMiddleware: limiter.SlidingWindow{},
			}),
			logger.New(logger.Config{
				Next:         nil,
				Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${body}\n",
				TimeFormat:   "15:04:05",
				TimeZone:     "Local",
				TimeInterval: 500 * time.Millisecond,
			}),
		)

	framework.Static("/media", config.UPLOAD_PATH)
	framework.Group("/api/v1/profile").Use(authorize.New())
	// framework.Get("/u/:name/inbox").Use(authorize.New())
	// framework.Post("/u/:name/outbox").Use(authorize.New())

	return &httpServer{
		framework: framework,
	}
}

func (server *httpServer) SetStorage(storage IStorage) {
	server.storage = storage
}

func (server *httpServer) SetLogger(logger ILogger) {
	server.logger = logger
}

func (server *httpServer) SetCache(cache ICache) {
	server.cache = cache
}

func (server *httpServer) Bind(routes ...IRoute) {
	for _, route := range routes {
		func(route IRoute) {
			switch route.Method() {
			case HttpGet:
				server.framework.Get(route.Path(), func(underlyingContext *fiber.Ctx) error {
					return route.Handler()(newContext(underlyingContext))
				})
			case HttpPost:
				server.framework.Post(route.Path(), func(underlyingContext *fiber.Ctx) error {
					return route.Handler()(newContext(underlyingContext))
				})
			case HttpPut:
				server.framework.Put(route.Path(), func(underlyingContext *fiber.Ctx) error {
					return route.Handler()(newContext(underlyingContext))
				})
			case HttpDelete:
				server.framework.Delete(route.Path(), func(underlyingContext *fiber.Ctx) error {
					return route.Handler()(newContext(underlyingContext))
				})
			default:
				panic("unsupported_method")
			}
		}(route)
	}
}

func (server *httpServer) Listen(address string) {
	server.framework.Listen(address)
}
