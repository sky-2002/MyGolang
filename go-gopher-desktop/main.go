package main

import (
	"fmt"
	"image/color"

	// "fyne.io/fyne" // this line solved the issue
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

// func loadJsonData() []string {
// 	fmt.Println("Loading data from JSON file")

// 	input, _ := ioutil.ReadFile("data.json")
// 	var data []string
// 	json.Unmarshal(input, &data)

// 	return data
// }

// func saveJsonData(data binding.StringList) {
// 	fmt.Println("Saving data to JSON file")
// 	d, _ := data.Get()
// 	jsonData, _ := json.Marshal(d)
// 	ioutil.WriteFile("data.json", jsonData, 0644)

// }

func greet() {
	fmt.Println("Hello there !!")
}

func main() {
	// a := app.New()
	// fmt.Println("App created")
	// w := a.NewWindow("Hello") // only creating a window works fine
	// fmt.Println("Window created")

	// w.SetContent(widget.NewButton("Button1", func() { return })) // this also works fine
	// w.SetContent(widget.NewButton("Button2", greet))             // now when this one is added, it is overlapping the first one
	// w.SetContent(widget.NewLabel("Hello, this is a label")) // set content sets the content of this window
	// fmt.Println("Widget added")

	// hello := widget.NewLabel("Hello Fyne!")
	// VBox is vertical box, HBox will be horizontal box
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))
	// w.ShowAndRun()

	myApp := app.New() // returns a new application instance
	fmt.Println("App created")

	// creates a new window for the application, also the first window is the master window
	// myWindow := myApp.NewWindow("Gopher") // Gopher is the title of the window
	myWindow := myApp.NewWindow("List Data")
	fmt.Println("Window created")

	// loadedData := loadJsonData()

	// data := binding.NewStringList()
	// data.Set(loadedData)

	// defer saveJsonData(data)

	// list := widget.NewListWithData(data,
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i binding.DataItem, o fyne.CanvasObject) {
	// 		o.(*widget.Label).Bind(i.(binding.String))
	// 	})

	// list.OnSelected = func(id widget.ListItemID) {
	// 	list.Unselect(id)
	// 	d, _ := data.GetValue(id)
	// 	w := myApp.NewWindow("Edit Data")

	// 	itemName := widget.NewEntry()
	// 	itemName.Text = d

	// 	updateData := widget.NewButton("Update", func() {
	// 		data.SetValue(id, itemName.Text)
	// 		w.Close()
	// 	})

	// 	cancel := widget.NewButton("Cancel", func() {
	// 		w.Close()
	// 	})

	// 	deleteData := widget.NewButton("Delete", func() {
	// 		var newData []string
	// 		dt, _ := data.Get()

	// 		for index, item := range dt {
	// 			if index != id {
	// 				newData = append(newData, item)
	// 			}
	// 		}

	// 		data.Set(newData)

	// 		w.Close()
	// 	})

	// 	w.SetContent(container.New(layout.NewVBoxLayout(), itemName, updateData, deleteData, cancel))
	// 	w.Resize(fyne.NewSize(400, 200))
	// 	w.CenterOnScreen()
	// 	w.Show()

	// }

	// add := widget.NewButton("Add", func() {
	// 	w := myApp.NewWindow("Add Data")

	// 	itemName := widget.NewEntry()

	// 	addData := widget.NewButton("Add", func() {
	// 		data.Append(itemName.Text)
	// 		w.Close()
	// 	})

	// 	cancel := widget.NewButton("Cancel", func() {
	// 		w.Close()
	// 	})

	// 	w.SetContent(container.New(layout.NewVBoxLayout(), itemName, addData, cancel))
	// 	w.Resize(fyne.NewSize(400, 200))
	// 	w.CenterOnScreen()
	// 	w.Show()

	// })

	// exit := widget.NewButton("Quit", func() {

	// 	myWindow.Close()
	// })

	// myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	// myWindow.Resize(fyne.NewSize(400, 600))
	// myWindow.SetMaster()
	// myWindow.CenterOnScreen()
	// myWindow.ShowAndRun()
	// // main menu
	fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() { myApp.Quit() })) // also used myApp as name
	fmt.Println("FileMenu created")

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close",
				container.NewVBox(widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
					widget.NewLabel("Version: v0.1"),
					widget.NewLabel("Author: Aurélie Vache"),
				), myWindow)
		}))
	fmt.Println("HelpMenu created")

	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu) // also used myWindow as name
	fmt.Println("MainMenu set")

	// // Define a welcome text centered
	text := canvas.NewText("Display a random Gopher!", color.White)
	text.Alignment = fyne.TextAlignCenter
	fmt.Println("Text added")

	// // Define a Gopher image
	var resource, _ = fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0
	fmt.Println("Image added")

	// // Define a "random" button
	randomBtn := widget.NewButton("Random", func() {
		resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
		gopherImg.Resource = resource

		//Redrawn the image with the new path
		gopherImg.Refresh()
	})
	randomBtn.Importance = widget.HighImportance
	fmt.Println("Button added")

	// // Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		randomBtn,
	)

	// Display our content
	myWindow.SetContent(box)
	fmt.Println("COntent set")

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
