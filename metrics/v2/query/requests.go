package query

import (
	"github.com/rackspace/gophercloud"
	"errors"
	"github.com/mitchellh/mapstructure"
)

const(
	SUM = "sum"
	MAX = "max"
	LATEST = "latest"
)

const(
	FULL = "full"
	MIN5 = "min5"
	MIN20 = "min20"
	MIN60 = "min60"
	MIN240 = "min240"
	MIN1440 = "min1440"
)

// GetListByPoints retrieve data for a list of metrics for the specified tenant associated with RackspaceMetrics.
//TODO: Change Request Payload Structure in Blueflood
/*func GetDataByPoints(c *gophercloud.ServiceClient, from int64, to int64, points int32, metrics ...string) GetResult {
	var res GetResult
	slc, _ := json.Marshal(metrics)
	reqBody := string(slc)
	fmt.Println(reqBody)
	_, res.Err = c.Post(getDataForList(c, from, to, points),reqBody, &res.Body, nil)
	return res
}*/

func GetDataByPoints(c *gophercloud.ServiceClient, metric string, from int64, to int64, points int32, sel ...string) GetResult {
	var res GetResult
	if len(sel) >1{
		 res.Err = errors.New("Only one select parameter is allowed.")
    return res
	}else {
		if len(sel) == 1{
			_, res.Err = c.Get(getURLForPointsWithSelect(c, metric, from, to, points,sel[0]), &res.Body, nil)
			return res
		}
		_, res.Err = c.Get(getURLForPoints(c, metric, from, to, points), &res.Body, nil)
		return res
    }
}

func GetDataByResolution(c *gophercloud.ServiceClient, metric string, from int64, to int64, resolution string, sel ...string) GetResult {
	var res GetResult
	if len(sel) >1{
		res.Err = errors.New("Only one select parameter is allowed.")
		return res
	}else {
		if len(sel) == 1{
			_, res.Err = c.Get(getURLForResolutionWithSelect(c, metric, from, to, resolution, sel[0]), &res.Body, nil)
			return res
		}
		_, res.Err = c.Get(getURLForResolution(c, metric, from, to, resolution), &res.Body, nil)
		return res
	}
}

// Limits retrieves the number of API transaction that are available for the specified tenant associated with RackspaceMetrics.
func Limits(c *gophercloud.ServiceClient) GetResult {
	var res GetResult
	_, res.Err = c.Get(getLimits(c), &res.Body, nil)
	return res
}

// SearchMetric retrieves a list of available metrics for the specified tenant associated with RackspaceMetrics.
func SearchMetric(c *gophercloud.ServiceClient, metric string) GetResult {
	var res GetResult
	_, res.Err = c.Get(getSearchURL(c, metric), &res.Body, nil)
	return res
}

//GetEvents retrieves a list of events for the specified tenant associated with RackspaceMetrics.
func GetEvents(c *gophercloud.ServiceClient, from int64, until int64, tag ...string) ([]Event, error) {
	var res GetResult
	if len(tag) == 1 {
		_, res.Err = c.Get(getEventURLForTag(c, from, until, tag[0]), &res.Body, nil)
	} else {
		_, res.Err = c.Get(getEventURL(c, from, until), &res.Body, nil)
	}
	b := res.Body.(interface{})
	var events []Event
	err := mapstructure.Decode(b,&events)
	return events, err
}