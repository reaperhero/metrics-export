package export

import "github.com/prometheus/client_golang/prometheus"

// Counter 一个累加指标数据，这个值初始值为0，随着时间只会逐渐的增加，比如程序完成的总任务数量，运行错误发生的总次数。常见的还有交换机中snmp采集的数据流量也属于该类型，代表了持续增加的数据包或者传输字节累加值
// 只有 Inc() 和 Add() 两个函数

var (
	totalRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "The total number of handled HTTP requests.",
		ConstLabels: map[string]string{"label_key": "label_value"},
	})
)
