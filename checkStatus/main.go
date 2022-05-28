package main

import "fmt"

const (
	passStatus = "pass"
	failStatus = "fail"
)

type HealthCheck struct {
	serviceId int
	status    string
}

func main() {
	list := GenerateCheck()
	for _, health := range list {
		if health.status == passStatus {
			fmt.Println(health.serviceId, health.status)
		}
	}
}

func GenerateCheck() []HealthCheck {
	var listOut []HealthCheck
	var status = ""
	for i := 0; i < 5; i++ {
		switch i % 2 {
		case 0:
			status = passStatus
		default:
			status = failStatus
		}
		check := HealthCheck{
			serviceId: i,
			status:    status,
		}
		listOut = append(listOut, check)
	}

	return listOut
}
