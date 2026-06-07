package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// Order represents an order record.
type Order struct {
	ID        int64     `json:"id"`
	Item      string    `json:"item"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	mu       sync.Mutex
	inMemory = []Order{}
	db       *sql.DB
)

// ConnectDB initializes the package-level DB connection if DATABASE_URL is set.
func ConnectDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL not set")
	}
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	// set conservative pool settings suitable for serverless + small demo workloads
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5 * time.Minute)

	// ping with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return fmt.Errorf("ping: %w", err)
	}
	db = conn
	return nil
}

// Handler serves GET and POST for /api/orders.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if db != nil {
			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()
			rows, err := db.QueryContext(ctx, "select id,item,address,status,created_at from orders order by id desc limit 100")
			if err != nil {
				http.Error(w, `{"error":"db query failed"}`, http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			var res []Order
			for rows.Next() {
				var o Order
				if err := rows.Scan(&o.ID, &o.Item, &o.Address, &o.Status, &o.CreatedAt); err != nil {
					// skip scan errors for resilience
					continue
				}
				res = append(res, o)
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		mu.Lock()
		json.NewEncoder(w).Encode(inMemory)
		mu.Unlock()
	case http.MethodPost:
		var payload struct {
			Item    string `json:"item"`
			Address string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
		if payload.Item == "" || payload.Address == "" {
			http.Error(w, `{"error":"missing fields"}`, http.StatusBadRequest)
			return
		}
		if db != nil {
			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()
			var id int64
			var status string
			var created time.Time
			err := db.QueryRowContext(ctx, "insert into orders(item,address,status,created_at) values($1,$2,$3,now()) returning id,status,created_at", payload.Item, payload.Address, "received").Scan(&id, &status, &created)
			if err != nil {
				http.Error(w, `{"error":"db insert failed"}`, http.StatusInternalServerError)
				return
			}
			o := Order{ID: id, Item: payload.Item, Address: payload.Address, Status: status, CreatedAt: created}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(o)
			return
		}
		o := Order{ID: time.Now().Unix(), Item: payload.Item, Address: payload.Address, Status: "received", CreatedAt: time.Now()}
		mu.Lock()
		inMemory = append([]Order{o}, inMemory...)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(o)
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

// DB is exposed for advanced usage or tests.
func DB() *sql.DB { return db }
