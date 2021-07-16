// Package constant contains the common literals used by
// the service.
package constant

const (

	// ServiceId is the unique identifier of this service.
	ServiceId string = "082"

	GetScheduleByChannelId string = "01"

	GetScheduleSource string = ServiceId + "-" + GetScheduleByChannelId

	// XTrackingId is the name of the request header that contains
	// the tracking identifier.
	XTrackingId string = "X-Tracking-Id"

	// XClientId is the name of the request header that contains
	// the client identifier.
	XClientId string = "X-Client-Id"

	// XAuthorization is the name of the request header that contains
	// the access token.
	XAuthorization string = "X-Authorization"

	// ContentType is the name of the request header that contains
	// the request content type.
	ContentType string = "content-type"

	// ApplicationJSON is the name of the request header that contains
	// the content type value.
	ApplicationJSON string = "application/json"

	// SuccessCode represents successful processing of the
	// incoming request.
	SuccessCode int = 0

	// FailureCode represents failure processing of the
	// incoming request.
	FailureCode int = -1

	// SuccessMessage represents successful processing
	// of the incoming request.
	SuccessMessage string = "Success"

	// FailureMessage represents failure processing
	// of the incoming request.
	FailureMessage string = "Failure"

	// Request payload constants
	UserId = "userId"
	ItemId = "itemId"
	Offset = "offset"

	// New Relic Constants
	// NRTransaction represents the context key for NewRelic transaction.
	NewRelicTransaction string = "NewRelicTransaction"
	HANDLER             string = "handler"
	SERVICE             string = "service"
	REPOSITORY          string = "repository"
)
