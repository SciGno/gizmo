package gizmo

import "testing"

func Test_step_String(t *testing.T) {
	v := VAR("t")
	v.AddE()
	g := G()
	g.AddV("v")

	type fields struct {
		name   string
		params []interface{}
		open   string
		close  string
		head   *step
		next   *step
		tail   *step
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"String", fields{"addE", nil, "(", ")", nil, nil, nil}, AddE().String()},
		{"String", fields{"addE", []interface{}{"edge"}, "(", ")", nil, nil, nil}, AddE("edge").String()},
		{"String", fields{"g", nil, "", "", nil, AddV("v").(*step), nil}, g.String()},
		{"String", fields{"within", []interface{}{"a", "b"}, "(", ")", nil, nil, nil}, Within([]string{"a", "b"}).String()},
		{"String", fields{"within", []interface{}{1, 2}, "(", ")", nil, nil, nil}, Within([]int{1, 2}).String()},
		{"String", fields{"within", []interface{}{1.0, 2.0}, "(", ")", nil, nil, nil}, Within([]float32{1.0, 2.0}).String()},
		{"String", fields{"within", []interface{}{1.0, 2.0}, "(", ")", nil, nil, nil}, Within([]float64{1.0, 2.0}).String()},
		{"String", fields{"t", nil, "=", "", nil, AddE().(*step), nil}, v.String()},
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
			if got := s.String(); got != tt.want {
				t.Errorf("step.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
