package tasks

import (
	"reflect"
	"testing"
)

func Test_task_AddTask(t *testing.T) {
	type fields struct {
		name   string
		status bool
	}
	type args struct {
		tk *task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &task{
				name:   tt.fields.name,
				status: tt.fields.status,
			}
			if got := tr.AddTask(tt.args.tk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("task.AddTask() = %v, want %v", got, tt.want)
			}
		})
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
