package main

var books []Book = []Book{
	{
		ID:     "1",
		ISBN:   "223321",
		Title:  "BookOne",
		Author: &Author{Firstname: "Jahid", Lastname: "Hasan"},
	},
	{
		ID:     "2",
		ISBN:   "234523",
		Title:  "BookTwo",
		Author: &Author{Firstname: "Arman", Lastname: "Hasan"},
	},
	{
		ID:     "3",
		ISBN:   "998338",
		Title:  "ThreeBook",
		Author: &Author{Firstname: "Selim", Lastname: "Reza"},
	},
	{
		ID:     "4",
		ISBN:   "394839",
		Title:  "FourBook",
		Author: &Author{Firstname: "Md.", Lastname: "Mijan"},
	},
}
