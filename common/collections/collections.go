// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package collections contains common Hugo functionality related to collection
// handling.
package collections

import (
	"reflect"
	//"fmt"
	//"github.com/spf13/cast"
)

// Grouper defines a very generic way to group items by a given key.
type Grouper interface {
	Group(key interface{}, items interface{}) (interface{}, error)
}

var _emptyList []interface{}
var _emptyAny interface{}
var anyType = reflect.TypeOf(_emptyAny)
var anyListType = reflect.TypeOf(_emptyList)
var stringType = reflect.TypeOf("")

func SeqItems(seq interface{}) []interface{} {

	seqv := reflect.ValueOf(seq)

	res := make([]interface{}, 0)

	switch seqv.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < seqv.Len(); i++ {
			res = append(res, seqv.Index(i).Interface())
		}
	case reflect.Map:
		for _, k := range seqv.MapKeys() {
            v := seqv.MapIndex(k)
            res = append(res, v.Interface())
        }
	default:
		return nil
	}

	return res
}

func Group2(key interface{}, seq interface{}) (interface{}, error) {

	// helpers.DistinctFeedbackLog.Printf("Getting sequence")
	// items := SeqItems(seq)
	// if items == nil {
	// 	return nil, fmt.Errorf("expected a slice, sequence, or map of items, got %T", seq)
	// }

	// helpers.DistinctFeedbackLog.Printf("Getting key str")
	// keyStr, err := cast.ToStringE(key)
	// if err != nil {
	// 	return nil, fmt.Errorf("expected a string key, got %T", key)
	// }

	// helpers.DistinctFeedbackLog.Printf("Making map")
	// mapv := reflect.MakeMap(reflect.MapOf(anyType, anyListType))

	// helpers.DistinctFeedbackLog.Printf("Iterating")
	// for _, e := range items {
	// 	ppv := reflect.ValueOf(e)
	// 	fv := ppv.Elem().FieldByName(keyStr)
	// 	if !fv.IsValid() {
	// 		fmt.Errorf("Group: no field by name of %s in sequence value, skipping item", keyStr)
	// 		continue
	// 	}
	// 		fmt.Printf("Inserting")
	// 	if !mapv.MapIndex(fv).IsValid() {
	// 		mapv.SetMapIndex(fv, reflect.MakeSlice(anyListType, 0, 0))
	// 	}
	// 	mapv.SetMapIndex(fv, reflect.Append(mapv.MapIndex(fv), ppv))
	// }

	// helpers.DistinctFeedbackLog.Printf("Done")
	return nil, nil
}