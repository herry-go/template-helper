package tmpl

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"template-helper/gui/common"
	"template-helper/internal"
)

func MakeBasicTab(_ fyne.Window) fyne.CanvasObject {
	basic := common.GetStorage().GetTmpl()
	name := widget.NewEntry()
	name.SetText(basic.Name)
	name.OnChanged = func(text string) {
		basic.Name = text
		common.GetStorage().SetTmpl(basic)
	}

	kind := widget.NewEntry()
	kind.SetText(basic.Kind)
	kind.OnChanged = func(text string) {
		basic.Kind = text
		common.GetStorage().SetTmpl(basic)
	}

	description := widget.NewEntry()
	description.SetText(basic.Description)
	description.OnChanged = func(text string) {
		basic.Description = text
		common.GetStorage().SetTmpl(basic)
	}

	pluginName := widget.NewEntry()
	pluginName.SetText(basic.PluginName)
	pluginName.OnChanged = func(text string) {
		basic.PluginName = text
		common.GetStorage().SetTmpl(basic)
	}

	version := widget.NewEntry()
	version.SetText(basic.Version)
	version.OnChanged = func(text string) {
		basic.Version = text
		common.GetStorage().SetTmpl(basic)
	}

	apiVersion := widget.NewEntry()
	apiVersion.SetText(basic.ApiVersion)
	apiVersion.OnChanged = func(text string) {
		basic.ApiVersion = text
		common.GetStorage().SetTmpl(basic)
	}

	clusterMode := widget.NewEntry()
	clusterMode.SetText(basic.ClusterMode)
	clusterMode.OnChanged = func(text string) {
		basic.ClusterMode = text
		common.GetStorage().SetTmpl(basic)
	}

	readme := widget.NewMultiLineEntry()
	readme.SetText(basic.Readme)
	readme.OnChanged = func(text string) {
		basic.Readme = text
		common.GetStorage().SetTmpl(basic)
	}

	form := widget.NewForm(
		widget.NewFormItem("name", name),
		widget.NewFormItem("kind", kind),
		widget.NewFormItem("version", version),
		widget.NewFormItem("apiVersion", apiVersion),
		widget.NewFormItem("description", description),
		widget.NewFormItem("pluginName", pluginName),
		widget.NewFormItem("clusterMode", clusterMode),
		widget.NewFormItem("readme", readme),
	)

	return form
}

func MakeNodeListTab(w fyne.Window) fyne.CanvasObject {
	data := common.GetStorage().GetTmpl().Spec.Host
	headers := []string{"name", "label", "description", "edit"}
	currentIndex := -1
	t := widget.NewTable(nil, nil, nil)
	SetColumnWidth(w, t, len(headers))
	t.Length = func() (int, int) {
		return len(data) + 1, len(headers)
	}
	t.CreateCell = func() fyne.CanvasObject {
		return container.NewStack(widget.NewEntry())
	}

	t.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		stack := cell.(*fyne.Container)
		entry := widget.NewEntry()
		if id.Row == 0 {
			entry.SetText(headers[id.Col])
			entry.TextStyle = fyne.TextStyle{Bold: true}
			entry.Disable()
			stack.Objects = []fyne.CanvasObject{entry}
		} else {
			switch id.Col {
			case 0:
				entry.SetText(data[id.Row-1].Name)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Name = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 1:
				entry.SetText(data[id.Row-1].Label)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Label = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 2:
				entry.SetText(data[id.Row-1].Description)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Description = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 3:
				addButton := widget.NewButton("+", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						// Insert a new row after the current row
						newData := make([]internal.Variable, len(data)+1)
						copy(newData[:currentIndex+1], data[:currentIndex+1])
						newData[currentIndex+1] = data[currentIndex]
						copy(newData[currentIndex+2:], data[currentIndex+1:])
						data = newData
					} else {
						// If no row is selected, add to the end
						data = append(data, internal.Variable{})
					}
					t.Refresh()
				})
				removeButton := widget.NewButton("-", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						data = append(data[:currentIndex], data[currentIndex+1:]...)
					}
					t.Refresh()
				})
				buttonContainer := container.NewGridWithColumns(2, addButton, removeButton)
				//buttonContainer := container.NewHBox(addButton, removeButton)
				stack.Objects = []fyne.CanvasObject{buttonContainer}

			}
			common.GetStorage().SetTmplHost(data)
		}

	}
	return t
}

