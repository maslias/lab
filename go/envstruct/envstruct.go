package envstruct

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	env "github.com/maslias/envstruct/pkg"
)

var (
	Delimeter             string = "_"
	Capitalize            bool   = true
	Preterm               string = ""
	TagKeyForValue        string = "env"
	TagKeyForDefaultValue string = "default"

	ErrStructStrconf  = errors.New("envstruct")
	ErrStructenvParse = errors.New("invalid structure err: envstruct")
)

func Parse(s interface{}) error {
	if reflect.TypeOf(s) == nil {
		return ErrStructenvParse
	}

	if reflect.TypeOf(s).Kind() != reflect.Ptr {
		return ErrStructenvParse
	}

	var buf bytes.Buffer
	err := parseStruct(reflect.ValueOf(s).Elem(), reflect.StructField{}, buf)
	if err != nil {
		return err
	}

	return nil
}

func parseStruct(rv reflect.Value, pRsf reflect.StructField, buf bytes.Buffer) error {
	if len(pRsf.Name) != 0 {
		if buf.Len() != 0 {
			buf.WriteString(Delimeter)
		}
		buf.WriteString(pRsf.Name)
	}

	if rv.Kind() != reflect.Struct {
		return setEnvForStruct(rv, pRsf, buf)
	}

	for i := range rv.NumField() {
		rvf := rv.Field(i)
		rsf := rv.Type().Field(i)

		if err := parseStruct(rvf, rsf, buf); err != nil {
			return err
		}
	}
	return nil
}

func setEnvForStruct(rv reflect.Value, rsf reflect.StructField, buf bytes.Buffer) error {
	envTag, ok := rsf.Tag.Lookup(TagKeyForValue)
	if !ok {
		envTag = buf.String()
	}

	if len(Preterm) != 0 {
		envTag = Preterm + Delimeter + envTag
	}

	if Capitalize {
		envTag = strings.ToUpper(envTag)
	}

	var envVal string
	var err error

	envVal, err = env.GetEnv(envTag)
	if err != nil {
		if !errors.Is(err, env.ErrEnvLookup) {
			return err
		}
		defValue, ok := rsf.Tag.Lookup(TagKeyForDefaultValue)
		if ok {
			envVal = defValue
		}
	}

	switch rv.Kind() {
	case reflect.Int:
		intEnvVal, err := strconv.Atoi(envVal)
		if err != nil {
			return fmt.Errorf(
				"could not convert key: %s msg: %s err: %w",
				envTag,
				err.Error(),
				ErrStructStrconf,
			)
		}
		rv.Set(reflect.ValueOf(intEnvVal))

	default:
		rv.Set(reflect.ValueOf(envVal))
	}

	return nil
}
