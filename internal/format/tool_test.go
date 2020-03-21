package format

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func load(t *testing.T, name string) io.ReadCloser {
	f, err := os.Open(filepath.Join("resources", name))
	require.NoError(t, err)
	return f
}

func TestTool_Process(t *testing.T) {
	t.Run("basic.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "basic.hcl"))
		require.NoError(t, err)
		require.False(t, result.Diagnostics.Problematic)
		require.Equal(t, "no diagnostics", result.Diagnostics.Result)
		require.Empty(t, result.Diagnostics.Body)
		require.NotEmpty(t, result.Fmt)
	})

	t.Run("complex.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "complex.hcl"))
		require.NoError(t, err)
		require.False(t, result.Diagnostics.Problematic)
		require.Equal(t, "no diagnostics", result.Diagnostics.Result)
		require.Empty(t, result.Diagnostics.Body)
		require.NotEmpty(t, result.Fmt)
	})

	t.Run("problem.hcl", func(t *testing.T) {
		result, err := NewTool(defaultMaxLen).Process(load(t, "problem.hcl"))
		require.NoError(t, err)
		require.True(t, result.Diagnostics.Problematic)
		require.True(t, strings.HasPrefix(result.Diagnostics.Result, "<input>:2,18-22: Missing"))
		require.NotEmpty(t, result.Diagnostics.Body)
		require.Empty(t, result.Fmt)
	})
}
