package textbox

import "testing"

func TestDependencies(t *testing.T) {
	for i := 0; i < 128; i++ {
		win := New(NewDumbBackend())
		defer win.Close()

		deps := NewDependencies()
		b1 := win.Box()
		b2 := win.Box()
		b3 := win.Box()
		deps.Add(b1, win.Root)
		deps.Add(b2, b1)
		deps.Add(b3, b2)
		deps.Add(b3, b1)
		deps.Add(b1, b3)

		visited := make(map[*Box]int)
		deps.Iter(win.Root, func(box *Box) {
			visited[box]++
		})
		if len(visited) != 4 {
			t.Fatal("visited")
		}
		expected := map[*Box]int{
			win.Root: 1,
			b1:       1,
			b2:       1,
			b3:       1,
		}
		for key, value := range expected {
			if n, ok := visited[key]; !ok || n != value {
				t.Fatal("visited")
			}
		}
	}
}
