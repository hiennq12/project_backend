package dms_utils

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ScanRowsToStruct quét dữ liệu từ rows vào dest, nơi dest là một con trỏ đến slice của struct.
func ScanRowsToStruct(rows *sql.Rows, dest interface{}) error {
	destVal := reflect.ValueOf(dest)
	// check dest phai la con tro de thay doi data trong ham nay thi ben ngoai cung thay doi
	if destVal.Kind() != reflect.Ptr || destVal.IsNil() {
		return fmt.Errorf("dest must be a non-nil pointer")
	}

	//get nhieu row => nen phai la slice de luu data
	sliceVal := destVal.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a slice")
	}

	// kiem tra kieu cua dest, phai la struct vi scan vao struct ma. Vi du []Product thi elemType la Product
	elemType := sliceVal.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs")
	}

	for rows.Next() {
		//Tạo một giá trị mới của kiểu phần tử (elemType). Ví dụ, nếu phần tử là Person, dòng này sẽ tạo một đối tượng mới kiểu Person.
		elem := reflect.New(elemType).Elem()

		fields := prepareFields(elem)
		//Quét dữ liệu từ hàng hiện tại vào các field đã chuẩn bị. Nếu có lỗi, hàm trả về lỗi.
		if err := rows.Scan(fields...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		populateStruct(elem, fields)

		//Thêm phần tử elem mới vào slice sliceVal.
		sliceVal.Set(reflect.Append(sliceVal, elem))
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return nil
}

// prepareFields tạo danh sách các giá trị để quét từ cơ sở dữ liệu dựa trên kiểu của struct.
func prepareFields(elem reflect.Value) []interface{} {
	//Tạo một slice fields có kích thước bằng số lượng trường của struct, với mỗi phần tử là kiểu interface{}.
	//Slice này sẽ được sử dụng để chứa các giá trị của các trường trong struct.
	fields := make([]interface{}, elem.NumField()) //elem.NumField(): Trả về số lượng trường (field) của struct được đại diện bởi elem.
	for i := 0; i < elem.NumField(); i++ {
		//elem.Field(i).Type(): Lấy kiểu của trường thứ i trong struct.
		//fieldType là đối tượng reflect.Type đại diện cho kiểu dữ liệu của trường.
		fieldType := elem.Field(i).Type()

		//tạo ra một giá trị mặc định tương ứng với kiểu dữ liệu của trường.
		fields[i] = createDefaultValue(fieldType)
	}
	return fields
}

// createDefaultValue tạo giá trị mặc định cho các kiểu dữ liệu.
// co may case scan data null vao field string trong struct bi loi unsupported
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
	default:
		return new(interface{})
	}
	return new(interface{})
}

// populateStruct gán giá trị từ các con trỏ vào các trường của struct.
func populateStruct(elem reflect.Value, fields []interface{}) {
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			//Kiểm tra nếu trường có kiểu là số nguyên (Int, Int8, Int16, Int32, Int64).
			//fields[i].(*sql.NullInt64): Kiểm tra và ép kiểu phần tử thứ i trong fields thành con trỏ kiểu *sql.NullInt64.
			//val.Valid: Kiểm tra xem giá trị trong cơ sở dữ liệu có phải là NULL không. Nếu không phải là NULL, thì val.Valid sẽ là true.
			//field.SetInt(val.Int64): Thiết lập giá trị cho trường của struct với giá trị nguyên (Int64) từ cơ sở dữ liệu.
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

// GetQueryColumnList
func GetQueryColumnList(ignoreColumns []string, obj interface{}) string {
	columns := parseColumns(obj)
	if len(ignoreColumns) > 0 {
		var newCol = make([]string, 0)
		for _, col := range columns {
			if !checkStringInSlice(col, ignoreColumns) {
				newCol = append(newCol, col)
			}
		}
		return strings.Join(newCol, ",")
	}
	return strings.Join(columns, ",")
}

func checkStringInSlice(str string, sliceStr []string) bool {
	for _, val := range sliceStr {
		if str == val {
			return true
		}
	}
	return false
}

// parse columns
func parseColumns(i interface{}) []string {
	columns := make([]string, 0)
	t := reflect.TypeOf(i).Elem()
	for index := 0; index < t.NumField(); index++ {
		f := t.Field(index)
		name := f.Name
		typeV := f.Type
		kind := f.Type.Kind()
		//jsonKey := getKeyFromJSONTag(f.Tag.Get("json"))
		//if jsonKey == "-" || utils.IsIgnoreThisField(name) {
		//	continue
		//}
		//st.addColumnFieldNameAndKind(jsonKey, name, kind, typeV)

		fmt.Println("+++++++++Test data: ", name, typeV, kind)
	}
	return columns
}
