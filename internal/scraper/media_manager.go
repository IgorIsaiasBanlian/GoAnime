// Package scraper provides unified media handling for anime, movies, and TV shows
package scraper

import (
	"fmt"
	"strings"

	"github.com/alvarorichard/Goanime/internal/models"
	"github.com/alvarorichard/Goanime/internal/util"
)

// MediaManager provides a unified interface for anime sources.
// Movie/TV scrapers (FlixHQ, SFlix) have been removed.
type MediaManager struct {
	scraperManager *ScraperManager
}

// NewMediaManager creates a new MediaManager
func NewMediaManager() *MediaManager {
	return &MediaManager{scraperManager: NewScraperManager()}
}

// SearchAll searches across all sources.
func (mm *MediaManager) SearchAll(query string) ([]*models.Anime, error) {
	return mm.scraperManager.SearchAnime(query, nil)
}

// SearchAnimeOnly searches only anime sources concurrently
func (mm *MediaManager) SearchAnimeOnly(query string) ([]*models.Anime, error) {
	type sourceResult struct {
		results []*models.Anime
		err     error
	}

	ch := make(chan sourceResult, 2)

	go func() {
		t := AllAnimeType
		results, err := mm.scraperManager.SearchAnime(query, &t)
		ch <- sourceResult{results: results, err: err}
	}()

	go func() {
		t := AnimefireType
		results, err := mm.scraperManager.SearchAnime(query, &t)
		ch <- sourceResult{results: results, err: err}
	}()

	var allResults []*models.Anime
	for range 2 {
		res := <-ch
		if res.err == nil {
			allResults = append(allResults, res.results...)
		}
	}

	if len(allResults) == 0 {
		return nil, fmt.Errorf("no anime found with name: %s", query)
	}

	return allResults, nil
}

// GetAnimeStreamURL gets stream URL for anime episodes
func (mm *MediaManager) GetAnimeStreamURL(anime *models.Anime, episodeNum string, quality, mode string) (string, map[string]string, error) {
	source := strings.ToLower(anime.Source)

	util.Debug("Getting stream URL", "source", source, "anime", anime.Name, "episode", episodeNum)

	switch {
	case strings.Contains(source, "allanime"):
		scraper, err := mm.scraperManager.GetScraper(AllAnimeType)
		if err != nil {
			return "", nil, err
		}
		return scraper.GetStreamURL(anime.URL, episodeNum, quality, mode)

	case strings.Contains(source, "animefire"):
		scraper, err := mm.scraperManager.GetScraper(AnimefireType)
		if err != nil {
			return "", nil, err
		}
		return scraper.GetStreamURL(anime.URL, episodeNum, quality, mode)

	default:
		return "", nil, fmt.Errorf("unknown source: %s", anime.Source)
	}
}

// GetScraperManager returns the underlying scraper manager for advanced usage
func (mm *MediaManager) GetScraperManager() *ScraperManager {
	return mm.scraperManager
}
