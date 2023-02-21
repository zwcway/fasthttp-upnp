package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/mewkiz/pkg/pathutil"
	"github.com/zwcway/fasthttp-upnp/scpd"
)

func checkVaribale(scpds *scpd.SCPD) {
	sort.SliceStable(scpds.ServiceStateTable, func(i, j int) bool {
		return strings.Compare(scpds.ServiceStateTable[i].Name, scpds.ServiceStateTable[j].Name) <= 0
	})
	for _, s := range scpds.ServiceStateTable {
		if s.AllowedRange != nil {
			if s.AllowedRange.Step == 0 {
				s.AllowedRange.Step = 1
			}
		}
		if s.AllowedValues != nil {
			av := []string{}
			for _, r := range *s.AllowedValues {
				if strings.Contains(strings.ToLower(r), "vendor") {
					continue
				}
				av = append(av, r)
			}
			s.AllowedValues = &av
		}
		if s.Default == "NOT_IMPLEMENTED" {
			s.Default = ""
		}
	}

	return
}

func printArguments(scpds *scpd.SCPD, pwd, srvName string, version string) {
	file := filepath.Join(filepath.Dir(pwd), strings.ToLower(srvName)+version, "arguments.go")
	fp, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open failed", file)
		return
	}
	defer func() {
		fp.Close()
		exec.Command("gofmt", "-e", "-w", file).Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fp.Truncate(0)
	fp.WriteString(fmt.Sprintf(`package %s%s
import 	"encoding/xml"
`, strings.ToLower(srvName), version))

	sort.SliceStable(scpds.ActionList, func(i, j int) bool {
		return strings.Compare(scpds.ActionList[i].Name, scpds.ActionList[j].Name) <= 0
	})

	for _, s := range scpds.ActionList {
		sort.SliceStable(s.Arguments, func(i, j int) bool {
			return strings.Compare(s.Arguments[i].Name, s.Arguments[j].Name) <= 0
		})
		argInCode := ""
		argOutCode := ""
		for _, a := range s.Arguments {
			var sv *scpd.Variable
			for _, v := range scpds.ServiceStateTable {
				if v.Name == a.RelatedStateVar {
					sv = v
					break
				}
			}
			if sv == nil {
				fmt.Printf("error not found variable '%s'\n", a.RelatedStateVar)
				continue
			}
			an := sv.Name
			if sv.Default != "" {
				an += " " + sv.Default
			}
			at := ""
			switch sv.DataType {
			case "i1":
				at = "int8"
			case "i2":
				at = "int16"
			case "i4":
				at = "int32"
			case "ui1":
				at = "uint8"
			case "ui2":
				at = "uint16"
			case "ui4":
				at = "uint32"
			case "boolean":
				at = "bool"
			case "string":
				at = "string"
			}
			if sv.SendEvents == "yes" {
				an += ",sendevent"
			}
			if an != "" {
				an = fmt.Sprintf("soap:\"%s\"", an)
			}
			if sv.AllowedRange != nil {
				an = fmt.Sprintf("%s range:\"%d,%d,%d\"", an, sv.AllowedRange.Max, sv.AllowedRange.Min, sv.AllowedRange.Step)
			}
			if sv.AllowedValues != nil {
				an = fmt.Sprintf("%s allowed:\"%s\"", an, strings.Join(*sv.AllowedValues, ","))
			}
			if an != "" {
				an = fmt.Sprintf("`%s`", an)
			}
			if a.Direction == "in" {
				argInCode += fmt.Sprintf("	%s %s %s\n", a.Name, at, an)
			} else if a.Direction == "out" {
				argOutCode += fmt.Sprintf("	%s %s %s\n", a.Name, at, an)
			} else {
				fmt.Printf("argument '%s' direction '%s' invalid for action '%s' service '%s'\n", a.Name, a.Direction, s.Name, srvName)
			}
		}

		fp.WriteString(fmt.Sprintf(`
type ArgIn%[2]s struct {
	XMLName    xml.Name %[1]sxml:"urn:schemas-upnp-org:service:%[3]s:%[4]s %[2]s"%[1]s
%[5]s
}`, "`", s.Name, srvName, version, argInCode))
		fp.WriteString(fmt.Sprintf(`
type ArgOut%[2]s struct {
	XMLName    xml.Name %[1]sxml:"urn:schemas-upnp-org:service:%[3]s:%[4]s %[2]sResponse"%[1]s
%[5]s
}`, "`", s.Name, srvName, version, argOutCode))
	}
}

