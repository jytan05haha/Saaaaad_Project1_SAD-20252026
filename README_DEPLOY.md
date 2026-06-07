# Deploying to Vercel (instructions)

This repository contains a `frontend/` single-page app (React + Vite + Mantine) and a `backend/` folder with a Vercel-compatible Go serverless handler under `backend/api`.

To deploy and get a public URL you can run (you must have `vercel` installed and be logged in):

```powershell
cd frontend; npm install; npm run build
cd ..
vercel login
vercel --prod
```

Notes:
- The included `backend/api/orders.go` is a serverless function that uses an in-memory store if `DATABASE_URL` is not set. For production, provision a managed Postgres and set `DATABASE_URL` in Vercel project environment variables.
- For local DB-backed development, start Postgres with `cd backend; docker-compose up -d` and run `go run ./cmd/server`.