func MakeResourceListTab(w fyne.Window) fyne.CanvasObject {
	data := common.GetStorage().GetTmpl().Spec.Resource
	headers := []string{"name", "label", "description", "edit"}
	currentIndex := -1
	t := widget.NewTable(nil, nil, nil)
	SetColumnWidth(w, t, len(headers))
	t.Length = func() (int, int) {
		return len(data) + 1, len(headers)
	}
	t.CreateCell = func() fyne.CanvasObject {
		return container.NewStack(widget.NewEntry())
	}
	t.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		stack := cell.(*fyne.Container)
		entry := widget.NewEntry()
		if id.Row == 0 {
			entry.SetText(headers[id.Col])
			entry.TextStyle = fyne.TextStyle{Bold: true}
			entry.Disable()
			stack.Objects = []fyne.CanvasObject{entry}
		} else {
			switch id.Col {
			case 0:
				entry.SetText(data[id.Row-1].Name)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Name = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 1:
				entry.SetText(data[id.Row-1].Label)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Label = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 2:
				entry.SetText(data[id.Row-1].Description)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Description = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 3:
				addButton := widget.NewButton("+", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						// Insert a new row after the current row
						newData := make([]internal.Variable, len(data)+1)
						copy(newData[:currentIndex+1], data[:currentIndex+1])
						newData[currentIndex+1] = data[currentIndex]
						copy(newData[currentIndex+2:], data[currentIndex+1:])
						data = newData
					} else {
						// If no row is selected, add to the end
						data = append(data, internal.Variable{})
					}
					t.Refresh()
				})
				removeButton := widget.NewButton("-", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						data = append(data[:currentIndex], data[currentIndex+1:]...)
					}
					t.Refresh()
				})
				buttonContainer := container.NewGridWithColumns(2, addButton, removeButton)
				//buttonContainer := container.NewHBox(addButton, removeButton)
				stack.Objects = []fyne.CanvasObject{buttonContainer}

			}
			common.GetStorage().SetTmplResource(data)
		}

	}
	return t
}

