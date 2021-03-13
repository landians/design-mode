package command

import "testing"

func Test_Command(t *testing.T) {
	canvas := newMockCanvas()
	canvas.Command(newColorCmd("Blue"))
	canvas.Command(newDotCmd(1, 2))
	canvas.Command(newLineCmd(1, 2, 1, 2))

	canvas.Undo()
	canvas.Undo()
}
