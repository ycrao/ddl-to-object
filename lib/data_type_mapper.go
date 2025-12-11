package lib

import (
	"fmt"
	"strings"
)

// MapToGoType map mysql data type to golang data type
func MapToGoType(attr AdditionalAttr) (string, string) {
	dataType := attr.DataType
	nullable := attr.Nullable

	// Handle nullable types with proper Go types
	if IsInteger(dataType) {
		isUnsigned := attr.IsUnsigned
		prefix := ""
		if isUnsigned {
			prefix = "u"
		}

		var baseType string
		switch dataType {
		case TinyIntType:
			baseType = prefix + "int8"
		case SmallIntType:
			baseType = prefix + "int16"
		case IntegerType, IntType:
			baseType = prefix + "int32"
		case MediumIntType, BigIntType:
			baseType = prefix + "int64"
		default:
			baseType = prefix + "int64"
		}

		if nullable {
			return fmt.Sprintf("sql.Null%s", strings.Title(baseType)), `import "database/sql"`
		}
		return baseType, ""

	} else if IsHighPrecisionNumber(dataType) {
		var baseType string
		switch dataType {
		case FloatType:
			baseType = "float32"
		case DoubleType, DecimalType:
			baseType = "float64"
		default:
			baseType = "float64"
		}

		if nullable {
			return fmt.Sprintf("sql.Null%s", strings.Title(baseType)), `import "database/sql"`
		}
		return baseType, ""

	} else if IsDateAndTime(dataType) {
		switch dataType {
		case DateType, DateTimeType, TimestampType:
			if nullable {
				return `sql.NullTime`, `import "database/sql"`
			}
			return `time.Time`, `import "time"`
		case TimeType:
			if nullable {
				return `sql.NullString`, `import "database/sql"`
			}
			return "string", ""
		default:
			if nullable {
				return `sql.NullString`, `import "database/sql"`
			}
			return "string", ""
		}
	} else if dataType == JsonType {
		return "json.RawMessage", `import "encoding/json"`
	} else if IsBlob(dataType) || IsBinary(dataType) {
		return "[]byte", ""
	} else if CanCastAsString(dataType) {
		if nullable {
			return `sql.NullString`, `import "database/sql"`
		}
		return "string", ""
	}

	// Fallback for unknown types
	return "interface{}", ""
}

// MapToJavaType map mysql data type to java type
func MapToJavaType(attr AdditionalAttr) (string, string) {
	dataType := attr.DataType
	if IsBlob(dataType) || IsBinary(dataType) || dataType == BitType {
		return "byte[]", ""
	} else if dataType == BigIntType {
		return "Long", ""
	} else if IsInteger(dataType) == true {
		return "Integer", ""
	} else if IsDateAndTime(dataType) == true {
		switch dataType {
		case DateType:
			return "Date", "import java.sql.Date;"
		case DateTimeType, TimestampType:
			return "Timestamp", "import java.sql.Timestamp;"
		case TimeType:
			return "Time", "import java.sql.Time;"
		default:
			return "String", ""
		}
	} else if IsHighPrecisionNumber(dataType) == true {
		switch dataType {
		case FloatType:
			return "Float", ""
		case DoubleType:
			return "Double", ""
		case DecimalType:
			return "BigDecimal", "import java.math.BigDecimal;"
		default:
			return "Double", ""
		}
	}
	return "String", ""
}

// MapToPhpType map mysql data type to php data type (only for comment)
func MapToPhpType(attr AdditionalAttr) (string, string) {
	dataType := attr.DataType
	if dataType == JsonType {
		return "array|object", ""
	} else if dataType == TinyIntType {
		return "boolean|integer", ""
	} else if IsInteger(dataType) == true {
		return "integer", ""
	} else if IsHighPrecisionNumber(dataType) == true {
		return "numeric", ""
	} else if CanCastAsString(dataType) == true {
		return "string", ""
	} else if IsDateAndTime(dataType) == true {
		return "mixed|string", ""
	} else {
		return "mixed", ""
	}
}

// MapToPythonType seems no need to map
func MapToPythonType(attr AdditionalAttr) (string, string) {
	return "", ""
}
