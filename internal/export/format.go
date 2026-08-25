package export

import "encoding/json"

func Encode(v Snapshot) ([]byte, error) { return json.MarshalIndent(v, "", "  ") }
func IsAuditable(v Snapshot) bool       { return v.Document.ID != "" && len(v.Annotations) > 0 }
func CountWords(v Snapshot) int {
	n := 0
	for _, a := range v.Annotations {
		for range a.Content {
			n++
		}
	}
	return n
}
