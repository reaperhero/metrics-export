package export

import "github.com/prometheus/client_golang/prometheus"

// gauge 代表了采集的一个单一数据，这个数据可以增加也可以减少，比如CPU使用情况，内存使用量，硬盘当前的空间容量等等
// gauge 指标对象暴露了 Set()、Inc()、Dec()、Add() 和 Sub() 这些函数来更改指标值

var (
	TempGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	TempGaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// 两个标签名称，通过它们来分割指标。
		[]string{"house", "room"},
	)
	// TempGaugeVec.WithLabelValues("ydzs", "living-room").Set(27) == home_temperature_celsius{house="cnych",room="bedroom"} 25.3
)
