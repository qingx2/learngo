package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xingzuo    string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	bytes, e := json.Marshal(o)

	if e != nil {
		return profile, e
	}

	e = json.Unmarshal(bytes, &profile)

	return profile, e
}
