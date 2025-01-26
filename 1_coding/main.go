package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// SlidingWindowLog represents the structure for the log-based sliding window
type SlidingWindowLog struct {
	requests []time.Time // List of request timestamps
	limit    int         // Maximum number of requests allowed
	window   time.Duration // Time window for rate limiting
}

// NewSlidingWindowLog initializes a new SlidingWindowLog
func NewSlidingWindowLog(limit int, window time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		requests: []time.Time{},
		limit:    limit,
		window:   window,
	}
}

// AllowRequest checks if the incoming request is allowed
func (s *SlidingWindowLog) AllowRequest(requestTime time.Time) bool {
	// Calculate the start of the valid window
	windowStart := requestTime.Add(-s.window)

	// Remove outdated requests
	filteredRequests := []time.Time{}
	for _, t := range s.requests {
		if t.After(windowStart) {
			filteredRequests = append(filteredRequests, t)
		}
	}
	s.requests = filteredRequests

	// Allow the request if the number of requests in the window is within the limit
	if len(s.requests) < s.limit {
		s.requests = append(s.requests, requestTime)
		return true
	}

	// Reject the request if the limit is exceeded
	return false
}

func main() {
	var scanner *bufio.Scanner

	// Check if input file is provided
	if len(os.Args) < 2 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		inputFile := os.Args[1]
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	// Read the first line for N and R
	if !scanner.Scan() {
		fmt.Println("Invalid input. Missing first line with N and R.")
		return
	}
	firstLine := strings.TrimSpace(scanner.Text())
	parts := strings.Split(firstLine, " ")

	if len(parts) != 2 {
		fmt.Println("Invalid input. Please provide two numbers: N and R.")
		return
	}

	N, err1 := strconv.Atoi(parts[0]) // Number of requests
	R, err2 := strconv.Atoi(parts[1]) // Rate limit

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. N and R must be integers.")
		return
	}

	// Initialize the SlidingWindowLog
	window := time.Hour // 1-hour time window
	rateLimiter := NewSlidingWindowLog(R, window)

	if scanner == bufio.NewScanner(os.Stdin) {
		fmt.Printf("Enter %d timestamps (ISO-8601 format):\n", N)
	}

	// Process N timestamps
	for i := 0; i < N; i++ {
		if !scanner.Scan() {
			fmt.Println("Invalid input. Missing timestamps.")
			return
		}
		line := strings.TrimSpace(scanner.Text())

		// Parse the timestamp
		requestTime, err := time.Parse(time.RFC3339, line)
		if err != nil {
			fmt.Printf("Invalid timestamp format: %s. Use ISO-8601 format.\n", line)
			return
		}

		// Check if the request is allowed and print the result immediately
		allowed := rateLimiter.AllowRequest(requestTime)
		fmt.Println(allowed)
	}
}
