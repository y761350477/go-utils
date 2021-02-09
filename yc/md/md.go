package md

import "fmt"

// @Summary md 的 1 级标题
// @Param s string "标题名称"
// @Return md 1 级标题
func Title1(s string) string {
	return "# " + s + "\n"
}

// @Summary md 的 2 级标题
// @Param s string "标题名称"
// @Return md 2 级标题
func Title2(s string) string {
	return "## " + s + "\n"
}

// @Summary md 的 3 级标题
// @Param s string "标题名称"
// @Return md 3 级标题
func Title3(s string) string {
	return "### " + s + "\n"
}

// @Summary md 的 4 级标题
// @Param s string "标题名称"
// @Return md 4 级标题
func Title4(s string) string {
	return "#### " + s + "\n"
}

// @Summary md 的图片链接
// @Param url string "图片链接“
// @Param text string "描述信息“ 可选
// @Return md 的图片链接信息
func Img(url string, text ...string) string {
	if len(text) > 0 {
		return fmt.Sprintf(`![%s](%s)`, text[0], url)
	}
	return fmt.Sprintf(`![](%s)`, url)
}

// @Summary md 的超链接
// @Param url string "地址链接“
// @Param text string "描述信息“ 可选
// @Return md 的超链接信息
func Link(url string, text ...string) string {
	if len(text) > 0 {
		return fmt.Sprintf(`[%s](%s)`, text[0], url)
	}
	return fmt.Sprintf(`[](%s)`, url)

}
