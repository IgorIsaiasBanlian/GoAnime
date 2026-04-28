package api

import (
	"strings"
	"testing"

	"github.com/alvarorichard/Goanime/internal/models"
	"github.com/stretchr/testify/assert"
)

// TestSourceBreakdown_LegacyCaseSensitiveBug documents the original case-sensitive
// regression: scrapers emit anime.Source = "Animefire.io" (lowercase 'f'), but the
// debug breakdown searched for the substring "AnimeFire" (capital 'F'). Because
// strings.Contains is byte-exact in Go, the AnimeFire counter was always zero
// even when the search returned dozens of AnimeFire results — making the
// "Source breakdown" diagnostic line lie to operators.
//
// This test reproduces the buggy comparison in isolation. The expectation
// matches the broken behaviour so that if anyone ever "fixes" the bug by
// hand without updating both the production code AND this regression test,
// the failure makes the intent obvious.
func TestSourceBreakdown_LegacyCaseSensitiveBug(t *testing.T) {
	source := "Animefire.io"

	// The exact pre-fix predicate from internal/api/enhanced.go.
	matchesLegacyPredicate := strings.Contains(source, "AnimeFire")

	assert.False(t, matchesLegacyPredicate,
		"sanity check: 'Animefire.io' must NOT match 'AnimeFire' under case-sensitive Contains — "+
			"if this assertion changes, Go's strings package semantics changed and the rest of the suite needs review")
}

// TestCountSourceBreakdown_AnimeFireCaseInsensitive is the positive regression
// test for the fix: the breakdown helper must count "Animefire.io" results
// regardless of capitalisation. Future scrapers (or upstream renames) that
// emit "AnimeFire", "ANIMEFIRE", or "animefire" must all be counted.
func TestCountSourceBreakdown_AnimeFireCaseInsensitive(t *testing.T) {
	animes := []*models.Anime{
		{Source: "Animefire.io"},
		{Source: "Animefire.io"},
		{Source: "AnimeFire"},
		{Source: "ANIMEFIRE"},
		{Source: "animefire"},
		{Source: "AllAnime"},
		{Source: "Goyabu"},
		{Source: "FlixHQ"},
		{Source: "9Anime"},
		{Source: "SuperFlix"},
		{Source: "AnimeDrive"},
	}

	got := countSourceBreakdown(animes)

	assert.Equal(t, 5, got.AnimeFire, "all AnimeFire spellings must be counted")
	assert.Equal(t, 1, got.AllAnime)
	assert.Equal(t, 1, got.Goyabu)
	assert.Equal(t, 1, got.FlixHQ)
	assert.Equal(t, 1, got.NineAnime)
	assert.Equal(t, 1, got.SuperFlix)
	assert.Equal(t, 1, got.AnimeDrive)
}

// TestCountSourceBreakdown_RealisticPayload mirrors the user-reported log:
// 10 AnimeFire results, 5 AllAnime, 8 Goyabu — and asserts that Goyabu is
// reported (the original breakdown silently dropped it).
func TestCountSourceBreakdown_RealisticPayload(t *testing.T) {
	var animes []*models.Anime
	for i := 0; i < 10; i++ {
		animes = append(animes, &models.Anime{Source: "Animefire.io"})
	}
	for i := 0; i < 5; i++ {
		animes = append(animes, &models.Anime{Source: "AllAnime"})
	}
	for i := 0; i < 8; i++ {
		animes = append(animes, &models.Anime{Source: "Goyabu"})
	}

	got := countSourceBreakdown(animes)

	assert.Equal(t, 10, got.AnimeFire, "AnimeFire breakdown must equal what the scraper returned")
	assert.Equal(t, 5, got.AllAnime)
	assert.Equal(t, 8, got.Goyabu, "Goyabu must appear in the breakdown")
	assert.Equal(t, 0, got.AnimeDrive)
	assert.Equal(t, 0, got.FlixHQ)
	assert.Equal(t, 0, got.NineAnime)
	assert.Equal(t, 0, got.SuperFlix)
}
