package tasks

import (
	"reflect"
	"testing"
)

func Test_task_AddTask(t *testing.T) {
	task := task{
		name:   "open",
		status: false,
	}
	got := task.AddTask(task)
	want := tasks

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %p expected %p", got, want)
	}
}

func Test_task_Display(t *testing.T) {
	type fields struct {
		name   string
		status bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &task{
				name:   tt.fields.name,
				status: tt.fields.status,
			}
			tr.Display()
		})
	}
}
