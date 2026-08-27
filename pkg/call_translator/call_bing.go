/*
Package calltranslator
QueryBingDict 使用 bing 翻译查询
*/
package calltranslator

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// QueryBingDict 使用 bing 翻译
func QueryBingDict(query string) (string, error) {
	// 1. 构造 GET 请求 URL（含 urlencoded 参数）
	baseURL := "https://cn.bing.com/dict/search"
	u, _ := url.Parse(baseURL)
	u.RawQuery = url.Values{"q": {query}}.Encode()

	// 2. 创建 HTTP 客户端（5秒超时 + 跟随重定向）
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	// 3. 模拟 Chrome 浏览器 User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP 状态码: %d", resp.StatusCode)
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	html := string(bodyBytes)

	// 4. 提取 <meta name="description" content="...">
	metaRe := regexp.MustCompile(`<meta\s+name="description"\s+content="([^"]+)"`)
	metaMatch := metaRe.FindStringSubmatch(html)
	if len(metaMatch) < 2 {
		return "", fmt.Errorf("未找到 description meta 标签")
	}
	content := metaMatch[1]

	// 5. 提取真正释义内容（等价于第一个 sed）
	descRe := regexp.MustCompile(`必应词典为您提供.+的释义，(.+)`)
	descMatch := descRe.FindStringSubmatch(content)
	if len(descMatch) < 2 {
		return "", fmt.Errorf("未找到释义内容")
	}
	description := descMatch[1]

	// 6. 替换第一个中文逗号为空格（等价于第二个 sed）
	description = strings.Replace(description, "，", " ", 1)

	// 7. 按空白分割字段（等价于 awk 的 NF 循环）
	fields := strings.Fields(description)

	var out strings.Builder
	out.WriteString(query + "\n")

	// 词性标注和“网络释义：”匹配正则（预先编译）
	pattern := regexp.MustCompile(`^[a-z]+\.`)

	c := 0
	for _, field := range fields {
		// 等价于 awk 的 match($i, "^[a-z]+\\.|网络释义：$")
		if pattern.MatchString(field) || field == "网络释义：" {
			if c == 0 {
				out.WriteString("\n")
			}
			out.WriteString("\n")
			c++
		}
		out.WriteString(field)
		out.WriteString(" ")
	}
	out.WriteString("\n")

	return out.String(), nil
}
