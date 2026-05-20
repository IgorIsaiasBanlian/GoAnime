package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMediaManager(t *testing.T) {
	t.Parallel()
	mm := NewMediaManager()
	require.NotNil(t, mm)
	assert.NotNil(t, mm.scraperManager)
}

func TestMediaManager_SearchAnimeOnly_NoMatch(t *testing.T) {
	t.Parallel()
	mm := NewMediaManager()
	_, _ = mm.SearchAnimeOnly("zzz_definitely_not_an_anime_xyz_999")
}

func TestMediaManager_SearchAll_NoPanic(t *testing.T) {
	t.Parallel()
	mm := NewMediaManager()
	_, _ = mm.SearchAll("zzz")
}
