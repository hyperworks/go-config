package config

import (
	"strings"
	"errors"
	"reflect"
	"os"
)

var ErrNotPointer = errors.New("read target must be a pointer to a struct.")

func ReadEnv(target interface{}, prefix string) error {
	value := reflect.ValueOf(target)
	if v := reflect.Indirect(value); value == v {
		return ErrNotPointer
	} else {
		value = v
	}

	t := value.Type()
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		readConfigField(prefix, t.Field(i), value.Field(i))
	}

	return nil
}

func readConfigField(prefix string, field reflect.StructField, value reflect.Value) {
	tagName := field.Tag.Get("config")
	if tagName == "" {
		tagName = strings.ToUpper(field.Name)
	}

	envName := prefix + tagName
	if envValue := os.Getenv(envName); envValue != "" {
		value.Set(reflect.ValueOf(envValue))
	}
}

