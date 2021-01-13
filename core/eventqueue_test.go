package core_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"gotest.tools/v3/assert"
)

type TestEvent int

func (te TestEvent) Action() string {
	return strconv.Itoa(int(te))
}

func TestQueueBasic(t *testing.T) {
	qq := core.NewQueue()
	output := make([]core.Event, 0)
	key := qq.AddListener(func(event core.Event, tick units.Tick) {
		output = append(output, event)
	})
	defer qq.RemoveListener(key)
	qq.Schedule(TestEvent(0), 1)
	qq.Schedule(TestEvent(1), 2)
	qq.Schedule(TestEvent(2), 3)
	assert.Equal(t, len(output), 0)
	qq.Tick()
	assert.Equal(t, len(output), 1)
	qq.Tick()
	assert.Equal(t, len(output), 2)
	qq.Tick()
	assert.Equal(t, len(output), 3)
	assert.DeepEqual(t, output, []core.Event{TestEvent(0), TestEvent(1), TestEvent(2)})
}

func TestQueueEmpty(t *testing.T) {
	qq := core.NewQueue()
	output := make([]core.Event, 0)
	key := qq.AddListener(func(event core.Event, tick units.Tick) {
		output = append(output, event)
	})
	defer qq.RemoveListener(key)
	for i := 0; i < 10; i++ {
		qq.Tick()
	}
	assert.Equal(t, len(output), 0)
}

func TestQueueInsert(t *testing.T) {
	qq := core.NewQueue()
	output := make([]core.Event, 0)
	key := qq.AddListener(func(event core.Event, tick units.Tick) {
		output = append(output, event)
	})
	defer qq.RemoveListener(key)
	in := []int{0, 1, 2, 3, 4, 5}
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
	for _, v := range in {
		qq.Schedule(TestEvent(v), units.Tick(v))
	}
	for range in {
		qq.Tick()
	}
	assert.DeepEqual(t, output, []core.Event{
		TestEvent(0),
		TestEvent(1),
		TestEvent(2),
		TestEvent(3),
		TestEvent(4),
		TestEvent(5),
	})
}

func TestHandlerForget(t *testing.T) {
	qq := core.NewQueue()
	key := qq.AddListener(func(event core.Event, tick units.Tick) {
		assert.Equal(t, event, TestEvent(1))
	})
	qq.Schedule(TestEvent(1), 1)
	qq.Schedule(TestEvent(2), 2)
	qq.Tick()
	qq.RemoveListener(key)
	qq.Tick()
}

// func _append() {
// 	var b byte
// 	var s []byte
// 	_ = append(s, b, s... /* ERROR can only use ... with matching parameter */ )
// }
