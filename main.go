package main

import (
	"go_csv_excel/controllers"
	"log"
	"net/http"
)

func main() {
	// Register the handler function for the "/hello" endpoint
	http.HandleFunc("/report", controllers.ReportHandler)
	http.HandleFunc("/report/csv", controllers.ReportCSVHandler)
	http.HandleFunc("/upload/csv", controllers.ReportCSVDownloadHandler)
	http.HandleFunc("/report/excel", controllers.ReportExcelHandler)
	http.HandleFunc("/upload/excel", controllers.UploadExcelHandler)

	// Define the server port
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)

	// Start the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
