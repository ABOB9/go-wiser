package document

type (
	Document interface {
		Name() string
		Location() string
		Content() string
	}

	Extractor interface {
		FromFile(path string) (Document, error)
		FromString(content string) (Document, error)
	}
)
