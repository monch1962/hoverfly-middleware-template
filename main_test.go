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

func TestPayload1(t *testing.T) {
	var payload RequestResponsePairViewV1
	var expected RequestResponsePairViewV1

	actual := transform(payload)

	expected.Response.Body = "abc123"
	if actual.Response.Body != expected.Response.Body {
		t.Fatalf("Expected Response.Body (%v) doesn't match Actual Response.Body: (%s)", expected.Response.Body, actual.Response.Body)
	}
}
