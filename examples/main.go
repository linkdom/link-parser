package main

import (
	"awesomeProject/link-parser"
	"fmt"
	"strings"
)

var exampleHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <h1>Servus</h1>
    <a href="www.google.at">A link to google.</a>
</body>
</html>
`
func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", links)
}