func MakeEnvBasicTab(w fyne.Window) fyne.CanvasObject {
	data := common.GetStorage().GetTmpl().Spec.Env.Basic
	headers := []string{"type", "name", "label", "description", "required", "value", "default", "options", "dataType", "edit"}
	currentIndex := -1
	t := widget.NewTable(nil, nil, nil)
	SetColumnWidth(w, t, len(headers))
	t.Length = func() (int, int) {
		return len(data) + 1, len(headers)
	}
	t.CreateCell = func() fyne.CanvasObject {
		return container.NewStack(widget.NewEntry())
	}
	t.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		stack := cell.(*fyne.Container)

		if id.Row == 0 {
			entry := widget.NewEntry()
			entry.SetText(headers[id.Col])
			entry.TextStyle = fyne.TextStyle{Bold: true}
			entry.Disable()
			stack.Objects = []fyne.CanvasObject{entry}
		} else {
			switch id.Col {
			case 0:
				entry := widget.NewSelectEntry([]string{"input", "select", "multi-select", "switch", "file-upload"})
				entry.SetText(data[id.Row-1].Type)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Type = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 1:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Name)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Name = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 2:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Label)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Label = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 3:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Description)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Description = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 4:
				entry := widget.NewSelectEntry([]string{"true", "false"})
				entry.SetText(DisplayValue(data[id.Row-1].Required))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Required = true
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 5:
				entry := widget.NewEntry()
				entry.SetText(DisplayValue(data[id.Row-1].Value))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Value = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 6:
				entry := widget.NewEntry()
				entry.SetText(DisplayValue(data[id.Row-1].Default))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Default = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 7:
				opts := []byte{}
				if data[id.Row-1].Options != nil {
					opts, _ = json.Marshal(data[id.Row-1].Options)
				}
				entry := widget.NewEntry()
				entry.SetText(string(opts))
				entry.OnChanged = func(text string) {
					opts := []internal.Options{}
					json.Unmarshal([]byte(text), &opts)
					data[id.Row-1].Options = opts
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 8:
				entry := widget.NewSelectEntry([]string{"TEXT", "INTEGER", "BOOLEAN", "ARRAY"})
				entry.SetText(data[id.Row-1].DataType)
				entry.OnChanged = func(text string) {
					data[id.Row-1].DataType = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case len(headers) - 1:
				addButton := widget.NewButton("+", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						// Insert a new row after the current row
						newData := make([]internal.Variable, len(data)+1)
						copy(newData[:currentIndex+1], data[:currentIndex+1])
						newData[currentIndex+1] = data[currentIndex]
						copy(newData[currentIndex+2:], data[currentIndex+1:])
						data = newData
					} else {
						// If no row is selected, add to the end
						data = append(data, internal.Variable{})
					}
					t.Refresh()
				})
				removeButton := widget.NewButton("-", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						data = append(data[:currentIndex], data[currentIndex+1:]...)
					}
					t.Refresh()
				})
				buttonContainer := container.NewGridWithColumns(2, addButton, removeButton)
				//buttonContainer := container.NewHBox(addButton, removeButton)
				stack.Objects = []fyne.CanvasObject{buttonContainer}

			}
			common.GetStorage().SetTmplBasicEnv(data)
		}
	}
	return t
}

func MakeEnvAdvancedTab(w fyne.Window) fyne.CanvasObject {
	data := common.GetStorage().GetTmpl().Spec.Env.Advanced
	if len(data) == 0 {
		data = append(data, internal.Variable{})
	}
	headers := []string{"type", "name", "label", "description", "required", "value", "default", "options", "dataType", "edit"}
	currentIndex := -1
	t := widget.NewTable(nil, nil, nil)
	SetColumnWidth(w, t, len(headers))
	t.Length = func() (int, int) {
		return len(data) + 1, len(headers)
	}
	t.CreateCell = func() fyne.CanvasObject {
		return container.NewStack(widget.NewEntry())
	}
	t.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		stack := cell.(*fyne.Container)

		if id.Row == 0 {
			entry := widget.NewEntry()
			entry.SetText(headers[id.Col])
			entry.TextStyle = fyne.TextStyle{Bold: true}
			entry.Disable()
			stack.Objects = []fyne.CanvasObject{entry}
		} else {
			switch id.Col {
			case 0:
				entry := widget.NewSelectEntry([]string{"input", "select", "multi-select", "switch", "file-upload"})
				entry.SetText(data[id.Row-1].Type)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Type = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 1:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Name)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Name = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 2:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Label)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Label = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 3:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Description)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Description = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 4:
				entry := widget.NewSelectEntry([]string{"true", "false"})
				entry.SetText(DisplayValue(data[id.Row-1].Required))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Required = true
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 5:
				entry := widget.NewEntry()
				entry.SetText(DisplayValue(data[id.Row-1].Value))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Value = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 6:
				entry := widget.NewEntry()
				entry.SetText(DisplayValue(data[id.Row-1].Default))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Default = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 7:
				opts := []byte{}
				if data[id.Row-1].Options != nil {
					opts, _ = json.Marshal(data[id.Row-1].Options)
				}
				entry := widget.NewEntry()
				entry.SetText(string(opts))
				entry.OnChanged = func(text string) {
					opts := []internal.Options{}
					json.Unmarshal([]byte(text), &opts)
					data[id.Row-1].Options = opts
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 8:
				entry := widget.NewSelectEntry([]string{"TEXT", "INTEGER", "BOOLEAN", "ARRAY"})
				entry.SetText(data[id.Row-1].DataType)
				entry.OnChanged = func(text string) {
					data[id.Row-1].DataType = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case len(headers) - 1:
				addButton := widget.NewButton("+", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						// Insert a new row after the current row
						newData := make([]internal.Variable, len(data)+1)
						copy(newData[:currentIndex+1], data[:currentIndex+1])
						newData[currentIndex+1] = data[currentIndex]
						copy(newData[currentIndex+2:], data[currentIndex+1:])
						data = newData
					} else {
						// If no row is selected, add to the end
						data = append(data, internal.Variable{})
					}
					t.Refresh()
				})
				removeButton := widget.NewButton("-", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						data = append(data[:currentIndex], data[currentIndex+1:]...)
					}
					t.Refresh()
				})
				buttonContainer := container.NewGridWithColumns(2, addButton, removeButton)
				//buttonContainer := container.NewHBox(addButton, removeButton)
				stack.Objects = []fyne.CanvasObject{buttonContainer}

			}
			common.GetStorage().SetTmplAdvancedEnv(data)
		}
	}
	return t
}

