package main

type OSMRResponse struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

type MyResponse struct {
	Source string      `json:"source"`
	Routes RoutesSlice `json:"routes"`
	Err string `json:"err"`
}

type Route struct {
	Destination string  `json:"destination"`
	Duration    float32 `json:"duration"`
	Distance    float32 `json:"distance"`
}

type RoutesSlice []Route

func (r RoutesSlice) Len() int { return len(r) }
func (r RoutesSlice) Less(i, j int) bool {
	if r[i].Duration == r[j].Duration {
		return r[i].Distance < r[j].Distance
	} else {
		return r[i].Duration < r[j].Duration
	}
}
func (r RoutesSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
