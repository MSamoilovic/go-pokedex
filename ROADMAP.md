# Go Pokedex — Roadmap

## Current State

The app is a functional CLI Pokedex. Users can explore location areas, catch Pokemon with a probability-based system, and inspect their collection. All API responses are cached in memory.

---

## v1.1 — Quality of Life

Small improvements that directly address friction in the current experience.

- **Persistent Pokedex** — Save caught Pokemon to a local JSON file so progress is not lost between sessions. This is the single most important missing feature; losing your Pokedex on exit kills replayability.
- **Release command** — `release <pokemon>` removes a Pokemon from the Pokedex. Mirrors the actual game and gives users control over their collection.
- **Colored output** — Highlight Pokemon types, stats, and caught/escaped messages with ANSI colors. Low effort, high perceived polish.
- **Case-insensitive input** — `catch Pikachu` and `catch pikachu` should behave identically. Currently fails silently with a not-found error.

---

## v1.2 — Richer Pokemon Data

Expand what the user can learn about each Pokemon using data already available in the API.

- **Abilities** — Show a Pokemon's abilities in `inspect`. Abilities are a core part of Pokemon identity and are missing from the current output.
- **Evolution chain** — `evolutions <pokemon>` shows the full evolution line (e.g. Charmander -> Charmeleon -> Charizard). Uses the evolution-chain endpoint.
- **Moves** — List learnable moves in `inspect`. Can be behind a `--verbose` flag to keep default output clean.
- **Sprites as ASCII art** — Render the Pokemon's official sprite as ASCII/Unicode block art in the terminal. A strong visual differentiator for this kind of CLI app.

---

## v1.3 — Gameplay Depth

Features that add meaningful decisions and longer play sessions.

- **Pokeball types** — Introduce Great Ball and Ultra Ball as commands. Each costs a "token" earned by exploring new areas. Higher-tier balls increase catch rate, adding a resource loop.
- **Shiny encounters** — Small random chance (1/128) that a Pokemon encountered via `explore` is shiny. Shiny Pokemon are harder to catch but display differently in the Pokedex. Pure cosmetic, but historically drives engagement.
- **Catch streak** — Track how many unique Pokemon the user has caught and display a counter in the prompt (e.g. `pokedex [12] >`). Creates a visible progression hook.
- **Pokedex completion percentage** — `pokedex` command shows what percentage of the total 898 Pokemon have been caught.

---

## v1.4 — Battle System

The most complex milestone. Introduces combat as a reason to care about which Pokemon you catch and how you build your team.

- **Simple turn-based battle** — `battle <pokemon1> <pokemon2>` simulates a fight between two of your caught Pokemon using base stats. No moves or type matchups in the first iteration — just attack vs defense math.
- **Type effectiveness** — Apply the official type chart to battles (e.g. Water beats Fire). This makes type diversity in your Pokedex strategically relevant.
- **Wild battles** — `explore` occasionally triggers a forced battle with a wild Pokemon before you can attempt to catch it. Weakening it first increases catch rate.

---

## Technical Debt to Address

These are not features but they will block the roadmap if left unresolved.

- **Disk-based cache** — The in-memory cache is wiped on exit. A file-backed cache (e.g. a local SQLite or flat JSON) would eliminate redundant API calls across sessions.
- **Structured error handling** — Most errors currently surface as raw strings. Define typed errors per package so the REPL can respond differently (e.g. retry on network error vs. print message on not-found).
- **Test coverage** — `cleanInput` is the only tested function. The API client and catch logic need table-driven unit tests before the codebase grows further.
