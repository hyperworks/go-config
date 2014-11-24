package config_test

import (
	. "github.com/hyperworks/go-config"
	a "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadEnv_NotPointer(t *testing.T) {
	str, num := "str", 123
	cases := []interface{}{
		str,
		&str,
		num,
		&num,
		[]interface{}{str, num},
		struct{}{},
		&struct{}{},
	}

	for c := range cases {
		a.Equal(t, ReadEnv(c, ""), ErrNotPointer)
	}

	a.Equal(t, ReadEnv(nil, ""), ErrNotPointer)
}

func TestReadEnv_EmptyStruct(t *testing.T) {
	config := &struct{}{}
	a.Nil(t, ReadEnv(config, ""))
}

func TestReadEnv_Simple(t *testing.T) {
	config := &struct {
		One string
		Two string
	}{}

	os.Setenv("ONE", "123")
	os.Setenv("TWO", "456")
	a.NoError(t, ReadEnv(config, ""))
	a.Equal(t, config.One, "123")
	a.Equal(t, config.Two, "456")
}

func TestReadEnv_Tagged(t *testing.T) {
	config := &struct {
		One string `config:"THREE"`
		Two string `config:"FOUR"`
	}{}

	os.Setenv("THREE", "third times a charm")
	os.Setenv("FOUR", "four five six")
	a.NoError(t, ReadEnv(config, ""))
	a.Equal(t, config.One, "third times a charm")
	a.Equal(t, config.Two, "four five six")
}

func TestReadEnv_Prefix(t *testing.T) {
	config := &struct {
		One string
		Two string
	}{}

	os.Setenv("PREFIX_ONE", "000123")
	os.Setenv("PREFIX_TWO", "000456")
	a.NoError(t, ReadEnv(config, "PREFIX_"))
	a.Equal(t, config.One, "000123")
	a.Equal(t, config.Two, "000456")
}

func TestReadEnv_TaggedPrefix(t *testing.T) {
	config := &struct {
		One string `config:"THREE"`
		Two string `config:"FOUR"`
	}{}

	os.Setenv("PREFIX_THREE", "000777")
	os.Setenv("PREFIX_FOUR", "000888")
	a.NoError(t, ReadEnv(config, "PREFIX_"))
	a.Equal(t, config.One, "000777")
	a.Equal(t, config.Two, "000888")
}
