# Changelog

## v1.0.3 - 2026-08-17
Added:
- Many previously missing DTO fields
- A String() method for TextEntity that returns text according to the entity type; for example, text links are converted to "text: link"

Changed:
- Refactored the DTO structure
- Refactored the Text and DateTime types
- Removed duplicate enum constants

Other changes:
- Domain model: renamed History to ExportDTO and improved message type definitions
- Parsing: improved handling of polls, reactions, message text, and optional fields
- Project configuration: updated output paths, ignored generated exports
- Moved code under internal directory
- Updated golanci-lint config


## v1.0.2 - 2026-08-12
Added:
- New mode: extract to json (in addition to extract to plain text)

Fixed:
- Skip "service" message (pin_message, "user1 Added user2", etc)

Changed:
- Moved domain types into dedicated packages

## v1.0.1 - 2026-03-10
Fixed:
- Added missing types: bank_card and unknown
- Stopwatch message printing
- Tested on 1.5 GB of chats' data

## v1.0.0 - 2026-03-10
Added:
- README.md in English and Russian
- golangci-lint + fix all issues
- slog
- input and output files stats printing
- elapsed time printing

Changed:
- linesToWrite = 0 - no limit for processing

Fixed:
- Fix go mod module path
- some refactoring

## v0.0.1
- Initial release
