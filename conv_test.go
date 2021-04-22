package egu
import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMap2Struct(t *testing.T) {
	type args struct {
		mapInstance map[string]interface{}
		obj         interface{}
	}

	type target struct {
		Hello string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				mapInstance: map[string]interface{}{"hello": "world"},
				obj:         &target{},
			},
			wantErr: false,
		}, {
			name: "failed",
			args: args{
				mapInstance: map[string]interface{}{"hello": 1},
				obj:         &target{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Map2Struct(tt.args.mapInstance, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Map2Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStruct2Map(t *testing.T) {
	type args struct {
		obj interface{}
	}

	type target struct {
		Hello string `json:"hello"`
	}

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "success",
			args: args{obj: target{Hello: "world"}},
			want: map[string]interface{}{"Hello": "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Struct2Map(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformInterface(t *testing.T) {
	type args struct {
		source interface{}
		target interface{}
	}

	type target struct {
		Hello string `json:"hello"`
	}

	type sourceTarget struct {
		Hello interface{} `json:"hello"`
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				source: sourceTarget{Hello: "1"},
				target: &target{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TransformInterface(tt.args.source, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("TransformInterface() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStr2Byte(t *testing.T) {
	s := "hello,world"
	assert.Equal(t, []byte(s), Str2Byte(s))
}

func TestByte2Str(t *testing.T) {
	s := "hello,world"
	b := Str2Byte(s)
	assert.Equal(t, s, Byte2Str(b))
}

func TestInt2String(t *testing.T) {
	i := 1
	s := Int2String(i)
	assert.Equal(t, "1", s)
}
func TestString2Int(t *testing.T) {
	s := "1"
	i := String2Int(s)
	assert.Equal(t, 1, i)
}