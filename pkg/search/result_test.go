package search

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepositoryExportData(t *testing.T) {
	var createdAt = time.Date(2021, 2, 28, 12, 30, 0, 0, time.UTC)
	tests := []struct {
		name   string
		fields []string
		repo   Repository
		output string
	}{
		{
			name:   "exports requested fields",
			fields: RepositoryFields,
			repo: Repository{
				Archived:    true,
				CreatedAt:   createdAt,
				Description: "description",
				Fork:        false,
				FullName:    "cli/cli",
				Private:     false,
				PushedAt:    createdAt,
			},
			output: `{"archived":true,"createdAt":"2021-02-28T12:30:00Z","description":"description","fork":false,"fullName":"cli/cli","private":false,"pushedAt":"2021-02-28T12:30:00Z"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exported := tt.repo.ExportData(tt.fields)
			buf := bytes.Buffer{}
			enc := json.NewEncoder(&buf)
			require.NoError(t, enc.Encode(exported))
			assert.Equal(t, tt.output, strings.TrimSpace(buf.String()))
		})
	}
}
