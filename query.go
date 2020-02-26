package query

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Q contains query parameters
type Q struct {
	Where   map[string]interface{}
	Include []string
	Limit   int
	Offset  int
}

// Default return a Q with default parameters
func Default() Q {
	return Q{
		Where:   map[string]interface{}{},
		Include: []string{},
		Limit:   20,
		Offset:  0,
	}
}

// FromURL returns a Q given a URL
func FromURL(url *url.URL) Q {
	q := Default()

	for key, values := range url.Query() {
		switch key {
		case "include":
			q.Include = strings.Split(values[0], ",")
		case "limit":
			q.Limit, _ = strconv.Atoi(values[0])
		case "offset":
			q.Offset, _ = strconv.Atoi(values[0])
		default:
			q.Where[key] = values[0]
		}
	}

	fmt.Printf("from url: %+v\n", q)

	return q
}
