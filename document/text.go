package document

type TextDocument struct {
	title    string
	location string
	content  string
}

func NewTextDocument(title string, location string, content string) *TextDocument {
	return &TextDocument{title: title, location: location, content: content}
}

func (t *TextDocument) Name() string {
	return t.title
}

func (t *TextDocument) Location() string {
	return t.location
}

func (t *TextDocument) Content() string {
	return t.content
}
