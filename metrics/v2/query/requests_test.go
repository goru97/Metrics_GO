package query
import (
	"testing"
	"net/http"
	"fmt"
	fake "github.com/rackspace/gophercloud/testhelper/client"
	th "github.com/rackspace/gophercloud/testhelper"
)

const DUMMY_METRIC  = "rollup01.iad.blueflood.rollup.com.rackspacecloud.blueflood.service.RollupService.Rollup-Execution-Timer.count"

func TestGetDataByPoints(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/views/"+DUMMY_METRIC, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
  "metadata": {
    "count": 4,
    "limit": null,
    "marker": null,
    "next_href": null
  },
  "unit": "unknown",
  "values": [
    {
      "average": 20015507706.50,
      "numPoints": 32,
      "timestamp": 1444353600000
    },
    {
      "average": 20045095111.91,
      "numPoints": 32,
      "timestamp": 1444354800000
    },
    {
      "average": 20077360883.33,
      "numPoints": 33,
      "timestamp": 1444356000000
    },
    {
      "average": 20110963258.48,
      "numPoints": 25,
      "timestamp": 1444357200000
    }]
    }`)
	})

	expected := MetricData{
		MetaData{
			Count: 4,
		},
		Unit: "unknown",
		[]Value: {
			Value{
				Average: 2.00155e+10,
				NumPoints: 32,
				TimeStamp: 1444353600000,
			},
			Value{
				Average: 2.00451e+10,
				NumPoints: 32,
				TimeStamp: 1444354800000,
			},
			Value{
				Average: 2.00774e+10,
				NumPoints: 33,
				TimeStamp: 1444356000000,
			},
			Value{
				Average: 2.0111e+10,
				NumPoints: 25,
				TimeStamp: 1444357200000,
			},
		},
	}
	actual, err := GetDataByPoints(fake.ServiceClient() ,DUMMY_METRIC, 1444354736000, 1444355736000, 200).GetMetricsData()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetEvents(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/events/getEvents", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `[{
	  "what": "Test Event 1",
	  "when": 1444354736003,
	  "data": "Test Data 1",
	  "tags": "Restart"
	  },
	  {
	  "what": "Test Event 2",
	  "when": 1444354736104,
	  "data": "Test Data 2",
	  "tags": "Shutdown"
	  }]`)
	})

	expected := []Event{
		Event{
			What:   "Test Event 1",
			When: 1444354736003,
			Data: "Test Data 1",
			Tags: "Restart",
		},
		Event{
			What:   "Test Event 2",
			When: 1444354736104,
			Data: "Test Data 2",
			Tags: "Shutdown",
		},
	}
	actual, err := GetEvents(fake.ServiceClient(),1444354736000, 1444355736000)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expected, actual)
}