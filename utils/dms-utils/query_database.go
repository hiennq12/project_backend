package dms_utils

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

// ScanRowsToStruct quét dữ liệu từ rows vào dest, nơi dest là một con trỏ đến slice của struct.
func ScanRowsToStruct(rows *sql.Rows, dest interface{}) error {
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.IsNil() {
		return fmt.Errorf("dest must be a non-nil pointer")
	}

	sliceVal := destVal.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a slice")
	}

	elemType := sliceVal.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs")
	}

	for rows.Next() {
		elem := reflect.New(elemType).Elem()
		fields := prepareFields(elem)
		if err := rows.Scan(fields...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		populateStruct(elem, fields)
		sliceVal.Set(reflect.Append(sliceVal, elem))
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return nil
}

// prepareFields tạo danh sách các giá trị để quét từ cơ sở dữ liệu dựa trên kiểu của struct.
func prepareFields(elem reflect.Value) []interface{} {
	fields := make([]interface{}, elem.NumField())
	for i := 0; i < elem.NumField(); i++ {
		fieldType := elem.Field(i).Type()
		fields[i] = createDefaultValue(fieldType)
	}
	return fields
}

// createDefaultValue tạo giá trị mặc định cho các kiểu dữ liệu.
func createDefaultValue(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return new(sql.NullInt64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return new(sql.NullInt64)
	case reflect.Float32, reflect.Float64:
		return new(sql.NullFloat64)
	case reflect.String:
		return new(sql.NullString)
	case reflect.Bool:
		return new(sql.NullBool)
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return new(time.Time)
		}
	}
	return new(interface{})
}

// populateStruct gán giá trị từ các con trỏ vào các trường của struct.
func populateStruct(elem reflect.Value, fields []interface{}) {
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val, ok := fields[i].(*sql.NullInt64); ok && val.Valid {
				field.SetInt(val.Int64)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val, ok := fields[i].(*sql.NullInt64); ok && val.Valid {
				field.SetUint(uint64(val.Int64))
			}
		case reflect.Float32, reflect.Float64:
			if val, ok := fields[i].(*sql.NullFloat64); ok && val.Valid {
				field.SetFloat(val.Float64)
			}
		case reflect.String:
			if val, ok := fields[i].(*sql.NullString); ok && val.Valid {
				field.SetString(val.String)
			}
		case reflect.Bool:
			if val, ok := fields[i].(*sql.NullBool); ok && val.Valid {
				field.SetBool(val.Bool)
			}
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				if val, ok := fields[i].(*time.Time); ok {
					field.Set(reflect.ValueOf(*val))
				}
			}
		}
	}
}
