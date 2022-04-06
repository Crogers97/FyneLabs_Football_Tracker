package main

// import fyne
import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//Fixture structure

	type Fixture struct {
		Name1 string
		Name2 string
		Score string
	}
	// Now create a slice/ array to store the data
	var myFixtureData []Fixture

	//lets read data from file and display it in a list
	data_from_file, _ := ioutil.ReadFile("Fixtures.txt")

	json.Unmarshal(data_from_file, &myFixtureData)

	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("Football Score Tracker")
	// resize window
	w.Resize(fyne.NewSize(400, 400))

	//New label to display team names and score
	//team1
	t_name1 := widget.NewLabel("")
	t_name1.TextStyle = fyne.TextStyle{Bold: true}

	//team2
	t_name2 := widget.NewLabel("")
	t_name2.TextStyle = fyne.TextStyle{Bold: true}

	//score
	t_score := widget.NewLabel("")

	// entry widget for team1 name
	e_name1 := widget.NewEntry()
	e_name1.SetPlaceHolder("Enter team1 name here...")

	// entry widget for team2 name
	e_name2 := widget.NewEntry()
	e_name2.SetPlaceHolder("Enter team2 name here...")
	// entry widget for score
	e_score := widget.NewEntry()
	e_score.SetPlaceHolder("Enter Score here...")
	// submit button
	submit_btn := widget.NewButton("Submit", func() {

		//logic part- store data in json format
		//let create a struct for our data
		//Get data from entry widget and push to slice
		myData1 := &Fixture{
			Name1: e_name1.Text, //data from name entry widget
			Name2: e_name2.Text, //data from name entry widget
			Score: e_score.Text, //data from score entry widget
		}

		myFixtureData = append(myFixtureData, *myData1)
		// convert data to json
		final_data, _ := json.MarshalIndent(myFixtureData, "", " ")
		//writing data to file
		// it takes 3 things
		// file name .txt or .json

		ioutil.WriteFile("Fixtures.txt", final_data, 0644)

		e_name1.Text = ""
		e_name2.Text = ""
		e_score.Text = ""

		e_name1.Refresh()
		e_name2.Refresh()
		e_score.Refresh()

	})

	//Delete button
	delete_btn := widget.NewButton("Delete", func() {
		//Create a new slice
		var TemporaryData []Fixture
		// loop through the slice, push all the data except the one to be deleted
		// i is the indexs and e is the element of the slice, i will push all the data to temporary data
		for _, e := range myFixtureData {

			//t_name is the label created to show details

			if t_name1.Text != e.Name1 {
				TemporaryData = append(TemporaryData, e)
			}
		}

		// add all the data back to the main slice myFixtureData
		myFixtureData = TemporaryData
		//convert to json
		result, _ := json.MarshalIndent(myFixtureData, "", " ")
		//write to file
		// first argument is the file name, second is the data, third is the permission

		ioutil.WriteFile("Fixtures.txt", result, 0644)
	})

	//list widget
	list := widget.NewList(
		// first argument is item count
		// len() function to get myFixtureData length

		func() int { return len(myFixtureData) },
		// 2nd argument is for widget choice. I want to use label
		func() fyne.CanvasObject { return widget.NewLabel("") },
		//3rd argument is to update widget with our new data
		func(x widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myFixtureData[x].Name1 + " vs " + myFixtureData[x].Name2 + " " + myFixtureData[x].Score)
		},
	)

	// update on clicked/ selected

	list.OnSelected = func(id widget.ListItemID) {
		t_name1.Text = myFixtureData[id].Name1
		t_name1.Refresh()
		t_name2.Text = myFixtureData[id].Name2
		t_name2.Refresh()
		t_score.Text = myFixtureData[id].Score
		t_score.Refresh()

	}
	//Update button
	update_button := widget.NewButton("Update", func() {
		// Temporary slice
		var TemporaryData []Fixture
		// the data that needs to be updated
		update := &Fixture{
			Name1: e_name1.Text,
			Name2: e_name2.Text,
			Score: e_score.Text,
		}

		//loop through the slice and update the data
		//_ ignore the index

		for _, e := range myFixtureData {
			// checking data criteria
			if t_name1.Text == e.Name1 {
				//if the data matches, update the data
				TemporaryData = append(TemporaryData, *update)
			} else {
				//if the data doesn't match, push the data to temporary data
				TemporaryData = append(TemporaryData, e)
			}

		}

		myFixtureData = TemporaryData
		// convert dtata to json
		// 1st argument is data source, 2nd is prefix, 3rd is indent
		result, _ := json.MarshalIndent(myFixtureData, "", " ")
		// write data to file
		ioutil.WriteFile("Fixtures.txt", result, 0644)

		//refresh the data & empty box and refresh list
		e_name1.Text = ""
		e_name2.Text = ""
		e_score.Text = ""
		e_name1.Refresh()
		e_name2.Refresh()
		e_score.Refresh()

		list.Refresh() //refresh the list

	})

	// show and run
	w.SetContent(
		// container
		container.NewHSplit(
			// first argument is list of data
			list,

			container.NewVBox(t_name1, t_name2, t_score, e_name1, e_name2, e_score, submit_btn, delete_btn, update_button),
		),
	)
	w.ShowAndRun()
}
