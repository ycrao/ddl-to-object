package lib

// MapToGoType map mysql data type to golang data type
func MapToGoType(attr AdditionalAttr) (string, string) {
	dataType := attr.DataType
	nullable := attr.Nullable
	if IsInteger(dataType) {
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
		case IntegerType:
		case IntType:
			return prefix + "int32", ""
		case MediumIntType:
		case BigIntType:
		default:
			return prefix + "int64", ""
		}
	} else if IsHighPrecisionNumber(dataType) {
		switch dataType {
		case FloatType:
			return "float32", ""
		case DoubleType:
		case DecimalType:
		default:
			return "float64", ""
		}
	} else if IsDateAndTime(dataType) {
		switch dataType {
		case DateType:
		case DateTimeType:
			if nullable {
				return `sql.NullTime`, `import "database/sql"`
			} else {
				return `time.Time`, `import "time"`
			}
		default:
			if nullable {
				return `sql.NullString`, `import "database/sql"`
			} else {
				return "string", ""
			}
		}
	} else if CanCastAsString(dataType) {
		if nullable {
			return `sql.NullString`, `import "database/sql"`
		} else {
			return "string", ""
		}
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
	} else if IsInteger(dataType) {
		return "Integer", ""
	} else if CanCastAsString(dataType) {
		return "String", ""
	} else if IsDateAndTime(dataType) {
		switch dataType {
		case DateType:
			return "Date", "import java.sql.Date;"
		case DateTimeType:
		case TimestampType:
			return "Timestamp", "import java.sql.Timestamp;"
		case TimeType:
			return "Time", "import java.sql.Time;"
		default:
			return "String", ""
		}
	} else if IsHighPrecisionNumber(dataType) {
		switch dataType {
		case FloatType:
			return "Float", ""
		case DoubleType:
			return "Double", ""
		case DecimalType:
			return "BigDecimal", "import java.math.BigDecimal;"
		}
	}
	return "String", ""
}
