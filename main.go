package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/getlantern/systray"
)

// Program: GoTrayLogger - A Go application with system tray and logging features
func main() {
	// Open or create log file
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Set up logger
	logger := log.New(logFile, "", log.LstdFlags)

	// Start message logging goroutine with file output
	go func() {
		for {
			message := "Hello! This message appears every 10 seconds."
			fmt.Println(message) // For initial console (if visible)
			logger.Println(message)
			time.Sleep(10 * time.Second)
		}
	}()

	// Start shutdown signal checker
	go func() {
		for {
			if _, err := os.Stat("shutdown.txt"); err == nil {
				logger.Println("Shutdown signal detected (shutdown.txt found). Exiting...")
				os.Exit(0)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Start system tray
	if err := startSystemTray(logger); err != nil {
		logger.Printf("Failed to start system tray: %v", err)
		fmt.Printf("System tray failed to start: %v\nRunning in background. Create 'shutdown.txt' in the program directory to exit.\n", err)
		// Keep the program running even if system tray fails
		select {}
	}
}

func startSystemTray(logger *log.Logger) error {
	// Read icon file
	iconData, err := os.ReadFile("icon.ico")
	if err != nil {
		return fmt.Errorf("error reading icon file: %v", err)
	}

	// Start system tray
	systray.Run(func() {
		// Set system tray icon
		systray.SetIcon(iconData)
		systray.SetTitle("GoTrayLogger") // Updated to match the suggested name
		systray.SetTooltip("GoTrayLogger: Showing message every 10 seconds")

		// Add menu items
		mShowConsole := systray.AddMenuItem("Show Console", "Show the console window")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit the application")

		// Handle menu item clicks
		go func() {
			for {
				select {
				case <-mShowConsole.ClickedCh:
					// Open a new PowerShell window that tails the log file in real-time
					cmd := exec.Command("powershell.exe", "-NoExit", "-Command", "Get-Content app.log -Wait")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Start(); err != nil {
						fmt.Printf("Error opening console: %v\n", err)
						logger.Printf("Error opening console: %v", err)
					} else {
						logger.Println("Console opened successfully")
					}
				case <-mQuit.ClickedCh:
					logger.Println("Quit selected from system tray. Exiting...")
					systray.Quit()
				}
			}
		}()
	}, func() {
		logger.Println("Application exiting...")
	})

	return nil
}
