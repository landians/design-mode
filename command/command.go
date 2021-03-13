package command

import (
	"errors"
	"fmt"
)

// 画板接口，供上层调用
type ICanvas interface {
	Command(cmd IDrawCommand)
	Undo()
}

// 绘图命令接口
type IDrawCommand interface {
	Draw(g IGraphics)
}

// 绘图上下文接口
type IGraphics interface {
	Color(color string)
	Clear()
	DrawDot(x int, y int)
	DrawLine(x0 int, y0 int, x1 int, y1 int)
}

// 模拟绘图上下文
type xMockGraphics struct {
	color string
}

func newMockGraphics() IGraphics {
	return &xMockGraphics{}
}

func (m *xMockGraphics) Clear() {
	fmt.Printf("xMockGraphics.Clear\n")
}

func (m *xMockGraphics) Color(color string) {
	m.color = color
	fmt.Printf("xMockGraphics.Color, %v\n", color)
}

func (m *xMockGraphics) DrawDot(x int, y int) {
	fmt.Printf("xMockGraphics.ColDrawDot, (%v, %v)\n", x, y)
}

func (m *xMockGraphics) DrawLine(x0 int, y0 int, x1 int, y1 int) {
	fmt.Printf("xMockGraphics.DrawLine, (%v, %v, %v, %v)\n", x0, y0, x1, y1)
}

// 模拟画板
type xMockCanvas struct {
	commands []IDrawCommand
	graphics IGraphics
}

func newMockCanvas() ICanvas {
	return &xMockCanvas{
		commands: make([]IDrawCommand, 0),
		graphics: newMockGraphics(),
	}
}

func (m *xMockCanvas) Command(cmd IDrawCommand) {
	m.push(cmd)
	m.update()
}

func (m *xMockCanvas) Undo() {
	err, _ := m.pop()
	if err != nil {
		return
	}

	m.update()
}

func (m *xMockCanvas) push(cmd IDrawCommand) {
	m.commands = append(m.commands, cmd)
}

func (m *xMockCanvas) pop() (error, IDrawCommand) {
	size := len(m.commands)
	if size <= 0 {
		return errors.New("no more commands"), nil
	}

	it := m.commands[size-1]
	m.commands = m.commands[:size-1]
	return nil, it
}

func (m *xMockCanvas) update() {
	m.graphics.Clear()
	for _, it := range m.commands {
		it.Draw(m.graphics)
	}
}

// 绘制颜色的命令
type colorCmd struct {
	color string
}

func newColorCmd(color string) IDrawCommand {
	return &colorCmd{color: color}
}

func (c *colorCmd) Draw(g IGraphics) {
	g.Color(c.color)
}

// 画点的命令
type dotCmd struct {
	x int
	y int
}

func newDotCmd(x int, y int) *dotCmd {
	return &dotCmd{
		x: x,
		y: y,
	}
}

func (d *dotCmd) Draw(g IGraphics) {
	g.DrawDot(d.x, d.y)
}

// 画线的命令
type lineCmd struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

func newLineCmd(x0 int, y0 int, x1 int, y1 int) IDrawCommand {
	return &lineCmd{
		x0: x0,
		y0: y0,
		x1: x1,
		y1: y1,
	}
}

func (l *lineCmd) Draw(g IGraphics) {
	g.DrawLine(l.x0, l.y0, l.x1, l.y1)
}
