package main

import (
	//"flag"
	"fmt"
	"log"
	"net/http"
	//"reflect"
	//"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	sync.RWMutex

	duration     prometheus.Gauge
	scrapeError  prometheus.Gauge
	totalScrapes prometheus.Counter
	gaugeMetrics map[string]*prometheus.GaugeVec

	mux          *http.ServeMux
	registry     *prometheus.Registry
	exporterPath ExporterPath
	options      Options
}

type Options struct {
	namespace string
}

func NewOptions(namespace string) Options {
	return Options{
		namespace,
	}
}

type ExporterPath struct {
	indexPath  string
	metricPath string
	healthPath string
}

func NewExporterPath(indexPath string, metricPath string, healthPath string) *ExporterPath {
	return &ExporterPath{
		indexPath,
		metricPath,
		healthPath,
	}
}

// addr是string类型的数组用，分割
func NewPixiuMetricsExporter(opt Options, path *ExporterPath) *Exporter {
	fmt.Printf("opts: %#v \n", opt)

	e := &Exporter{
		options: opt,
		//新的注册器用来注册本身
		registry: prometheus.NewRegistry(),

		duration: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: opt.namespace,
			Name:      "exporter_last_scrape_duration_seconds",
			Help:      "The last scrape duration.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: opt.namespace,
			Name:      "exporter_scrapes_total",
			Help:      "Current total pixiu scrapes.",
		}),
		scrapeError: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: opt.namespace,
			Name:      "exporter_last_scrape_error",
			Help:      "The last scrape error status.",
		}),
		gaugeMetrics: map[string]*prometheus.GaugeVec{},
	}
	if path != nil {
		e.exporterPath = *path
	} else {
		e.exporterPath = *NewExporterPath("/", "metrics", "/health")
	}
	//e需要实现prometheus.Collect接口，Describe、Collect方法
	e.registry.MustRegister(e)
	e.mux = http.NewServeMux()
	e.mux.HandleFunc(e.exporterPath.indexPath, e.indexHandler)
	e.mux.HandleFunc(e.exporterPath.healthPath, e.healthHandler)
	e.mux.Handle(e.exporterPath.metricPath, promhttp.HandlerFor(e.registry, promhttp.HandlerOpts{}))
	return e
}

func (e *Exporter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.mux.ServeHTTP(w, r)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.gaugeMetrics {
		m.Describe(ch)
	}
	ch <- e.duration.Desc()
	ch <- e.totalScrapes.Desc()
	ch <- e.scrapeError.Desc()
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	//收集到的数据都会被写入scrapes这个变量
	scrapes := make(chan scrapeResult)

	e.Lock()
	defer e.Unlock()

	// todo: need to clear metrics cause of eg. deleted tables.
	// but can we do better? delete selectively ?
	e.gaugeMetrics = map[string]*prometheus.GaugeVec{}
	go e.scrape(scrapes)
	e.setMetrics(scrapes)

	ch <- e.duration
	ch <- e.totalScrapes
	ch <- e.scrapeError
	e.collectMetrics(ch)
}

func (e *Exporter) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html>
<head><title>RethinkDB exporter</title></head>
<body>
<h1>Pixiu exporter</h1>
<p><a href='` + e.exporterPath.metricPath + `'>Metrics</a></p>
</body>
</html>
`))
}
func (e *Exporter) healthHandler(w http.ResponseWriter, r *http.Request) {

	_, _ = w.Write([]byte(`ok`))
}

type scrapeResult struct {
	Name   string
	Value  float64
	Server string
}

//开始收集数据 一方面统计收集数据时出现的错误的个数，然后赋值给e.scrapeError
//另外一方面把收集到的数据写入到scrapes channel
//e.scrapeError被赋值
//e.duration被赋值
func (e *Exporter) scrape(scrapes chan<- scrapeResult) {

	defer close(scrapes)
	now := time.Now().UnixNano()

	e.totalScrapes.Inc()

	errCount := 0
	if err := extractAllMetrics(scrapes); err != nil {
		errCount++
	}

	e.scrapeError.Set(float64(errCount))
	e.duration.Set(float64(time.Now().UnixNano()-now) / 1000000000)
}

func extractAllMetrics(scrapes chan<- scrapeResult) error {

	//数据的初始化
	countServers := 0
	countServerErrors := 0

	/*
		数据的计算
	*/

	//数据的写入
	scrapes <- scrapeResult{Name: "cluster_server_errors_total", Value: float64(countServerErrors)}
	scrapes <- scrapeResult{Name: "cluster_servers_total", Value: float64(countServers)}

	return nil
}

func (e *Exporter) setMetrics(scrapes <-chan scrapeResult) {

	for s := range scrapes {

		name := s.Name
		value := s.Value
		var labels prometheus.Labels = map[string]string{}

		if _, ok := e.gaugeMetrics[name]; !ok {

			asArray := make([]string, 0, len(labels))
			for k := range labels {
				asArray = append(asArray, k)
			}

			e.gaugeMetrics[name] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: e.options.namespace,
				Name:      name,
			}, asArray)
		}

		//e.metrics才被赋值
		e.gaugeMetrics[name].With(labels).Set(float64(value))
	}
}

func (e *Exporter) collectMetrics(metrics chan<- prometheus.Metric) {
	for _, m := range e.gaugeMetrics {
		m.Collect(metrics)
	}
}

func main() {
	var listenAddress = ":8899"
	options := NewOptions("pixiu")
	path := NewExporterPath("/", "metrics", "health")
	exporter := NewPixiuMetricsExporter(options, path)
	prometheus.MustRegister(exporter)
	log.Printf("listening at %s", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, exporter))

}
