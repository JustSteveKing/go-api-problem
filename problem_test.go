package problem

import (
	"encoding/json"
	"testing"
)

func TestAPIProblem(t *testing.T) {
	t.Run("test API Problem Title", func(t *testing.T) {
		problem := APIProblem{
			Title:  "Test Problem",
			Detail: "This is a test problem",
			Status: "123",
			Code:   "PRO-123",
			Meta:   &map[string]interface{}{"key": "val"},
		}

		want := "Test Problem"

		if want != problem.Title {
			t.Error("want does not equal problem.Title")
		}
	})

	t.Run("test API Problem Detail", func(t *testing.T) {
		problem := APIProblem{
			Title:  "Test Problem",
			Detail: "This is a test problem",
			Status: "123",
			Code:   "PRO-123",
			Meta:   &map[string]interface{}{"key": "val"},
		}

		want := "This is a test problem"

		if want != problem.Detail {
			t.Error("want does not equal problem.Title")
		}
	})

	t.Run("test API Problem Status", func(t *testing.T) {
		problem := APIProblem{
			Title:  "Test Problem",
			Detail: "This is a test problem",
			Status: "123",
			Code:   "PRO-123",
			Meta:   &map[string]interface{}{"key": "val"},
		}

		want := "123"

		if want != problem.Status {
			t.Error("want does not equal problem.Title")
		}
	})

	t.Run("test API Problem Code", func(t *testing.T) {
		problem := APIProblem{
			Title:  "Test Problem",
			Detail: "This is a test problem",
			Status: "123",
			Code:   "PRO-123",
			Meta:   &map[string]interface{}{"key": "val"},
		}

		want := "PRO-123"

		if want != problem.Code {
			t.Error("want does not equal problem.Title")
		}
	})

	t.Run("test Marshal API Problem to JSON", func(t *testing.T) {
		problem := APIProblem{
			Title:  "Test Problem",
			Detail: "This is a test problem",
			Status: "123",
			Code:   "PRO-123",
			Meta:   &map[string]interface{}{"key": "val"},
		}

		_, err := json.Marshal(&problem)
		if err != nil {
			t.Error("Failed to Marshal JSON")
		}
	})
}
