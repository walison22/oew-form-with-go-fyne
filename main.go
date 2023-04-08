package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app" 
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	res, _ := ioutil.ReadFile("myfilename.txt")

	type registrant struct {
		Name  string
		Email string
		Phone string
		Msg   string
	}

	var registrantData []registrant

	json.Unmarshal(res, &registrantData)

	a := app.New()

	w := a.NewWindow("OEW Registration Form")
	w.Resize(fyne.NewSize(400, 400))

	list := widget.NewList(
		func() int { return len(registrantData) },
		func() fyne.CanvasObject {
			return widget.NewLabel("OEW Registrant item widget")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(registrantData[lii].Name)

		},
	)
	label1 := widget.NewLabel("...")
	label2 := widget.NewLabel("...")
	label3 := widget.NewLabel("...")
	label4 := widget.NewLabel("...")

	list.OnSelected = func(id widget.ListItemID) {
		label1.Text = registrantData[id].Name
		label1.TextStyle = fyne.TextStyle{Bold: true}

		label2.Text = registrantData[id].Email
		label2.TextStyle = fyne.TextStyle{Bold: true}

		label3.Text = registrantData[id].Phone
		label3.TextStyle = fyne.TextStyle{Bold: true}

		label4.Text = registrantData[id].Msg
		label4.TextStyle = fyne.TextStyle{Bold: true}

		label1.Refresh()
		label2.Refresh()
		label3.Refresh()
		label4.Refresh()

	}

	title := canvas.NewText("INPUT YOUR DETAILS BELOW", color.White)
	title.TextSize = 12
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Resize(fyne.NewSize(300, 35))
	title.Move(fyne.NewPos(550, 100))

	e_name := widget.NewEntry()
	e_name.SetPlaceHolder("Enter Your Name..")
	e_name.Resize(fyne.NewSize(300, 35))
	e_name.Move(fyne.NewPos(500, 150))

	e_email := widget.NewEntry()
	e_email.SetPlaceHolder("Enter Your Email..")
	e_email.Resize(fyne.NewSize(300, 35))
	e_email.Move(fyne.NewPos(500, 200))

	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter Your Phone..")
	e_phone.Resize(fyne.NewSize(300, 35))
	e_phone.Move(fyne.NewPos(500, 250))

	e_msg := widget.NewMultiLineEntry()
	e_msg.SetPlaceHolder("leave a message for us..")
	e_msg.MultiLine = true
	e_msg.Resize(fyne.NewSize(300, 140))
	e_msg.Move(fyne.NewPos(500, 300))

	submit_btn := widget.NewButton("SUBMIT", func() {
		if e_name.Text != "" && e_email.Text != "" && e_phone.Text != "" && e_msg.Text != "" {

			obj1 := &registrant{
				Name:  e_name.Text,
				Email: e_email.Text,
				Phone: e_phone.Text,
				Msg:   e_msg.Text,
			}
			registrantData = append(registrantData, *obj1)
			b, _ := json.MarshalIndent(registrantData, "", "")
			os.WriteFile("myfilename.txt", b, 0644)

			e_name.Text = ""
			e_email.Text = ""
			e_phone.Text = ""
			e_msg.Text = ""

			e_name.Refresh()
			e_email.Refresh()
			e_phone.Refresh()
			e_msg.Refresh()

		} else {
			label1.Text = "Incorrect data"
			label2.Text = ""
			label3.Text = ""
			label4.Text = ""
			label2.Refresh()
			label3.Refresh()
			label4.Refresh()
			label1.Refresh()

		}
	})
	submit_btn.Resize(fyne.NewSize(90, 50))
	submit_btn.Move(fyne.NewPos(500, 445))

	image := canvas.NewImageFromFile("C:/Users/Emmanuel Wali Okocha/Desktop/bg/sunset-g6b6d1af0e_1920.jpg")
	image.Resize(fyne.NewSize(1400, 850))

	w.SetContent(
		container.NewWithoutLayout(
			image,
			title,
			e_name,
			e_email,
			e_phone,
			e_msg,
			submit_btn,
		),
	)

	w.ShowAndRun()
}
