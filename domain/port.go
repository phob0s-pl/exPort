package domain

type Port struct {
	Key         string
	Name        string
	Coordinates []float64
	City        string
	Province    string
	Country     string
	Alias       []string
	Regions     []string
	Timezone    string
	Unlocs      []string
	Code        string
}
