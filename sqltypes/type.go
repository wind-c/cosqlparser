/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqltypes

import (
	"fmt"
)

type Flag int32

const (
	Flag_NONE       Flag = 0
	Flag_ISINTEGRAL Flag = 256
	Flag_ISUNSIGNED Flag = 512
	Flag_ISFLOAT    Flag = 1024
	Flag_ISQUOTED   Flag = 2048
	Flag_ISTEXT     Flag = 4096
	Flag_ISBINARY   Flag = 8192
)

// This file provides wrappers and support
// functions for Type.

// These bit flags can be used to query on the
// common properties of types.
const (
	flagIsIntegral = int(Flag_ISINTEGRAL)
	flagIsUnsigned = int(Flag_ISUNSIGNED)
	flagIsFloat    = int(Flag_ISFLOAT)
	flagIsQuoted   = int(Flag_ISQUOTED)
	flagIsText     = int(Flag_ISTEXT)
	flagIsBinary   = int(Flag_ISBINARY)
)

const (
	TimestampFormat           = "2006-01-02 15:04:05"
	TimestampFormatPrecision3 = "2006-01-02 15:04:05.000"
	TimestampFormatPrecision6 = "2006-01-02 15:04:05.000000"
)

type Type int32

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return Type_name[int32(x)]
}

// IsIntegral returns true if Type is an integral
// (signed/unsigned) that can be represented using
// up to 64 binary bits.
// If you have a Value object, use its member function.
func IsIntegral(t Type) bool {
	return int(t)&flagIsIntegral == flagIsIntegral
}

// IsSigned returns true if Type is a signed integral.
// If you have a Value object, use its member function.
func IsSigned(t Type) bool {
	return int(t)&(flagIsIntegral|flagIsUnsigned) == flagIsIntegral
}

// IsUnsigned returns true if Type is an unsigned integral.
// Caution: this is not the same as !IsSigned.
// If you have a Value object, use its member function.
func IsUnsigned(t Type) bool {
	return int(t)&(flagIsIntegral|flagIsUnsigned) == flagIsIntegral|flagIsUnsigned
}

// IsFloat returns true is Type is a floating point.
// If you have a Value object, use its member function.
func IsFloat(t Type) bool {
	return int(t)&flagIsFloat == flagIsFloat
}

// IsQuoted returns true if Type is a quoted text or binary.
// If you have a Value object, use its member function.
func IsQuoted(t Type) bool {
	return (int(t)&flagIsQuoted == flagIsQuoted) && t != Bit
}

// IsText returns true if Type is a text.
// If you have a Value object, use its member function.
func IsText(t Type) bool {
	return int(t)&flagIsText == flagIsText
}

// IsBinary returns true if Type is a binary.
// If you have a Value object, use its member function.
func IsBinary(t Type) bool {
	return int(t)&flagIsBinary == flagIsBinary
}

// IsNumber returns true if the type is any type of number.
func IsNumber(t Type) bool {
	return IsIntegral(t) || IsFloat(t) || t == Decimal
}

// IsDate returns true if the type represents a date and/or time.
func IsDate(t Type) bool {
	return t == Datetime || t == Date || t == Timestamp || t == Time
}

// IsNull returns true if the type is NULL type
func IsNull(t Type) bool {
	return t == Null
}

