package main

import (
	"encoding/json"
	"fmt"
	"github.com/yzbuilder/json-fly/jsonutil"
	"reflect"
)

func main() {
	movie := new(jsonutil.Movie)

	jsonutil := jsonutil.NewJson(nil)
	jsonutil.JsonSetUp(reflect.ValueOf(movie))

	data, _ := json.MarshalIndent(movie, "", " ")
	fmt.Println(string(data))
}
