package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	str := `

sdfdfdf<br>

sdffaff</br>
<ul>
<li><p class="line">测试框架</p>
<ul>
<li><a href="https://apitest.dev" target="_blank">apitest</a> - 简单和可扩展的行为测试库，用于基于REST的服务或HTTP处理程序，支持模拟外部http调用和序列图的呈现。</li><li><a href="https://github.com/go-playground/assert" target="_blank">assert</a> - 与本地本机go测试一起使用的基本断言库，带有用于自定义断言的构建块。</li><li><a href="https://github.com/cavaliercoder/badio" target="_blank">badio</a> - Go testing/iotest包的扩展。</li><li><a href="https://github.com/h2non/baloo" target="_blank">baloo</a> - 表达性和通用性的端到端HTTP API测试变得容易。</li><li><a href="https://github.com/fulldump/biff" target="_blank">biff</a> - 分叉测试框架，与BDD兼容。</li><li><a href="https://github.com/percolate/charlatan" target="_blank">charlatan</a> - 生成用于测试的虚假接口实现的工具。</li><li><a href="https://github.com/SimonBaeumer/commander" target="_blank">commander</a> - 用于在Windows，Linux和osx上测试cli应用程序的工具。</li><li><a href="https://github.com/bradleyjkemp/cupaloy" target="_blank">cupaloy</a> - 适用于您的测试框架的简单快照测试插件。</li><li><a href="https://github.com/khaiql/dbcleaner" target="_blank">dbcleaner</a> - 受database_cleanerRuby 启发，清理数据库以进行测试。</li><li><a href="https://github.com/viant/dsunit" target="_blank">dsunit</a> - SQL，NoSQL，结构化文件的数据存储区测试。</li><li><a href="https://github.com/fergusstrange/embedded-postgres" target="_blank">embedded-postgres</a> - Postgres-作为另一个Go应用程序或测试的一部分，在Linux，OSX或Windows上本地运行真实的Postgres数据库。</li><li><a href="https://github.com/viant/endly" target="_blank">endly</a> - 声明式端到端功能测试。</li><li><a href="https://github.com/suzuki-shunsuke/flute" target="_blank">flute</a> - HTTP客户端测试框架。</li><li><a href="https://github.com/verdverm/frisby" target="_blank">frisby</a> - REST API测试框架。</li><li><a href="http://onsi.github.io/ginkgo/" target="_blank">ginkgo</a> - Go的BDD测试框架。</li><li><a href="https://github.com/msoap/go-carpet" target="_blank">go-carpet</a> - 用于在终端中查看测试覆盖率的工具。</li><li><a href="https://github.com/google/go-cmp" target="_blank">go-cmp</a> - 用于比较测试中Go值的软件包。</li><li><a href="https://github.com/zimmski/go-mutesting" target="_blank">go-mutesting</a> - Go源代码的变异测试。</li><li><a href="https://github.com/maxatome/go-testdeep" target="_blank">go-testdeep</a> - 极其灵活的golang深度比较，扩展了go测试包。</li><li><a href="https://github.com/dnaeon/go-vcr" target="_blank">go-vcr</a> - 记录并重放您的HTTP交互，以进行快速，确定性和准确的测试。</li><li><a href="https://github.com/franela/goblin" target="_blank">goblin</a> - 摩卡（Mocha）像Go的测试框架。</li><li><a href="http://labix.org/gocheck" target="_blank">gocheck</a> - 比getest更高级的测试框架。</li><li><a href="https://github.com/smartystreets/goconvey/" target="_blank">GoConvey</a> - 具有Web UI和实时重载的BDD样式的框架。</li><li><a href="https://github.com/corbym/gocrest" target="_blank">gocrest</a> - Go断言的可组合的类似于hamcrest的匹配器。</li><li><a href="https://github.com/DATA-DOG/godog" target="_blank">godog</a> - Cucumber或Behat像Go的BDD框架。</li><li><a href="https://github.com/appleboy/gofight" target="_blank">gofight</a> - Golang路由器框架的API处理程序测试。</li><li><a href="https://github.com/corbym/gogiven" target="_blank">gogiven</a> - Go的类似YATSPEC的BDD测试框架。</li><li><a href="https://github.com/jfilipczyk/gomatch" target="_blank">gomatch</a> - 创建用于针对模式测试JSON的库。</li><li><a href="http://onsi.github.io/gomega/" target="_blank">gomega</a> - Rspec，例如匹配器/断言库。</li><li><a href="https://github.com/orfjackal/gospec" target="_blank">GoSpec</a> - 用于Go编程语言的BDD样式测试框架。</li><li><a href="https://github.com/stesla/gospecify" target="_blank">gospecify</a> - 这提供了BDD语法来测试您的Go代码。使用过rspec之类的库的任何人都应该熟悉。</li><li><a href="https://github.com/pavlo/gosuite" target="_blank">gosuite</a> - testing利用Go1.7的子测试，使带有设置/拆卸功能的轻量级测试套件成为可能。</li><li><a href="https://github.com/gotestyourself/gotest.tools" target="_blank">gotest.tools</a> - 一组软件包的集合，以增强go测试软件包并支持常见模式。</li><li><a href="https://github.com/rdrdr/hamcrest" target="_blank">Hamcrest</a> - 用于声明式Matcher对象的流利框架，将其应用于输入值时，会产生自描述结果。</li><li><a href="https://github.com/gavv/httpexpect" target="_blank">httpexpect</a> - 简洁，声明性且易于使用的端到端HTTP和REST API测试。</li><li><a href="https://github.com/kinbiko/jsonassert" target="_blank">jsonassert</a> - 用于验证JSON有效负载已正确序列化的软件包。</li><li><a href="https://github.com/yookoala/restit" target="_blank">restit</a> - 一个微框架，可帮助编写RESTful API集成测试。</li><li><a href="https://github.com/jgroeneveld/schema" target="_blank">schema</a> - 在请求和响应中使用JSON模式方便快捷则表达式匹配。</li><li><a href="https://github.com/adamluzsi/testcase" target="_blank">testcase</a> - 行为驱动开发的惯用测试框架。</li><li><a href="https://github.com/go-testfixtures/testfixtures" target="_blank">testfixtures</a> - Rails的测试夹具，用于测试数据库应用程序。</li><li><a href="https://github.com/stretchr/testify" target="_blank">Testify</a> - 对标准go测试包的神圣扩展。</li><li><a href="https://godoc.org/github.com/tvastar/test/cmd/testmd" target="_blank">testmd</a> - 将markdown代码片段转换为可测试的go代码。</li><li><a href="https://github.com/zhulongcheng/testsql" target="_blank">testsql</a> - 在测试之前从SQL文件生成测试数据，并在完成后将其清除。</li><li><a href="https://github.com/jgroeneveld/trial" target="_blank">trial</a> - 快速简单的可扩展断言，无需引入太多样板。</li><li><a href="https://github.com/vcaesar/tt" target="_blank">Tt</a> - 简单而丰富多彩的测试工具。</li><li><a href="https://github.com/posener/wstest" target="_blank">wstest</a> - Websocket客户端，用于对websocket http.Handler进行单元测试。</li></ul>
</li><li><p class="line">Mock</p>
<ul>
<li><a href="https://github.com/maxbrunsfeld/counterfeiter" target="_blank">counterfeiter</a> -生成独立模拟对象的工具。</li><li><a href="https://github.com/DATA-DOG/go-sqlmock" target="_blank">go-sqlmock</a> - 模拟SQL驱动程序，用于测试数据库交互。</li><li><a href="https://github.com/DATA-DOG/go-txdb" target="_blank">go-txdb</a> - 基于单事务的数据库驱动程序，主要用于测试目的。</li><li><a href="https://github.com/h2non/gock" target="_blank">gock</a> - 多种HTTP 模拟变得容易。</li><li><a href="https://github.com/golang/mock" target="_blank">gomock</a> - Go编程语言的模拟框架。</li><li><a href="https://github.com/seborama/govcr" target="_blank">govcr</a> - 用于Golang的HTTP模拟：记录和重放HTTP交互以进行脱机测试。</li><li><a href="https://github.com/SpectoLabs/hoverfly" target="_blank">hoverfly</a> - HTTP（S）代理，用于通过可扩展的中间件和易于使用的CLI记录和模拟REST / SOAP API。</li><li><a href="https://github.com/jarcoal/httpmock" target="_blank">httpmock</a> - 轻松模拟来自外部资源的HTTP响应。</li><li><a href="https://github.com/gojuno/minimock" target="_blank">minimock</a> - Go接口的模拟生成器。</li><li><a href="https://github.com/tv42/mockhttp" target="_blank">mockhttp</a> - Go http.ResponseWriter的模拟对象。</li></ul>
</li><li><p class="line">Fuzzing and delta-debugging/reducing/shrinking.</p>
<ul>
<li><a href="https://github.com/dvyukov/go-fuzz" target="_blank">go-fuzz</a> - 随机测试系统。</li><li><a href="https://github.com/google/gofuzz" target="_blank">gofuzz</a> - 用于填充带有随机值的go对象的库。</li><li><a href="https://github.com/zimmski/tavor" target="_blank">Tavor</a> - 通用的模糊测试和增量调试框架。</li></ul>
</li><li><p class="line">Selenium and browser control tools.</p>
<ul>
<li><a href="https://github.com/mafredri/cdp" target="_blank">cdp</a> - Chrome调试协议的类型安全绑定，可与实现该协议的浏览器或其他调试目标一起使用。</li><li><a href="https://github.com/knq/chromedp" target="_blank">chromedp</a> - 一种驱动/测试Chrome，Safari，Edge，Android Webview和其他支持Chrome调试协议的浏览器的方法。</li><li><a href="https://github.com/aerokube/ggr" target="_blank">ggr</a> - 轻量级服务器，将Selenium WebDriver请求路由和代理到多个Selenium集线器。</li><li><a href="https://github.com/aerokube/selenoid" target="_blank">selenoid</a> - 替代的Selenium集线器服务器，可在容器内启动浏览器。</li></ul>
</li><li><p class="line">Fail injection</p>
<ul>
<li><a href="https://github.com/pingcap/failpoint" target="_blank">failpoint</a> - failpoint -的实现failpoints为Golang。</li></ul>
</li></ul>


`

	Content := HtmlToMd(str)
	file, err := os.Create("test.md")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(Content)
	fmt.Println(Content)
}

