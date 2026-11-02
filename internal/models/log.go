package models

import (
	// "encoding/json"
	"fmt"
)
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level string `json:"level"`
	Service string `json:"service"`
	Message string `json:"message"`
}


func Models() {
	testLog := LogEntry{Timestamp: "2026-11-02 10:00:00", Level: "INFO", Message: "Application started", Service: "Onebank"}
	// When sending the data as json, I will use Marshal and when i want to work with the data i got back from json I will use the Unmarshal
	// data, _ := json.Marshal(testLog)
	fmt.Println(testLog)
}