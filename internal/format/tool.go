package format

import (
	"bytes"
	"fmt"
	"io"

	hcl2 "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/shoenig/ignore"
)

const (
	defaultMaxLen = 1 << 20
)

type Tool struct {
	maxRequestLen int
}

func NewTool(max int) *Tool {
	if max <= 0 {
		max = defaultMaxLen
	}
	return &Tool{
		maxRequestLen: max,
	}
}

type Result struct {
	Diagnostics *Diagnostics
	Fmt         string
}

func (t *Tool) Process(r io.ReadCloser) (*Result, error) {
	input, err := t.read(r)
	if err != nil {
		return nil, fmt.Errorf("unable to read payload: %w", err)
	}

	diagnostics, err := t.check(input)
	if err != nil {
		return nil, err
	}

	if diagnostics.Problematic {
		return &Result{
			Diagnostics: diagnostics,
		}, nil
	}

	return &Result{
		Diagnostics: diagnostics,
		Fmt:         t.format(input),
	}, nil
}

func (t *Tool) format(input []byte) string {
	return string(hclwrite.Format(input))
}

type Diagnostics struct {
	Body        string
	Result      string
	Problematic bool
}

func (t *Tool) check(body []byte) (*Diagnostics, error) {
	parser := hclparse.NewParser()
	_, diagnostics := parser.ParseHCL(body, "<input>")

	var buf bytes.Buffer
	w := hcl2.NewDiagnosticTextWriter(&buf, parser.Files(), 120, false)
	if err := w.WriteDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("unable to write diagnostics: %w", err)
	}

	return &Diagnostics{
		Body:        buf.String(),
		Result:      diagnostics.Error(),
		Problematic: diagnostics.HasErrors(),
	}, nil
}

func (t *Tool) read(r io.ReadCloser) ([]byte, error) {
	defer ignore.Close(r)
	limit := io.LimitReader(r, int64(t.maxRequestLen))
	return io.ReadAll(limit)
}
