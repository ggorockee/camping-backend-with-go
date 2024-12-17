package serializer

import "reflect"

func GeneralSerializer(input any, output any) error {
	inputValue := reflect.ValueOf(input)
	outputValue := reflect.ValueOf(output).Elem()

	for i := 0; i < outputValue.NumField(); i++ {
		field := outputValue.Type().Field(i)
		if inputValue.FieldByName(field.Name).IsValid() {
			outputValue.Field(i).Set(inputValue.FieldByName(field.Name))
		}
	}

	return nil
}
