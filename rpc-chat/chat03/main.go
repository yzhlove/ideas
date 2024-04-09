package main

import (
	"bytes"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"strings"
)

var buf *bytes.Buffer

func init() {
	buf = bytes.NewBuffer([]byte{})
}

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			generateGoFile(plugin, f)
		}
		return nil
	})

	if buf.Len() > 0 {
		f, err := os.OpenFile("proto.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
			return
		}
		defer func() {
			f.Sync()
			f.Close()
		}()

		f.Write(buf.Bytes())
	}
}

func generateGoFile(plugin *protogen.Plugin, f *protogen.File) {
	name := f.GeneratedFilenamePrefix + ".to_yzh_struct.go"
	log("GeneratedFilenamePrefix:", f.GeneratedFilenamePrefix)
	log("ImportPath:", f.GoImportPath.String())
	log("Ident:", f.GoDescriptorIdent.String())
	log("PackageName:", f.GoPackageName)

	gen := plugin.NewGeneratedFile(name, f.GoImportPath)
	gen.P("package ", f.GoPackageName)
	gen.P()

	for _, m := range f.Messages {
		log("GoIdent --> ", m.GoIdent)
		gen.P("type ", m.GoIdent, " struct { //")
		for _, field := range m.Fields {
			log("field.Comments.Leading.String() => ", field.Comments.Leading.String())
			log("LeadingDetached ==> ")
			for _, d := range field.Comments.LeadingDetached {
				log("LeadingDetached.String() ==> ", d.String())
			}
			log("field.Comments.Trailing.String() => ", field.Comments.Trailing.String())

			content := fmt.Sprintf("%s %s `json:\"%s\"` %s",
				field.GoName,
				field.Desc.Kind(),
				field.Desc.JSONName(),
				field.Comments.Leading.String(),
			)

			gen.P(field.Comments.Trailing.String())
			gen.P(content)
		}
		gen.P("}")
	}
	gen.P()
}

func log(s ...any) {
	var content []string
	for _, k := range s {
		content = append(content, fmt.Sprint(k))
	}
	buf.WriteString(strings.Join(content, " "))
	buf.WriteString("\n")
}
