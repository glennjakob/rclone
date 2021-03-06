//go:generate go run assets_generate.go
// The "go:generate" directive compiles static assets by running assets_generate.go

package data

import (
	"html/template"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
)

// GetTemplate returns the HTML template for serving directories via HTTP/Webdav
func GetTemplate(tmpl string) (tpl *template.Template, err error) {
	var templateString string
	if tmpl == "" {
		templateFile, err := Assets.Open("index.html")
		if err != nil {
			return nil, errors.Wrap(err, "get template open")
		}

		defer fs.CheckClose(templateFile, &err)

		templateBytes, err := ioutil.ReadAll(templateFile)
		if err != nil {
			return nil, errors.Wrap(err, "get template read")
		}

		templateString = string(templateBytes)

	} else {
		templateFile, err := ioutil.ReadFile(tmpl)
		if err != nil {
			return nil, errors.Wrap(err, "get template open")
		}

		templateString = string(templateFile)
	}

	tpl, err = template.New("index").Parse(templateString)
	if err != nil {
		return nil, errors.Wrap(err, "get template parse")
	}

	return
}
