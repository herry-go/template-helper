package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
	"template-helper/gui/common"
	"template-helper/internal"
)

var Win *Window

type Window struct {
	app fyne.App
	win fyne.Window
	tabs *common.CustomDocTabs
}

func (w *Window) Run() {

	w.app = app.NewWithID(common.APP_KEY)
	w.app.Settings().SetTheme(theme.LightTheme())

	w.win = w.app.NewWindow(common.APP_NAME)
	w.initUI()

	w.win.Resize(fyne.NewSize(1000, 600))
	w.win.CenterOnScreen()
	w.win.ShowAndRun()
}

func (w *Window) initUI() {

	//toolBar := w.makeToolBar()




	tabs := common.NewCustomDocTabs()
	tabs.SetTabLocation(container.TabLocationTop)
	tabs.OnClosed = func(tab *container.TabItem) {
		if len(tabs.Items) == 0 {
			w.win.SetContent(w.makeWelcome())
			common.GetStorage().SetPath("")
		}
	}

	tabs.OnSelected = func(tab *container.TabItem) {
		filePath := tabs.GetFilePath(tab)
		if filePath != "" {
			fmt.Println("Selected file path:", filePath)
			common.GetStorage().SetPath(filePath)
		}

	}

	// 使用 VBox 将 toolBar 和 tabs 垂直布局
	//content := container.NewVBox(
	//	//toolBar,
	//	tabs,
	//)
	w.setMenu()
	w.tabs = tabs
	w.win.SetContent(w.makeWelcome())

}

func (w *Window) setMenu()  {
	w.win.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open File...", func() {
				w.showOpenYamlFileDialog(w.win)
			}),
			fyne.NewMenuItem("Save", func() {
				w.saveYamlFile()
			}),
		)))
}

func (w *Window) makeWelcome() fyne.CanvasObject {
	//logo := canvas.NewImageFromResource(data.FyneLogoTransparent)
	//logo.FillMode = canvas.ImageFillContain
	//logo.SetMinSize(fyne.NewSize(128, 128))
	welcome := container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the template helper", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		//logo,
		container.NewHBox(
			widget.NewHyperlink("部署运维助手", parseURL("https://doc.uniontech.com/file/KrkEVzYrvQSGJ2AJ")),
			widget.NewLabel("-"),
			widget.NewHyperlink("模板指导手册", parseURL("https://doc.uniontech.com/docs/vVqRVeN2vZhyrQqy/ \"统信集中域管平台_运维助手模板编写_指导手册_V1.0.0\"")),
		),
		widget.NewLabel(""),
	))
	return welcome
}

func (w *Window) makeNavSplit(key string) fyne.CanvasObject {
	content := container.NewStack()
	header := widget.NewCard(key, "", widget.NewSeparator())
	headerBox := container.NewVBox(header)
	containerRright := container.NewBorder(
		headerBox, nil, nil, nil, content)

	setMenu := func(t Menu) {
		header.Title = t.Title
		header.Refresh()

		content.Objects = []fyne.CanvasObject{t.View(w.win)}
		content.Refresh()
	}

	split := container.NewHSplit(w.makeNav(setMenu, false, key), containerRright)
	split.Offset = common.SplitLeft
	return split
}

func (w *Window) makeNav(setMenu func(menu Menu), loadPrevious bool, key string) fyne.CanvasObject {
	a := fyne.CurrentApp()
	menuIndex := MenuIndexMap[key]

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return menuIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := menuIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := Menus[uid]
			if !ok {
				fyne.LogError("Missing panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
		},
		OnSelected: func(uid string) {
			if t, ok := Menus[uid]; ok {
				a.Preferences().SetString(common.PreferenceCurrent, uid)
				setMenu(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(common.PreferenceCurrent, menuIndex[""][0])
		tree.Select(currentPref)
	} else {
		tree.Select(menuIndex[""][0])
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}

func (w *Window) makeToolBar() fyne.CanvasObject {

	t := widget.NewToolbar(
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			w.showAboutDialog()
		}),
		//widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
	)

	return container.NewBorder(t, nil, nil, nil)
}

func (w *Window) showUserInfoDialog(win fyne.Window) {

	userAvatar := widget.NewCard("", "", nil)
	userAvatar.Image = canvas.NewImageFromFile("data/images/app.jpg")
	userAvatar.Image.FillMode = canvas.ImageFillContain
	userAvatar.Resize(fyne.NewSize(5,5))

	username := widget.NewLabel("用户名：herry")
	username.TextStyle = fyne.TextStyle{Bold: true}

	email := widget.NewLabel("用户邮箱：herry@163.com")
	email.TextStyle = fyne.TextStyle{Italic: true}

	dlg := container.NewVBox(
		container.NewGridWithColumns(2, container.NewVBox(userAvatar), container.NewVBox(username,email)))

	dialog.ShowCustom("个人信息", "关闭", dlg, win)
}

func (w *Window) showAboutDialog() {
	dialog.NewInformation(common.APP_NAME, "TemplateHelper is a simple GUI client", w.win).Show()
}


func (w *Window) showOpenProjectDialog(fw fyne.Window) {
	dialog.ShowFolderOpen(func(u fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, w.win)
			return
		}
		if u == nil {
			return
		}

		w.win.Show()
		fw.Close()

	}, w.win)
}

func (w *Window) showOpenYamlFileDialog(fw fyne.Window) {
	// 创建文件选择对话框
	fileDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
		if err != nil {
			log.Println("文件选择错误:", err)
			return
		}
		if r == nil {
			return
		}
		defer r.Close()
		log.Println("文件:", r.URI().Path())


		// 处理读取到的 YAML 数据
		w.handleYamlData(r.URI().Path())
	}, fw)

	// 设置文件过滤器，只显示 YAML 文件
	fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))

	// 显示文件选择对话框
	fileDialog.Show()
}

func (w *Window) saveYamlFile() {
	err := common.GetStorage().SaveTmpl()
	if err != nil {
		dialog.ShowError(err, w.win)
		return
	}
	dialog.ShowInformation("保存成功", "模板文件保存成功", w.win)
}


func (w *Window) handleYamlData(uri string) {
	// 在这里处理读取到的 YAML 数据
	template, err := internal.ParsePlugin(uri)
	if err != nil {
		dialog.ShowError(err, w.win)
		return
	}
	common.GetStorage().SetPath(uri)
	common.GetStorage().SetTmpl(template)
	w.win.SetContent(w.tabs)

	tab := container.NewTabItem(template.Name, w.makeNavSplit(common.YamlTab))
	w.tabs.AppendWithFilePath(tab, uri)
	w.win.SetContent(w.tabs)
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}