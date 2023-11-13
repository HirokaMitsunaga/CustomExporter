package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    db *sql.DB
    schemaSize = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "mysql_schema_size_bytes",
            Help: "MySQL schema size in bytes.",
        },
        []string{"schema_name"},
    )
)

func init() {
    prometheus.MustRegister(schemaSize)
}

func collectSchemaSizes() {
    rows, err := db.Query("SELECT table_schema, SUM(data_length + index_length) FROM information_schema.tables GROUP BY table_schema")
    if err != nil {
        log.Println("Error querying schema sizes:", err)
        return
    }
    defer rows.Close()

    var schema string
    var size float64
    for rows.Next() {
        if err := rows.Scan(&schema, &size); err != nil {
            log.Println("Error scanning row:", err)
            continue
        }
        schemaSize.WithLabelValues(schema).Set(size)
    }
}

func main() {
    var err error
    db, err = sql.Open("mysql", "root:my-secret-pw@tcp(mysql.default.svc.cluster.local:3306)/")
    if err != nil {
        log.Fatal("Error opening database:", err)
    }

    go func() {
        for {
            collectSchemaSizes()
            time.Sleep(5 * time.Minute) // 5分間隔で待機
        }
    }()

    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
