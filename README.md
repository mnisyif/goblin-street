# Goblin Street

A terminal UI tool for Old School RuneScape Grand Exchange flipping. Suggests profitable items, tracks your trades, and helps you make smarter GE investments.

## Features

- **Market tab** — Live table of all tradeable items with buy/sell prices, spread, ROI%, and volume. Scroll through thousands of items with `j`/`k`.
- **History tab** — Log of past trades with profit tracking.
- **Cached API requests** — Per-endpoint TTL caching so you don't hammer the OSRS Wiki API.
- **GE tax** — All profit/ROI calculations account for the 1% GE tax.

## Usage

```bash
go run .
```

| Key | Action |
|---|---|
| `j` / `k` | Move cursor up/down |
| `Tab` | Switch between Market and History |
| `q` / `Ctrl+C` | Quit |

## Project structure

```
goblin-street/
├── main.go                     # Entry point, fetches data and launches TUI
├── internal/
│   ├── goblinapi/              # OSRS Wiki API client (mapping, latest, 5m, 1h)
│   ├── goblincache/            # In-memory TTL cache with per-entry expiry
│   ├── goblinengine/           # Scoring functions (tax, profit, ROI, margin)
│   └── goblintui/              # Bubble Tea TUI (model, update, view, scroll)
├── docs/problems.md            # Problem log with solutions
├── LICENSE                     # Apache 2.0
└── NOTICE                      # Required attribution notice
```

## License

Apache 2.0. See `LICENSE` and `NOTICE`.
