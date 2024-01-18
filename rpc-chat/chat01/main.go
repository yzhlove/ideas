package main

import (
	"fmt"
	"github.com/jhump/protoreflect/desc/protoparse"
	"log"
)

/*
FileDescriptor:
	GetMessageTypes() []*MessageDescriptor			// 获取文件内所有的 message 类型
	GetEnumTypes() []*EnumDescriptor				// 获取文件内所有的 enum 类型
	FindMessage(msgName string) *MessageDescriptor	// 根据名字来获取 message 类型
	FindEnum(enumName string) *EnumDescriptor		// 根据名字来获取 enum 类型

MessageDescriptor:
	GetName() string								// message 的名字
	GetFullyQualifiedName() string					// message 的全限定名
	AsDescriptorProto() *descriptor.DescriptorProto	// AsDescriptorProto returns the underlying descriptor proto
	GetFields() []*FieldDescriptor					// 获取 message 内的所有字段
	GetNestedMessageTypes() []*MessageDescriptor	// 获取在 message 内内嵌的 message
	GetNestedEnumTypes() []*EnumDescriptor
	FindFieldByName(fieldName string) *FieldDescriptor
	FindFieldByNumber(tagNumber int32) *FieldDescriptor
	FindFieldByJSONName(jsonName string) *FieldDescriptor

FieldDescriptor:
	GetName() string
	GetFullyQualifiedName() string
	AsEnumDescriptorProto() *descriptor.EnumDescriptorProto
	String() string
	GetValues() []*EnumValueDescriptor			// 获取该枚举所有的枚举值
	FindValueByName(name string) *EnumValueDescriptor
	FindValueByNumber(num int32) *EnumValueDescriptor
*/

func main() {

	var prase protoparse.Parser
	file := "/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/rpc-chat/chat01/proto/test.proto"
	ds, err := prase.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, ret := range ds {
		fmt.Println("GetName => ", ret.GetName())
		//fmt.Println("String => ", ret.String())
		//fmt.Println("AsProto => ", ret.AsProto())
		fmt.Println("FullyQualifiedName => ", ret.GetFullyQualifiedName())
		//fmt.Println("File => ", ret.GetFile())
		fmt.Println("MessageTypes => ", ret.GetMessageTypes())
		fmt.Println("EnumTypes => ", ret.GetEnumTypes())
	}

}
