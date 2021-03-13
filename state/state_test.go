package state

import "testing"

func Test_State(t *testing.T) {
	fnTestFileStream := func(fs IFileStream, readonly bool) {
		if readonly {
			if err := fs.OpenRead(); err != nil {
				t.Log(err)
			}

			if err := fs.OpenWrite(); err != nil {
				t.Log(err)
			}
		} else {
			if err := fs.OpenWrite(); err != nil {
				t.Log(err)
			}

			if err := fs.OpenRead(); err != nil {
				t.Log(err)
			}
		}

		buffer := make([]byte, 8192)
		n, err := fs.Read(buffer)
		if err != nil {
			t.Log(err)
		} else {
			t.Logf("%v bytes read", n)
		}

		n, err = fs.Write(buffer)
		if err != nil {
			t.Log(err)
		} else {
			t.Logf("%v bytes written", n)
		}

		err = fs.Close()
		if err != nil {
			t.Log(err)
		}
	}

	fnTestFileStream(newMockFileStream("read-only.txt"), true)
	fnTestFileStream(newMockFileStream("write-only.txt"), false)
}
