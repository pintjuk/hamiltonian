package osrm

type Route struct {
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`
}

type APIResponse struct {
	Routes []Route `json:"routes"`
	Code   string  `json:"code"`
}
