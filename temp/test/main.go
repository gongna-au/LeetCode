package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ClusterManager struct {
	Zone string
	// Contains many more fields not listed in this example.
}

//把每个ClusterManager实例对应的ClusterManagerCollector收集器都用传进来的reg注册器注册
func NewClusterManager(zone string, reg prometheus.Registerer) *ClusterManager {
	c := &ClusterManager{
		Zone: zone,
	}
	cc := ClusterManagerCollector{ClusterManager: c}
	//把传进来的注册器reg 用包装label包装，然后用得到的注册器注册我们自己定义的收集器
	prometheus.WrapRegistererWith(prometheus.Labels{"zone": zone}, reg).MustRegister(cc)
	return c
}

//ReallyExpensiveAssessmentOfTheSystemState 是数据收集的模拟，真正的集群管理器必须这样做
func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (oomCountByHost map[string]int, ramUsageByHost map[string]float64) {
	// do something about data

	// Just example fake data.
	//oom内存溢出(Out Of Memory，简称OOM)是指应用系统中存在无法回收的内存
	oomCountByHost = map[string]int{
		"foo.example.org": 42,
		"bar.example.org": 2001,
	}
	//RAM（发音同ram），是指随机存取存储器
	ramUsageByHost = map[string]float64{
		"foo.example.org": 6.023e23,
		"bar.example.org": 3.14,
	}
	return
}

// ClusterManagerCollector implements the Collector interface.
type ClusterManagerCollector struct {
	ClusterManager *ClusterManager
}

// Descriptors used by the ClusterManagerCollector below.
var (
	oomCountDesc = prometheus.NewDesc(
		"clustermanager_oom_crashes_total",
		"Number of OOM crashes.",
		[]string{"host"}, nil,
	)
	ramUsageDesc = prometheus.NewDesc(
		"clustermanager_ram_usage_bytes",
		"RAM usage as reported to the cluster manager.",
		[]string{"host"}, nil,
	)
)

//实现 prometheus Collector 接口
func (cc ClusterManagerCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(cc, ch)
}

func (cc ClusterManagerCollector) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := cc.ClusterManager.ReallyExpensiveAssessmentOfTheSystemState()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(
			oomCountDesc,
			prometheus.CounterValue,
			float64(oomCount),
			host,
		)
	}
	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(
			ramUsageDesc,
			prometheus.GaugeValue,
			ramUsage,
			host,
		)
	}
}

func main() {
	// Since we are dealing with custom Collector implementations, it might
	// be a good idea to try it out with a pedantic registry.
	reg := prometheus.NewPedanticRegistry()

	//构造集群管理器。
	//在实际代码中，可以它们分配给变量然后对它们做一些事情。
	NewClusterManager("db", reg)
	NewClusterManager("ca", reg)

	// Add the standard process and Go metrics to the custom registry.
	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
