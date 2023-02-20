package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

func printVaribale(scpds *scpd.SCPD, prefix string, fp *os.File) (vars map[string]string) {
	sort.SliceStable(scpds.ServiceStateTable, func(i, j int) bool {
		return strings.Compare(scpds.ServiceStateTable[i].Name, scpds.ServiceStateTable[j].Name) <= 0
	})
	vars = make(map[string]string)

	for _, s := range scpds.ServiceStateTable {
		ar := "nil"
		if s.AllowedRange != nil {
			ar = fmt.Sprintf("scpd.AllowRange{%d, %d,%d}", s.AllowedRange.Min, s.AllowedRange.Max, s.AllowedRange.Step)
		}
		al := "nil"
		if s.AllowedValues != nil {
			al = "[]string{"
			for _, r := range *s.AllowedValues {
				if strings.Contains(strings.ToLower(r), "vendor") {
					continue
				}
				al += `"` + r + `",`
			}
			al += "}"
		}
		if s.Default == "NOT_IMPLEMENTED" {
			s.Default = ""
		}
		if s.SendEvents == "yes" {
			s.SendEvents = "BoolYes"
		} else if s.SendEvents == "no" {
			s.SendEvents = "BoolNo"
		}

		if s.DataType == "string" {
			vars[s.Name] = "string"
			s.DataType = "DataTypeStr"
		} else if s.DataType == "ui4" {
			vars[s.Name] = "uint32"
			s.DataType = "DataTypeUint32"
		} else if s.DataType == "i4" {
			vars[s.Name] = "int32"
			s.DataType = "DataTypeInt32"
		} else if s.DataType == "i2" {
			vars[s.Name] = "int16"
			s.DataType = "DataTypeInt16"
		} else if s.DataType == "ui2" {
			vars[s.Name] = "uint16"
			s.DataType = "DataTypeUInt16"
		} else if s.DataType == "i1" {
			vars[s.Name] = "int8"
			s.DataType = "DataTypeInt8"
		} else if s.DataType == "ui1" {
			vars[s.Name] = "uint8"
			s.DataType = "DataTypeUInt8"
		} else if s.DataType == "boolean" {
			vars[s.Name] = "bool"
			s.DataType = "DataTypeBool"
		}
		def := ""
		if s.Default != "" {
			def = fmt.Sprintf("Default:\"%s\",\n", s.Default)
		}
		if al != "nil" {
			al = fmt.Sprintf("AllowedValues:&%s,\n", al)
		} else {
			al = ""
		}
		if ar != "nil" {
			ar = fmt.Sprintf("AllowedRange:&%s,\n", ar)
		} else {
			ar = ""
		}

		fp.WriteString(fmt.Sprintf(`var %s_%s = scpd.Variable{
	Name:"%s",
	DataType:%s,
	SendEvents:%s,
%s%s%s
}
`, prefix, s.Name, s.Name, s.DataType, s.SendEvents, def, al, ar))
	}

	return
}

