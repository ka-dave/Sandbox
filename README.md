# Sandbox

A personal playground for learning, experimenting, and testing different technologies — one focused spike at a time. The current focus is Redis; new technologies get added as their own experiments over time.

> [!NOTE]
> This is a throwaway-friendly space, not a production project. Code here optimizes for fast iteration and understanding, not polish. Expect rough edges.

## Why this exists

Trying out a new technology usually means the same friction every time: spinning up a scratch repo, wiring up a runtime, remembering how to run it. This repo removes that friction. Each technology gets its own self-contained directory so experiments stay isolated and easy to throw away.

## Layout

Each technology lives in its own top-level directory with its own setup and notes:

```
Sandbox/
├── redis/          # current focus
└── README.md
```

Every experiment folder is meant to stand on its own — its own dependencies, its own run instructions, and ideally a short `README.md` capturing what you were trying to learn.

## Getting started

1. Clone the repo:
   ```sh
   git clone git@github.com:ka-dave/Sandbox.git
   cd Sandbox
   ```
2. Enter the directory for whatever you're exploring (e.g. `cd redis`).
3. Follow that directory's own README for setup and run steps.

Most experiments assume [Docker](https://www.docker.com/) is available for spinning up services locally.

## Adding a new experiment

1. Create a new top-level directory named after the technology.
2. Add a short `README.md` describing what you're testing and how to run it.
3. Keep dependencies scoped to that directory so experiments never collide.

## Conventions

- **One technology per directory** — keep spikes isolated.
- **Self-documenting** — each experiment explains how to run itself.
- **Disposable** — nothing here needs to be maintained forever; delete freely once you've learned what you came for.
