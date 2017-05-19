//
// location.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

package instagram

import "encoding/json"

// A Location describes an Instagram location info.
type Location struct {
	ID         string
	Name       string
	PublicPage bool
	Lat        float64
	Lng        float64
	Slug       string
}

func getFromLocationPage(data []byte) (Location, bool) {
	var locationJson struct {
		Location struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			HasPublicPage bool `json:"has_public_page"`
			Lat           float64 `json:"lat"`
			Lng           float64 `json:"lng"`
			Slug          string `json:"slug"`
		} `json:"location"`
	}

	err := json.Unmarshal(data, &locationJson)
	if err != nil {
		return Location{}, false
	}

	location := Location{}
	location.ID = locationJson.Location.ID
	location.Name = locationJson.Location.Name
	location.PublicPage = locationJson.Location.HasPublicPage
	location.Lat = locationJson.Location.Lat
	location.Lng = locationJson.Location.Lng
	location.Slug = locationJson.Location.Slug

	return location, true
}
