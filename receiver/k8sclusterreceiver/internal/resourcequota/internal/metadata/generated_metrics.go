// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

type metricK8sResourceQuotaHardLimit struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.resource_quota.hard_limit metric with initial data.
func (m *metricK8sResourceQuotaHardLimit) init() {
	m.data.SetName("k8s.resource_quota.hard_limit")
	m.data.SetDescription("The upper limit for a particular resource in a specific namespace. Will only be sent if a quota is specified. CPU requests/limits will be sent as millicores")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricK8sResourceQuotaHardLimit) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, resourceAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("resource", resourceAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sResourceQuotaHardLimit) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sResourceQuotaHardLimit) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sResourceQuotaHardLimit(cfg MetricConfig) metricK8sResourceQuotaHardLimit {
	m := metricK8sResourceQuotaHardLimit{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sResourceQuotaUsed struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.resource_quota.used metric with initial data.
func (m *metricK8sResourceQuotaUsed) init() {
	m.data.SetName("k8s.resource_quota.used")
	m.data.SetDescription("The usage for a particular resource in a specific namespace. Will only be sent if a quota is specified. CPU requests/limits will be sent as millicores")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricK8sResourceQuotaUsed) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, resourceAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("resource", resourceAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sResourceQuotaUsed) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sResourceQuotaUsed) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sResourceQuotaUsed(cfg MetricConfig) metricK8sResourceQuotaUsed {
	m := metricK8sResourceQuotaUsed{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	startTime                       pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                 int                 // maximum observed number of metrics per resource.
	resourceCapacity                int                 // maximum observed number of resource attributes.
	metricsBuffer                   pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                       component.BuildInfo // contains version information
	resourceAttributesConfig        ResourceAttributesConfig
	metricK8sResourceQuotaHardLimit metricK8sResourceQuotaHardLimit
	metricK8sResourceQuotaUsed      metricK8sResourceQuotaUsed
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                       pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                   pmetric.NewMetrics(),
		buildInfo:                       settings.BuildInfo,
		resourceAttributesConfig:        mbc.ResourceAttributes,
		metricK8sResourceQuotaHardLimit: newMetricK8sResourceQuotaHardLimit(mbc.Metrics.K8sResourceQuotaHardLimit),
		metricK8sResourceQuotaUsed:      newMetricK8sResourceQuotaUsed(mbc.Metrics.K8sResourceQuotaUsed),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(ResourceAttributesConfig, pmetric.ResourceMetrics)

// WithK8sNamespaceName sets provided value as "k8s.namespace.name" attribute for current resource.
func WithK8sNamespaceName(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sNamespaceName.Enabled {
			rm.Resource().Attributes().PutStr("k8s.namespace.name", val)
		}
	}
}

// WithK8sResourcequotaName sets provided value as "k8s.resourcequota.name" attribute for current resource.
func WithK8sResourcequotaName(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sResourcequotaName.Enabled {
			rm.Resource().Attributes().PutStr("k8s.resourcequota.name", val)
		}
	}
}

// WithK8sResourcequotaUID sets provided value as "k8s.resourcequota.uid" attribute for current resource.
func WithK8sResourcequotaUID(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sResourcequotaUID.Enabled {
			rm.Resource().Attributes().PutStr("k8s.resourcequota.uid", val)
		}
	}
}

// WithOpencensusResourcetype sets provided value as "opencensus.resourcetype" attribute for current resource.
func WithOpencensusResourcetype(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.OpencensusResourcetype.Enabled {
			rm.Resource().Attributes().PutStr("opencensus.resourcetype", val)
		}
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(_ ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/k8sclusterreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricK8sResourceQuotaHardLimit.emit(ils.Metrics())
	mb.metricK8sResourceQuotaUsed.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesConfig, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordK8sResourceQuotaHardLimitDataPoint adds a data point to k8s.resource_quota.hard_limit metric.
func (mb *MetricsBuilder) RecordK8sResourceQuotaHardLimitDataPoint(ts pcommon.Timestamp, val int64, resourceAttributeValue string) {
	mb.metricK8sResourceQuotaHardLimit.recordDataPoint(mb.startTime, ts, val, resourceAttributeValue)
}

// RecordK8sResourceQuotaUsedDataPoint adds a data point to k8s.resource_quota.used metric.
func (mb *MetricsBuilder) RecordK8sResourceQuotaUsedDataPoint(ts pcommon.Timestamp, val int64, resourceAttributeValue string) {
	mb.metricK8sResourceQuotaUsed.recordDataPoint(mb.startTime, ts, val, resourceAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
