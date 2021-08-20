package apiv1

import (
	"encoding/json"
	"fmt"
)

func MapDatader(dict map[int]string) (jons string)  {
	jsonstr, err := json.Marshal(dict)
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonstr)
}