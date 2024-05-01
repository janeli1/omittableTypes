package omittable

import "encoding/json"

// Omittable is a wrapper around a value that also stores whether it is set
// or not.

type OmittableInt struct {
	value *int
	set   bool
}

func OmittableOfInt(value *int) OmittableInt {
	return OmittableInt{
		value: value,
		set:   true,
	}
}

func (o OmittableInt) Value() *int {
	if !o.set {
		var zero *int
		return zero
	}
	return o.value
}

func (o OmittableInt) ValueOK() (*int, bool) {
	if !o.set {
		var zero *int
		return zero, false
	}
	return o.value, true
}

func (o OmittableInt) IsSet() bool {
	return o.set
}

func (o OmittableInt) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

func (o *OmittableInt) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &o.value)
	if err != nil {
		return err
	}
	o.set = true
	return nil
}
