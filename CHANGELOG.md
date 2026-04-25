# GoAnime Release Notes - Version 1.8.2

Release date: 2026-04-25

## Highlights

- **Source Health Diagnostics**: Introduced health checks and diagnostics for scraper sources to proactively monitor and ensure source reliability.
- **Enhanced Stream & Download Stability**: Fixed playback and download issues for direct streams by properly passing required `Referer` headers for AllAnime and AnimeFire CDNs.
- **Windows Auto-Updater Fix**: Resolved a critical issue on Windows where the updater failed to extract the `.exe` binary from the downloaded zip archive.
- **AllAnime Key Rotation Fix**: Updated the AES-GCM decryption key handling to accommodate AllAnime's latest API changes, restoring access to its catalog.

## Features

- Implement source diagnostics and health checks for scrapers to monitor their status.

## Bug Fixes

- Fix AllAnime AES-GCM key rotation to maintain stream extraction functionality.
- Fix Goyabu `batchexecute` index payload structure for search queries.
- Fix the player to properly pass the `Referer` header to direct streams, resolving playback failures.
- Fix AllAnime source selection to preserve the `fast4speed` direct source as a fallback.
- Stop waiting unnecessarily after failed stream sources in AllAnime.
- Fix downloads by adding the `Referer` header for AllAnime CDN URLs.
- Fix downloads by setting the `Referer` header for AnimeFire's CDN (`lightspeedst.net`).
- Fix the auto-updater on Windows to extract the `.exe` file from the zip archive before replacing the application binary.
- Fix potential crashes by adding a nil guard to the tracker.
- Fix flaky FlixHQ tests and unblock Codacy CI checks.

## Improvements

- Update comments and error messages throughout the codebase for improved clarity and consistency.
- Update Linux installation instructions.
- Add test coverage for download retry exhaustion and referer handling.
