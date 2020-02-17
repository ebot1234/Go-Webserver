package web

import (
	"fmt"
	"text/template"
)

const (
	sesstionTolkenCookie = "session_cookie"
	adminUser            = "admin"
)

type Web struct {
	dmv             *dmv.dmv
	templateHelpers template.FuncMap
}

func NewWeb(dmv dmv.dmv) *Web {
	web := &Web{dmv: dmv}

	//Helper functions that can be used inside templates
	web.templateHelpers = template.FuncMap{
		// Allows sub-templates to be invoked with multiple arguments.
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("Invalid dict call.")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("Dict keys must be strings.")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"add": func(a, b int) int {
			return a + b
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"seq": func(count int) []int {
			seq := make([]int, count)
			for i := 0; i < count; i++ {
				seq[i] = i + 1
			}
			return seq
		},
	}

	return web
}

//Starts that webserver
func ServeWebInterface(port int) {

}
