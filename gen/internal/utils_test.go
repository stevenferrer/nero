package internal

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func TestGetTypeC(t *testing.T) {
	type args struct {
		typ *Typ
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				typ: NewTyp(int(0)),
			},
			want: jen.Int(),
		},
		{
			name: "*int",
			args: args{
				typ: NewTyp(intPtr(0)),
			},
			want: jen.Op("*").Int(),
		},
		{
			name: "int32",
			args: args{
				typ: NewTyp(int32(0)),
			},
			want: jen.Int32(),
		},
		{
			name: "*int32",
			args: args{
				typ: NewTyp(int32Ptr(0)),
			},
			want: jen.Op("*").Int32(),
		},
		{
			name: "int64",
			args: args{
				typ: NewTyp(int64(0)),
			},
			want: jen.Int64(),
		},
		{
			name: "*int64",
			args: args{
				typ: NewTyp(int64Ptr(0)),
			},
			want: jen.Op("*").Int64(),
		},
		{
			name: "uint",
			args: args{
				typ: NewTyp(uint(0)),
			},
			want: jen.Uint(),
		},
		{
			name: "*uint",
			args: args{
				typ: NewTyp(uintPtr(0)),
			},
			want: jen.Op("*").Uint(),
		},
		{
			name: "uint32",
			args: args{
				typ: NewTyp(uint32(0)),
			},
			want: jen.Uint32(),
		},
		{
			name: "*uint32",
			args: args{
				typ: NewTyp(uint32Ptr(0)),
			},
			want: jen.Op("*").Uint32(),
		},
		{
			name: "uint64",
			args: args{
				typ: NewTyp(uint64(0)),
			},
			want: jen.Uint64(),
		},
		{
			name: "*uint64",
			args: args{
				typ: NewTyp(uint64Ptr(0)),
			},
			want: jen.Op("*").Uint64(),
		},
		{
			name: "float32",
			args: args{
				typ: NewTyp(float32(0)),
			},
			want: jen.Float32(),
		},
		{
			name: "*float32",
			args: args{
				typ: NewTyp(float32Ptr(0)),
			},
			want: jen.Op("*").Float32(),
		},
		{
			name: "float64",
			args: args{
				typ: NewTyp(float64(0)),
			},
			want: jen.Float64(),
		},
		{
			name: "*float64",
			args: args{
				typ: NewTyp(float64Ptr(0)),
			},
			want: jen.Op("*").Float64(),
		},
		{
			name: "string",
			args: args{
				typ: NewTyp(""),
			},
			want: jen.String(),
		},
		{
			name: "*string",
			args: args{
				typ: NewTyp(strPtr("")),
			},
			want: jen.Op("*").String(),
		},
		{
			name: "big.Int",
			args: args{
				typ: NewTyp(big.Int{}),
			},
			want: jen.Qual("math/big", "Int"),
		},

		{
			name: "*big.Int",
			args: args{
				typ: NewTyp(big.NewInt(0)),
			},
			want: jen.Op("*").Qual("math/big", "Int"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTypeC(tt.args.typ)
			assert.Equal(t, fmt.Sprintf("%#v", tt.want),
				fmt.Sprintf("%#v", got))
		})
	}
}

func TestGetZeroValC(t *testing.T) {
	type args struct {
		typ *Typ
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				typ: NewTyp(int(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int",
			args: args{
				typ: NewTyp(intPtr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "int32",
			args: args{
				typ: NewTyp(int32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int32",
			args: args{
				typ: NewTyp(int32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "int64",
			args: args{
				typ: NewTyp(int64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int64",
			args: args{
				typ: NewTyp(int64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint",
			args: args{
				typ: NewTyp(uint(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint",
			args: args{
				typ: NewTyp(uintPtr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint32",
			args: args{
				typ: NewTyp(uint32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint32",
			args: args{
				typ: NewTyp(uint32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint64",
			args: args{
				typ: NewTyp(uint64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint64",
			args: args{
				typ: NewTyp(uint64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "float32",
			args: args{
				typ: NewTyp(float32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float32",
			args: args{
				typ: NewTyp(float32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "float64",
			args: args{
				typ: NewTyp(float64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float64",
			args: args{
				typ: NewTyp(float64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "string",
			args: args{
				typ: NewTyp(""),
			},
			want: jen.Lit(""),
		},
		{
			name: "*string",
			args: args{
				typ: NewTyp(strPtr("")),
			},
			want: jen.Nil(),
		},
		{
			name: "big.Int",
			args: args{
				typ: NewTyp(big.Int{}),
			},
			want: jen.Qual("math/big", "Int").Op("{").Op("}"),
		},
		{
			name: "*big.Int",
			args: args{
				typ: NewTyp(big.NewInt(0)),
			},
			want: jen.Nil(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetZeroValC(tt.args.typ)
			assert.Equal(t, fmt.Sprintf("%#v", tt.want),
				fmt.Sprintf("%#v", got))
		})
	}
}
