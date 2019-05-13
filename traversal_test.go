package gizmo

import (
	"reflect"
	"testing"
)

func Test_graphTraversal(t *testing.T) {
	type args struct {
		name []string
	}
	tests := []struct {
		name string
		args args
		want *Traversal
	}{
		{"", args{[]string{"g"}}, graphTraversal("g")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := graphTraversal(tt.args.name...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graphTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraversal_New(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		name []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Traversal
	}{
		{"", fields{""}, args{[]string{"V()"}}, graphTraversal().New("V()")},
		{"", fields{""}, args{[]string{}}, graphTraversal().New()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Traversal{
				value: tt.fields.value,
			}
			if got := b.New(tt.args.name...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Traversal.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
