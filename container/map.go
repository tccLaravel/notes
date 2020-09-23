package container

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       []string
	Position  string
	Salary    int
	ManagerID int
}
func Map()  {
	f := map[string][]string{
		"Name":{"tcc"},
		"Address":{"Address_1"},
		"Position":{"Position_1"},
		"Salary":{"36541"},
		"DoB":{"DoB1","DoB2"},
	}
	e := &Employee{}
	MapForm(e,f)
	fmt.Printf("Employee is %+v \n",e)
	fmt.Printf("Employee.Salary is %+v,  and type is %T \n",e.Salary,e.Salary)
}

func MapForm(ptr interface{}, params map[string][]string) error {
	//获取ptr的Type
	typ := reflect.TypeOf(ptr).Elem()
	//获取ptr的Value
	val := reflect.ValueOf(ptr).Elem()
	//遍历整个ptr结构字段
	for i := 0; i < typ.NumField(); i++ {
		//类型字段
		typeField := typ.Field(i)
		//值字段
		structField := val.Field(i)
		//是否可修改,就是说是否可寻址
		if !structField.CanSet() {
			continue
		}
		//获取值类型
		structFieldKind := structField.Kind()
		//获取字段的名称
		inputFieldName := typeField.Tag.Get("form")
		if inputFieldName == "" {
			//获取字段的名称
			inputFieldName = typeField.Name
			//字段类型是结构体的话 递归继续检索
			if structFieldKind == reflect.Struct {
				err := MapForm(structField.Addr().Interface(), params)
				if err != nil {
					return err
				}
				continue
			}
		}
		//判定参数里是否选在该字段
		inputValue, exists := params[inputFieldName]
		if !exists {
			continue
		}

		numElems := len(inputValue)
		//进行字段类型转换
		if structFieldKind == reflect.Slice && numElems > 0 {
			sliceOf := structField.Type().Elem().Kind()
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
			for i := 0; i < numElems; i++ {
				if err := setWithProperType(sliceOf, inputValue[i], slice.Index(i)); err != nil {
					return err
				}
			}
			val.Field(i).Set(slice)
		} else {
			if _, isTime := structField.Interface().(time.Time); isTime {
				if err := setTimeField(inputValue[0], typeField, structField); err != nil {
					return err
				}
				continue
			}
			if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
				return err
			}
		}
	}
	return nil
}



func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return nil
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	//2018-01-02 01:02:03

	if timeFormat == "" {
		timeFormat = "2006-01-02 15:04:05"
		val = strings.Replace(val,"/","-",0)
		num := len(strings.Split(val," "))
		if num==1{
			val = val +" 00:00:00"
		}else{
			//2018-01-02 00
			num =len(strings.Split(val,":"))

			if num==1{
				val = val +":00:00"
			}else if num==2{
				val = val +":00"
			}
		}

	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}