// Package requests is a useful library for constructing HTTP requests
// and retrieving responses.
//
// For example, if we wanted to search on CNN for articles on the US
// elections:
//
//   resp, err := requests.Get("http://www.cnn.com/search/").
//           Params(map[string]string{"text":"elections"}).
//           Do()
//
//
package requests
