package client

import (
	"testing"
)

func TestTM0(t *testing.T) {
	{
		tm0 := &TimeLine{1, 0}
		tm1 := &TimeLine{6, 3}
		if _, can := tm0.Merge(tm1); can {
			t.Fail()
		}
	}

	{
		tm0 := &TimeLine{3, 0}
		tm1 := &TimeLine{6, 3}
		if _, can := tm0.Merge(tm1); !can {
			t.Fail()
		}
	}

	{
		tm0 := &TimeLine{5, 0}
		tm1 := &TimeLine{6, 3}
		if _, can := tm0.Merge(tm1); !can {
			t.Fail()
		}
	}

	{
		tm0 := &TimeLine{5, 1}
		tm1 := &TimeLine{5, 1}
		if _, can := tm0.Merge(tm1); !can {
			t.Fail()
		}
	}

	{
		tm0 := &TimeLine{1, 5}
		tm1 := &TimeLine{1, 5}
		if _, can := tm0.Merge(tm1); can {
			t.Fail()
		}
	}

}
