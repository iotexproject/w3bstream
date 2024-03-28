package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func parseEnvTag(tag string) (key string, require bool) {
	if tag == "" || tag == "-" {
		return "", false
	}
	tagKeys := strings.Split(tag, ",")
	key = tagKeys[0]
	if len(tagKeys) > 1 && tagKeys[1] == "optional" {
		return key, false
	}
	return key, true
}

func ParseEnv(c any) error {
	rv := reflect.ValueOf(c).Elem()
	rt := reflect.TypeOf(c).Elem()

	for i := 0; i < rt.NumField(); i++ {
		fi := rt.Field(i)
		fv := rv.Field(i)
		key, require := parseEnvTag(fi.Tag.Get("env"))
		if key == "" {
			continue
		}
		viper.MustBindEnv(key)

		v := viper.Get(key)
		if require && v == nil && fv.IsZero() {
			panic(fmt.Sprintf("env `%s` is require but got empty", key))
		}
		if v == nil {
			continue
		}

		switch fv.Kind() {
		case reflect.String:
			fv.Set(reflect.ValueOf(viper.GetString(key)))
		case reflect.Int:
			fv.Set(reflect.ValueOf(viper.GetInt(key)))
		case reflect.Uint64:
			fv.Set(reflect.ValueOf(viper.GetUint64(key)))
		}
	}
	return nil
}

func Print(c any) {
	rt := reflect.TypeOf(c).Elem()
	rv := reflect.ValueOf(c).Elem()

	if env, ok := c.(interface{ Env() string }); ok {
		fmt.Println(color.RedString("ENV: %s", env.Env()))
	}

	for i := 0; i < rt.NumField(); i++ {
		fi := rt.Field(i)
		fv := rv.Field(i)
		key, _ := parseEnvTag(fi.Tag.Get("env"))
		if key == "" {
			continue
		}
		fmt.Printf("%s: %v\n", color.GreenString(key), fv.Interface())
	}
}
