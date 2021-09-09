package main

//////////////////////////////////////////////////////////
//
// DON'T CHANGE THIS FILE
//
// It contains the request-response type format used by Hoverfly middleware
//////////////////////////////////////////////////////////

type RequestResponsePairViewV1 struct {
	Response ResponseDetailsView `json:"response"`
	Request  RequestDetailsView  `json:"request"`
}

type ResponseDetailsView struct {
	Status      int                 `json:"status"`
	Body        string              `json:"body"`
	EncodedBody bool                `json:"encodedBody"`
	Headers     map[string][]string `json:"headers,omitempty"`
}

// RequestDetailsView is used when marshalling and unmarshalling RequestDetails
type RequestDetailsView struct {
	RequestType *string             `json:"requestType,omitempty"`
	Path        *string             `json:"path"`
	Method      *string             `json:"method"`
	Destination *string             `json:"destination"`
	Scheme      *string             `json:"scheme"`
	Query       *string             `json:"query"`
	QueryMap    map[string][]string `json:"-"`
	Body        *string             `json:"body"`
	Headers     map[string][]string `json:"headers"`
}
