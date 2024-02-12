package app

type Environment struct {
	Development string
	Production  string
}

var Env = Environment{
	Development: "development",
	Production:  "production",
}

type Metadata struct {
	Environment string
	Name        string
	Version     string
}
