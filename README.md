# GoSystemTrayLogger
**GoTrayLogger** is a lightweight Go application that displays a message every 10 seconds, integrates with the system tray, and logs output to a file. Designed for Windows (with potential cross-platform support), it’s ideal for developers looking to create or experiment with system tray applications, periodic tasks, or logging utilities in Go.

## Features
- **System Tray Integration**: Adds a tray icon with options to show the console or quit the application.
- **Periodic Messaging**: Displays "Hello! This message appears every 10 seconds." at 10-second intervals.
- **Logging**: Writes logs with timestamps to `app.log` for easy monitoring.
- **Console Access**: Opens a PowerShell window to view live logs via the tray menu.
- **Background Execution**: Runs silently in the background with a shutdown option via `shutdown.txt`.

## Installation

### Prerequisites
- **Go**: Version 1.20 or later (check with `go version`).
- **Git**: For cloning the repository.
- **Windows**: Currently optimized for Windows (Linux/macOS support may require additional setup).
- **CGO**: Enabled for system tray functionality (`export CGO_ENABLED=1`).

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/GoTrayLogger.git
   cd GoTrayLogger

**2. Steps Install Dependencies:**
go get github.com/getlantern/systray

**3. Build the executable:**
go build -ldflags "-H=windowsgui" -o GoTrayLogger.exe

**Run the application:**
.\GoTrayLogger.exe

**Usage**

Launch: Run GoTrayLogger.exe to start the application. A system tray icon will appear.
Show Console: Right-click the tray icon and select "Show Console" to open a PowerShell window displaying live logs from app.log.
Quit: Right-click the tray icon and select "Quit" to exit the application.
Shutdown: Create an empty shutdown.txt file in the same directory to stop the program gracefully:
echo. > shutdown.txt

**File Structure**

main.go: The main Go source file containing the application logic.
icon.ico: The system tray icon file (replace with your own .ico file if desired).
app.log: The log file generated during execution (created automatically).

**Configuration**

Icon: Replace icon.ico with a custom 32x32 or 64x64 .ico file for the tray icon.
Message: Modify the message variable in main.go to change the displayed text.
Interval: Adjust the time.Sleep(10 * time.Second) value to change the message interval.
