package common

import "fyne.io/fyne/v2/container"

// CustomDocTabs 自定义结构体，封装 container.DocTabs 并添加额外字段
type CustomDocTabs struct {
	*container.DocTabs
	FilePaths map[*container.TabItem]string // 存储每个标签页对应的文件路径
}

// NewCustomDocTabs 创建一个新的 CustomDocTabs 实例
func NewCustomDocTabs() *CustomDocTabs {
	return &CustomDocTabs{
		DocTabs:   container.NewDocTabs(),
		FilePaths: make(map[*container.TabItem]string),
	}
}

// AppendWithFilePath 添加一个新的标签页，并关联文件路径
func (cdt *CustomDocTabs) AppendWithFilePath(tab *container.TabItem, filePath string) {
	cdt.DocTabs.Append(tab)
	cdt.FilePaths[tab] = filePath
}

// GetFilePath 获取指定标签页的文件路径
func (cdt *CustomDocTabs) GetFilePath(tab *container.TabItem) string {
	return cdt.FilePaths[tab]
}
