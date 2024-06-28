package main

import (
	"errors"
	"fmt"
	"reflect"
)

func reflectOver(t any) error {
	typeOfT := reflect.TypeOf(t)
	if typeOfT.Kind() != reflect.Struct {
		return fmt.Errorf("can't reflect the fields of non-struct type %s", typeOfT.Elem().Name())
	}

	fields := reflect.VisibleFields(reflect.TypeOf(t))
	fmt.Printf("Type '%s' fields:\n", typeOfT.Name())
	for _, f := range fields {
		fmt.Println(f.Name)
	}
	fmt.Println()
	return nil
}

type Foo struct {
	Bar string
	Baz string
}

type Bam struct {
	Bip int
	Bop float64
}

func main() {
	f := Foo{}
	b := Bam{}

	// These both work beacause the variables f and b contain struct values
	_ = reflectOver(f)
	_ = reflectOver(b)

	var e error
	e = errors.New("this is an error interface")
	// This won't work, because the variable e is an interface value
	err := reflectOver(e)
	if err != nil {
		fmt.Println(err)
	}
}