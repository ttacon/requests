package requests

import "testing"

var (
	yolo2 = "/hello/:Super/:Whooper"
	vals  = map[string]string{"Super": "hello", "Whooper": "ruhroh"}
)

func Test_Parse(t *testing.T) {
	templ := ParseURLTemplate(yolo2)
	got := templ.Build(vals)
	if got != "/hello/hello/ruhroh" {
		t.Errorf("did not get expected, got: %s", got)
	}
}