func printActions(scpds *scpd.SCPD, srvName, prefix string, fp *os.File, vars map[string]string) {
	sort.SliceStable(scpds.ActionList, func(i, j int) bool {
		return strings.Compare(scpds.ActionList[i].Name, scpds.ActionList[j].Name) <= 0
	})
	actions := ""
	for _, s := range scpds.ActionList {
		args := ""
		argStructIn := fmt.Sprintf("XMLName xml.Name `xml:\"%s %s\"`\n", soap.ActionNS, s.Name)
		argStructOut := fmt.Sprintf("XMLName xml.Name `xml:\"u:%sResponse\"`\nXMLPrefix string `xml:\"xmlns:u,attr\"`\n", s.Name)
		for _, sa := range s.Arguments {
			// sa.Direction = `"` + sa.Direction + `"`
			if sa.Direction == "in" {
				sa.Direction = "DirIn"
				argStructIn += fmt.Sprintf("%s %s\n", sa.Name, vars[sa.RelatedStateVar])
			} else if sa.Direction == "out" {
				sa.Direction = "DirOut"
				argStructOut += fmt.Sprintf("%s %s\n", sa.Name, vars[sa.RelatedStateVar])
			}

			args += fmt.Sprintf("{\"%s\", %s, &%s_%s},\n", sa.Name, sa.Direction, prefix, sa.RelatedStateVar)
		}

		fp.WriteString(fmt.Sprintf("type %sArgIn_%s struct{\n%s}\n", prefix, s.Name, argStructIn))
		fp.WriteString(fmt.Sprintf("type %sArgOut_%s struct{\n%s}\n", prefix, s.Name, argStructOut))

		fp.WriteString(fmt.Sprintf("var %[1]s_%[2]s = Action{\nHandler:nil,\nArgIn: %[1]sArgIn_%[2]s{},\nArgOut: %[1]sArgOut_%[2]s{XMLPrefix:soap.ActionNS},\narguments:[]Argument{\n%[3]s\n},\n}\n", prefix, s.Name, args))

		actions += fmt.Sprintf("\"%s\": &%s_%s,\n", s.Name, prefix, s.Name)
	}
	fp.WriteString(fmt.Sprintf("var %sV1= ActionMap{\n%s}\n", srvName, actions))
}

func printArguments(scpds *scpd.SCPD) {
	argsMap := map[string]scpd.Argument{}
	for _, s := range scpds.ActionList {
		for _, r := range s.Arguments {
			if a, ok := argsMap[r.Name]; ok && a.Direction != r.Direction {
				fmt.Println("action", s.Name, "arg", a.Name)
				panic("direction ")
			}
			argsMap[r.Name] = r
		}
	}
	args := []scpd.Argument{}
	for _, s := range argsMap {
		args = append(args, s)
	}
	sort.SliceStable(args, func(i, j int) bool {
		return strings.Compare(args[i].Name, args[j].Name) <= 0
	})
	for _, arg := range args {
		if arg.Direction == "in" {
			arg.Direction = "DirIn"
		} else if arg.Direction == "out" {
			arg.Direction = "DirOut"
		}
		fmt.Printf("var Arg%s = Argument{\nName:\"%s\",\nDirection:%s,\nRelatedStateVar:&%s,\n}\n", arg.Name, arg.Name, arg.Direction, arg.RelatedStateVar)
	}
}
func main() {
	var version uint64
	flag.Uint64Var(&version, "v", 1, "")
	flag.Parse()

	if version < 1 {
		version = 1
	}

	files := map[string]string{
		"AVT": "AVTransport",
		"CM":  "ConnectionManager",
		"RC":  "RenderingControl",
	}

	regexpRemoveVendor := regexp.MustCompile(`(?i) *<allowedValue> *vendor.*?\n`)
	regexpReplaceVendor := regexp.MustCompile(`(?i)(.+?>).*?Vendor.+?(<.+?)\n`)

	for p, f := range files {
		var scpd scpd.SCPD

		fileXMLBytes, err := os.ReadFile(fmt.Sprintf("%sV%d.xml", f, version))
		if err != nil {
			fmt.Println("file", f)
			panic(err)
		}

		fileXMLBytes = regexpRemoveVendor.ReplaceAll(fileXMLBytes, []byte{})
		fileXMLBytes = regexpReplaceVendor.ReplaceAll(fileXMLBytes, []byte("$1 0 $2"))

		err = xml.Unmarshal(fileXMLBytes, &scpd)
		if err != nil {
			fmt.Println("file", f)
			panic(err)
		}

		file := fmt.Sprintf("../service_%s.go", f)
		fp, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0)
		if err != nil {
			fmt.Println("file", file)
			panic(err)
		}
		fp.Truncate(0)
		fp.WriteString(`package upnp
import (
	"encoding/xml"
	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

`)

		vars := printVaribale(&scpd, p, fp)
		printActions(&scpd, f, p, fp, vars)

		fp.Close()

		err = exec.Command("gofmt", "-e", "-w", file).Start()
		if err != nil {
			fmt.Println(err)
		}
	}

}
