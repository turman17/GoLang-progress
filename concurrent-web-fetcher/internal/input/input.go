package input

import (
	"os"
	"strings"
)

func LinkSplitter(filePath string) ([]string, error){
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	rawLinks := strings.Fields(string(content))
	var links []string
	for _, l := range rawLinks {
		if strings.HasPrefix(l, "http://") || strings.HasPrefix(l, "https://"){
			links = append(links, l)
		}
	}

	return links, nil
}