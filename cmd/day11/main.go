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

func containsRequired(queue []string) (bool, bool) {
	hasDac, hasFft := false, false
	for _, name := range queue {
		switch name {
		case "dac":
			hasDac = true
		case "fft":
			hasFft = true
		}
	}

	return hasDac, hasFft
}

func travel(device *device, queue []string, visited *map[string]int) int {
	hasDac, hasFft := containsRequired(queue)

	key := fmt.Sprintf("%s-%t-%t", device.name, hasDac, hasFft)

	if device == out {
		if hasDac && hasFft {
			return 1
		} else {
			return 0
		}
	}

	if val, ok := (*visited)[key]; ok {
		return val
	}

	queue = append(queue, device.name)
	count := 0
	for _, conn := range device.connects {
		count += travel(conn, queue, visited)
	}

	(*visited)[key] = count

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

	paths := travel(devicemap["svr"], []string{}, &map[string]int{})

	fmt.Printf("total paths %d\n", paths)
}
