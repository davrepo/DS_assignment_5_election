package database

import (
	"encoding/csv"
	"os"
)

func WriteToCSV(data string) error {
	filePath := "database/database/data.csv"

	// Open the CSV file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the data to the CSV file
	err = writer.Write([]string{data})
	if err != nil {
		return err
	}

	// Flush any buffered data to the underlying writer (the file)
	writer.Flush()

	// Check for any errors during the flush
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
