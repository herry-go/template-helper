package user

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeAccountTab(_ fyne.Window) fyne.CanvasObject {
	//userAvatar := canvas.NewImageFromFile("data/images/app.jpg")
	//userAvatar.FillMode = canvas.ImageFillOriginal


	userAvatar := widget.NewCard("", "", nil)
	userAvatar.Image = canvas.NewImageFromFile("data/images/app.jpg")

	username := widget.NewLabel("用户名")
	username.TextStyle = fyne.TextStyle{Bold: true}

	email := widget.NewLabel("用户邮箱")
	email.TextStyle = fyne.TextStyle{Italic: true}

	return container.NewGridWithColumns(2,
		container.NewVBox(userAvatar),
		container.NewVBox(username,email,widget.NewLabel(""),  widget.NewLabel("")),
		)
}

func MakePasswordTab(_ fyne.Window) fyne.CanvasObject {
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("请输入密码")

	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "密码", Widget: password},
		},
	}
	submitButton := &widget.Button{
		Text:       "   确认  ",
		Importance: widget.HighImportance,
		OnTapped:   func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   password.Text,
				Content: largeText.Text,
			})
		},
	}

	cancelButton := &widget.Button{
		Text:       "   取消   ",
		Importance: widget.LowImportance,
		OnTapped: func() {
			password.SetText("")
		},
	}

	buttonRow := container.NewHBox(submitButton, cancelButton)
	centeredButtons := container.NewCenter(buttonRow)
	content := container.NewVBox(
		form,
		centeredButtons,
	)

	return content
}
