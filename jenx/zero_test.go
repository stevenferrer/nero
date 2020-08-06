package jenx

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/google/uuid"
	"github.com/sf9v/mira"
	"github.com/stretchr/testify/assert"

	"github.com/sf9v/nero/jenx/internal"
)

func TestZero(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				v: int(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int",
			args: args{
				v: mira.IntPtr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "int32",
			args: args{
				v: int32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int32",
			args: args{
				v: mira.Int32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "int64",
			args: args{
				v: int64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int64",
			args: args{
				v: mira.Int64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint",
			args: args{
				v: uint(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint",
			args: args{
				v: mira.UintPtr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint32",
			args: args{
				v: uint32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint32",
			args: args{
				v: mira.Uint32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint64",
			args: args{
				v: uint64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint64",
			args: args{
				v: mira.Uint64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "float32",
			args: args{
				v: float32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float32",
			args: args{
				v: mira.Float32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "float64",
			args: args{
				v: float64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float64",
			args: args{
				v: mira.Float64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "string",
			args: args{
				v: "",
			},
			want: jen.Lit(""),
		},
		{
			name: "*string",
			args: args{
				v: mira.StrPtr(""),
			},
			want: jen.Nil(),
		},
		{
			name: "bool",
			args: args{
				v: false,
			},
			want: jen.False(),
		},
		{
			name: "*bool",
			args: args{
				v: mira.BoolPtr(false),
			},
			want: jen.Nil(),
		},
		{
			name: "big.Int",
			args: args{
				v: big.Int{},
			},
			want: jen.Qual("math/big", "Int").Op("{").Op("}"),
		},
		{
			name: "*big.Int",
			args: args{
				v: big.NewInt(0),
			},
			want: jen.Nil(),
		},
		{
			name: "Str",
			args: args{
				v: internal.Str(""),
			},
			want: jen.Lit(""),
		},
		{
			name: "StrOne",
			args: args{
				v: internal.StrOne,
			},
			want: jen.Lit(""),
		},
		{
			name: "Int",
			args: args{
				v: internal.Int(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "IntOne",
			args: args{
				v: internal.IntOne,
			},
			want: jen.Lit(0),
		},
		{
			name: "[8]int",
			args: args{
				v: [8]int{},
			},
			want: jen.Index(jen.Lit(8)).Int().Values(),
		},
		{
			name: "uuid.UUID",
			args: args{
				v: uuid.UUID{},
			},
			want: jen.Qual("github.com/google/uuid", "UUID").Values(),
		},
		{
			name: "*uuid.UUID",
			args: args{
				v: &uuid.UUID{},
			},
			want: jen.Nil(),
		},
		{
			name: "[]uuid.UUID",
			args: args{
				v: []uuid.UUID{},
			},
			want: jen.Nil(),
		},
		{
			name: "[8]uuid.UUID",
			args: args{
				v: [8]uuid.UUID{},
			},
			want: jen.Index(jen.Lit(8)).Qual("github.com/google/uuid", "UUID").Values(),
		},
		{
			name: "Bool",
			args: args{
				v: internal.Bool(false),
			},
			want: jen.False(),
		},
		{
			name: "BoolFalse",
			args: args{
				v: internal.BoolFalse,
			},
			want: jen.False(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Zero(tt.args.v)
			assert.Equal(t, fmt.Sprintf("%#v", tt.want),
				fmt.Sprintf("%#v", got))
		})
	}
}
