package format

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
)

func load(t *testing.T, name string) io.ReadCloser {
	f, err := os.Open(filepath.Join("resources", name))
	must.NoError(t, err)
	return f
}

func TestTool_Process(t *testing.T) {
	t.Run("basic.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "basic.hcl"))
		must.NoError(t, err)
		must.False(t, result.Diagnostics.Problematic)
		must.Equal(t, "no diagnostics", result.Diagnostics.Result)
		must.Empty(t, result.Diagnostics.Body)
		must.NotEmpty(t, result.Fmt)
	})

	t.Run("complex.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "complex.hcl"))
		must.NoError(t, err)
		must.False(t, result.Diagnostics.Problematic)
		must.Equal(t, "no diagnostics", result.Diagnostics.Result)
		must.Empty(t, result.Diagnostics.Body)
		must.NotEmpty(t, result.Fmt)
	})

	t.Run("problem.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "problem.hcl"))
		must.NoError(t, err)
		must.True(t, result.Diagnostics.Problematic)
		must.True(t, strings.HasPrefix(result.Diagnostics.Result, "<input>:2,18-22: Missing"))
		must.NotEmpty(t, result.Diagnostics.Body)
		must.Empty(t, result.Fmt)
	})
}
