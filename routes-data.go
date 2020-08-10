package main

// where all structs, constants and global variables (if any) goes

// route response struct
type routesResponse struct {
	Source string   `json:"source"`
	Routes []Routes `json:"routes"`
}

type Routes struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}
