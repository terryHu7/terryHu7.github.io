package defaults

import (
	"errors"
	"reflect"
	"strings"

	"github.com/buger/jsonparser"
)

type FieldData struct {
	Field         reflect.StructField
	Value         reflect.Value
	TagValue      string
	JsonBytes     []byte
	JsonParserErr error
	Parent        *FieldData
}

type FillerFunc func(field *FieldData) error

// Filler contains all the functions to fill any struct field with any type
// allowing to define function by Kind, Type of field name
type Filler struct {
	FuncByKind map[reflect.Kind]FillerFunc
	Tag        string
}

// Fill apply all the functions contained on Filler, setting all the possible
// values
func (f *Filler) Fill(variable interface{}, jsonBytes []byte) error {
	fields := f.getFields(variable, jsonBytes)
	return f.SetDefaultValues(fields)
}

func (f *Filler) getFields(variable interface{}, jsonBytes []byte) []*FieldData {
	valueObject := reflect.ValueOf(variable).Elem()

	return f.GetFieldsFromValue(valueObject, jsonBytes, nil)
}

func (f *Filler) GetFieldsFromValue(valueObject reflect.Value, jsonBytes []byte, parent *FieldData) []*FieldData {
	var results []*FieldData
	typeObject := valueObject.Type()
	switch typeObject.Kind() {
	case reflect.Struct:
		{
			count := valueObject.NumField()

			for i := 0; i < count; i++ {
				structValue := valueObject.Field(i)
				field := typeObject.Field(i)

				value, _, _, err := jsonparser.Get(jsonBytes, field.Tag.Get("json"))

				if structValue.CanSet() {
					results = append(results, &FieldData{
						Value:         structValue,
						Field:         field,
						TagValue:      field.Tag.Get(f.Tag),
						Parent:        parent,
						JsonBytes:     value,
						JsonParserErr: err,
					})
				}
			}
		}
	case reflect.Slice: //[{},{}]
		{
			results = append(results, &FieldData{
				Value: valueObject,
				Field: reflect.StructField{
					Type: typeObject,
				},
				Parent:    parent, //nil
				JsonBytes: jsonBytes,
			})
		}

	}

	return results
}

func (f *Filler) SetDefaultValues(fields []*FieldData) error {
	errStrings := []string{}
	for _, field := range fields {
		err := f.SetDefaultValue(field)
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}
	if len(errStrings) != 0 {
		return errors.New(strings.Join(errStrings, ";"))
	}
	return nil
}

func (f *Filler) SetDefaultValue(field *FieldData) error {
	getters := []func(field *FieldData) FillerFunc{
		f.getFunctionByKind,
	}

	for _, getter := range getters {
		filler := getter(field)
		if filler != nil {
			return filler(field)
		}
	}
	return nil
}

func (f *Filler) getFunctionByKind(field *FieldData) FillerFunc {
	if f, ok := f.FuncByKind[field.Field.Type.Kind()]; ok {
		return f
	}

	return nil
}
