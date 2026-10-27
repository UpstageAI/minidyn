package server

import "encoding/json"

// AttributeValue is a concrete representation of DynamoDB's AttributeValue
// that works with standard json encoding/decoding.
type AttributeValue struct {
	B    []byte                     `json:"B,omitempty"`
	BOOL *bool                      `json:"BOOL,omitempty"`
	BS   [][]byte                   `json:"BS,omitempty"`
	L    []*AttributeValue          `json:"L,omitempty"`
	M    map[string]*AttributeValue `json:"M,omitempty"`
	N    *string                    `json:"N,omitempty"`
	NS   []*string                  `json:"NS,omitempty"`
	NULL *bool                      `json:"NULL,omitempty"`
	S    *string                    `json:"S,omitempty"`
	SS   []*string                  `json:"SS,omitempty"`
}

// MarshalJSON keeps empty L and M members. DynamoDB stores {"L": []} and
// {"M": {}} as values, while omitempty would drop them and return no attribute.
func (av AttributeValue) MarshalJSON() ([]byte, error) {
	type plain AttributeValue

	out := struct {
		plain
		L *[]*AttributeValue          `json:"L,omitempty"`
		M *map[string]*AttributeValue `json:"M,omitempty"`
	}{plain: plain(av)}

	if av.L != nil {
		out.L = &av.L
	}

	if av.M != nil {
		out.M = &av.M
	}

	return json.Marshal(out)
}
