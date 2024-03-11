package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// Function to execute shell commands
func executeCommand(command string) {
    cmd := exec.Command("sh", "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

// Function to get user input
func getUserInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: wifirand <SSID>")
        return
    }
    
    ssid := os.Args[1]
    password := getUserInput("Enter password for SSID '" + ssid + "': ")

    // Save SSID and password as environment variables
    os.Setenv("WIFI_SSID", ssid)
    os.Setenv("WIFI_PASSWORD", password)

    fmt.Println("SSID and password saved successfully.")
}

