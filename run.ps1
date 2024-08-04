# Navigate to the backend directory and run the Go application
Set-Location -Path "backend"
Start-Process -FilePath "go" -ArgumentList "run main.go"

# Navigate back to the root directory
Set-Location -Path ".."

# Navigate to the frontend directory and start the npm server
Set-Location -Path "frontend"
# npm run start BROWSER=none
Start-Process -FilePath "npm" -ArgumentList "run start BROWSER=none"

Set-Location -Path ".."