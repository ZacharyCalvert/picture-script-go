package main

import "strings"

const (
	movie   = "mov"
	picture = "pic"
)

// TypeMap is mean to facilitate a file extension (jpg, mp4, etc) to a file type of movie or pic
type TypeMap struct {
	Mappings map[string]string
}

// GetType is given a file extension and maps it to a type of file (pic, mov)
func (mapping *TypeMap) GetType(fileExtension string) (string, bool) {
	search := strings.ToLower(fileExtension)
	if val, ok := mapping.Mappings[search]; ok {
		return val, true
	}
	return "", false
}

// GetTypeMapping provides the default known file extension maps
func GetTypeMapping() *TypeMap {
	supportedTypes := make(map[string]string)
	supportedTypes["m4v"] = movie
	supportedTypes["mp4"] = movie
	supportedTypes["png"] = picture
	supportedTypes["gif"] = picture
	supportedTypes["bmp"] = picture
	supportedTypes["jpeg"] = picture
	supportedTypes["jpg"] = picture
	supportedTypes["mov"] = movie
	supportedTypes["cr2"] = movie
	supportedTypes["avi"] = movie
	supportedTypes["mpg"] = movie
	return &TypeMap{supportedTypes}
}
