package export

import "github.com/prometheus/client_golang/prometheus"

// Histogram和Summary使用的频率较少，两种都是基于采样的方式。
// 这两个类型对于某一些业务需求可能比较常见，比如查询单位时间内：总的响应时间低于300ms的占比，或者查询95%用户查询的门限值对应的响应时间是多少
// 使用Histogram和Summary指标的时候同时会产生多组数据，
// 		_count代表了采样的总数
//		_sum则代表采样值的和
//		_bucket则代表了落入此范围的数据

var (
	requestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "A histogram of the HTTP request durations in seconds.",
		// Bucket 配置：第一个 bucket 包括所有在 0.05s 内完成的请求，最后一个包括所有在10s内完成的请求。
		Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
	// 如果你刚刚处理了一个 HTTP 请求，花了 0.42 秒 = requestDurations.Observe(0.42)
)
