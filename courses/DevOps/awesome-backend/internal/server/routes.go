package server

import (
	"fmt"
	"net/http"
	"time"

	"git.itmo.su/go-modules/swaga/v4"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) initRoutes() {
	// healthcheck endpoint
	app := s.app.Group("/awesome-backend")
	app.Get("/healthcheck", s.HealthCheck)

	v1 := app.Group("/api/v1")
	swaga.Setup(v1, swaga.Config{
		URL:       "/awesome-backend/api/v1",
		MountPath: "/docs",
		Title:     "awesome-backend API",
	}, swaga.WithSpecFile("./docs/swagger.json"),
	)
	s.RmqRoutes(v1)
	s.RedisRoutes(v1)
}

func (s *Server) HealthCheck(c *fiber.Ctx) error {
	fmt.Println(fmt.Sprintf("{\n  \"_index\": \"dev-index\",\n  \"_id\": \"c_5yBp0BhBpvtXBTJjF5\",\n  \"_version\": 1,\n  \"_score\": null,\n  \"_source\": {\n    \"@timestamp\": \"2026-03-19T14:13:51.252005304Z\",\n    \"@version\": \"1\",\n    \"message\": \"generated log #91 - test\",\n    \"logger_name\": \"stc.example.ordermanagerservice.controller.DebugLoggingController\",\n    \"thread_name\": \"http-nio-8080-exec-2\",\n    \"level\": \"INFO\",\n    \"level_value\": 20000,\n    \"hostname\": \"fluentd\",\n    \"requestId\": \"a084cc93-7928-4e59-a638-8fff4b575167\",\n    \"module\": \"app\",\n    \"environment\": \"dev\"\n  },\n  \"fields\": {\n    \"@timestamp\": [\n      " + time.Now().String() + "\n    ]\n  },\n  \"sort\": [\n    1773929631252\n  ]\n}"))
	if err := s.svc.Ping(c.Context()); err != nil {
		return c.SendStatus(http.StatusServiceUnavailable)
	}
	return c.SendStatus(http.StatusOK)
}
