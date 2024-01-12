package util

import (
	"bytes"

	"github.com/alecthomas/chroma/v2"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"

	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

type ConvertMarkdownResult struct {
	Frontmatter
	Result string
}

func ConvertMarkdown(post []byte) (Frontmatter, string, error) {
	theme := chroma.MustNewStyle("custom", chroma.StyleEntries{
		chroma.Background:               "bg:#1e1e2e #cdd6f4",
		chroma.CodeLine:                 "#cdd6f4",
		chroma.Error:                    "#f38ba8",
		chroma.Other:                    "#cdd6f4",
		chroma.LineTableTD:              "",
		chroma.LineTable:                "",
		chroma.LineHighlight:            "bg:#45475a",
		chroma.LineNumbersTable:         "#7f849c",
		chroma.LineNumbers:              "#7f849c",
		chroma.Keyword:                  "#cba6f7",
		chroma.KeywordReserved:          "#cba6f7",
		chroma.KeywordPseudo:            "#cba6f7",
		chroma.KeywordConstant:          "#fab387",
		chroma.KeywordDeclaration:       "#f38ba8",
		chroma.KeywordNamespace:         "#94e2d5",
		chroma.KeywordType:              "#f38ba8",
		chroma.Name:                     "#cdd6f4",
		chroma.NameClass:                "#f9e2af",
		chroma.NameConstant:             "#f9e2af",
		chroma.NameDecorator:            "bold #89b4fa",
		chroma.NameEntity:               "#94e2d5",
		chroma.NameException:            "#fab387",
		chroma.NameFunction:             "#89b4fa",
		chroma.NameFunctionMagic:        "#89b4fa",
		chroma.NameLabel:                "#89dceb",
		chroma.NameNamespace:            "#fab387",
		chroma.NameProperty:             "#fab387",
		chroma.NameTag:                  "#cba6f7",
		chroma.NameVariable:             "#f5e0dc",
		chroma.NameVariableClass:        "#f5e0dc",
		chroma.NameVariableGlobal:       "#f5e0dc",
		chroma.NameVariableInstance:     "#f5e0dc",
		chroma.NameVariableMagic:        "#f5e0dc",
		chroma.NameAttribute:            "#89b4fa",
		chroma.NameBuiltin:              "#89dceb",
		chroma.NameBuiltinPseudo:        "#89dceb",
		chroma.NameOther:                "#cdd6f4",
		chroma.Literal:                  "#cdd6f4",
		chroma.LiteralDate:              "#cdd6f4",
		chroma.LiteralString:            "#a6e3a1",
		chroma.LiteralStringChar:        "#a6e3a1",
		chroma.LiteralStringSingle:      "#a6e3a1",
		chroma.LiteralStringDouble:      "#a6e3a1",
		chroma.LiteralStringBacktick:    "#a6e3a1",
		chroma.LiteralStringOther:       "#a6e3a1",
		chroma.LiteralStringSymbol:      "#a6e3a1",
		chroma.LiteralStringInterpol:    "#a6e3a1",
		chroma.LiteralStringAffix:       "#f38ba8",
		chroma.LiteralStringDelimiter:   "#89b4fa",
		chroma.LiteralStringEscape:      "#89b4fa",
		chroma.LiteralStringRegex:       "#94e2d5",
		chroma.LiteralStringDoc:         "#6c7086",
		chroma.LiteralStringHeredoc:     "#6c7086",
		chroma.LiteralNumber:            "#fab387",
		chroma.LiteralNumberBin:         "#fab387",
		chroma.LiteralNumberHex:         "#fab387",
		chroma.LiteralNumberInteger:     "#fab387",
		chroma.LiteralNumberFloat:       "#fab387",
		chroma.LiteralNumberIntegerLong: "#fab387",
		chroma.LiteralNumberOct:         "#fab387",
		chroma.Operator:                 "bold #89dceb",
		chroma.OperatorWord:             "bold #89dceb",
		chroma.Comment:                  "italic #6c7086",
		chroma.CommentSingle:            "italic #6c7086",
		chroma.CommentMultiline:         "italic #6c7086",
		chroma.CommentSpecial:           "italic #6c7086",
		chroma.CommentHashbang:          "italic #6c7086",
		chroma.CommentPreproc:           "italic #6c7086",
		chroma.CommentPreprocFile:       "bold #6c7086",
		chroma.Generic:                  "#cdd6f4",
		chroma.GenericInserted:          "bg:#313244 #a6e3a1",
		chroma.GenericDeleted:           "#f38ba8 bg:#313244",
		chroma.GenericEmph:              "italic #cdd6f4",
		chroma.GenericStrong:            "bold #cdd6f4",
		chroma.GenericUnderline:         "underline #cdd6f4",
		chroma.GenericHeading:           "bold #fab387",
		chroma.GenericSubheading:        "bold #fab387",
		chroma.GenericOutput:            "#cdd6f4",
		chroma.GenericPrompt:            "#cdd6f4",
		chroma.GenericError:             "#f38ba8",
		chroma.GenericTraceback:         "#f38ba8",
	})

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&frontmatter.Extender{},
			highlighting.NewHighlighting(highlighting.WithCustomStyle(theme)),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)

	ctx := parser.NewContext()
	var result bytes.Buffer
	if err := md.Convert(post, &result, parser.WithContext(ctx)); err != nil {
		return Frontmatter{}, "", err
	}

	d := frontmatter.Get(ctx)
	var meta Frontmatter

	if err := d.Decode(&meta); err != nil {
		return Frontmatter{}, "", err
	}

	return meta, result.String(), nil
}
