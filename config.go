package main

var Config = struct {
	APPName    string
	AppVersion string

	DB struct {
		Host     string
		Name     string
		User     string
		Password string
		Port     uint
	}
}{}
