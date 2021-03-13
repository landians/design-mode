package state

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

// 文件流API接口
type IFileStream interface {
	io.ReadWriteCloser

	OpenRead() error
	OpenWrite() error
}

// 文件流上下文接口，仅内部使用
type iFileStreamContext interface {
	File() string
	Switch(state iFileStreamState)
}

// 文件流状态接口, 仅内部使用
type iFileStreamState interface {
	IFileStream
}

// 模拟文件流API的实现
type xMockFileStream struct {
	state iFileStreamState
	file  string
}

func newMockFileStream(file string) IFileStream {
	fs := &xMockFileStream{
		state: nil,
		file:  file,
	}
	fs.state = newDefaultState(fs)
	return fs
}

func (m *xMockFileStream) File() string {
	return m.file
}

func (m *xMockFileStream) Switch(st iFileStreamState) {
	fmt.Printf("xMockFileStream.Switch, %s => %s\n", reflect.TypeOf(m.state).String(), reflect.TypeOf(st).String())
	m.state = st
}

func (m *xMockFileStream) OpenRead() error {
	return m.state.OpenRead()
}

func (m *xMockFileStream) OpenWrite() error {
	return m.state.OpenWrite()
}

func (m *xMockFileStream) Read(p []byte) (n int, err error) {
	return m.state.Read(p)
}

func (m *xMockFileStream) Write(p []byte) (n int, err error) {
	return m.state.Write(p)
}

func (m *xMockFileStream) Close() error {
	return m.state.Close()
}

// 默认状态-未打开状态，该状态下只允许打开/关闭操作
type defaultState struct {
	context iFileStreamContext
}

func newDefaultState(context iFileStreamContext) iFileStreamState {
	return &defaultState{context: context}
}

func (d *defaultState) OpenRead() error {
	fmt.Printf("defaultState.OpenRead, file=%s\n", d.context.File())
	d.context.Switch(newReadingState(d.context))
	return nil
}

func (d *defaultState) OpenWrite() error {
	fmt.Printf("defaultState.OpenRead, file=%s\n", d.context.File())
	d.context.Switch(newWritingState(d.context))
	return nil
}

func (d *defaultState) Read(p []byte) (n int, err error) {
	return 0, errors.New(fmt.Sprintf("defaultState.Read, file not opened, %s", d.context.File()))
}

func (d *defaultState) Write(p []byte) (n int, err error) {
	return 0, errors.New(fmt.Sprintf("defaultState.Write, file not opened, %s", d.context.File()))
}

func (d *defaultState) Close() error {
	fmt.Printf("defaultState.Close, file=%s\n", d.context.File())
	d.context.Switch(newClosedState(d.context))
	return nil
}

// 读取状态，该状态下只允许读取/关闭操作
type readingState struct {
	context    iFileStreamContext
	iBytesRead int
}

func newReadingState(context iFileStreamContext) iFileStreamState {
	return &readingState{
		context:    context,
		iBytesRead: 0,
	}
}

func (r *readingState) OpenRead() error {
	return errors.New(fmt.Sprintf("readingState.OpenRead, already reading %s", r.context.File()))
}

func (r *readingState) OpenWrite() error {
	return errors.New(fmt.Sprintf("readingState.OpenWrite, already reading %s", r.context.File()))
}

func (r *readingState) Read(p []byte) (n int, err error) {
	size := len(p)
	r.iBytesRead += size
	fmt.Printf("readingState.Read, file=%s, iBytesRead=%v\n", r.context.File(), r.iBytesRead)
	return size, nil
}

func (r *readingState) Write(p []byte) (n int, err error) {
	return 0, errors.New(fmt.Sprintf("readingState.Write, cannot write to %s", r.context.File()))
}

func (r *readingState) Close() error {
	fmt.Printf("readingState.Close, file=%s, iBytesRead=%v\n", r.context.File(), r.iBytesRead)
	r.context.Switch(newClosedState(r.context))
	return nil
}

// 写入状态，该状态下只允许写入/关闭操作
type writingState struct {
	context iFileStreamContext
	written int
}

func newWritingState(context iFileStreamContext) *writingState {
	return &writingState{
		context: context,
		written: 0,
	}
}

func (w *writingState) OpenRead() error {
	return errors.New(fmt.Sprintf("writingState.OpenRead, already writing %s", w.context.File()))
}

func (w *writingState) OpenWrite() error {
	return errors.New(fmt.Sprintf("writingState.OpenWrite, already writing %s", w.context.File()))
}

func (w *writingState) Read(p []byte) (n int, err error) {
	return 0, errors.New(fmt.Sprintf("writingState.Read, cannot read %s", w.context.File()))
}

func (w *writingState) Write(p []byte) (n int, err error) {
	size := len(p)
	w.written += size
	fmt.Printf("writingState.Write, file=%s, written=%v\n", w.context.File(), w.written)
	return size, nil
}

func (w *writingState) Close() error {
	fmt.Printf("writingState.Close, file=%s, written=%v\n", w.context.File(), w.written)
	w.context.Switch(newClosedState(w.context))
	return nil
}

// 关闭状态，该状态下只允许关闭操作
type closedState struct {
	context iFileStreamContext
}

func newClosedState(context iFileStreamContext) iFileStreamState {
	return &closedState{
		context: context,
	}
}

func (c * closedState) OpenRead() error {
	return errors.New(fmt.Sprintf("closedState.OpenRead, file(%s) already closed ", c.context.File()))
}

func (c * closedState) OpenWrite() error {
	return errors.New(fmt.Sprintf("closedState.OpenWrite, file(%s) already closed ", c.context.File()))
}

func (c * closedState) Read(p []byte) (n int, e error) {
	return 0, errors.New(fmt.Sprintf("closedState.Read, file(%s) already closed ", c.context.File()))
}

func (c * closedState) Write(p []byte) (n int, e error) {
	return 0, errors.New(fmt.Sprintf("closedState.Write, file(%s) already closed ", c.context.File()))
}

func (c * closedState) Close() error {
	return nil
}
