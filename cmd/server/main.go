package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	//"github.com/omrikiei/protoc-gen-mcp/examples/complex"
	//"github.com/omrikiei/protoc-gen-mcp/examples/helloworld"
)

var (
	// Metrics
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

type greeterServer struct {
	helloworld.GreeterServer
	logger *zap.Logger
}

func (s *greeterServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	s.logger.Info("Handling SayHello request",
		zap.String("name", req.Name),
		zap.String("language", req.Language),
		zap.Bool("formal", req.Formal),
	)

	message := fmt.Sprintf("Hello, %s!", req.Name)
	if req.Formal {
		message = fmt.Sprintf("Greetings, %s!", req.Name)
	}
	if req.Language != "" {
		message = fmt.Sprintf("%s (in %s)", message, req.Language)
	}
	return &helloworld.HelloReply{Message: message}, nil
}

func (s *greeterServer) SayHelloWithQuery(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return s.SayHello(ctx, req)
}

type complexServer struct {
	complex.ComplexServiceServer
	logger *zap.Logger
}

func (s *complexServer) GetUserProfile(ctx context.Context, req *complex.UserProfileRequest) (*complex.UserProfileResponse, error) {
	s.logger.Info("Handling GetUserProfile request",
		zap.String("user_id", req.UserId),
		zap.String("profile_id", req.ProfileId),
		zap.Bool("include_details", req.IncludeDetails),
		zap.Int32("max_depth", req.MaxDepth),
	)

	return &complex.UserProfileResponse{
		UserId:    req.UserId,
		ProfileId: req.ProfileId,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Metadata: map[string]string{
			"department": "Engineering",
			"role":       "Developer",
		},
	}, nil
}

func (s *complexServer) GetNestedResource(ctx context.Context, req *complex.NestedRequest) (*complex.NestedResponse, error) {
	s.logger.Info("Handling GetNestedResource request",
		zap.String("org_id", req.OrgId),
		zap.String("team_id", req.TeamId),
		zap.String("member_id", req.MemberId),
		zap.Bool("recursive", req.Recursive),
	)

	return &complex.NestedResponse{
		OrgId:       req.OrgId,
		TeamId:      req.TeamId,
		MemberId:    req.MemberId,
		Name:        "Jane Smith",
		Role:        "Team Lead",
		Permissions: []string{"read", "write", "admin"},
	}, nil
}

func (s *complexServer) SearchResources(ctx context.Context, req *complex.SearchRequest) (*complex.SearchResponse, error) {
	s.logger.Info("Handling SearchResources request",
		zap.String("query", req.Query),
		zap.Int32("page", req.Page),
		zap.Int32("page_size", req.PageSize),
		zap.Strings("filters", req.Filters),
		zap.Bool("include_deleted", req.IncludeDeleted),
		zap.String("sort_by", req.SortBy),
		zap.Bool("sort_desc", req.SortDesc),
	)

	now := time.Now().Unix()
	resources := []*complex.Resource{
		{
			Id:        "1",
			Name:      "Resource 1",
			Type:      "type1",
			Metadata:  map[string]string{"key1": "value1"},
			Deleted:   false,
			CreatedAt: now - 3600,
			UpdatedAt: now,
		},
		{
			Id:        "2",
			Name:      "Resource 2",
			Type:      "type2",
			Metadata:  map[string]string{"key2": "value2"},
			Deleted:   false,
			CreatedAt: now - 7200,
			UpdatedAt: now,
		},
	}

	return &complex.SearchResponse{
		Resources:  resources,
		TotalCount: int32(len(resources)),
		Page:       req.Page,
		PageSize:   req.PageSize,
	}, nil
}

// Health check handler
type healthHandler struct {
	isReady *atomic.Value
}

func newHealthHandler() *healthHandler {
	var isReady atomic.Value
	isReady.Store(false)
	return &healthHandler{
		isReady: &isReady,
	}
}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
		return
	}

	if r.URL.Path == "/readyz" {
		if h.isReady.Load().(bool) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "ok")
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "not ready")
		}
		return
	}

	http.NotFound(w, r)
}

func main() {
	// Parse command line flags
	port := flag.Int("port", 8080, "Port to listen on")
	metricsPort := flag.Int("metrics-port", 9090, "Port to listen on for metrics")
	flag.Parse()

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create server instances
	greeter := &greeterServer{
		logger: logger,
	}
	complex := &complexServer{
		logger: logger,
	}

	// Create MCP servers
	greeterServer := helloworld.NewGreeterMCPServer(greeter)
	complexServer := complex.NewComplexServiceMCPServer(complex)

	// Create health handler
	health := newHealthHandler()

	// Create main HTTP server
	mainMux := http.NewServeMux()
	mainMux.Handle("/v1/hello/", greeterServer)
	mainMux.Handle("/v1/users/", complexServer)
	mainMux.Handle("/v1/orgs/", complexServer)
	mainMux.Handle("/v1/resources", complexServer)
	mainMux.Handle("/healthz", health)
	mainMux.Handle("/readyz", health)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mainMux,
	}

	// Create metrics server
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())
	metricsServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *metricsPort),
		Handler: metricsMux,
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 2)

	// Start the main server
	go func() {
		logger.Info("Starting main server", zap.Int("port", *port))
		serverErrors <- httpServer.ListenAndServe()
	}()

	// Start the metrics server
	go func() {
		logger.Info("Starting metrics server", zap.Int("port", *metricsPort))
		serverErrors <- metricsServer.ListenAndServe()
	}()

	// Mark the service as ready
	health.isReady.Store(true)

	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		logger.Error("Server error", zap.Error(err))
		os.Exit(1)

	case sig := <-shutdown:
		logger.Info("Shutdown signal received", zap.String("signal", sig.String()))
		// Mark the service as not ready
		health.isReady.Store(false)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Shutdown both servers
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("Error shutting down main server", zap.Error(err))
		}
		if err := metricsServer.Shutdown(ctx); err != nil {
			logger.Error("Error shutting down metrics server", zap.Error(err))
		}
	}
}
