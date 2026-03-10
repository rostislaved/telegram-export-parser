# Telegram Export Parser

[[Русская версия]](README-ru.md)

Parses Telegram chat exports from `result.json`.

The main value of this project is the Go data model for Telegram export JSON. The export structure is not trivial, especially for message text entities and rich text parts.

---

## What it does

- Reads a Telegram export JSON file.
- Unmarshals the export into Go structs.
- Filters messages by date.
- Writes plain message text to an output file.
- Optionally prints messages to stdout in a table.
- Prints input and output file sizes.
