package main

import (
	"fmt"
	"html"
)

//Html包 提供了用于转义和修改 HTML 文本的功能。
func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
	const escapeString = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(escapeString))
}
