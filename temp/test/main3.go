package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

type ClusterManager struct {
	sync.Mutex
	Zone              string
	metricMapCounters map[string]string
	metricMapGauges   map[string]string
}

//Simulate prepare the data
func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
	metrics map[string]float64,
) {
	metrics = map[string]float64{
		"oom_crashes_total": 42.00,
		"ram_usage":         6.023e23,
	}
	return
}

//通过NewClusterManager方法创建结构体及对应的指标信息，代码如下所示。
// NewClusterManager creates the two Descs OOMCountDesc and RAMUsageDesc. Note
// that the zone is set as a ConstLabel. (It's different in each instance of the
// ClusterManager, but constant over the lifetime of an instance.) Then there is
// a variable label "host", since we want to partition the collected metrics by
// host. Since all Descs created in this way are consistent across instances,
// with a guaranteed distinction by the "zone" label, we can register different
// ClusterManager instances with the same registry.
func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		Zone: zone,
		metricMapGauges: map[string]string{
			"ram_usage": "ram_usage",
		},
		metricMapCounters: map[string]string{
			"oom_crashes": "oom_crashes_total",
		},
	}
}

func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	// prometheus.NewDesc(prometheus.BuildFQName(namespace, "", metricName), docString, labels, nil)
	for _, v := range c.metricMapGauges {
		ch <- prometheus.NewDesc(prometheus.BuildFQName(c.Zone, "", v), v, nil, nil)
	}

	for _, v := range c.metricMapCounters {
		ch <- prometheus.NewDesc(prometheus.BuildFQName(c.Zone, "", v), v, nil, nil)
	}
}

//Collect方法是核心，它会抓取你需要的所有数据，根据需求对其进行分析，然后将指标发送回客户端库。
// 用于传递所有可能指标的定义描述符
// 可以在程序运行期间添加新的描述，收集新的指标信息
// 重复的描述符将被忽略。两个不同的Collector不要设置相同的描述符
func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
	c.Lock()
	defer c.Unlock()
	m := c.ReallyExpensiveAssessmentOfTheSystemState()
	for k, v := range m {
		t := prometheus.GaugeValue
		if c.metricMapCounters[k] != "" {
			t = prometheus.CounterValue
		}
		c.registerConstMetric(ch, k, v, t)
	}
}

// 用于传递所有可能指标的定义描述符给指标
func (c *ClusterManager) registerConstMetric(ch chan<- prometheus.Metric, metric string, val float64, valType prometheus.ValueType, labelValues ...string) {
	descr := prometheus.NewDesc(prometheus.BuildFQName(c.Zone, "", metric), metric, nil, nil)
	if m, err := prometheus.NewConstMetric(descr, valType, val, labelValues...); err == nil {
		ch <- m
	}
}
func main() {
	workerCA := NewClusterManager("xiaodian")
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerCA)
	//当promhttp.Handler()被执行时，所有metric被序列化输出。题外话，其实输出的格式既可以是plain text，也可以是protocol Buffers。
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8888", nil)
}
