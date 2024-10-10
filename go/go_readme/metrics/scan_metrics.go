package metrics

import (
	"fmt"
	"time"

	"go_readme/db"
)

// InsertMetric logs a metric to the metrics table
func InsertMetric(metricName string, metricValue float64) error {
	timestamp := time.Now().Unix()
	_, err := db.DB.Exec(
		"INSERT INTO metrics (metric_name, metric_value, timestamp) VALUES (?, ?, ?)",
		metricName, metricValue, timestamp,
	)
	if err != nil {
		return fmt.Errorf("failed to insert metric: %v", err)
	}
	return nil
}
