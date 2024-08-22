package dms_utils

import (
	"database/sql"
	"fmt"
	log2 "github.com/hiennq12/project_backend/utils/log"
	"reflect"
	"strings"
	"time"
)

// ScanRowToStruct : func scan rows database to struct params, use reflection
//
// The function receives two parameters:
//  1. `rows` has type *sql.Rows: rows after query database
//  2. `dest` has type interface{}: struct want to pass data into it
//
// return error if it has error during scan data
func ScanRowToStruct(rows *sql.Rows, dest interface{}) error {
	destValue := reflect.ValueOf(dest).Elem()
	destType := reflect.TypeOf(dest)
	fmt.Println(destValue, destType)

	// get name columns from rows
	columns, err := rows.Columns()
	if err != nil {
		log2.LogErrorWithLine(err)
		return err
	}

	// Tạo một slice các giá trị để lưu kết quả các cột
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{}) // Khởi tạo với con trỏ đến interface{} để lưu giá trị của mỗi cột
	}

	// Lặp qua từng hàng và ánh xạ vào struct
	for rows.Next() {
		// Đọc giá trị của hàng hiện tại vào slice values
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		// Tạo một giá trị mới của struct đích
		structValue := reflect.New(destType.Elem()).Elem()

		// Ánh xạ các giá trị cột vào các trường của struct
		for i, column := range columns {
			fieldName := strings.Title(column) // Chuyển tên cột thành dạng Title để so sánh với tên trường struct
			field := structValue.FieldByName(fieldName)
			if field.IsValid() && field.CanSet() {
				// Lấy giá trị từ slice values và xử lý theo loại dữ liệu
				value := reflect.ValueOf(values[i]).Elem().Interface()

				// Ánh xạ giá trị vào trường tương ứng
				switch field.Kind() {
				case reflect.Int:
					if v, ok := value.(int64); ok {
						field.SetInt(v)
					}
				case reflect.String:
					if v, ok := value.(string); ok {
						field.SetString(v)
					}
				case reflect.Float64:
					if v, ok := value.(float64); ok {
						field.SetFloat(v)
					}
				case reflect.Bool:
					if v, ok := value.(bool); ok {
						field.SetBool(v)
					}
				case reflect.Struct:
					if field.Type() == reflect.TypeOf(time.Time{}) {
						if v, ok := value.(time.Time); ok {
							field.Set(reflect.ValueOf(v))
						}
					}
				default:
					return fmt.Errorf("unsupported field type: %s", field.Kind())
				}
			}
		}

		// Thêm giá trị struct vào slice đích
		destValue.Set(reflect.Append(destValue, structValue))
	}

	// Kiểm tra lỗi nếu có sau khi lặp qua tất cả các hàng
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return nil
}
