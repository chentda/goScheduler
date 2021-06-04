package config

type Details struct {
	Input input
	Time  time
}

type input struct {
	SourceScript string
}

type time struct {
	Hour   int
	Minute int
}
