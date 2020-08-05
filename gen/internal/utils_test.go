package internal

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/mira"
	"github.com/sf9v/nero/example"
	"github.com/stretchr/testify/assert"
)

func TestGetTypeC(t *testing.T) {
	type args struct {
		typ *mira.Type
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				typ: mira.NewType(int(0)),
			},
			want: jen.Int(),
		},
		{
			name: "*int",
			args: args{
				typ: mira.NewType(mira.IntPtr(0)),
			},
			want: jen.Op("*").Int(),
		},
		{
			name: "int32",
			args: args{
				typ: mira.NewType(int32(0)),
			},
			want: jen.Int32(),
		},
		{
			name: "*int32",
			args: args{
				typ: mira.NewType(mira.Int32Ptr(0)),
			},
			want: jen.Op("*").Int32(),
		},
		{
			name: "int64",
			args: args{
				typ: mira.NewType(int64(0)),
			},
			want: jen.Int64(),
		},
		{
			name: "*int64",
			args: args{
				typ: mira.NewType(mira.Int64Ptr(0)),
			},
			want: jen.Op("*").Int64(),
		},
		{
			name: "uint",
			args: args{
				typ: mira.NewType(uint(0)),
			},
			want: jen.Uint(),
		},
		{
			name: "*uint",
			args: args{
				typ: mira.NewType(mira.UintPtr(0)),
			},
			want: jen.Op("*").Uint(),
		},
		{
			name: "uint32",
			args: args{
				typ: mira.NewType(uint32(0)),
			},
			want: jen.Uint32(),
		},
		{
			name: "*uint32",
			args: args{
				typ: mira.NewType(mira.Uint32Ptr(0)),
			},
			want: jen.Op("*").Uint32(),
		},
		{
			name: "uint64",
			args: args{
				typ: mira.NewType(uint64(0)),
			},
			want: jen.Uint64(),
		},
		{
			name: "*uint64",
			args: args{
				typ: mira.NewType(mira.Uint64Ptr(0)),
			},
			want: jen.Op("*").Uint64(),
		},
		{
			name: "float32",
			args: args{
				typ: mira.NewType(float32(0)),
			},
			want: jen.Float32(),
		},
		{
			name: "*float32",
			args: args{
				typ: mira.NewType(mira.Float32Ptr(0)),
			},
			want: jen.Op("*").Float32(),
		},
		{
			name: "float64",
			args: args{
				typ: mira.NewType(float64(0)),
			},
			want: jen.Float64(),
		},
		{
			name: "*float64",
			args: args{
				typ: mira.NewType(mira.Float64Ptr(0)),
			},
			want: jen.Op("*").Float64(),
		},
		{
			name: "string",
			args: args{
				typ: mira.NewType(""),
			},
			want: jen.String(),
		},
		{
			name: "*string",
			args: args{
				typ: mira.NewType(mira.StrPtr("")),
			},
			want: jen.Op("*").String(),
		},
		{
			name: "big.Int",
			args: args{
				typ: mira.NewType(big.Int{}),
			},
			want: jen.Qual("math/big", "Int"),
		},
		{
			name: "*big.Int",
			args: args{
				typ: mira.NewType(big.NewInt(0)),
			},
			want: jen.Op("*").Qual("math/big", "Int"),
		},
		{
			name: "CustomString",
			args: args{
				typ: mira.NewType(example.CustomString("")),
			},
			want: jen.Qual("example", "CustomString"),
		},
		{
			name: "CustomStringOne",
			args: args{
				typ: mira.NewType(example.CustomStringOne),
			},
			want: jen.Qual("example", "CustomString"),
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
		typ *mira.Type
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				typ: mira.NewType(int(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int",
			args: args{
				typ: mira.NewType(mira.IntPtr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "int32",
			args: args{
				typ: mira.NewType(int32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int32",
			args: args{
				typ: mira.NewType(mira.Int32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "int64",
			args: args{
				typ: mira.NewType(int64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int64",
			args: args{
				typ: mira.NewType(mira.Int64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint",
			args: args{
				typ: mira.NewType(uint(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint",
			args: args{
				typ: mira.NewType(mira.UintPtr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint32",
			args: args{
				typ: mira.NewType(uint32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint32",
			args: args{
				typ: mira.NewType(mira.Uint32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "uint64",
			args: args{
				typ: mira.NewType(uint64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint64",
			args: args{
				typ: mira.NewType(mira.Uint64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "float32",
			args: args{
				typ: mira.NewType(float32(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float32",
			args: args{
				typ: mira.NewType(mira.Float32Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "float64",
			args: args{
				typ: mira.NewType(float64(0)),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float64",
			args: args{
				typ: mira.NewType(mira.Float64Ptr(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "string",
			args: args{
				typ: mira.NewType(""),
			},
			want: jen.Lit(""),
		},
		{
			name: "*string",
			args: args{
				typ: mira.NewType(mira.StrPtr("")),
			},
			want: jen.Nil(),
		},
		{
			name: "big.Int",
			args: args{
				typ: mira.NewType(big.Int{}),
			},
			want: jen.Qual("math/big", "Int").Op("{").Op("}"),
		},
		{
			name: "*big.Int",
			args: args{
				typ: mira.NewType(big.NewInt(0)),
			},
			want: jen.Nil(),
		},
		{
			name: "CustomString",
			args: args{
				typ: mira.NewType(example.CustomString("")),
			},
			want: jen.Lit(""),
		},
		{
			name: "CustomStringOne",
			args: args{
				typ: mira.NewType(example.CustomStringOne),
			},
			want: jen.Lit(""),
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
