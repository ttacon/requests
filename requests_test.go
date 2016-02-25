package requests

import (
	"fmt"
	"testing"
)

func Test_Json(t *testing.T) {
	resp, err := Post("http://localhost:18099").
		AddHeader("X-Mariner-ID", "test").
		JSONBody(map[string]string{"hello?": "why yes!"}).
		Do()
	if err != nil {
		t.Error("failed to send request, err:", err)
		t.Fail()
	}

	fmt.Println(resp.Status)
}
