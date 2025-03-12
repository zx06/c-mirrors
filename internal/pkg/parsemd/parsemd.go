package parsemd

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type MDInfo struct {
	From       string
	To         string
	Dockerfile string
}

func ParseMd(md string) *MDInfo {
	info := &MDInfo{}
	reader := text.NewReader([]byte(md))
	parser := goldmark.DefaultParser()
	doc := parser.Parse(reader)

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || n.Kind() != ast.KindFencedCodeBlock {
			return ast.WalkContinue, nil
		}

		fc := n.(*ast.FencedCodeBlock)
		language := string(fc.Language(reader.Source()))
		if strings.ToLower(language) != "dockerfile" {
			return ast.WalkContinue, nil
		}

		// 提取Dockerfile内容
		var buf bytes.Buffer
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			segment := lines.At(i)
			buf.Write(segment.Value(reader.Source()))
		}
		dockerfileContent := buf.String()

		// 解析 [to] 标记
		for _, line := range strings.Split(dockerfileContent, "\n") {
			if strings.HasPrefix(strings.ToLower(line), "# [to]:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					info.To = strings.TrimSpace(parts[1])
				}
			}
		}

		info.Dockerfile = dockerfileContent
		return ast.WalkStop, nil
	})
	info.To = handleTag(info.To)
	return info
}

func handleTag(tag string) string {
	// 在这里编写处理标签的逻辑
	// 例如，将标签转换为小写
	tag = strings.ToLower(tag)
	if !strings.Contains(tag, ":") {
		return tag + ":latest"
	}
	return tag
}
