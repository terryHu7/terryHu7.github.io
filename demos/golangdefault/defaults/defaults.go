package defaults

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

type RecordSpec struct {
	SplitMethod string `json:"SplitMethod" default:"NA" omit:"NA"`
	SplitUnit   int    `json:"SplitUnit" default:"65535" omit:"65535"`
}
type TMP struct {
	TMPRECORD RecordSpec
}

// Applies the default values to the struct object, the struct type must have
// the StructTag with name "default" and the directed value.
//
// Usage
//     type ExampleBasic struct {
//         Foo bool   `default:"true"`
//         Bar string `default:"33"`
//         Qux int8
//         Dur time.Duration `default:"2m3s"`
//     }
//
//      foo := &ExampleBasic{}
//      SetDefaults(foo)
func SetDefaults(variable interface{}, jsonBytes []byte) error {
	return getDefaultFiller().Fill(variable, jsonBytes)
}

var defaultFiller *Filler = nil

func getDefaultFiller() *Filler {
	if defaultFiller == nil {
		defaultFiller = newDefaultFiller()
	}

	return defaultFiller
}

func newDefaultFiller() *Filler {
	buildParseErr := func(field *FieldData, ParseErr error) error {
		return errors.New(fmt.Sprintf("Parse filed error %s! Name: %s, Expected Type: %s, Input Json: %s.", ParseErr.Error(), field.Field.Name, field.Field.Type, string(field.JsonBytes)))
	}
	funcs := make(map[reflect.Kind]FillerFunc)
	funcs[reflect.Bool] = func(field *FieldData) error {
		var setString string
		if field.JsonParserErr != nil {
			setString = field.TagValue
		} else {
			setString = string(field.JsonBytes)
		}
		value, err := strconv.ParseBool(setString)
		if err != nil {
			return buildParseErr(field, err)
		}
		field.Value.SetBool(value)
		return nil
	}

	funcs[reflect.Int] = func(field *FieldData) error {
		var setString string
		if field.JsonParserErr != nil {
			setString = field.TagValue
		} else {
			setString = string(field.JsonBytes)
		}
		value, err := strconv.ParseInt(setString, 10, 64)
		if err != nil {
			return buildParseErr(field, err)
		}
		field.Value.SetInt(value)
		return nil
	}

	funcs[reflect.Int8] = funcs[reflect.Int]
	funcs[reflect.Int16] = funcs[reflect.Int]
	funcs[reflect.Int32] = funcs[reflect.Int]
	funcs[reflect.Int64] = func(field *FieldData) error {
		if field.Field.Type == reflect.TypeOf(time.Second) {
			value, _ := time.ParseDuration(field.TagValue)
			field.Value.Set(reflect.ValueOf(value))
			return nil
		} else {
			var setString string
			if field.JsonParserErr != nil {
				setString = field.TagValue
			} else {
				setString = string(field.JsonBytes)
			}
			value, err := strconv.ParseInt(setString, 10, 64)
			if err != nil {
				return buildParseErr(field, err)
			}
			field.Value.SetInt(value)
			return nil
		}
	}

	funcs[reflect.Float32] = func(field *FieldData) error {
		var setString string
		if field.JsonParserErr != nil {
			setString = field.TagValue
		} else {
			setString = string(field.JsonBytes)
		}
		value, err := strconv.ParseFloat(setString, 64)
		if err != nil {
			return buildParseErr(field, err)
		}
		field.Value.SetFloat(value)
		return nil
	}

	funcs[reflect.Float64] = funcs[reflect.Float32]

	funcs[reflect.Uint] = func(field *FieldData) error {
		var setString string
		if field.JsonParserErr != nil {
			setString = field.TagValue
		} else {
			setString = string(field.JsonBytes)
		}
		value, err := strconv.ParseUint(setString, 10, 64)
		if err != nil {
			return buildParseErr(field, err)
		}
		field.Value.SetUint(value)
		return nil
	}

	funcs[reflect.Uint8] = funcs[reflect.Uint]
	funcs[reflect.Uint16] = funcs[reflect.Uint]
	funcs[reflect.Uint32] = funcs[reflect.Uint]
	funcs[reflect.Uint64] = funcs[reflect.Uint]

	funcs[reflect.String] = func(field *FieldData) error {
		if field.JsonParserErr != nil {
			field.Value.SetString(field.TagValue)
		} else {
			field.Value.SetString(string(field.JsonBytes))
		}
		return nil
	}

	funcs[reflect.Struct] = func(field *FieldData) error {
		fields := getDefaultFiller().GetFieldsFromValue(field.Value, field.JsonBytes, field)
		return getDefaultFiller().SetDefaultValues(fields)
	}

	funcs[reflect.Slice] = func(field *FieldData) error {
		errorStrings := []string{}

		elemType := field.Value.Type().Elem()
		k := elemType.Kind()
		switch k {
		case reflect.Struct:
			count := 0
			jsonparser.ArrayEach(field.JsonBytes, func(jsonBytes []byte, dataType jsonparser.ValueType, offset int, err error) {
				count++
			})
			for i := 0; i < count; i++ {
				jsonBytes, _, _, _ := jsonparser.Get(field.JsonBytes, "["+strconv.Itoa(i)+"]")
				field.Value.Set(reflect.Append(field.Value, reflect.New(elemType).Elem()))
				fields := getDefaultFiller().GetFieldsFromValue(field.Value.Index(i), jsonBytes, nil)
				err := getDefaultFiller().SetDefaultValues(fields)
				if err != nil {
					errorStrings = append(errorStrings, err.Error())
				}
			}
			if len(errorStrings) != 0 {
				return errors.New(strings.Join(errorStrings, ";"))
			}
			return nil
		default:
			//处理形如 [1,2,3,4]
			reg := regexp.MustCompile(`^\[((?s).*)\]$`)
			jsonMatchs := reg.FindStringSubmatch(string(field.JsonBytes))
			if len(jsonMatchs) != 2 {
				return nil
			}
			if jsonMatchs[1] == "" {
				field.Value.Set(reflect.MakeSlice(field.Value.Type(), 0, 0))
				return nil
			} else {
				defaultValue := strings.Split(jsonMatchs[1], ",")
				result := reflect.MakeSlice(field.Value.Type(), len(defaultValue), len(defaultValue))
				for i := 0; i < len(defaultValue); i++ {
					jsonBytes, _, _, err := jsonparser.Get(field.JsonBytes, "["+strconv.Itoa(i)+"]")
					itemValue := result.Index(i)
					item := &FieldData{
						Value:         itemValue,
						Field:         reflect.StructField{},
						TagValue:      defaultValue[i],
						Parent:        nil,
						JsonBytes:     jsonBytes,
						JsonParserErr: err,
					}
					fillerr := funcs[k](item)
					if fillerr != nil {
						errorStrings = append(errorStrings, err.Error())
					}
				}
				field.Value.Set(result)
				if len(errorStrings) != 0 {
					return errors.New(strings.Join(errorStrings, ";"))
				}
				return nil
			}
		}
	}

	return &Filler{FuncByKind: funcs, Tag: "default"}
}
