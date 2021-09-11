package main

import (
	"fmt"
	"testing"
)

//
func Compare(t *testing.T, expected RequestResponsePairViewV1, actual RequestResponsePairViewV1) {
	ar := actual.Response
	er := expected.Response
	var ok = true
	if ar.Body != er.Body {
		ok = false
		t.Logf("Expected Response.Body (%v) doesn't match Actual Response.Body: (%s)", er.Body, ar.Body)
	}
	if ar.Status != er.Status {
		ok = false
		t.Logf("Expected Response.Status (%d) doesn't match Actual Response.Status: (%d)", er.Status, ar.Status)
	}
	if !ok {
		t.Fail()
	}
}

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

	// Set all the request parameters prior to applying the middleware transformation
	//payload.Request.Body = ...

	// Now copy the input payload to expected, so we know only those fields that have been changed will be different
	expected := payload

	// Now run the middleware against the payload, and save the result into transformed
	transformed := transform(payload)

	// Set all the expected response parameters after the payload is transformed
	//expected.Response.Body = "abc123"
	expected.Response.Status = 200
	//expected.Response.Headers = ...

	// Now compare them
	Compare(t, expected, transformed)
}
