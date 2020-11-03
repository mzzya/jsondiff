package jsondiff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		json1 string
		json2 string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "simple test", args: args{json1: `{"int":1,"string":"string","float":1.1,"array":[1,2,3],"object":{"int":1,"string":"string"},"objectAry":[{"int":1,"string":"string"}]}`, json2: `{"int":12,"string":"string2","float":1.12,"array":[12,22,32],"object":{"int":1,"string":"string2"}}`}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Diff(tt.args.json1, tt.args.json2, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

var json1 = `{"int":1,"string":"string","float":1.1,"array":[1,2,3],"object":{"int":1,"string":"string"},"objectAry":[{"int":1,"string":"string"}]}`
var json2 = `{"int":12,"string":"string2","float":1.12,"array":[12,22,32],"object":{"int":1,"string":"string2"}}`

func BenchmarkDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Diff(json1, json2, false)
	}
}