func HtmlToMd(html string) string {
	var r = bytes.NewBufferString(html)
	doc, _ := goquery.NewDocumentFromReader(r)
	//fmt.Println(trimHtml(str))
	doc.Find("*").Each(func(i int, selector *goquery.Selection) {

		if selector.Is("h1") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("\n" + "# " + t + "\n")
		}
		if selector.Is("h2") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("\n" + "## " + t + "\n")
		}
		if selector.Is("h3") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("\n" + "### " + t + "\n")
		}
		if selector.Is("h4") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("\n" + "#### " + t + "\n")
		}
		if selector.Is("br") {
			selector.AfterHtml("\n\n")
		}
		if selector.Is("p") {
			selector.AfterHtml("\n\n")
		}
		if selector.Is("li") {
			selector.AfterHtml("\n\n")
		}
		if selector.Is("img") {
			t, _ := selector.Attr("src")
			selector.ReplaceWithHtml("![](" + t + ")")
		}
		if selector.Is("a") {
			l, _ := selector.Attr("href")
			t := selector.Text()
			selector.ReplaceWithHtml("[" + t + "](" + l + ")")
		}
		if selector.Is("blockquote") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("> " + t)
		}
		if selector.Is("code") {
			t, _ := selector.Html()
			selector.ReplaceWithHtml("**" + t + "**")
		}
		if selector.Is("pre") {
			t, _ := selector.Find("code").Html()
			selector.Find("code").ReplaceWithHtml("\n" + "```" + "\n" + t + "\n" + "```" + "\n")
		}
		if selector.Is("table") {
			th := "\n| "
			sp := "|"
			td := ""
			selector.Find("th").Each(func(i int, selector *goquery.Selection) {
				th = th + selector.Text() + " |"
				sp = sp + "---------------" + " |"
			})
			th = th + "\n"
			sp = sp + "\n"
			selector.Find("tbody").Find("tr").Each(func(i int, selector *goquery.Selection) {
				td = td + "| "
				selector.Find("td").Each(func(i int, selector *goquery.Selection) {
					td = td + selector.Text() + " | "
				})
				td = td + "\n"
			})

			table := th + sp + td
			selector.ReplaceWithHtml(table)
		}

	})
	content := doc.Text()
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ := regexp.Compile("<[A-Za-z//]{2,10}.*?>")
	content = re.ReplaceAllString(content, "")
	//取消html转义
	content = strings.ReplaceAll(content, "&#34;", "\"")
	content = strings.ReplaceAll(content, "&#38;", "&")
	content = strings.ReplaceAll(content, "&#60;", "<")
	content = strings.ReplaceAll(content, "&#62;", ">")
	return content
}

func UnEscapeHTML(s string) string {
	h := strings.ReplaceAll(s, "&#34;", "\"")
	h = strings.ReplaceAll(h, "&#38;", "&")
	h = strings.ReplaceAll(h, "&#60;", "<")
	h = strings.ReplaceAll(h, "&#62;", ">")
	return h
}
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	return strings.TrimSpace(src)
}
