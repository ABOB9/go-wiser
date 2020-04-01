package document

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"os"
	"pers.drcz.gowiser/common/stm"
	"pers.drcz.gowiser/utils"
	"strings"
)

// states of this parser
const (
	normal = iota
	title
	article
)

type WYZhBlogParser struct {
	stm stm.StateMachine

	/* for parse document */
	counter uint
	err     error
	/* store document information */
	title    *bytes.Buffer
	location string
	content  *bytes.Buffer
}

func NewWYZhBlogParser() *WYZhBlogParser {
	b, p := stm.NewBuilder(), &WYZhBlogParser{}
	b.SetInitState(normal)
	b.RegState(normal, p.normal)
	b.RegState(title, p.inTitle)
	b.RegState(article, p.inArticle)

	p.stm, _ = b.Build()
	return p
}

func (p *WYZhBlogParser) FromFile(path string) (Document, error) {
	if f, err := os.Open(path); os.IsNotExist(err) {
		return nil, err
	} else {
		p.location = path
		p.process(html.NewTokenizer(f))
		return p.checkResult()
	}
}

func (p *WYZhBlogParser) FromString(content string) (Document, error) {
	p.location = "unknown"
	p.process(html.NewTokenizer(strings.NewReader(content)))
	return p.checkResult()
}

/* Private Methods */

func (p *WYZhBlogParser) process(tk *html.Tokenizer) {
	p.stm.Reset()
	p.counter = 0
	p.content = new(bytes.Buffer)

	for p.err != nil {
		switch tk.Next() {
		case html.ErrorToken:
			p.err = tk.Err()
		default:
			_, p.err = p.stm.Process(tk.Token())
		}
	}
}

func (p *WYZhBlogParser) checkResult() (Document, error) {
	switch p.err {
	case io.EOF:
		return NewTextDocument(p.title.String(), p.location, p.content.String()), nil
	default:
		utils.Log.Printf("error while parse document: %s\n", p.err.Error())
		return nil, p.err
	}
}

func (p *WYZhBlogParser) normal(ctx stm.Context) (interface{}, error) {
	switch t := ctx.Event().(html.Token); {
	case isTitleStart(t):
		_ = ctx.Become(title)
	case isArticleStart(t):
		p.counter = 1
		_ = ctx.Become(article)
	}
	return nil, nil
}

func (p *WYZhBlogParser) inTitle(ctx stm.Context) (interface{}, error) {
	switch t := ctx.Event().(html.Token); {
	case t.Type == html.TextToken:
		p.title.WriteString(t.Data)
	case isTitleEnd(t):
		_ = ctx.Become(normal)
	}
	return nil, nil
}

func (p *WYZhBlogParser) inArticle(ctx stm.Context) (interface{}, error) {
	switch t := ctx.Event().(html.Token); {
	case t.Type == html.TextToken:
		p.content.WriteString(t.Data)
	case isArticleEnd(t):
		return nil, io.EOF
	}
	return nil, nil
}

/* Helper */

func isTitleStart(t html.Token) bool {
	return t.Type == html.StartTagToken && strings.ToLower(t.Data) == "title"
}

func isTitleEnd(t html.Token) bool {
	return t.Type == html.EndTagToken && strings.ToLower(t.Data) == "title"
}

func isArticleStart(t html.Token) bool {
	if t.Type == html.StartTagToken && strings.ToLower(t.Data) == "div" {
		for _, a := range t.Attr {
			if strings.ToLower(a.Key) == "class" && strings.ToLower(a.Val) == "inner" {
				return true
			}
		}
	}
	return false
}

func isArticleEnd(t html.Token) bool {

}
