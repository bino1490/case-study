package response

// Error describes the error condition.
type Error struct {

	// Code provides the numeric identifier of the
	// error condition occurred.
	Code string `json:"code"`

	// Description provides the short description of the
	// error condition occurred.
	Description string `json:"description"`
}

// Header describes the overall success or failure condition.
type Header struct {

	// Source provides the unique identifier of the service.
	Source string `json:"source"`

	// Code provides the numeric identifier of the overall response.
	// Returns 0 or -1
	Code int `json:"code"`

	// Message provides the short description of the overall response.
	// Returns success or failure
	Message string `json:"message"`

	// SystemTime provides the time in milliseconds when the response.
	// was served
	SystemTime int64 `json:"system_time"`

	// Start provides the start index of the records returned.
	Start int `json:"start,omitempty"`

	// Rows provides the number of records returned.
	Rows int `json:"rows,omitempty"`

	// Count provides the total number of records available.
	Count int `json:"count,omitempty"`

	// TrackingId provides the unique identifier of the request for
	// tracing purposes.
	TrackingId string `json:"tracking_id"`

	// Errors provides the array of error conditions.
	Errors []Error `json:"errors,omitempty"`
}

// Response describes a outgoing response structure.
type Response struct {

	// Header provides the overall success or failure condition.
	Header *Header `json:"header"`

	// Data provides the business payload
	Data interface{} `json:"data,omitempty"`
}

// Count describes the total count of entities returned
// for incoming condition.
type Count struct {

	// Count provides the number of records returned.
	Count *int `json:"count,omitempty"`
}