const (
	// NULL_TYPE specifies a NULL type.
	Type_NULL_TYPE Type = 0
	// INT8 specifies a TINYINT type.
	// Properties: 1, IsNumber.
	Type_INT8 Type = 257
	// UINT8 specifies a TINYINT UNSIGNED type.
	// Properties: 2, IsNumber, IsUnsigned.
	Type_UINT8 Type = 770
	// INT16 specifies a SMALLINT type.
	// Properties: 3, IsNumber.
	Type_INT16 Type = 259
	// UINT16 specifies a SMALLINT UNSIGNED type.
	// Properties: 4, IsNumber, IsUnsigned.
	Type_UINT16 Type = 772
	// INT24 specifies a MEDIUMINT type.
	// Properties: 5, IsNumber.
	Type_INT24 Type = 261
	// UINT24 specifies a MEDIUMINT UNSIGNED type.
	// Properties: 6, IsNumber, IsUnsigned.
	Type_UINT24 Type = 774
	// INT32 specifies a INTEGER type.
	// Properties: 7, IsNumber.
	Type_INT32 Type = 263
	// UINT32 specifies a INTEGER UNSIGNED type.
	// Properties: 8, IsNumber, IsUnsigned.
	Type_UINT32 Type = 776
	// INT64 specifies a BIGINT type.
	// Properties: 9, IsNumber.
	Type_INT64 Type = 265
	// UINT64 specifies a BIGINT UNSIGNED type.
	// Properties: 10, IsNumber, IsUnsigned.
	Type_UINT64 Type = 778
	// FLOAT32 specifies a FLOAT type.
	// Properties: 11, IsFloat.
	Type_FLOAT32 Type = 1035
	// FLOAT64 specifies a DOUBLE or REAL type.
	// Properties: 12, IsFloat.
	Type_FLOAT64 Type = 1036
	// TIMESTAMP specifies a TIMESTAMP type.
	// Properties: 13, IsQuoted.
	Type_TIMESTAMP Type = 2061
	// DATE specifies a DATE type.
	// Properties: 14, IsQuoted.
	Type_DATE Type = 2062
	// TIME specifies a TIME type.
	// Properties: 15, IsQuoted.
	Type_TIME Type = 2063
	// DATETIME specifies a DATETIME type.
	// Properties: 16, IsQuoted.
	Type_DATETIME Type = 2064
	// YEAR specifies a YEAR type.
	// Properties: 17, IsNumber, IsUnsigned.
	Type_YEAR Type = 785
	// DECIMAL specifies a DECIMAL or NUMERIC type.
	// Properties: 18, None.
	Type_DECIMAL Type = 18
	// TEXT specifies a TEXT type.
	// Properties: 19, IsQuoted, IsText.
	Type_TEXT Type = 6163
	// BLOB specifies a BLOB type.
	// Properties: 20, IsQuoted, IsBinary.
	Type_BLOB Type = 10260
	// VARCHAR specifies a VARCHAR type.
	// Properties: 21, IsQuoted, IsText.
	Type_VARCHAR Type = 6165
	// VARBINARY specifies a VARBINARY type.
	// Properties: 22, IsQuoted, IsBinary.
	Type_VARBINARY Type = 10262
	// CHAR specifies a CHAR type.
	// Properties: 23, IsQuoted, IsText.
	Type_CHAR Type = 6167
	// BINARY specifies a BINARY type.
	// Properties: 24, IsQuoted, IsBinary.
	Type_BINARY Type = 10264
	// BIT specifies a BIT type.
	// Properties: 25, IsQuoted.
	Type_BIT Type = 2073
	// ENUM specifies an ENUM type.
	// Properties: 26, IsQuoted.
	Type_ENUM Type = 2074
	// SET specifies a SET type.
	// Properties: 27, IsQuoted.
	Type_SET Type = 2075
	// TUPLE specifies a tuple. This cannot
	// be returned in a QueryResult, but it can
	// be sent as a bind var.
	// Properties: 28, None.
	Type_TUPLE Type = 28
	// GEOMETRY specifies a GEOMETRY type.
	// Properties: 29, IsQuoted.
	Type_GEOMETRY Type = 2077
	// JSON specifies a JSON type.
	// Properties: 30, IsQuoted.
	Type_JSON Type = 2078
	// EXPRESSION specifies a SQL expression.
	// This type is for internal use only.
	// Properties: 31, None.
	Type_EXPRESSION Type = 31
	// HEXNUM specifies a HEXNUM type (unquoted varbinary).
	// Properties: 32, IsText.
	Type_HEXNUM Type = 4128
	// HEXVAL specifies a HEXVAL type (unquoted varbinary).
	// Properties: 33, IsText.
	Type_HEXVAL Type = 4129
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:     "NULL_TYPE",
		257:   "INT8",
		770:   "UINT8",
		259:   "INT16",
		772:   "UINT16",
		261:   "INT24",
		774:   "UINT24",
		263:   "INT32",
		776:   "UINT32",
		265:   "INT64",
		778:   "UINT64",
		1035:  "FLOAT32",
		1036:  "FLOAT64",
		2061:  "TIMESTAMP",
		2062:  "DATE",
		2063:  "TIME",
		2064:  "DATETIME",
		785:   "YEAR",
		18:    "DECIMAL",
		6163:  "TEXT",
		10260: "BLOB",
		6165:  "VARCHAR",
		10262: "VARBINARY",
		6167:  "CHAR",
		10264: "BINARY",
		2073:  "BIT",
		2074:  "ENUM",
		2075:  "SET",
		28:    "TUPLE",
		2077:  "GEOMETRY",
		2078:  "JSON",
		31:    "EXPRESSION",
		4128:  "HEXNUM",
		4129:  "HEXVAL",
	}
	Type_value = map[string]int32{
		"NULL_TYPE":  0,
		"INT8":       257,
		"UINT8":      770,
		"INT16":      259,
		"UINT16":     772,
		"INT24":      261,
		"UINT24":     774,
		"INT32":      263,
		"UINT32":     776,
		"INT64":      265,
		"UINT64":     778,
		"FLOAT32":    1035,
		"FLOAT64":    1036,
		"TIMESTAMP":  2061,
		"DATE":       2062,
		"TIME":       2063,
		"DATETIME":   2064,
		"YEAR":       785,
		"DECIMAL":    18,
		"TEXT":       6163,
		"BLOB":       10260,
		"VARCHAR":    6165,
		"VARBINARY":  10262,
		"CHAR":       6167,
		"BINARY":     10264,
		"BIT":        2073,
		"ENUM":       2074,
		"SET":        2075,
		"TUPLE":      28,
		"GEOMETRY":   2077,
		"JSON":       2078,
		"EXPRESSION": 31,
		"HEXNUM":     4128,
		"HEXVAL":     4129,
	}
)

