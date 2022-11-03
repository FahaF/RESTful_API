package data

import "GoLangProject/model"

var Books = []model.Book{
	{
		ID:    "1",
		Isbn:  "438227",
		Title: "Amar Bondhu Rashed",
		Author: &model.Author{
			Firstname: "Md Jafor", Lastname: "Iqbal",
		},
	},
	{
		ID:    "2",
		Isbn:  "454555",
		Title: "Thousand Splendid Sun",
		Author: &model.Author{
			Firstname: "Khaled", Lastname: "Hosseini",
		},
	},
}
