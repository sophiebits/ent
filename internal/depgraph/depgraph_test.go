package depgraph

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func getKey(i int) string {
	return fmt.Sprintf("item%d", i)
}

type depgraphTest struct {
	Depgraph
	sum int
}

func (g *depgraphTest) sumFunc(i int) {
	g.sum = g.sum + i
}

func (g *depgraphTest) RunLoop() {
	for i := 0; i < 10; i++ {
		j := i // to capture the loop variable correctly
		g.CheckAndQueue(getKey(i), func(item interface{}) {
			passedFunc, ok := item.(func(int))
			if !ok {
				panic("invalid func passed")
			}
			passedFunc(j)
		})
	}
}

type depgraphTestSimple struct {
	Depgraph
	//	numberOfTimesCalled int
	//	numbers []int
	//	map[]
}

type object struct {
	field1 int
	field2 string
	field3 int
	field4 string
}

// func (g *depgraphTestSimple) execFn(exec func(interface{})) {
// 	g.numberOfTimesCalled++
// }

// func (g *depgraphTestSimple) sumFunc(i int) {
// 	g.numbers = append(g.numbers, i)
// }

func TestRunLoopNoDeps(t *testing.T) {
	g := &depgraphTest{}

	for i := 0; i < 10; i++ {
		g.AddItem(getKey(i), g.sumFunc)
	}

	if len(g.items) != 10 {
		t.Errorf("expected 10 items to be added. got %d instead", len(g.items))
	}

	g.RunLoop()

	if len(g.queue) != 0 {
		t.Errorf("expected no items queued up. %d items were queued up instead", len(g.queue))
	}

	expectedSum := (10 * 9) / 2
	if g.sum != expectedSum {
		t.Errorf(
			"expected the sum to be %d, it was %d instead implying every function wasn't called exactly once",
			expectedSum,
			g.sum,
		)
	}
}

func verifyRunLoopSimpleDeps(t *testing.T, g *depgraphTest, runQueuePanics bool) {
	if len(g.queue) != 5 {
		t.Errorf("expected 5 items queued up. %d items were queued up instead", len(g.queue))
	}

	expectedSum := 1 + 3 + 5 + 7 + 9
	if g.sum != expectedSum {
		t.Errorf("expected sum for the items run so far is not as expected")
	}

	if runQueuePanics {
		assert.Panics(t, g.RunQueuedUpItems)
		return
	}
	g.RunQueuedUpItems()

	if len(g.queue) != 0 {
		t.Errorf("expected no items queued up. %d items were queued up instead", len(g.queue))
	}

	expectedSum = (10 * 9) / 2
	if g.sum != expectedSum {
		t.Errorf(
			"expected the sum to be %d, it was %d instead implying every function wasn't called exactly once",
			expectedSum,
			g.sum,
		)
	}

}

func TestRunLoopWithSimpleDeps(t *testing.T) {
	g := &depgraphTest{}

	// for half of the items, add a dependency, for the other half run once
	// even numbered items have a dependency on the one after them
	for i := 0; i < 10; i++ {
		key := getKey(i)
		if i%2 == 0 {
			g.AddItem(key, g.sumFunc, getKey(i+1))
		} else {
			g.AddItem(key, g.sumFunc)
		}
	}

	if len(g.items) != 10 {
		t.Errorf("expected 10 items to be added. got %d instead", len(g.items))
	}

	g.RunLoop()

	verifyRunLoopSimpleDeps(t, g, false)
}

func TestRunLoopOptionalItemsNotCleared(t *testing.T) {
	g := &depgraphTest{}
	g.AddOptionalItem(getKey(11), g.sumFunc)

	// same as above. except add a dependency on 11 from 8
	for i := 0; i < 10; i++ {
		key := getKey(i)
		if i == 8 {
			g.AddItem(key, g.sumFunc, getKey(11))
		} else if i%2 == 0 {
			g.AddItem(key, g.sumFunc, getKey(i+1))
		} else {
			g.AddItem(key, g.sumFunc)
		}
	}

	if len(g.items) != 11 {
		t.Errorf("expected 11 items to be added. got %d instead", len(g.items))
	}

	g.RunLoop()
	verifyRunLoopSimpleDeps(t, g, true)
}

