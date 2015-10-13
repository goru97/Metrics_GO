package query

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
)

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

type MetricListData struct {
	Metrics []MetricList `mapstructure:"metrics"`
}

type MetricList struct {
		Unit string `mapstructure:"unit"`
		Metric string `mapstructure:"metric"`
		Data []Value `mapstructure:"data"`
		Type string `mapstructure:"type"`
}

type Metric struct {
	Metric string `mapstructure:"metric"`
}

type MetricData struct {
	Unit string `mapstructure:"unit"`
	Values []Value `mapstructure:"values"`
	MetaData MetaData `mapstructure:"metadata"`
}
type Value struct {
	NumPoints int32 `mapstructure:"numPoints"`
	TimeStamp int64 `mapstructure:"timestamp"`
	Average float64 `mapstructure:"average"`
	Sum float64 `mapstructure:"sum"`
	Max float64 `mapstructure:"max"`
	Min float64 `mapstructure:"min"`
	Latest float64 `mapstructure:"latest"`
	Variance float64 `mapstructure:"variance"`
}

type MetaData struct {
	Limit string `mapstructure:"limit"`
	Next_Href string `mapstructure:"next_href"`
	Count int32 `mapstructure:"count"`
	Marker string `mapstructure:"marker"`
}

type Event struct {
	What string `mapstructure:"what"`
	When int64 `mapstructure:"when"`
	Data string `mapstructure:"data"`
	Tags string `mapstructure:"tags"`
}

// ExtractMetadata is a function that takes a GetResult (of type *http.Response)
// and returns the custom metatdata associated with the account.
func (gr GetResult) ExtractMetadata() (MetaData, error) {
	var res MetaData
	b := gr.Body.(map[string]interface{})
	err := mapstructure.Decode(b["metadata"],&res)
	return res, err
}

func (gr GetResult) GetValues() ([]Value, error) {
	var res []Value
	b := gr.Body.(map[string]interface{})
	err := mapstructure.Decode(b["values"],&res)
	return res, err
}

func (gr GetResult) GetMetricsData() (MetricData, error) {
	var res MetricData
	b := gr.Body.(interface{})
	err := mapstructure.Decode(b,&res)
	return res, err
}
