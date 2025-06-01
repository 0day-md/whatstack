// whatstack.go
//
// A command-line tool to detect technologies used by a website.
// It uses the Wappalyzer library to fingerprint technologies
// based on HTTP headers and HTML content.
//
// Usage:
//   go run whatstack.go -d example.com [-o output.txt]
//
// Flags:
//   -d <domain>    Domain to analyze (required)
//   -o <file>      Output file to save detected stack (optional)
//
// Author: 0day.md (0day-md.github.io)
// Created: 2025-05-31


package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	// Define command-line flags
	domain := flag.String("d", "", "Domain to analyze")
	outputFile := flag.String("o", "", "Output file to save the stack")
	flag.Parse()

	// Check if the domain is provided
	if *domain == "" {
		log.Fatal("❌ Please provide a domain using the -d flag.")
	}

	// Make the HTTP request
	resp, err := http.Get("https://" + *domain)
	if err != nil {
		log.Fatalf("❌ Error fetching the domain: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("❌ Error: Received status code %d from the server", resp.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("❌ Error reading response body: %v", err)
	}

	// Initialize Wappalyzer client
	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Fatalf("❌ Error initializing Wappalyzer: %v", err)
	}

	// Fingerprint the technologies used
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)

	// Print the results
	fmt.Println("\n🔍 Technologies detected:")
	if len(fingerprints) == 0 {
		fmt.Println("No technologies detected.")
	} else {
		for tech := range fingerprints {
			fmt.Println("- " + tech)
		}
	}

	// Detect server information
	if server := resp.Header.Get("Server"); server != "" {
		fmt.Printf("🖥️  Server: %s\n", server)
	}

	// Save results to file if specified
	if *outputFile != "" {
		if err := saveToFile(*outputFile, fingerprints, resp.Header.Get("Server")); err != nil {
			log.Fatalf("❌ Error saving to file: %v", err)
		}
		fmt.Printf("✅ Results saved to %s\n", *outputFile)
	}
}

// saveToFile writes the detected technologies and server info to the specified file
func saveToFile(filename string, fingerprints map[string]struct{}, server string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for tech := range fingerprints {
		if _, err := file.WriteString(tech + "\n"); err != nil {
			return err
		}
	}

	if server != "" {
		if _, err := file.WriteString("Server: " + server + "\n"); err != nil {
			return err
		}
	}

	return nil
}
