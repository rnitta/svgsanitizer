package main

import (
	"flag"
	"fmt"
	"math/rand"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/beevik/etree"
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		absPath, _ := filepath.Abs(arg)
		convert(absPath)
	}
}

func convert(filepath string) {
	svg := loadXML(filepath)
	allElements, ids, classes := traverseAllElements(svg.Element)
	idTable := createCorresMap(ids)
	classTable := createCorresMap(classes)
	replaceWithMaps(svg, allElements,idTable, classTable)
	svg.Indent(2)
	rep := regexp.MustCompile(`(.+)/(.+?)`)
	outputPath := rep.ReplaceAllString(filepath, "$1/sanitized_$2")
	svg.WriteToFile(outputPath)
	fmt.Printf("Generated. %s\n", outputPath)
}

func loadXML(filepath string) (svg *etree.Document){
	svg = etree.NewDocument()
	if err := svg.ReadFromFile(filepath); err != nil {
		panic(err)
	}
	return
}

func traverseAllElements(e etree.Element) (allElements []*etree.Element, ids []string, classes[]string) {
	allElements = append(allElements, &e)
	for i := 0; true; i++ {
		if len(allElements) <= i { break }
		if children := (allElements[i]).ChildElements(); children != nil {
			allElements = append(allElements, children...)
		}
	}

	idsMap := map[string]bool{}
	classesMap := map[string]bool{}
	allElements = allElements[1:]
	for _, elm := range allElements {

		//FIXME: idとclassの一覧を作成　あとで分離する
		if idAttr := elm.SelectAttr("id"); idAttr != nil {
			for _, idName := range strings.Fields(idAttr.Value) {
				if !idsMap[idName] {
					idsMap[idName] = true
					ids = append(ids, idName)
				}
			}
		}
		if classAttr := elm.SelectAttr("class"); classAttr != nil {
			for _, className := range strings.Fields(classAttr.Value) {
				if !classesMap[className] {
					classesMap[className] = true
					classes = append(classes, className)
				}
			}
		}

		// カスタムdata属性を殺す
		match := regexp.MustCompile(`^data-`)
		for _, attr := range elm.Attr {
			if match.MatchString(attr.Key) {
				elm.RemoveAttr(attr.Key)
			}
		}
	}
	return
}

func createCorresMap(list []string) map[string]string {
	table := map[string]string{}
	for _, attrValue := range list {
		table[attrValue] = randString(20)
	}
	return table
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func replaceWithMaps(doc *etree.Document, allElements []*etree.Element, idTable map[string]string, classTable map[string]string) {
	replaceStyles(doc, idTable, classTable)
	replaceAttrs(allElements, idTable, classTable)
}

func replaceStyles(svg *etree.Document, idTable map[string]string, classTable map[string]string) {
	//FIXME: 遅い
	for _, styleElm := range svg.Element.FindElements("//style") {
		styleText := styleElm.Text()
		reg := regexp.MustCompile(`([^{}]+?)({[^{}]+?})`)
		matches := reg.FindAllStringSubmatch(styleText, -1)
		innerText := ""
		for _, match := range matches {
			selector := match[1]
			selector = replaceTextWithTable(selector, idTable, "#")
			selector = replaceTextWithTable(selector, classTable, ".")
			innerText += selector
			innerText += match[2]
		}
		styleElm.SetText(innerText)
	}
}

func replaceAttrs(allElements []*etree.Element, idTable map[string]string, classTable map[string]string) {
	for _, elm := range allElements {
		if idAttr := elm.SelectAttr("id"); idAttr != nil {
			idName := idAttr.Value
			elm.SelectAttr("id").Value = replaceTextWithTable(idName, idTable, "")
		}
		if classAttr := elm.SelectAttr("class"); classAttr != nil {
			className := classAttr.Value
			elm.SelectAttr("class").Value = replaceTextWithTable(className, classTable, "")
		}
	}
}

func replaceTextWithTable(text string, table map[string]string, prefix string) (replaced string) {
	replaced = text
	for key, val := range table {
		replaced = strings.Replace(replaced, prefix + key, prefix + val, -1)
	}
	return
}