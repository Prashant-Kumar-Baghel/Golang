# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working in this repository.

## What this repo is

A personal Go-learning sandbox, not an application. It's a single module (`golang-learning`, `go.mod` at the root) containing **23 independent, numbered lesson folders**, each demonstrating one Go language concept in isolation:

```
1_hello_world/          basic program structure, go run vs go build
2_simple_values/
3_variables/
4_constants/
5_for/
6_ifElse/
7_switch/
8_arrays/
9_slices/
10_maps/
11_range/
12_functions/
13_variadic_functions/
14_pointers/
15_structs/
16_closure/
17_StructEmbedding/     (note the space in "Struct Embedding.go")
18_interface/
19_enums/
20_generics/
21_goroutines/
22_waitgroups/
23_channels/
```

Each folder holds exactly one `.go` file, always `package main` with its own `func main()`. There is no shared code between folders and no internal imports across them — every folder is a self-contained example.

## Running a lesson

There's no build tooling beyond the Go toolchain itself. To run any single lesson from the repo root:

```bash
go run ./10_maps
```

or from inside the folder:

```bash
cd 10_maps
go run .
```

Because every folder is `package main`, you can only run one lesson at a time — running `go build ./...` from the root would try to produce 23 separate `main` binaries, which is rarely what's wanted; prefer targeting one folder.

There are no tests, no linter config, and no CI in this repo.

## Conventions to preserve when editing

- **Heavy inline commentary is intentional.** These files are personal study notes as much as code — comments explain *why* a concept works the way it does (e.g. deadlock walkthroughs in `23_channels`, pointer-vs-value receiver reasoning in `22_waitgroups`), not just what the line does. When editing a lesson file, keep this teaching-notes style rather than stripping comments down to idiomatic production-code minimalism.
- **Commented-out "earlier attempts" are kept deliberately.** Several files (e.g. `18_interface/interface.go`, `20_generics/generics.go`) show 2-3 progressively better approaches to the same problem, with earlier ones commented out rather than deleted, so the progression stays visible. Don't delete these blocks when adding a new approach — append instead, following the existing "Case N" comment-header pattern.
- **One folder = one concept.** If asked to add a new topic, create a new numbered folder (`24_<topic>/`) rather than adding to an existing one.
