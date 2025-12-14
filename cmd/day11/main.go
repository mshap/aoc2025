package main

import (
	"aoc2025/internal/files"
	"fmt"
	"strings"
)

type device struct {
	name     string
	connStr  []string
	connects []*device
}

var out *device = &device{name: "out"}

func buildConnections(devices []device) map[string]*device {
	connections := make(map[string]*device)

	for i := range devices {
		connections[devices[i].name] = &devices[i]
	}

	for i := range devices {
		for _, connName := range devices[i].connStr {
			if connName == "out" {
				devices[i].connects = append(devices[i].connects, out)
			} else {
				if connDevice, ok := connections[connName]; ok {
					devices[i].connects = append(devices[i].connects, connDevice)
				}
			}
		}
	}

	return connections
}

func travel(device *device) int {
	if device == out {
		return 1
	}

	count := 0
	for _, conn := range device.connects {
		count += travel(conn)
	}
	return count
}

func main() {
	devices := files.ProcessFile("../../internal/input/day11.txt", false, true, func(line string) (device, error) {
		parts := strings.Split(line, " ")

		name := parts[0][0 : len(parts[0])-1]
		connects := []*device{}
		names := parts[1:]

		return device{name: name, connects: connects, connStr: names}, nil
	})

	devicemap := buildConnections(devices)

	paths := travel(devicemap["you"])

	fmt.Printf("total paths %d\n", paths)
}
