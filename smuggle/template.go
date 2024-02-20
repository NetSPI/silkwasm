package smuggle

import (
	"io"
	"os"
	"text/template"
)

func PrintTemplateStr(Data, tmplstr string) {
	tmpl, err := template.New("StrTemplate").Parse(tmplstr)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, Data)
	if err != nil {
		panic(err)
	}
}

func (s *Smuggler) writeFinalTemplate(writer io.Writer) error {
	tmpl, err := template.New("newSmuggler").Parse(smuggleMain)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, s)
	if err != nil {
		return err
	}

	return nil
}

func (h *htmldata) writeHTMLTemplate(writer io.Writer) error {
	tmpl, err := template.New("newHTML").Parse(htmlExample)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, h)
	if err != nil {
		return err
	}

	return nil
}
