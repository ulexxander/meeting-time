package scalars

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type TimeOfDay time.Time

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (t *TimeOfDay) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TimeOfDay must be a string")
	}
	parsed, err := time.Parse(time.Kitchen, str)
	*t = TimeOfDay(parsed)
	return err
}

// MarshalGQL implements the graphql.Marshaler interface
func (t TimeOfDay) MarshalGQL(w io.Writer) {
	formatted := time.Time(t).Format(time.Kitchen)
	marshaled, _ := json.Marshal(formatted)
	w.Write(marshaled)
}
