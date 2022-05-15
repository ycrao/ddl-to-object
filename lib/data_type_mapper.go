package lib

// MapToGoType map mysql data type to golang data type
func MapToGoType(attr AdditionalAttr) (string, string) {
	dataType := attr.DataType
	nullable := attr.Nullable
	if IsInteger(dataType) == true {
		isUnsigned := attr.IsUnsigned
		prefix := ""
		if isUnsigned {
			prefix = "u"
		}
		switch dataType {
		case TinyIntType:
			return prefix + "int8", ""
		case SmallIntType:
			return prefix + "int16", ""
		case IntegerType, IntType:
			return prefix + "int32", ""
		case MediumIntType, BigIntType:
			return prefix + "int64", ""
		default:
			return prefix + "int64", ""
		}
	} else if IsHighPrecisionNumber(dataType) == true {
		switch dataType {
		case FloatType:
			return "float32", ""
		case DoubleType, DecimalType:
			return "float64", ""
		default:
			return "float64", ""
		}
	} else if IsDateAndTime(dataType) == true {
		switch dataType {
		case DateType, DateTimeType:
			if nullable {
				return `sql.NullTime`, `import "database/sql"`
			}
			return `time.Time`, `import "time"`
		default:
			if nullable {
				return `sql.NullString`, `import "database/sql"`
			}
			return "string", ""
		}
	} else if CanCastAsString(dataType) == true {
		if nullable {
			return `sql.NullString`, `import "database/sql"`
		}
		return "string", ""
	} else if IsBlob(dataType) || IsBinary(dataType) {
		return "[]byte", ""
	}
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
