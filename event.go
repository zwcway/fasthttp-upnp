package upnp

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Subscribe struct {
	Sid     string
	URLs    []*url.URL
	Timeout time.Time
}

func (s *Subscribe) IsExpired() bool {
	return s.Timeout.Before(time.Now())
}

type Event struct {
	subscribers map[string]*Subscribe
}

func (e *Event) HasSubscribed(uuid string) bool {
	if e.subscribers == nil {
		return false
	}
	sub, ok := e.subscribers[uuid]
	return ok && !sub.IsExpired()
}

func (e *Event) Subscribe(uuid string, callback []*url.URL, timeout time.Time) *Subscribe {

	if e.subscribers == nil {
		e.subscribers = make(map[string]*Subscribe)
	}

	sb := &Subscribe{
		Sid:     uuid,
		URLs:    callback,
		Timeout: timeout,
	}
	e.subscribers[uuid] = sb

	return sb
}

func (e *Event) MakeTimeout(uuid string) string {
	sub, ok := e.subscribers[uuid]
	if !ok {
		return ""
	}
	second := int(time.Until(sub.Timeout).Seconds())
	return fmt.Sprintf("Second-%d", second)
}

var (
	callbackRegexp = regexp.MustCompile(`<([^<>]+)>`)
)

func ParseCallback(callback string) (ret []*url.URL, err error) {
	list := callbackRegexp.FindAllStringSubmatch(callback, -1)
	for _, match := range list {
		var url *url.URL
		url, err = url.Parse(match[1])
		if err != nil {
			return
		}

		ret = append(ret, url)
	}
	return
}

func ParseTimeout(timeout string) (ret time.Time, err error) {
	times := strings.Split(timeout, "-")
	if len(times) == 2 {
		if times[0] == "Second" {
			var ti int64
			ti, err = strconv.ParseInt(times[1], 0, 32)
			if err != nil {
				return
			}
			ret = time.Now().Add(time.Duration(ti) * time.Second)
			return
		}
	}

	ret = time.Now()
	return
}
