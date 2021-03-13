package memento

import (
	"errors"
	"fmt"
	"time"
)

// 编辑器接口
type IEditor interface {
	Title(title string)
	Content(content string)
	Save()
	Undo() error
	Redo() error

	Show()
}

// 编辑器备忘录实现
type editorMemento struct {
	title      string
	content    string
	createTime int64
}

func newEditorMemento(title string, content string) *editorMemento {
	return &editorMemento{
		title:      title,
		content:    content,
		createTime: time.Now().Unix(),
	}
}

// 模拟编辑器
type xMockEditor struct {
	title    string
	content  string
	versions []*editorMemento
	index    int
}

func newMockEditor() IEditor {
	return &xMockEditor{
		versions: make([]*editorMemento, 0),
	}
}

func (m *xMockEditor) Title(title string) {
	m.title = title
}

func (m *xMockEditor) Content(content string) {
	m.content = content
}

func (m *xMockEditor) Save() {
	it := newEditorMemento(m.title, m.content)
	m.versions = append(m.versions, it)
	m.index = len(m.versions) - 1
}

func (m *xMockEditor) Undo() error {
	return m.load(m.index - 1)
}

func (m *xMockEditor) load(i int) error {
	size := len(m.versions)
	if size <= 0 {
		return errors.New("no history versions")
	}

	if i < 0 || i >= size {
		return errors.New("no more history versions")
	}

	it := m.versions[i]
	m.title = it.title
	m.content = it.content
	m.index = i
	return nil
}

func (m *xMockEditor) Redo() error {
	return m.load(m.index + 1)
}

func (m *xMockEditor) Show() {
	fmt.Printf("xMockEditor.Show, title=%s, content=%s\n", m.title, m.content)
}
