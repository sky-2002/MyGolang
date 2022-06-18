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

// const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
const KuteGoAPIURL = "data:image/jpeg;base64,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rbwK8mSjPm/IsPkZB9cSO1luo3ppCoxuU6lhmmth4J/evnlyyq89xyNp0fZTUUqnwYVmq2sbrK2JZrNzHNq2EDvVZskt1ycMFPKVatePpi9zVtLOWVUmtuxowJzESuagiIgCIiAIiIAiIgHEzPaK57bV0quUr7s3al1JVzWW2pWrDmgYh9zDmFQgYLZGmmd7U9njql3VvtsCNWVP7rUVl1dqreRIVtu3cOYDt1yQeZJtbHq53I3Cra7KUapNtW3FS7Qq7ASFZVHRGADD2MOQldxiik2DZ3o1jKCo0+03uoPI2q39WUzkZs5DmFIMsBVrrfRSldMOjW2Mjlf8KuskP7CxUDl6LdJd8K4XXplKpuLMd1ljHdbc2MbnbxPgOgAwAAABIKdKSlqexLUlGUdOMmSq0nEK0zfplfmedDqXAzyLVORg4xna7884E6V4nQWCM4Rz0rsDVWn3JYFY/UJ6JOnUUJYpV1VlPVWAKn3gy9GtJcmZUsact1sY3ES7fspoTnbQteeZ7lnqOfP+qZec6/6K0j1bdWo8u9d/zs3EyVXC7oqy6a+zRURLP+ia/wDm9YP/AI/60z7HZWvGG1Grb7arn8Cie/eF4PPw6flFSJH1Wtqqx3tlaEnADuqsx8lBOSfdNCnZLRg+ktz56h773X8DOV/KWGg4TptOMUUU1/QRVz79o59Zy7jwiSPTf5Mx9DX3fw+muYc/6xwaah9qwb2HtVGEkarsjqLUzZqELg7hpwhGkceKXHPeWA+eVXx2HodtBkUqkpclula06e6W/kynDdUtildndvWRXZScZpYDIXlyKkYKkciCCPZWcY1uBe4Upfo1OqrOci2nazEZx6jqjoyn1WUHnhWmm4twUXMLa37rUKu0WgbldM52Wpkb0zkjmCpJwRk5rk7MPbalursrcIpVKqleut8sjnvizsbFDIpC8gMHO7MpKi1PK4L7qZjhmoRsgHzAP3z7iJZIRERAEREAREQBERAEREAREQBERAEREAREQBERAE4nMQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQD//Z"

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
	// var resource, _ = fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
	var resource, _ = fyne.LoadResourceFromURLString(KuteGoAPIURL)
	gopherImg := canvas.NewImageFromResource(resource)
	fmt.Println(resource)
	fmt.Println(gopherImg)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0
	fmt.Println("Image added")

	// // Define a "random" button
	randomBtn := widget.NewButton("Random", func() {
		// resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
		resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL)
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
