package instagram

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	sharedDataSelector = "window._sharedData = "
	scriptSelector     = "script"
	endSelector        = ";"

	//Errors
	ScriptNotFound = "Script window._sharedData not found"
	DontReadHtmlDoc = "Don`t read HTML document"
)

//Find window._sharedData script from HTML document
func findJsonOnHtml(data []byte) (outData []byte, err error) {
	reader := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []byte{}, fmt.Errorf(DontReadHtmlDoc)
	}
	items := doc.Find(scriptSelector)

	items.Each(func(i int, selection *goquery.Selection) {

		if !func(str string) bool {
			return strings.Contains(str, sharedDataSelector)
		}(selection.Text()) {
			selection.Next()
		} else {
			outData = func(str string) []byte {
				str = strings.Replace(str, sharedDataSelector, "", -1)
				str = strings.Replace(str, endSelector, "", -1)
				return []byte(str)
			}(selection.Text())
			selection.End()
		}
	})

	if len(outData) == 0 {
		return []byte{}, fmt.Errorf(ScriptNotFound)
	}
	return outData, nil
}
