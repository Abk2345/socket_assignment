package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

//definig structurw of book
type Book struct {
    Id      string   
    Title   string 
    Desc    string 
    Author  string 
}

func main() {
	//opens file
	file, err := os.Open("books.csv")
	//error handling
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	//after the work, closing the file
	defer file.Close()

	//reading csv file
	reader := csv.NewReader(file)
	books, err := reader.ReadAll()
	//handling error
	if err != nil {
		fmt.Println("Failed to read CSV data:", err)
		return
	}

	//slice struct to store all the details
	var data []Book
	for _, r := range books {
		//adding to data slice sruct
		data = append(data, Book{Id: r[0], Title: r[1], Desc: r[2], Author:r[3]})
	}

	//printing read data to the output stream
	fmt.Println("Id\tTitle\tTitle\tAuthor")
	fmt.Println("====\t===\t====\t====")
	for _, d := range data {
		fmt.Printf("%s\t%s\t%s\t%s\n", d.Id, d.Title, d.Desc, d.Author)
	}
}
