package composite

import "testing"

func Test_Composite(t *testing.T) {
	user := newMockUser(1, "张三")

	sc := newSimpleCourse(11, "golang入门", 100)
	user.Learn(sc)

	user = newMockUser(2, "李四")
	cc := newCompositeCourse(21, "golang架构师", 500)
	cc.Append(newSimpleCourse(11, "golang入门", 100))
	cc.Append(newSimpleCourse(12, "golang基础", 100))
	cc.Append(newSimpleCourse(13, "golang进阶", 100))
	cc.Append(newSimpleCourse(14, "golang高级", 100))
	user.Learn(cc)
}
