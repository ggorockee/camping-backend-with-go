package serializer

import (
	"fmt"
	"reflect"
)

func GeneralSerializer(input any, output any) error {
	// input, output의 reflect.Value 가져옴
	inputValue := reflect.ValueOf(input)
	outputValue := reflect.ValueOf(output)

	// input이 포인터면 실제 값으로 변경
	if inputValue.Kind() == reflect.Ptr {
		inputValue = inputValue.Elem()
	}

	// output 포인터 아니면 에러
	if outputValue.Kind() != reflect.Ptr {
		return fmt.Errorf("output must be a pointer to a struct")
	}
	// output 실제 값으로 변경
	outputValue = outputValue.Elem()

	// output 구조체 아니면 에러
	if outputValue.Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	// output 구조체 필드 순회
	for i := 0; i < outputValue.NumField(); i++ {
		// 현재 output 필드 정보
		outputField := outputValue.Type().Field(i)
		// 같은 이름의 input 필드 찾기
		inputField := inputValue.FieldByName(outputField.Name)

		// input 필드 유효하고 output 필드 설정 가능하면
		if inputField.IsValid() && outputValue.Field(i).CanSet() {
			switch {
			// 타입 같으면 그대로 설정
			case outputField.Type == inputField.Type():
				outputValue.Field(i).Set(inputField)
			// output int, input uint면 변환해서 설정
			case outputField.Type.Kind() == reflect.Int && inputField.Kind() == reflect.Uint:
				outputValue.Field(i).SetInt(int64(inputField.Uint()))
			// output uint, input int면 변환해서 설정
			case outputField.Type.Kind() == reflect.Uint && inputField.Kind() == reflect.Int:
				outputValue.Field(i).SetUint(uint64(inputField.Int()))
			// 그 외
			default:
				// 할당 가능하면 설정
				if outputValue.Field(i).Type().AssignableTo(inputField.Type()) {
					outputValue.Field(i).Set(inputField)
				} else {
					// 타입 안 맞으면 에러
					return fmt.Errorf("cannot assign field %s: type mismatch", outputField.Name)
				}
			}
		}
	}

	// 다 잘되면 nil 반환
	return nil
}
