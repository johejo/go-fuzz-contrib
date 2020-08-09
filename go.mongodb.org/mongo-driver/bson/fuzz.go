package bson

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func Fuzz(data []byte) int {
	var ok bool
	_, _, ok = bsoncore.ReadArray(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadBinary(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadBoolean(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadCodeWithScope(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadDBPointer(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadDateTime(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadDecimal128(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadDocument(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadDouble(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadElement(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadHeader(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadHeaderBytes(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadInt32(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadInt64(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadJavaScript(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadKey(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadKeyBytes(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadLength(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadObjectID(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadRegex(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadString(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadSymbol(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadTime(data)
	if !ok {
		return 0
	}
	_, _, _, ok = bsoncore.ReadTimestamp(data)
	if !ok {
		return 0
	}
	_, _, ok = bsoncore.ReadType(data)
	if !ok {
		return 0
	}
	return 1
}
