# Backend (Go) — Local dev

This backend includes a Vercel-compatible serverless handler at `backend/api/orders.go` and local runner at `backend/cmd/server`.

Run Postgres locally (PowerShell):

```powershell
cd backend; docker-compose up -d
```

Apply migration:

```powershell
docker exec -it <postgres-container> psql -U sa -d safood -f /path/to/backend/migrations/init.sql
```

Run local server (connects to `DATABASE_URL` if set):

```powershell
cd backend
go mod download
go run ./cmd/server
```

To deploy serverless functions to Vercel, set `DATABASE_URL` in Vercel project env vars and deploy — the handler will use the DB when available, otherwise it falls back to an in-memory store.
