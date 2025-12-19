package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	processed := make(map[uintptr]bool)
	trim(reflect.ValueOf(a), processed)
}

func trim(v reflect.Value, processed map[uintptr]bool) {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		if processed[v.Pointer()] {
			return
		}
		processed[v.Pointer()] = true

		trim(v.Elem(), processed)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			trim(v.Field(i), processed)
		}
	case reflect.String:
		if v.CanSet() {
			v.SetString(strings.Trim(v.String(), " "))
		}
	default:
		return
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
