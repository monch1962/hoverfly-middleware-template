package main

import (
	"fmt"
	"testing"
)

func TestNothing(t *testing.T) {
	var payload RequestResponsePairViewV1
	sample_req_body := "abc123"
	payload.Request.Body = &sample_req_body
	//s := "s"
	//payload.Request.Destination = &s
	//payload.Request.Method = &s
	//payload.Request.RequestType = &s
	//payload.Request.Path = &s
	//payload.Request.Headers
	//payload.Request.Query
	//payload.Request.QueryMap
	//payload.Request.Scheme

	payload.Response.Body = "abc" + "def" + fmt.Sprint(4) + *payload.Request.Body
	payload.Response.Status = 200
	payload.Response.Headers = map[string][]string{"header": {"value"}}
	payload.Response.EncodedBody = false

	t.Log("Check this test for how to use each of the parameters inside middleware payload")
}

// Use the following test as a template and copy/paste/rename/adapt it to create however many test cases
// you need to adequately test your middleware
func TestPayload1(t *testing.T) {

	// Declare two RequestResponsePairViewV1 variables - one to hold the request payload to be transformed, and another to hold the expected
	// response payload after transformation. We're going to transform the 1st and compare with the 2nd
	var payload RequestResponsePairViewV1
	var expected RequestResponsePairViewV1

	// Set all the request parameters prior to applying the middleware transformation
	//payload.Request.Body = ...

	actual := transform(payload)

	// Set all the expected response parameters after the payload is transformed
	//expected.Response.Body = "abc123"
	//expected.Response.Status = 200
	//expected.Response.Headers = ...

	// Now compare them, and log any discrepancies
	if actual.Response.Body != expected.Response.Body {
		t.Fatalf("Expected Response.Body (%v) doesn't match Actual Response.Body: (%s)", expected.Response.Body, actual.Response.Body)
	}
}
