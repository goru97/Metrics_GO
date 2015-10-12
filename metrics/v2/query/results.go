package query

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
)

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

type MetricData struct {
	Unit string `mapstructure:"unit"`
	Values []Value `mapstructure:"values"`
	Metadata MetaData `mapstructure:"metadata"`
}
type Value struct {
	NumPoints int32 `mapstructure:"numPoints"`
	TimeStamp int64 `mapstructure:"timestamp"`
	Average int64 `mapstructure:"average"`
	Sum int64 `mapstructure:"sum"`
	Max int64 `mapstructure:"max"`
	Latest int64 `mapstructure:"latest"`
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
