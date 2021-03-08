package composite

import (
	"fmt"
	"time"
)

// 学员接口
type IUser interface {
	ID() int
	Name() string
	Learn(course ICourse)
}

// 课程接口
type ICourse interface {
	ID() int
	Name() string
	Price() float64

	SetUser(user IUser)
	Learn() LearningStates
}

// 课程的学习状态
type LearningStates int

const MORE LearningStates = 1
const DONE LearningStates = 2

// 组合课程接口, 从 ICourse 继承
type ICompositeCourse interface {
	ICourse
	Append(course ICourse)
}

// mock 学员实体
type xMockUser struct {
	xID   int
	xName string
}

func newMockUser(id int, name string) IUser {
	return &xMockUser{
		xID:   id,
		xName: name,
	}
}

func (m *xMockUser) ID() int {
	return m.xID
}

func (m *xMockUser) Name() string {
	return m.xName
}

func (m *xMockUser) Learn(course ICourse) {
	course.SetUser(m)
	for {
		state := course.Learn()
		if state == DONE {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// 简单课程实体, 实现 ICourse 接口
type simpleCourse struct {
	xID    int
	xName  string
	xPrice float64
	xUser  IUser
}

func newSimpleCourse(id int, name string, price float64) ICourse {
	return &simpleCourse{
		xID:    id,
		xName:  name,
		xPrice: price,
		xUser:  nil,
	}
}

func (s *simpleCourse) ID() int {
	return s.xID
}

func (s *simpleCourse) Name() string {
	return s.xName
}

func (s *simpleCourse) Price() float64 {
	return s.xPrice
}

func (s *simpleCourse) SetUser(user IUser) {
	s.xUser = user
}

func (s *simpleCourse) Learn() LearningStates {
	fmt.Printf("%s is learing %s\n", s.xUser.Name(), s.xName)
	return DONE
}

// 组合课程，实现 ICompositeCourse 接口
type xCompositeCourse struct {
	simpleCourse

	courseList  []ICourse
	courseIndex int
}

func newCompositeCourse(id int, name string, price float64) ICompositeCourse {
	return &xCompositeCourse{
		simpleCourse: simpleCourse{
			xID:    id,
			xName:  name,
			xPrice: price,
			xUser:  nil,
		},
		courseList:  make([]ICourse, 0),
		courseIndex: 0,
	}
}

func (c *xCompositeCourse) Append(course ICourse) {
	c.courseList = append(c.courseList, course)
}

func (c *xCompositeCourse) Learn() LearningStates {
	if c.IsDone() {
		fmt.Printf("%s is learning %s: no more courses\n", c.xUser.Name(), c.Name())
		return DONE
	}

	course := c.courseList[c.courseIndex]
	fmt.Printf("%s is learning %s.%s\n", c.xUser.Name(), c.Name(), course.Name())
	c.courseIndex++

	if c.IsDone() {
		return DONE
	}
	return MORE
}

func (c *xCompositeCourse) IsDone() bool {
	return c.courseIndex >= len(c.courseList)
}
