package main

import (
    "fmt"
    "embed"
    "io"
    "log"
    "net/http"
    "path"
    "strings"
    "time"

    "com-seek/backend/database"

    "github.com/joho/godotenv"
)

//go:embed public/*
var embeddedStatic embed.FS

func main() {
    godotenv.Load()

    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    fmt.Println("Connected to the database", db != nil)

    mux := http.NewServeMux()

    // --- API routes ---
    mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        _, _ = io.WriteString(w, "ok")
    })

    mux.HandleFunc("/api/echo", func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        io.Copy(w, r.Body)
    })

	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    _, _ = io.WriteString(w, `{"message":"Hello from Go backend!"}`)
	})



    // --- Static + SPA fallback ---
    fsys := http.FS(embeddedStatic)  // ✅ FIX
    spa := spaHandler(fsys)

    // Root handler: serve SPA and static
    mux.Handle("/", spa)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      withCORS(mux),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Println("Server running at http://localhost:8080")
    log.Fatal(srv.ListenAndServe())
}

func withCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func spaHandler(fsys http.FileSystem) http.Handler {
    fileServer := http.FileServer(fsys)

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Let API routes pass through
        if strings.HasPrefix(r.URL.Path, "/api/") {
            http.NotFound(w, r)
            return
        }

        // Map URL path to /public/*
        p := strings.TrimPrefix(r.URL.Path, "/")
        if p == "" {
            p = "public/index.html"
        } else {
            p = path.Join("public", p)
        }

        // If file exists, serve it; else fallback to index.html
        if f, err := fsys.Open(p); err == nil {
            f.Close()
            r.URL.Path = "/" + p
        } else {
            r.URL.Path = "/public/index.html"
        }
        fileServer.ServeHTTP(w, r)
    })
}
