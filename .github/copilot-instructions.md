<!-- Copied guidance for AI coding agents in this repository -->
# Copilot Instructions — Saaaaad_Project1_SAD-20252026

This repository contains a small React + Vite single-page app under the `Saaaaad/` folder. The notes below are focused on practical, discoverable patterns and workflows an AI coding agent should know to be productive immediately.

**Run & Build**
- `dev`: Start local dev server with Vite. From the repo root run: `cd Saaaaad; npm install; npm run dev` (PowerShell friendly).
- `build`: Create production build: `cd Saaaaad; npm run build`.
- `preview`: Preview a built bundle: `cd Saaaaad; npm run preview`.
- `lint`: Run ESLint: `cd Saaaaad; npm run lint`.

**Project Layout (important files)**
- `Saaaaad/package.json`: scripts + dependencies (use this to discover dev commands).
- `Saaaaad/vite.config.js`: Vite plugins include React plugin and a RollDown/Babel setup that enables the React Compiler preset. Be cautious when changing it — it affects dev/build semantics.
- `Saaaaad/src/main.jsx`: app entry; mounts React root.
- `Saaaaad/src/App.jsx`: top-level app logic — uses `BrowserRouter`, `MantineProvider`, and a simple auth toggle based on `localStorage.getItem('token')` to switch between `pages/auth` and `pages/app`.
- `Saaaaad/eslint.config.js`: ESLint is configured with specific rules; note `no-unused-vars` ignores names matching `^[A-Z_]`.
- `Saaaaad/index.html`: HTML entry referencing `/src/main.jsx`.

**Architecture & Patterns**
- SPA using React Router (`BrowserRouter` in `Saaaaad/src/App.jsx`). Routing and page composition live under `Saaaaad/src/pages/*` (the app expects `pages/auth` and `pages/app`).
- Auth state is derived solely from `localStorage.token` during initial render (see `Saaaaad/src/App.jsx` `useEffect`). To simulate login in dev, set `localStorage.setItem('token', '<token>')` in the browser console.
- UI library: code imports Mantine (`@mantine/core`) in `App.jsx` but `package.json` currently does not list Mantine — verify and add missing dependencies before working on Mantine components.
- Tooling: React Compiler is enabled via Babel + `reactCompilerPreset()` in `vite.config.js`. This may change component transforms, so tests and HMR behavior can differ from a standard React setup.

**Developer conventions / gotchas**
- File extensions: `.jsx` is used consistently for React components.
- ESLint rule: `no-unused-vars` uses `varsIgnorePattern: '^[A-Z_]'`, so capitalized identifiers or those starting with `_` may be intentionally exempted — respect this when cleaning code.
- Keep changes scoped to `Saaaaad/` unless adding repo-level CI or metadata.

**Common tasks & examples**
- Add a new route: create `Saaaaad/src/pages/myPage.jsx`, export a default React component, then import & route it from the routing component (likely under `pages/app`).
- Fix missing dependency: if code imports `@mantine/core`, add it with `cd Saaaaad; npm install @mantine/core` and run `npm run dev`.
- Reproduce login UI locally: open dev server and run `localStorage.setItem('token','x'); location.reload();` in the browser console.

**Tests & CI**
- No test framework or CI configs were detected. If adding tests, prefer lightweight setups (Jest/Vitest) and ensure `npm run build` still succeeds.

**When editing build/tooling**
- Updating `vite.config.js` or `.eslintrc` can affect all contributors — run `npm run dev` and `npm run build` locally after changes.

**Where to look for more context**
- `Saaaaad/README.md` — template notes about React + Vite used in this project.
- `Saaaaad/src/*` — primary application code.

If anything above is unclear or you want me to expand a specific area (detailed DB migrations, connection pooling guidance, or a fuller frontend checkout/cart flow), tell me which area to expand and I'll iterate.