// Vitess data types. These are idiomatically named synonyms for the Type values.
// Although these constants are interchangeable, they should be treated as different from Type.
// Use the synonyms only to refer to the type in Value. For proto variables, use the Type constants instead.
// The following is a complete listing of types that match each classification function in this API:
//
//    IsSigned(): INT8, INT16, INT24, INT32, INT64
//    IsFloat(): FLOAT32, FLOAT64
//    IsUnsigned(): UINT8, UINT16, UINT24, UINT32, UINT64, YEAR
//    IsIntegral(): INT8, UINT8, INT16, UINT16, INT24, UINT24, INT32, UINT32, INT64, UINT64, YEAR
//    IsText(): TEXT, VARCHAR, CHAR, HEXNUM, HEXVAL
//    IsNumber(): INT8, UINT8, INT16, UINT16, INT24, UINT24, INT32, UINT32, INT64, UINT64, FLOAT32, FLOAT64, YEAR, DECIMAL
//    IsQuoted(): TIMESTAMP, DATE, TIME, DATETIME, TEXT, BLOB, VARCHAR, VARBINARY, CHAR, BINARY, ENUM, SET, GEOMETRY, JSON
//    IsBinary(): BLOB, VARBINARY, BINARY
//    IsDate(): TIMESTAMP, DATE, TIME, DATETIME
//    IsNull(): NULL_TYPE
//
// TODO(sougou): provide a categorization function
// that returns enums, which will allow for cleaner
// switch statements for those who want to cover types
// by their category.
const (
	Null       = Type_NULL_TYPE
	Int8       = Type_INT8
	Uint8      = Type_UINT8
	Int16      = Type_INT16
	Uint16     = Type_UINT16
	Int24      = Type_INT24
	Uint24     = Type_UINT24
	Int32      = Type_INT32
	Uint32     = Type_UINT32
	Int64      = Type_INT64
	Uint64     = Type_UINT64
	Float32    = Type_FLOAT32
	Float64    = Type_FLOAT64
	Timestamp  = Type_TIMESTAMP
	Date       = Type_DATE
	Time       = Type_TIME
	Datetime   = Type_DATETIME
	Year       = Type_YEAR
	Decimal    = Type_DECIMAL
	Text       = Type_TEXT
	Blob       = Type_BLOB
	VarChar    = Type_VARCHAR
	VarBinary  = Type_VARBINARY
	Char       = Type_CHAR
	Binary     = Type_BINARY
	Bit        = Type_BIT
	Enum       = Type_ENUM
	Set        = Type_SET
	Geometry   = Type_GEOMETRY
	TypeJSON   = Type_JSON
	Expression = Type_EXPRESSION
	HexNum     = Type_HEXNUM
	HexVal     = Type_HEXVAL
	Tuple      = Type_TUPLE
)

// bit-shift the mysql flags by two byte so we
// can merge them with the mysql or vitess types.
const (
	mysqlUnsigned = 32
	mysqlBinary   = 128
	mysqlEnum     = 256
	mysqlSet      = 2048
)

// If you add to this map, make sure you add a test case
// in tabletserver/endtoend.
var mysqlToType = map[int64]Type{
	0:   Decimal,
	1:   Int8,
	2:   Int16,
	3:   Int32,
	4:   Float32,
	5:   Float64,
	6:   Null,
	7:   Timestamp,
	8:   Int64,
	9:   Int24,
	10:  Date,
	11:  Time,
	12:  Datetime,
	13:  Year,
	15:  VarChar,
	16:  Bit,
	17:  Timestamp,
	18:  Datetime,
	19:  Time,
	245: TypeJSON,
	246: Decimal,
	247: Enum,
	248: Set,
	249: Text,
	250: Text,
	251: Text,
	252: Text,
	253: VarChar,
	254: Char,
	255: Geometry,
}

