package upnp

import (
	"fmt"
	"strconv"
)

type ActionArgument struct {
	vals map[string]string
	args map[string]*Argument
}

func (a *ActionArgument) check(key string, t string) (g *Argument, err error) {
	var ok bool
	g, ok = a.args[key]
	if !ok {
		err = (fmt.Errorf("argument '%s' not declared", key))
		return
	}
	if g.RelatedStateVar.DataType != t {
		err = (fmt.Errorf("argument %s[%s] not '%s'", key, g.RelatedStateVar.DataType, t))
		return
	}

	return
}

func (a *ActionArgument) String(key string) (ret string, err error) {
	var (
		g  *Argument
		ok bool
	)
	g, err = a.check(key, DataTypeStr)
	if err != nil {
		return
	}
	if ret, ok = a.vals[key]; ok {
		return
	}
	ret = g.RelatedStateVar.Default
	return
}

func (a *ActionArgument) Uint(key string) (ret uint, err error) {
	var (
		g  *Argument
		ui int64
	)
	g, err = a.check(key, DataTypeUint32)
	if err != nil {
		return
	}

	if r, ok := a.vals[key]; ok {
		ui, err = strconv.ParseInt(r, 0, 64)
		if err != nil {
			return
		}
		if ui < 0 {
			err = fmt.Errorf("'%s' is not a unsigned integer", r)
			return
		}
		ret = uint(ui)
		return
	}
	ui, err = strconv.ParseInt(g.RelatedStateVar.Default, 0, 64)
	if err != nil {
		return
	}
	ret = uint(ui)

	return
}

func (a *ActionArgument) Int(key string) (ret int, err error) {
	var (
		g  *Argument
		ui int64
	)
	g, err = a.check(key, DataTypeInt32)
	if err != nil {
		return
	}

	if r, ok := a.vals[key]; ok {
		ui, err = strconv.ParseInt(r, 0, 64)
		if err != nil {
			return
		}
		ret = int(ui)
		return
	}
	ui, err = strconv.ParseInt(g.RelatedStateVar.Default, 0, 64)
	if err != nil {
		return
	}
	ret = int(ui)

	return
}

func (a *ActionArgument) SetString(key string, val string) {
	a.vals[key] = val
}
func (a *ActionArgument) SetUint(key string, val uint) {
	a.vals[key] = fmt.Sprintf("%d", val)
}
func (a *ActionArgument) SetInt(key string, val int) {
	a.vals[key] = fmt.Sprintf("%d", val)
}
