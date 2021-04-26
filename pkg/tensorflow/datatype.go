package tensorflow

import (
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type DataType tf.DataType

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	Float      DataType = DataType(tf.Float)
	Double     DataType = DataType(tf.Double)
	Int32      DataType = DataType(tf.Int32)
	Uint32     DataType = DataType(tf.Uint32)
	Uint8      DataType = DataType(tf.Uint8)
	Int16      DataType = DataType(tf.Int16)
	Int8       DataType = DataType(tf.Int8)
	String     DataType = DataType(tf.String)
	Complex64  DataType = DataType(tf.Complex64)
	Int64      DataType = DataType(tf.Int64)
	Uint64     DataType = DataType(tf.Uint64)
	Bool       DataType = DataType(tf.Bool)
	Qint8      DataType = DataType(tf.Qint8)
	Quint8     DataType = DataType(tf.Quint8)
	Qint32     DataType = DataType(tf.Qint32)
	Bfloat16   DataType = DataType(tf.Bfloat16)
	Qint16     DataType = DataType(tf.Qint16)
	Quint16    DataType = DataType(tf.Quint16)
	Uint16     DataType = DataType(tf.Uint16)
	Complex128 DataType = DataType(tf.Complex128)
	Half       DataType = DataType(tf.Half)
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (t DataType) String() string {
	switch tf.DataType(t) {
	case tf.Float:
		return "float"
	case tf.Double:
		return "double"
	case tf.Int32:
		return "int32"
	case tf.Uint32:
		return "uint32"
	case tf.Uint8:
		return "uint8"
	case tf.Int16:
		return "int16"
	case tf.Int8:
		return "int8"
	case tf.String:
		return "string"
	case tf.Complex64:
		return "complex64"
	case tf.Int64:
		return "int64"
	case tf.Uint64:
		return "uint64"
	case tf.Bool:
		return "bool"
	case tf.Qint8:
		return "qint8"
	case tf.Quint8:
		return "quint8"
	case tf.Qint32:
		return "qint32"
	case tf.Bfloat16:
		return "bfloat16"
	case tf.Qint16:
		return "qint16"
	case tf.Quint16:
		return "quint16"
	case tf.Uint16:
		return "uint16"
	case tf.Complex128:
		return "complex128"
	case tf.Half:
		return "half"
	default:
		return "??"
	}
}