// modifyType modifies the vitess type based on the
// mysql flag. The function checks specific flags based
// on the type. This allows us to ignore stray flags
// that MySQL occasionally sets.
func modifyType(typ Type, flags int64) Type {
	switch typ {
	case Int8:
		if flags&mysqlUnsigned != 0 {
			return Uint8
		}
	case Int16:
		if flags&mysqlUnsigned != 0 {
			return Uint16
		}
	case Int32:
		if flags&mysqlUnsigned != 0 {
			return Uint32
		}
	case Int64:
		if flags&mysqlUnsigned != 0 {
			return Uint64
		}
	case Int24:
		if flags&mysqlUnsigned != 0 {
			return Uint24
		}
	case Text:
		if flags&mysqlBinary != 0 {
			return Blob
		}
	case VarChar:
		if flags&mysqlBinary != 0 {
			return VarBinary
		}
	case Char:
		if flags&mysqlBinary != 0 {
			return Binary
		}
		if flags&mysqlEnum != 0 {
			return Enum
		}
		if flags&mysqlSet != 0 {
			return Set
		}
	case Year:
		if flags&mysqlBinary != 0 {
			return VarBinary
		}
	}
	return typ
}

// MySQLToType computes the vitess type from mysql type and flags.
func MySQLToType(mysqlType, flags int64) (typ Type, err error) {
	result, ok := mysqlToType[mysqlType]
	if !ok {
		return 0, fmt.Errorf("unsupported type: %d", mysqlType)
	}
	return modifyType(result, flags), nil
}

// AreTypesEquivalent returns whether two types are equivalent.
func AreTypesEquivalent(mysqlTypeFromBinlog, mysqlTypeFromSchema Type) bool {
	return (mysqlTypeFromBinlog == mysqlTypeFromSchema) ||
		(mysqlTypeFromBinlog == VarChar && mysqlTypeFromSchema == VarBinary) ||
		// Binlog only has base type. But doesn't have per-column-flags to differentiate
		// various logical types. For Binary, Enum, Set types, binlog only returns Char
		// as data type.
		(mysqlTypeFromBinlog == Char && mysqlTypeFromSchema == Binary) ||
		(mysqlTypeFromBinlog == Char && mysqlTypeFromSchema == Enum) ||
		(mysqlTypeFromBinlog == Char && mysqlTypeFromSchema == Set) ||
		(mysqlTypeFromBinlog == Text && mysqlTypeFromSchema == Blob) ||
		(mysqlTypeFromBinlog == Int8 && mysqlTypeFromSchema == Uint8) ||
		(mysqlTypeFromBinlog == Int16 && mysqlTypeFromSchema == Uint16) ||
		(mysqlTypeFromBinlog == Int24 && mysqlTypeFromSchema == Uint24) ||
		(mysqlTypeFromBinlog == Int32 && mysqlTypeFromSchema == Uint32) ||
		(mysqlTypeFromBinlog == Int64 && mysqlTypeFromSchema == Uint64)
}

// typeToMySQL is the reverse of mysqlToType.
var typeToMySQL = map[Type]struct {
	typ   int64
	flags int64
}{
	Int8:      {typ: 1},
	Uint8:     {typ: 1, flags: mysqlUnsigned},
	Int16:     {typ: 2},
	Uint16:    {typ: 2, flags: mysqlUnsigned},
	Int32:     {typ: 3},
	Uint32:    {typ: 3, flags: mysqlUnsigned},
	Float32:   {typ: 4},
	Float64:   {typ: 5},
	Null:      {typ: 6, flags: mysqlBinary},
	Timestamp: {typ: 7},
	Int64:     {typ: 8},
	Uint64:    {typ: 8, flags: mysqlUnsigned},
	Int24:     {typ: 9},
	Uint24:    {typ: 9, flags: mysqlUnsigned},
	Date:      {typ: 10, flags: mysqlBinary},
	Time:      {typ: 11, flags: mysqlBinary},
	Datetime:  {typ: 12, flags: mysqlBinary},
	Year:      {typ: 13, flags: mysqlUnsigned},
	Bit:       {typ: 16, flags: mysqlUnsigned},
	TypeJSON:  {typ: 245},
	Decimal:   {typ: 246},
	Text:      {typ: 252},
	Blob:      {typ: 252, flags: mysqlBinary},
	VarChar:   {typ: 253},
	VarBinary: {typ: 253, flags: mysqlBinary},
	Char:      {typ: 254},
	Binary:    {typ: 254, flags: mysqlBinary},
	Enum:      {typ: 254, flags: mysqlEnum},
	Set:       {typ: 254, flags: mysqlSet},
	Geometry:  {typ: 255},
}

// TypeToMySQL returns the equivalent mysql type and flag for a vitess type.
func TypeToMySQL(typ Type) (mysqlType, flags int64) {
	val := typeToMySQL[typ]
	return val.typ, val.flags
}
