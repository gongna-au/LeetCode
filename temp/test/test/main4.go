package main

type Metrics interface {
	// Type returns metrics logical type, e.g. 'downstream'/'upstream', this is more like the Subsystem concept
	Type() string

	// Labels used to distinguish the metrics' owner for same metrics key set, like 'cluster: local_service'
	Labels() map[string]string

	// SortedLabels return keys and vals in stable order
	SortedLabels() (keys, vals []string)

	// Counter creates or returns a go-metrics counter by key
	// if the key is registered by other interface, it will be panic
	Counter(key string) metrics.Counter

	// Gauge creates or returns a go-metrics gauge by key
	// if the key is registered by other interface, it will be panic
	Gauge(key string) metrics.Gauge

	// Histogram creates or returns a go-metrics histogram by key
	// if the key is registered by other interface, it will be panic
	Histogram(key string) metrics.Histogram

	// Each call the given function for each registered metric.
	Each(func(string, interface{}))

	// UnregisterAll unregister all metrics.  (Mostly for testing.)
	UnregisterAll()
}

// MetricsSink flush metrics to backend storage
type MetricsSink interface {
	// Flush flush given metrics
	Flush(writer io.Writer, metrics []Metrics)
}
