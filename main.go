package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stderr, "", 0)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		var payload RequestResponsePairViewV1

		err := json.Unmarshal(s.Bytes(), &payload)
		if err != nil {
			l.Println("Failed to unmarshal payload from hoverfly")
		}

		payload = transform(payload)

		bts, err := json.Marshal(payload)
		if err != nil {
			l.Println("Failed to marshal new payload")
		}

		os.Stdout.Write(bts)
	}
}

//////////////////////////////////////////////////////////////////////////////
// DON'T CHANGE ANYTHING ABOVE THIS LINE!!
//
// Just replace the transform function with your own middleware logic
//////////////////////////////////////////////////////////////////////////////

func transform(payload RequestResponsePairViewV1) RequestResponsePairViewV1 {
	// Set the payload.Request.Body to a value simply so we can show how to
	// use it in the response later
	sample_request_body := "abc"
	payload.Request.Body = &sample_request_body

	payload.Response.Body = "abc" + "def" + fmt.Sprint(4) + *payload.Request.Body
	payload.Response.Body = fmt.Sprint(time.Now())

	payload.Response.Status = 200

	payload.Response.Headers = map[string][]string{"header": {"value"}}
	payload.Response.EncodedBody = false

	return payload
}
