package metric

//vec：创建具有标签维度的指标

// CounterVec counter vec.
// （计数器）：counter类型代表一种样本数据单调递增的指标，即只增不减，除非监控系统发生了重置
type CounterVec interface {
	// Inc increments the counter by 1. Use Add to increment it by arbitrary
	// non-negative values.
	Inc(labels ...string)
	// Add adds the given value to the counter. It panics if the value is <
	// 0.
	Add(v float64, labels ...string)
}

// GaugeVec gauge vec.
// （仪表盘）：Gauge类型代表一种样本数据可以任意变化的指标，即可增可减
type GaugeVec interface {
	// Set sets the Gauge to an arbitrary value.
	Set(v float64, labels ...string)
	// Inc increments the Gauge by 1. Use Add to increment it by arbitrary
	// values.
	Inc(labels ...string)
	// Dec decrements the Gauge by 1. Use Sub to decrement it by arbitrary
	// values.
	Dec(labels ...string)
	// Add adds the given value to the Gauge. (The value can be negative,
	// resulting in a decrease of the Gauge.)
	Add(v float64, labels ...string)
	// Sub subtracts the given value from the Gauge. (The value can be
	// negative, resulting in an increase of the Gauge.)
	Sub(v float64, labels ...string)
}

// HistogramVec histogram vec.
// （直方图）
type HistogramVec interface {
	// Observe adds a single observation to the histogram.
	Observe(v float64, labels ...string)
}

// SummaryVec summary vec.
//（摘要）：与Histogram类似类型，用于表示一段时间内的数据采样结果 （通常是请求持续时间或响应大小等）
// 但它直接存储了分位数（通过客户端计算，然后展示出来），而不是通过区间计算
//type SummaryVec interface {
//	Observe(v float64, labels ...string)
//}

// Opts contains the common arguments for creating vec Metric.
type Opts struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}
