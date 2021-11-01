package storage

//Male
type Male struct {
	Name    string
	Surname string
	Age     int
}

//Female
type Female struct {
	Name    string
	Surname string
	Age     int
}

//People
type Team struct {
	Name    string
	Surname string
	Age     int
}


//Males. Array of utils. Male type
var Males = []Male{
	{"Javier", "Miguez", 44},
	{"Luis", "Alonso", 32},
	{"Mariano", "Gomez", 47},
}

//Females. Array of utils. Female type
var Females = []Female{
	{"Cristina", "Perez", 21},
	{"Alicia", "Domenech", 32},
	{"Susana", "Gomez", 19},
}