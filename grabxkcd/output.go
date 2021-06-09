package grabxkcd

import (
	"encoding/json"
	"fmt"
)

// Output is the JSON output of this app
type Output struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func printJSON(ar APIResponse) error {
	o := Output{
		Title:       ar.Title,
		Number:      ar.Number,
		Date:        ar.Date().Format("2006-01-02"),
		Description: ar.Description,
		Image:       ar.Image,
	}
	b, err := json.MarshalIndent(&o, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func prettyPrint(ar APIResponse) error {
	_, err := fmt.Printf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		ar.Title,
		ar.Number,
		ar.Date().Format("02-01-2006"),
		ar.Description,
		ar.Image,
	)
	return err
}
