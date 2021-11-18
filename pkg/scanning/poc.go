package scanning

import (
	"io/ioutil"
	"net/http"

	"github.com/hahwul/dalfox/v2/pkg/model"
)

func MakePoC(poc string, req *http.Request, options model.Options) string {
	if req.Body != nil {
		body, err := req.GetBody()
		if err == nil {
			reqBody, err := ioutil.ReadAll(body)
			if err == nil {
				if string(reqBody) != "" {
					return poc + " -d " + string(reqBody)
				}
			}
		}
	} else {
		return poc
	}
	return poc
}