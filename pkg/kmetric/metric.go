package kmetric

var (
	// TypeHTTP ...
	TypeHTTP = "http"
	// TypeGRPC ...
	TypeGRPC = "grpc"
	// TypeRedis ...
	TypeRedis = "redis"
	TypeGorm  = "gorm"
	TypeMongo = "mongo"
	// TypeRocketMQ ...
	TypeRocketMQ = "rocketmq"

	// TypeMySQL ...
	TypeMySQL = "mysql"

	// CodeJob
	CodeJobSuccess = "ok"
	// CodeJobFail ...
	CodeJobFail = "fail"
	// CodeCache
	CodeCacheMiss = "miss"
	// CodeCacheHit ...
	CodeCacheHit = "hit"

	// Namespace
	DefaultNamespace = "kepler"
)

var (
	// ServerHandleCounter ...
	ServerHandleCounter = NewCounterVec(&CounterVecoption{
		Namespace: DefaultNamespace,
		Name:      "server_handle_total",
		Labels:    []string{"type", "name", "method", "code"},
	})

	// ServerHandleHistogram ...
	ServerHandleHistogram = NewHistogramVec(&HistogramVecoption{
		Namespace: DefaultNamespace,
		Name:      "server_handle_seconds",
		Labels:    []string{"type", "name", "method"},
	})

	// ClientHandleCounter ...
	ClientHandleCounter = NewCounterVec(&CounterVecoption{
		Namespace: DefaultNamespace,
		Name:      "client_handle_total",
		Labels:    []string{"type", "name", "method", "code"},
	})

	// ClientHandleHistogram ...
	ClientHandleHistogram = NewHistogramVec(&HistogramVecoption{
		Namespace: DefaultNamespace,
		Name:      "client_handle_seconds",
		Labels:    []string{"type", "name", "method"},
	})

	//GormHandleCounter = NewCounterVec(&CounterVecoption{
	//	Namespace: DefaultNamespace,
	//	Name:      "gorm_handle_total",
	//	Labels:    []string{"type", "name", "method", "code"},
	//})
	//
	//// ClientHandleHistogram ...
	//GormHandleHistogram = NewHistogramVec(&HistogramVecoption{
	//	Namespace: DefaultNamespace,
	//	Name:      "gorm_handle_seconds",
	//	Labels:    []string{"type", "name", "method"},
	//})

	// // JobHandleCounter ...
	// JobHandleCounter = NewCounterVec(&CounterVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "job_handle_total",
	// 	Labels:    []string{"type", "name", "code"},
	// })

	// // JobHandleHistogram ...
	// JobHandleHistogram = NewHistogramVec(&HistogramVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "job_handle_seconds",
	// 	Labels:    []string{"type", "name"},
	// })

	// LibHandleHistogram = NewHistogramVec(&HistogramVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "lib_handle_seconds",
	// 	Labels:    []string{"type", "method", "address"},
	// })
	// // LibHandleCounter ...
	// LibHandleCounter = NewCounterVec(&CounterVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "lib_handle_total",
	// 	Labels:    []string{"type", "method", "address", "code"},
	// })

	// // CacheHandleCounter ...
	// CacheHandleCounter = NewCounterVec(&CounterVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "cache_handle_total",
	// 	Labels:    []string{"type", "name", "action", "code"},
	// })

	// // CacheHandleHistogram ...
	// CacheHandleHistogram = NewHistogramVec(&HistogramVecoption{
	// 	Namespace: DefaultNamespace,
	// 	Name:      "cache_handle_seconds",
	// 	Labels:    []string{"type", "name", "action"},
	// })
)
