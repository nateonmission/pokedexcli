# 🧭 Pokedex CLI (Go)

A terminal-based Pokédex built in Go that lets you explore locations, inspect Pokémon, and try to catch them — all from your command line.

---

## 🚀 Features

* 🌍 Browse location areas (`map`, `mapb`)
* 🔎 Inspect Pokémon stats and details
* 🎯 Catch Pokémon using a probability system based on base experience
* 🧠 In-memory cache to reduce redundant API calls
* 📦 Clean internal package structure (`pokeapi`, `pokecache`)

---

## 🛠️ Tech Stack

* Go (1.22+)
* REST API: https://pokeapi.co/
* Standard library (`net/http`, `encoding/json`, `math/rand`, etc.)

---

## 📁 Project Structure

```
pokedexcli/
├── main.go
├── repl.go
├── commands.go
├── go.mod
└── internal/
    ├── pokeapi/
    │   ├── apiCalls.go
    │   ├── pokemonStruct.go
    │   └── ...
    └── pokecache/
        ├── structs.go
        ├── functions.go
        └── ...
```

---

## ⚙️ Installation

```bash
git clone https://github.com/yourusername/pokedexcli.git
cd pokedexcli
go build
```

Run:

```bash
./pokedexcli
```

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 💻 Usage

Once running, you’ll enter an interactive CLI.

### Commands

| Command          | Description                     |
| ---------------- | ------------------------------- |
| `help`           | Show available commands         |
| `exit`           | Exit the CLI                    |
| `map`            | Show next 20 location areas     |
| `mapb`           | Show previous 20 location areas |
| `catch <name>`   | Attempt to catch a Pokémon      |
| `inspect <name>` | Show details about a Pokémon    |
| `inventory     ` | Show a list of your Pokémon     |


---

## 🎯 Catch System

Catching a Pokémon is probabilistic:

* Based on `base_experience`
* Higher experience → harder to catch
* Clamped between 10% and 90% success rate
* Uses Go’s `math/rand` for randomness

---

## ⚡ Caching

* API responses cached in-memory
* Reduces duplicate network calls
* Background goroutine reaps expired entries

---

## 🧠 Design Philosophy

* Keep CLI logic separate from API logic
* Use internal packages for encapsulation
* Favor simple, readable code over over-engineering
* Build incrementally (Boot.dev style)

---

## 📌 Future Improvements

* Persistent Pokédex (save caught Pokémon)
* Better catch mechanics (items, retries, status effects)
* Command history
* Pagination improvements
* Richer inspect output (sprites, moves, etc.)

---




