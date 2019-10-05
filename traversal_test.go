package gizmo

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		n []string
	}
	tests := []struct {
		name string
		args args
		want Traversal
	}{
		// TODO: Add test cases.
		{"NewGraph", args{nil}, NewGraph()},
		{"NewGraph", args{[]string{"x"}}, NewGraph("x")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.n...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestG(t *testing.T) {
	type args struct {
		n []string
	}
	tests := []struct {
		name string
		args args
		want Traversal
	}{
		// TODO: Add test cases.
		{"G", args{nil}, G()},
		{"G", args{[]string{"x"}}, G("x")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := G(tt.args.n...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("G() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVAR(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want Traversal
	}{
		// TODO: Add test cases.
		{"VAR", args{"t"}, VAR("t")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VAR(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VAR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args struct {
		d Traversal
		s Traversal
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Append", args{G(), AddE("test")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Append(tt.args.d, tt.args.s)
		})
	}
}

func Test_step_V(t *testing.T) {
	var st step
	st = step{"g", nil, "", "", &st, nil, &st}

	type fields struct {
		name   string
		params []interface{}
		open   string
		close  string
		head   *step
		next   *step
		tail   *step
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Traversal
	}{
		// TODO: Add test cases.
		{
			"V",
			fields{"V", nil, "(", ")", st.head, nil, nil},
			args{},
			st.V(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &step{
				name:   tt.fields.name,
				params: tt.fields.params,
				open:   tt.fields.open,
				close:  tt.fields.close,
				head:   tt.fields.head,
				next:   tt.fields.next,
				tail:   tt.fields.tail,
			}
			if got := s.V(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("step.V() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_step_AddE(t *testing.T) {
	var st step
	st = step{"g", nil, "", "", &st, nil, &st}

	type fields struct {
		name   string
		params []interface{}
		open   string
		close  string
		head   *step
		next   *step
		tail   *step
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Traversal
	}{
		// TODO: Add test cases.
		{
			"AddE",
			fields{"addE", nil, "(", ")", st.head, nil, nil},
			args{},
			st.AddE(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &step{
				name:   tt.fields.name,
				params: tt.fields.params,
				open:   tt.fields.open,
				close:  tt.fields.close,
				head:   tt.fields.head,
				next:   tt.fields.next,
				tail:   tt.fields.tail,
			}
			if got := s.AddE(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("step.AddE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_step_AddV(t *testing.T) {
	var st step
	st = step{"g", nil, "", "", &st, nil, &st}

	type fields struct {
		name   string
		params []interface{}
		open   string
		close  string
		head   *step
		next   *step
		tail   *step
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Traversal
	}{
		// TODO: Add test cases.
		{
			"AddV",
			fields{"AddV", nil, "(", ")", st.head, nil, nil},
			args{},
			st.AddV(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &step{
				name:   tt.fields.name,
				params: tt.fields.params,
				open:   tt.fields.open,
				close:  tt.fields.close,
				head:   tt.fields.head,
				next:   tt.fields.next,
				tail:   tt.fields.tail,
			}
			if got := s.AddV(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("step.AddV() = %v, want %v", got, tt.want)
			}
		})
	}
}
