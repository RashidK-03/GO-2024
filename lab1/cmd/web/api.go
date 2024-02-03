package api

type Character struct {
	Id       int    `json:"id"`
	Name     string `json:"fname"`
	LastName string `json:"lname"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Actor    string `json:"actor"`
}

var Characters = []Character{
	{1, "Walter", "White", 52, "09-07-1958", "Bryan Cranston"},
	{2, "Jesse", "Pinkman", 25, "09-24-1984", "Aaron Paul"},
	{3, "Skyler", "White", 50, "08-11-1970", "Anna Gunn"},
	{4, "Henry", "Schrader", 50, "unknown", "Dean Norris"},
	{5, "Mike", "Ehrmantraut", 60, "unknown", "Jhonathan Banks"},
	{6, "Saul", "Goodman", 42, "unknown", "Bob Odenkirk"},
	{7, "Gustavo", "Fring", 58, "unknown", "Giancarlo Esposito"},
	{8, "Hector", "Salamanca", 42, "unknown", "Mark Margolis"},
	{9, "WalterJr", "White", 17, "07-08-1993", "RJ Mitte"},
	{10, "Marie", "Schreder", 42, "unknown", "Betsy Brandit"},
}