func MakeCmdTab(w fyne.Window) fyne.CanvasObject {
	data := common.GetStorage().GetTmpl().Spec.Actions
	headers := []string{"name", "name", "description", "required", "command", "args", "edit"}
	currentIndex := -1
	t := widget.NewTable(nil, nil, nil)
	SetColumnWidth(w, t, len(headers))
	t.Length = func() (int, int) {
		return len(data) + 1, len(headers)
	}
	t.CreateCell = func() fyne.CanvasObject {
		return container.NewStack(widget.NewEntry())
	}
	t.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		stack := cell.(*fyne.Container)

		if id.Row == 0 {
			entry := widget.NewEntry()
			entry.SetText(headers[id.Col])
			entry.TextStyle = fyne.TextStyle{Bold: true}
			entry.Disable()
			stack.Objects = []fyne.CanvasObject{entry}
		} else {
			switch id.Col {
			case 0:
				entry := widget.NewSelectEntry([]string{"script_shell", "ansible_shell"})
				entry.SetText(data[id.Row-1].Type)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Type = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 1:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Name)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Name = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 2:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Description)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Description = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 3:
				entry := widget.NewSelectEntry([]string{"true", "false"})
				entry.SetText(DisplayValue(data[id.Row-1].Required))
				entry.OnChanged = func(text string) {
					data[id.Row-1].Required = true
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 4:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Command)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Command = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case 5:
				entry := widget.NewEntry()
				entry.SetText(data[id.Row-1].Args)
				entry.OnChanged = func(text string) {
					data[id.Row-1].Args = text
				}
				stack.Objects = []fyne.CanvasObject{entry}
			case len(headers) - 1:
				addButton := widget.NewButton("+", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						// Insert a new row after the current row
						newData := make([]internal.Actions, len(data)+1)
						copy(newData[:currentIndex+1], data[:currentIndex+1])
						newData[currentIndex+1] = data[currentIndex]
						copy(newData[currentIndex+2:], data[currentIndex+1:])
						data = newData
					} else {
						// If no row is selected, add to the end
						data = append(data, internal.Actions{})
					}
					t.Refresh()
				})
				removeButton := widget.NewButton("-", func() {
					currentIndex = id.Row - 1
					if currentIndex >= 0 && currentIndex < len(data) {
						data = append(data[:currentIndex], data[currentIndex+1:]...)
					}
					t.Refresh()
				})
				buttonContainer := container.NewGridWithColumns(2, addButton, removeButton)
				//buttonContainer := container.NewHBox(addButton, removeButton)
				stack.Objects = []fyne.CanvasObject{buttonContainer}

			}
			common.GetStorage().SetTmplAction(data)
		}
	}

	return t
}

func SetColumnWidth(w fyne.Window, t *widget.Table, col int) {
	for i := 0; i < col; i++ {
		t.SetColumnWidth(i, common.WindowWidthRight(w)/float32(col))
	}
}

func DisplayValue(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%s", v)
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
