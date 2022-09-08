# Anomaly Detection - Golang 

This project is an example of an anomaly detection system based on data collected in CSV.

The anomaly detection route returns all values ​​read from the CSV with the Level of the log:
- Info
- Warn
- Alert

### Instalation and Testing

Runing project:
```
go mod vendor
```
```
go run main.go
```

Local Routes
```
http://localhost:8080/
http://localhost:8080/anomaly-detection
```

Example Response Api
```
{
  timestamp: "2014-02-09 10:30:00",
  value: "30.72721278",
  level: "Alert"
}
```
