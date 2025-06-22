package models

type Route struct {
	Prefix  string   `json:"prefix"`
	Methods []string `json:"methods"`
	Target  string   `json:"target"`
}

type Config struct {
	Rotues []Route `json:"routes"`
}
