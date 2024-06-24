package PKG

import (
	"fmt"
	"os"
	"strings"
)

func Text(text string, art string) (string, error) {
	fmt.Println(text)
	fmt.Println(art)

	var fileName string
	switch art {
	case "standard":
		fileName = ("../banners/standard.txt")
	case "shadow":
		fileName = ("../banners/shadow.txt")
	case "thinkertoy":
		fileName = ("../banners/thinkertoy.txt")
	default:
		// fmt.Println("Wrong banner name !!!!")
		return "500", nil
		// link it with strings.go
	}
	file, err := os.Open(fileName)
	//fmt.Println("Openinging <====", fileName)
	if err != nil {
		// need to change this as breaks code link to text.go
		// log.Fatal("error reading banner file!!! ", err)
		fmt.Println("1")
		return "500", fmt.Errorf("error reading banner file!!! ")
	}
	defer file.Close()
	for i, v := range text {
		if v < 32 || v > 126 {
			if v == 10 {
				fmt.Println(v, i)
				// text = text[:i] + "<br>" + text[i+1:]
				continue
			}
			if v == 13 {
				fmt.Println(text[:i], v, i)
				text = text[:i] + "\\n" + text[i+2:]
				i += 2
				continue
			}
			return "400", nil
			// log.Fatal("wrong text entered!!! ", v)
		}
	}
	if len(text) == 0 {
		return "", nil
	}

	sections := strings.Split(text, "\\n")
	if len(sections) >= 2 && sections[0] == "" && sections[1] == "" {
		fmt.Println()
		sections = sections[2:]

	}
	textArt := "<pre>"
	for _, s := range sections {
		if s == "" {

			textArt += "<br>"
			continue
		}
		cache := [8]string{}
		for _, r := range s {
			Strings(&fileName, r, &cache)
		}
		PrintA(&cache, &textArt)
	}

	// fmt.Println(textArt)
	fmt.Println("3")
	return textArt + "</pre>", nil
}
