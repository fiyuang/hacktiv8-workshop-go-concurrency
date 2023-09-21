<div align="center">
      <h1><br/>Go Concurrency (Hacktiv8 Workshop)</h1>
     </div>


# Description
A simple update stock demo with goroutine, sync.Mutex, sync.Waitgroup implementation
 
# Tech Used
 ![GO](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
      
# Getting Start:
Before you running the program, make sure you've run this command:
- `go mod tidy`

### Run the program
- `go run main.go`
- `go run -race main.go` for build-in race detector

The program will run on http://localhost:9090

### API Route List
| Method | URL | Description |
| ----------- | ----------- | ----------- | 
| GET | localhost:8000/update-stock  | Update Stock Demo |
| GET | localhost:8000/update-stock-race  | Update Stock Demo with Race Condition Example |
      
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->
