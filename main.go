package main

import (
	"diff-algorithm/diff"
	"fmt"
	"strings"
)

var File1 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"utf-8\" />\n  <title>Base UI Project</title>\n</head>\n<body>\n  <div id=\"root\"></div>\n</body>\n</html>"
var File2 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"utf-8\" />\n  <title>Base new UI Project</title>\n</head>\n<body>\n  <div id=\"root\"></div>\n</body>\n</html>"

func main() {
	file1Lines := strings.Split(File1, "\n")
	file2Lines := strings.Split(File2, "\n")

	delta := diff.Diff(file1Lines, file2Lines)

	for _, line := range delta {
		fmt.Println(line)
	}
}
