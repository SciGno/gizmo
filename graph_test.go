package gizmo

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		name string
		want *GraphAPI
	}{
		{"", Graph()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Graph(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Graph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Graph()
	}
}

func TestGraphAPI_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"graph."}, Graph().String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.String(); got != tt.want {
				t.Errorf("GraphAPI.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_Traversal(t *testing.T) {
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
		{"", fields{""}, args{[]string{"g"}}, Graph().Traversal("g")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.Traversal(tt.args.name...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.Traversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_GraphAPI_Traversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Graph().Traversal("g")
	}
}

func TestGraphAPI_IO(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphAPI
	}{
		{"", fields{"graph."}, args{GRYO}, Graph().IO(GRYO)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.IO(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.IO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_WriteGraph(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphAPI
	}{
		{"", fields{"graph."}, args{"/tmp/test.gryo"}, Graph().WriteGraph("/tmp/test.gryo")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.WriteGraph(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.WriteGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_ReadGraph(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphAPI
	}{
		{"", fields{"graph."}, args{"/tmp/test.gryo"}, Graph().ReadGraph("/tmp/test.gryo")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.ReadGraph(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.ReadGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_Config(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   *GraphAPI
	}{
		{"", fields{"graph."}, Graph().Config()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_Option(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		o string
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphAPI
	}{
		{"", fields{"graph."}, args{"allow_scan", true}, Graph().Option("allow_scan", true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.Option(tt.args.o, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.Option() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_Open(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   *GraphAPI
	}{
		{"", fields{"graph."}, Graph().Open()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.Open(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_TX(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   *GraphAPI
	}{
		{"", fields{"graph."}, Graph().TX()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.TX(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.TX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphAPI_Raw(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphAPI
	}{
		{"", fields{"graph."}, args{"allow_scan"}, Graph().Raw("allow_scan")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphAPI{
				value: tt.fields.value,
			}
			if got := g.Raw(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GraphAPI.Raw() = %v, want %v", got, tt.want)
			}
		})
	}
}