func printActions(scpds *scpd.SCPD, pwd, srvName string, version string) {
	file := filepath.Join(filepath.Dir(pwd), strings.ToLower(srvName)+version, "actions.go")

	fp, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open failed", file)
		return
	}
	defer func() {
		fp.Close()

		exec.Command("gofmt", "-e", "-w", file).Start()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fp.Truncate(0)
	fp.WriteString(fmt.Sprintf(`package %s%s
import (
	"github.com/zwcway/fasthttp-upnp/service"
)
`, strings.ToLower(srvName), version))

	for _, s := range scpds.ActionList {
		fp.WriteString(fmt.Sprintf(`func %[1]s(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "%[1]s",
		Handler: handle,
		ArgIn:   &ArgIn%[1]s{},
		ArgOut:  &ArgOut%[1]s{},
	}
}
	`, s.Name))
	}
}
func printController(scpds *scpd.SCPD, pwd, srvName string, version string) {
	file := filepath.Join(filepath.Dir(pwd), strings.ToLower(srvName)+version, "controller.go")

	fp, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open failed", file)
		return
	}
	defer func() {
		fp.Close()

		exec.Command("gofmt", "-e", "-w", file).Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	actions := ""
	for _, s := range scpds.ActionList {
		actions += (fmt.Sprintf("			%s(service.DefaultActionHandler),\n", s.Name))
	}
	fp.Truncate(0)
	fp.WriteString(fmt.Sprintf(`package %s%s
import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func ServiceController() *service.Controller {
	return &service.Controller{
		ServiceName: NAME,
		Actions:     []*service.Action{
%s			
		},
	}
}

`, strings.ToLower(srvName), version, actions))

}

func printTypes(scpds *scpd.SCPD, pwd, srvName string, version string) {
	file := filepath.Join(filepath.Dir(pwd), strings.ToLower(srvName)+version, "types.go")

	fp, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open failed", file)
		return
	}
	defer func() {
		fp.Close()

		exec.Command("gofmt", "-e", "-w", file).Start()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fp.Truncate(0)
	fp.WriteString(fmt.Sprintf(`package %s%s
import (
	"github.com/zwcway/fasthttp-upnp/service"
)

const (
	NAME    = service.ServiceName_%[3]s
	VERSION = %[2]s
)

`, strings.ToLower(srvName), version, srvName))
}
func main() {
	var version uint64
	flag.Uint64Var(&version, "v", 1, "")
	flag.Parse()

	if version < 1 {
		version = 1
	}

	regexpRemoveVendor := regexp.MustCompile(`(?i) *<allowedValue> *vendor.*?\n`)
	regexpReplaceVendor := regexp.MustCompile(`(?i)(.+?>).*?Vendor.+?(<.+?)\n`)

	pwd, _ := os.Getwd()
	serviceDir := filepath.Join(pwd, "services")
	files, err := ioutil.ReadDir(serviceDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() || strings.ToLower(path.Ext(file.Name())) != ".xml" {
			continue
		}
		f := pathutil.FileName(file.Name())
		version := f[len(f)-1:]
		f = f[:len(f)-1]

		var scpds scpd.SCPD

		fileXMLBytes, err := os.ReadFile(filepath.Join(serviceDir, file.Name()))
		if err != nil {
			fmt.Println("file", f)
			panic(err)
		}

		fileXMLBytes = regexpRemoveVendor.ReplaceAll(fileXMLBytes, []byte{})
		fileXMLBytes = regexpReplaceVendor.ReplaceAll(fileXMLBytes, []byte("$1 0 $2"))

		err = xml.Unmarshal(fileXMLBytes, &scpds)
		if err != nil {
			fmt.Println("file", f)
			panic(err)
		}

		path := fmt.Sprintf("../%s%s", strings.ToLower(f), version)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed %s", err.Error())
				continue
			}
		}

		checkVaribale(&scpds)
		printArguments(&scpds, pwd, f, (version))
		printActions(&scpds, pwd, f, (version))
		printController(&scpds, pwd, f, (version))
		printTypes(&scpds, pwd, f, (version))
	}

}
