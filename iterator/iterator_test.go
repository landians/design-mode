package iterator

import "testing"

func Test_Iterator(t *testing.T) {
	queue := newLinkedList()
	if queue.Size() != 0 {
		t.Fatal("expecting queue.size == 0")
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	err, value := queue.Poll()
	if err != nil || value != 1 {
		t.Fatal("expecting queue.poll 1")
	}
	err, value = queue.Poll()
	if err != nil || value != 2 {
		t.Fatal("expecting queue.poll 2")
	}
	err, value = queue.Poll()
	if err != nil || value != 3 {
		t.Fatal("expecting queue.poll 3")
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	iter := queue.Iterator()
	for iter.More() {
		err, value := iter.Next()
		if err != nil {
			t.Error(err)
		} else {
			t.Log(value)
		}
	}
}