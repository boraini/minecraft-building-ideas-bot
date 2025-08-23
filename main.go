package main

import "math/rand/v2"

func main() {
	config := LoadConfig()
	totalFrequency, ideas := LoadData(config)

	var selected *BuildingIdea

	point := rand.Int64N(totalFrequency)

	currentFrequency := int64(0)

	for _, idea := range ideas {
		currentFrequency += idea.Frequency

		if currentFrequency >= point {
			selected = &idea
			break
		}
	}

	if selected == nil {
		selected = &ideas[len(ideas)-1]
	}

	PostBuildingIdea(*selected, config)
}
