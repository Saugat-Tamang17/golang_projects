package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

// emailRegex is used to perform a basic syntax validation
// on an email address before deeper verification.
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

// main is the entry point of the program.
// It performs email verification using syntax checking,
// MX record lookup, and SMTP probing.
func main() {
	email := "saugatm814@gmai.com"

	// Step 1: Validate email syntax using regex.
	if !isValidSyntax(email) {
		fmt.Println("Invalid email syntax")
		return
	}

	// Step 2: Split email into local and domain parts.
	local, domain, err := splitEmail(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Local part:", local)
	fmt.Println("Domain part:", domain)

	// Step 3: Lookup MX records for the domain.
	mxRecords, err := lookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 4: Attempt SMTP verification using the highest-priority MX.
	verifySMTP(email, mxRecords[0].Host)
}

// isValidSyntax checks whether the email matches
// a basic RFC-compliant pattern.
func isValidSyntax(email string) bool {
	return emailRegex.MatchString(email)
}

// splitEmail separates an email address into local and domain parts.
// It ensures exactly one '@' symbol exists.
func splitEmail(email string) (string, string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid email structure")
	}
	return parts[0], parts[1], nil
}

// lookupMX retrieves MX records for a given domain.
// It ensures at least one mail server is available.
func lookupMX(domain string) ([]*net.MX, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return nil, fmt.Errorf("no MX records found for domain")
	}

	fmt.Println("MX records found:")
	for _, mx := range mxRecords {
		fmt.Printf("Host: %s | Preference: %d\n", mx.Host, mx.Pref)
	}

	return mxRecords, nil
}

// verifySMTP attempts to validate the email address by
// communicating with the domain's mail server via SMTP.
func verifySMTP(email, mxHost string) {
	conn, err := net.DialTimeout("tcp", mxHost+":25", 10*time.Second)
	if err != nil {
		fmt.Println("SMTP connection failed")
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read server greeting.
	reader.ReadString('\n')

	// Helper function to send SMTP commands.
	send := func(cmd string) {
		fmt.Fprintf(conn, "%s\r\n", cmd)
	}

	send("HELO example.com")
	reader.ReadString('\n')

	send("MAIL FROM:<test@example.com>")
	reader.ReadString('\n')

	send("RCPT TO:<" + email + ">")
	response, _ := reader.ReadString('\n')

	// Interpret SMTP response codes.
	if strings.HasPrefix(response, "250") {
		fmt.Println("Email is likely valid")
	} else if strings.HasPrefix(response, "550") {
		fmt.Println("Email does NOT exist")
	} else {
		fmt.Println("Email verification inconclusive")
	}
}
