package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quoutedJSONValue := strconv.Quote(jsonValue)

	return []byte(quoutedJSONValue), nil
}
