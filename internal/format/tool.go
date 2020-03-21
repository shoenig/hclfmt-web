package format

import (
	"bytes"
	"io"
	"io/ioutil"

	hcl2 "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/pkg/errors"
	"gophers.dev/pkgs/ignore"
)

const (
	defaultMaxLen = 1 << 20
)

type Tool struct {
	maxRequestLen int64
}

func NewTool(max int64) *Tool {
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
		return nil, errors.Wrap(err, "unable to read payload")
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
		return nil, errors.Wrap(err, "unable to write diagnostics")
	}

	return &Diagnostics{
		Body:        buf.String(),
		Result:      diagnostics.Error(),
		Problematic: diagnostics.HasErrors(),
	}, nil
}

func (t *Tool) read(r io.ReadCloser) ([]byte, error) {
	defer ignore.Close(r)
	limit := io.LimitReader(r, t.maxRequestLen)
	return ioutil.ReadAll(limit)
}
