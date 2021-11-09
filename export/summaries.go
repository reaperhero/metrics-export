package export

import "github.com/prometheus/client_golang/prometheus"

// 指定要跟踪的 quantiles 分位数值, 比如我们想要跟踪 HTTP 请求延迟的第 50、90 和 99 个百分位数

var (
	requestSummaryDurations = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds",
		Help: "A summary of the HTTP request durations in seconds.",
		Objectives: map[float64]float64{
			0.5:  0.05,  // 第50个百分位数，最大绝对误差为0.05。
			0.9:  0.01,  // 第90个百分位数，最大绝对误差为0.01。
			0.99: 0.001, // 第90个百分位数，最大绝对误差为0.001。
		},
	})
	// 跟踪持续时间的方式和直方图是完全一样的，使用一个 Observe() 函数即可：requestDurations.Observe(0.42)
)
