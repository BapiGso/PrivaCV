package main

import (
	"bytes"
	fences "github.com/stefanfritsch/goldmark-fences"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
	"syscall/js"
)

func main() {
	mdLiteParse := goldmark.New(
		goldmark.WithExtensions(
			extension.Strikethrough,
			extension.Linkify,
			extension.GFM,
			emoji.Emoji,
			&fences.Extender{},
			&frontmatter.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		))
	js.Global().Set("md2html", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return "Error: No Markdown text provided"
		}
		markdown := []byte(args[0].String())
		var buf bytes.Buffer
		if err := mdLiteParse.Convert(markdown, &buf); err != nil {
			return err.Error()
		}
		return buf.String()
	}))
	select {} //阻塞
}

// var mdParse = goldmark.New(
//
//	goldmark.WithExtensions(
//		extension.Strikethrough,
//		extension.Linkify,
//		extension.TaskList,
//		extension.GFM,
//		extension.Footnote,
//		&d2.Extender{},
//		mathjax.MathJax,
//		//&katex.Extender{},
//		&mermaid.Extender{
//			MermaidURL: "/assets/js/mermaid.min.js",
//			//NoScript: true,
//		},
//		&fences.Extender{},
//		&frontmatter.Extender{},
//		emoji.Emoji,
//		//latex.NewLatex(),
//		highlighting.NewHighlighting(
//			highlighting.WithStyle("github")),
//	),
//	goldmark.WithParserOptions(
//		parser.WithAutoHeadingID(),
//	),
//	goldmark.WithRendererOptions(
//		html.WithHardWraps(),
//		html.WithUnsafe(),
//		//html.WithXHTML(),
//	))
