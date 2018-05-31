package echo

import (
	"io"
	"github.com/labstack/echo"
	"github.com/ezaurum/boongeoppang"
	"html/template"
)

var _ echo.Renderer = &Template{}

type Template struct {
	templateContainer *boongeoppang.TemplateContainer
}

func NewDebug(templateDir string, funcMap template.FuncMap) *Template {
	b, c := boongeoppang.LoadDebug(templateDir, funcMap)
	t := &Template{
		templateContainer: b,
	}
	go func() {
		for {
			tc := <-c
			t.templateContainer = tc
		}
	}()
	return t
}

// New instance
func New(templateDir string, funcMap template.FuncMap) *Template {
	return &Template{
		templateContainer: boongeoppang.Load(templateDir, funcMap),
	}
}

func (t *Template) Render(w io.Writer, name string,
	data interface{}, c echo.Context) error {
	layout, isExist := t.templateContainer.Get(name)
	if !isExist {
		panic("not exist template " + name)
	}
	return layout.Layout.ExecuteTemplate(w, name, data)
}