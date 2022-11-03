package data

import "GoLangProject/model"

var Books = []model.Book{
	{
		ID:    "1",
		Isbn:  "438227",
		Title: "Book One",
		Author: &model.Author{
			Firstname: "John", Lastname: "Doe",
		},
	},
	{
		ID:    "2",
		Isbn:  "454555",
		Title: "Book Two",
		Author: &model.Author{
			Firstname: "Steve", Lastname: "Smith",
		},
	},
}
