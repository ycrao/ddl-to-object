package lib

import "strings"

const Signed = "SIGNED"
const Unsigned = "UNSIGNED"

// integer
const (
	TinyIntType       = "TINYINT"
	SmallIntType      = "SMALLINT"
	IntType           = "INT"
	IntegerType       = "INTEGER"
	MediumIntType     = "MEDIUMINT"
	BigIntType        = "BIGINT"
	IntegerSeriesType = "TINYINT|SMALLINT|INT|INTEGER|MEDIUMINT|BIGINT"
)

// float/double/decimal
const (
	FloatType   = "FLOAT"
	DoubleType  = "DOUBLE"
	DecimalType = "DECIMAL"
)

// date and time
const (
	DateType              = "DATE"
	TimeType              = "TIME"
	DateTimeType          = "DATETIME"
	TimestampType         = "TIMESTAMP"
	YearType              = "YEAR"
	DateAndTimeSeriesType = "DATE|TIME|DATETIME|TIMESTAMP|YEAR"
)

// char, string and text
const (
	CharType       = "CHAR"
	VarCharType    = "VARCHAR"
	CharSeriesType = "CHAR|VARCHAR"

	TinyTextType   = "TINYTEXT"
	TextType       = "TEXT"
	MediumTextType = "MEDIUMTEXT"
	LongTextType   = "LONGTEXT"
	TextSeriesType = "TINYTEXT|TEXT|MEDIUMTEXT|LONGTEXT"
)

// binary
const (
	BinaryType       = "BINARY"
	VarBinaryType    = "VARBINARY"
	BinarySeriesType = "BINARY|VARBINARY"
)

// blob
const (
	TypeBlob       = "BLOB"
	TypeTinyBlob   = "TINYBLOB"
	TypeMediumBlob = "MEDIUMBLOB"
	TypeLongBlob   = "LONGBLOB"
	BlobSeriesType = "TINYBLOB|BLOB|MEDIUMBLOB|LONGBLOB"
)

// others: such as bit/enum/set/json and spatial series type
const (
	BitType        = "BIT"
	EnumType       = "ENUM"
	SetType        = "SET"
	JsonType       = "JSON"
	BooleanType    = "TINYINT"
	BoolType       = "TINYINT"
	ExpressionType = "EXPRESSION"

	SpatialSeriesType = "GEOMETRY|POINT|LINESTRING|POLYGON|MULTIPOINT|MULTILINESTRING|MULTIPOLYGON|GEOMETRYCOLLECTION"
)

// IsText is text?
func IsText(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	textTypesArr := strings.Split(TextSeriesType, "|")
	for _, textType := range textTypesArr {
		if upperDataType == textType {
			return true
		}
	}
	return false
}

// IsChar is string?
func IsChar(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	charTypesArr := strings.Split(CharSeriesType, "|")
	for _, charType := range charTypesArr {
		if upperDataType == charType {
			return true
		}
	}
	return false
}

// IsBlob is blob?
func IsBlob(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	blobTypesArr := strings.Split(BlobSeriesType, "|")
	for _, blobType := range blobTypesArr {
		if upperDataType == blobType {
			return true
		}
	}
	return false
}

// IsInteger is integer
func IsInteger(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	integerTypesArr := strings.Split(IntegerSeriesType, "|")
	for _, integerType := range integerTypesArr {
		if upperDataType == integerType {
			return true
		}
	}
	return false
}

// IsDateAndTime is date or time or datetime?
func IsDateAndTime(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	dateAndTimeTypesArr := strings.Split(DateAndTimeSeriesType, "|")
	for _, dateAndTimeType := range dateAndTimeTypesArr {
		if upperDataType == dateAndTimeType {
			return true
		}
	}
	return false
}

// IsHighPrecisionNumber is float or double or decimal
func IsHighPrecisionNumber(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	return upperDataType == FloatType || upperDataType == DoubleType || upperDataType == DecimalType
}

// IsSpatialType is spatial data type
func IsSpatialType(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	spatialTypesArr := strings.Split(SpatialSeriesType, "|")
	for _, spatialType := range spatialTypesArr {
		if upperDataType == spatialType {
			return true
		}
	}
	return false
}

// IsBinary is binary
func IsBinary(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	binaryTypesArr := strings.Split(BinarySeriesType, "|")
	for _, binaryType := range binaryTypesArr {
		if upperDataType == binaryType {
			return true
		}
	}
	return false
}

// CanCastAsString can cast as string?
func CanCastAsString(dataType string) bool {
	upperDataType := strings.ToUpper(dataType)
	isText := IsText(upperDataType)
	isChar := IsChar(upperDataType)
	return isText || isChar || (upperDataType == EnumType) || (upperDataType == SetType) || (upperDataType == JsonType)
}
