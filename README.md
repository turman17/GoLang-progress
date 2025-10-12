
# GoLang Progress

This repository documents my journey of learning and mastering the Go programming language — from beginner-level exercises to advanced concurrent and networked systems.  
Each folder contains a standalone Go module representing a practical project focused on a specific concept.

---

## Overview

Through this repo, I explore:
- Core Go syntax and structure
- Concurrency and synchronization
- Networking and REST APIs
- CLI development and JSON handling
- Real-world Go project structure (cmd/internal separation)

---

## Project List

### **1. Calculator**
A simple CLI calculator for basic operations (add, subtract, multiply, divide).  
**Concepts:** CLI flags, error handling, testing.

---

### **2. Concurrent Counter**
A concurrency exercise demonstrating race conditions and synchronization.  
**Features:**
- Safe and unsafe counter implementations  
- Uses `sync.Mutex` and `sync.WaitGroup`  
- Flags to control number of goroutines and iterations  

---

### **3. Concurrent Image Downloader**
Downloads multiple images concurrently using worker pools.  
**Concepts:** Channels, goroutines, I/O, file operations.

---

### **4. Concurrent Web Fetcher**
Fetches multiple URLs in parallel and reports response times.  
**Concepts:** `net/http`, concurrency patterns, benchmarking.

---

### **5. Worker Pool**
Implements a configurable pool of workers processing integer tasks concurrently.  
**Concepts:** Fan-in/fan-out patterns, channels, synchronization.  
**Flags:** control workers, tasks, and simulated delay.

---

### **6. Weather CLI**
Command-line weather app using the OpenWeatherMap API.  
**Concepts:** HTTP requests, JSON decoding, environment variables.  
**Usage Example:**
```bash
go run ./cmd/weather -city "Lisbon" -units metric
```

---

### **7. File Analyzer**
Analyzes text files to count words, lines, and characters.  
**Concepts:** File I/O, scanners, CLI flags.

---

### **8. Mini REST Server**
A minimal in-memory REST API for user management.  
**Endpoints:**
- `GET /users`
- `POST /users`
- `DELETE /users/{id}`  
**Concepts:** `net/http`, JSON I/O, concurrency-safe data structures.

---

## 🧠 Skills Practiced
- **Go Modules** and project structure (`cmd/`, `internal/`)
- **Goroutines** and **Channels**
- **Mutexes**, **WaitGroups**, and concurrency safety
- **HTTP servers** and **clients**
- **JSON serialization/deserialization**
- **Error handling** and **testing**

---

## Repository Structure
```
GoLang-progress/
├── calculator/
├── concurrent-counter/
├── concurrent-image-downloader/
├── concurrent-web-fetcher/
├── file-analyzer/
├── weather-cli/
├── worker-pool/
├── mini-rest/
└── README.md
```

---

## Goal
This repository is my growing portfolio of Go projects — from foundational tools to concurrent systems.  
My goal is to continue expanding it with more advanced projects, covering topics like databases, REST APIs, distributed systems, and Go-based DevOps tooling.

---