func TestRunLoopOptionalItemsCleared(t *testing.T) {
	g := &depgraphTest{}
	g.AddOptionalItem(getKey(11), g.sumFunc)

	// same as above. except add a dependency on 11 from 8
	for i := 0; i < 10; i++ {
		key := getKey(i)
		if i == 8 {
			g.AddItem(key, g.sumFunc, getKey(11))
		} else if i%2 == 0 {
			g.AddItem(key, g.sumFunc, getKey(i+1))
		} else {
			g.AddItem(key, g.sumFunc)
		}
	}

	if len(g.items) != 11 {
		t.Errorf("expected 11 items to be added. got %d instead", len(g.items))
	}

	g.RunLoop()
	g.ClearOptionalItems()
	verifyRunLoopSimpleDeps(t, g, false)
}

func TestRunLoopTooManyDeps(t *testing.T) {
	g := &depgraphTest{}

	// we have 1 depends on 2, 3; 2 depends on 3, 4- and so on and since it's simple and not a
	// real graph, it breaks
	for i := 0; i < 10; i++ {
		key := getKey(i)
		if i == 9 {
			g.AddItem(key, g.sumFunc)
		} else {
			g.AddItem(key, g.sumFunc, getKey(i+1), getKey((i+2)%10))
		}
	}

	if len(g.items) != 10 {
		t.Errorf("expected 10 items to be added. got %d instead", len(g.items))
	}

	g.RunLoop()

	if len(g.queue) != 9 {
		t.Errorf("expected 9 items queued up. %d items were queued up instead", len(g.queue))
	}

	assert.Panics(t, g.RunQueuedUpItems)
}

func TestInvalidDep(t *testing.T) {
	g := &depgraphTest{}

	for i := 0; i < 10; i++ {
		key := getKey(i)
		if i == 0 {
			g.AddItem(key, g.sumFunc, getKey(10))
		} else {
			g.AddItem(key, g.sumFunc)
		}
	}

	// panics here
	assert.Panics(t, g.RunLoop)
}

func TestRunNoDeps(t *testing.T) {
	g := &depgraphTestSimple{}

	g.AddItem("field1", func(obj *object) {
		obj.field1 = 1
	})

	g.AddItem("field2", func(obj *object) {
		obj.field2 = "field2"
	})

	g.AddItem("field3", func(obj *object) {
		obj.field3 = 3
	})

	g.AddItem("field4", func(obj *object) {
		obj.field4 = "field4"
	})

	if len(g.items) != 4 {
		t.Errorf("expected 4 items to be added. got %d instead", len(g.items))
	}

	var hardToCalObj object

	g.Run(func(item interface{}) {
		execFn, ok := item.(func(*object))
		if !ok {
			panic("invalid object passed")
		}
		execFn(&hardToCalObj)
	})

	if len(g.queue) != 0 {
		t.Errorf("expected no items queued up. %d items were queued up instead", len(g.queue))
	}

	if hardToCalObj.field1 != 1 {
		t.Errorf("field1 was not set when Run() was called")
	}
	if hardToCalObj.field2 != "field2" {
		t.Errorf("field2 was not set when Run() was called")
	}
	if hardToCalObj.field3 != 3 {
		t.Errorf("field3 was not set when Run() was called")
	}
	if hardToCalObj.field4 != "field4" {
		t.Errorf("field4 was not set when Run() was called")
	}
}

func TestRunWithDeps(t *testing.T) {
	g := &depgraphTestSimple{}

	g.AddItem("field1", func(obj *object) {
		obj.field1 = uuid.MustParse(obj.field2).ClockSequence()
	}, "field2")

	g.AddItem("field2", func(obj *object) {
		obj.field2 = uuid.New().String()
	})

	g.AddItem("field3", func(obj *object) {
		obj.field3 = 3
	}, "field4")

	g.AddItem("field4", func(obj *object) {
		obj.field4 = "field4"
	})

	if len(g.items) != 4 {
		t.Errorf("expected 4 items to be added. got %d instead", len(g.items))
	}

	var hardToCalObj object

	g.Run(func(item interface{}) {
		execFn, ok := item.(func(*object))
		if !ok {
			panic("invalid function passed")
		}
		execFn(&hardToCalObj)
	})

	if len(g.queue) != 0 {
		t.Errorf("expected no items queued up. %d items were queued up instead", len(g.queue))
	}

	// shows the dependency btw 1 and 2 since 2 has to be set before 1 can be
	// 4 and 3 is random right now and doesn't necessarily prove anything but that's fine
	field2Uuid := uuid.MustParse(hardToCalObj.field2)

	if hardToCalObj.field1 != field2Uuid.ClockSequence() {
		t.Errorf("field1 was not set when Run() was called")
	}
	if hardToCalObj.field3 != 3 {
		t.Errorf("field3 was not set when Run() was called")
	}
	if hardToCalObj.field4 != "field4" {
		t.Errorf("field4 was not set when Run() was called")
	}
}
