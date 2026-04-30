# GoAnime Release Notes - Version 1.8.3

Release date: 2026-04-30

## Highlights

- **AllAnime Encryption Update**: Updated AllAnime stream decryption from AES-GCM to AES-256-CTR cipher mode to maintain compatibility with their latest API changes.
- **Zenpen/Kouhen Title Fix**: Fixed critical bug where anime titles with meaningful suffixes (e.g., "Jujutsu Kaisen - Zenpen") were being incorrectly stripped, causing incorrect metadata binding to sibling entries. Now preserves legitimate episode part markers while still removing PT-BR noise.
- **Distinct ID Tracking**: Added support for keeping MyAnimeList (MAL) and AniList IDs separate throughout the playback pipeline. Previously they were conflated, causing incorrect skip intros and tracking data.
- **AnimeFire Quality Menu Improvements**: Fixed three UX bugs in the quality picker: items now sort in descending resolution order, duplicate qualities are disambiguated with mirror labels, and pressing Esc now routes back instead of silently playing the first source.
- **Source Health Diagnostics Enhancement**: Added post-timeout origin probing to distinguish between "client timeout waiting" and "upstream actively down" (Cloudflare 521-524), producing more actionable diagnostic messages.
- **Cloudflare Challenge Detection Refinement**: Fixed false-positive challenge detection that was blocking legitimate Goyabu episode pages mentioning "cloudflare" in footers/comments/ToS. Now targets specific challenge markers instead of the bare word.

## Features

- Add regression tests for AllAnime cipher-mode and transport changes.
- Implement post-timeout origin probing for enhanced source diagnostics (FlixHQ, SFlix, 9Anime).
- Add support for distinct MAL and AniList IDs in playback and tracking pipeline.
- Normalize hyphenated search queries in SuperFlix search for better results.
- Improve AnimeFire quality menu with descending sort, mirror disambiguation, and proper abort handling.

## Bug Fixes

- **Critical**: Fix Zenpen/Kouhen title mismatch by preserving meaningful anime title suffixes. CleanTitle now only strips known PT-BR noise patterns, preventing data corruption where "Jujutsu Kaisen: Shimetsu Kaiyuu - Zenpen" was misbound to the Kouhen (Part 2) entry instead of Zenpen (Part 1). This affected Discord RPC, AniSkip, file naming, and tracking.
- Fix AllAnime cipher-mode decryption: update from AES-GCM to AES-256-CTR to maintain stream extraction compatibility.
- Fix case-sensitive source matching for AnimeFire ("Animefire.io" vs "AnimeFire") causing source breakdown diagnostics to report zero AnimeFire results despite successful searches.
- Fix false-positive Cloudflare challenge detection that was blocking legitimate Goyabu episode pages. Relaxed body-text scan from bare "cloudflare" match to specific challenge phrases ("cf-error", "checking your browser", "ray id", etc.) while preserving all real challenge detection.
- Fix AnimeFire quality selection bug: properly index into sorted sources by user selection instead of matching on rendered labels, fixing silent loss of selection when duplicates exist.
- Fix AnimeFire quality picker handling of Esc/Ctrl-C to return ErrBackRequested instead of silently playing the first source.
- Fix dead Blogger token fast-fail by detecting empty batchexecute responses (anti-hijacking prefix only) instead of wasting ~6 seconds on doomed retries.
- Add nil guard to tracker to prevent crashes on update tracking attempts.
- Improve Blogger video URL extraction robustness to handle edge cases and malformed responses.
- Update error messages and log output throughout the codebase for improved clarity and consistency.

## Improvements

- Refactor quality menu sorting and labeling logic into buildQualityMenu helper for testability and maintainability.
- Add comprehensive regression test suites for all major bug fixes (Zenpen/Kouhen, source breakdown, season selection abort, quality menu, Blogger unavailable token, Cloudflare challenges, origin probing).
- Enhance debug logging throughout Goyabu scraper to aid diagnostics of extraction failures.
- Add episode page diagnostic dumps (HTML snapshot) when all Goyabu extraction strategies fail in debug mode.
- Improve download header handling with centralized applyDownloadAuthHeaders function for consistency.
- Add MarginTop/Priority fields to SuperFlix media results for better UI presentation.
- Use slices.Sort instead of deprecated sort.Slice in unified scraper manager.

---

