package jenx

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/google/uuid"
	"github.com/sf9v/mira"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
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
			want: jen.Int(),
		},
		{
			name: "*int",
			args: args{
				v: mira.IntPtr(0),
			},
			want: jen.Op("*").Int(),
		},
		{
			name: "int32",
			args: args{
				v: int32(0),
			},
			want: jen.Int32(),
		},
		{
			name: "*int32",
			args: args{
				v: mira.Int32Ptr(0),
			},
			want: jen.Op("*").Int32(),
		},
		{
			name: "int64",
			args: args{
				v: int64(0),
			},
			want: jen.Int64(),
		},
		{
			name: "*int64",
			args: args{
				v: mira.Int64Ptr(0),
			},
			want: jen.Op("*").Int64(),
		},
		{
			name: "uint",
			args: args{
				v: uint(0),
			},
			want: jen.Uint(),
		},
		{
			name: "*uint",
			args: args{
				v: mira.UintPtr(0),
			},
			want: jen.Op("*").Uint(),
		},
		{
			name: "uint32",
			args: args{
				v: uint32(0),
			},
			want: jen.Uint32(),
		},
		{
			name: "*uint32",
			args: args{
				v: mira.Uint32Ptr(0),
			},
			want: jen.Op("*").Uint32(),
		},
		{
			name: "uint64",
			args: args{
				v: uint64(0),
			},
			want: jen.Uint64(),
		},
		{
			name: "*uint64",
			args: args{
				v: mira.Uint64Ptr(0),
			},
			want: jen.Op("*").Uint64(),
		},
		{
			name: "float32",
			args: args{
				v: float32(0),
			},
			want: jen.Float32(),
		},
		{
			name: "*float32",
			args: args{
				v: mira.Float32Ptr(0),
			},
			want: jen.Op("*").Float32(),
		},
		{
			name: "float64",
			args: args{
				v: float64(0),
			},
			want: jen.Float64(),
		},
		{
			name: "*float64",
			args: args{
				v: mira.Float64Ptr(0),
			},
			want: jen.Op("*").Float64(),
		},
		{
			name: "string",
			args: args{
				v: "",
			},
			want: jen.String(),
		},
		{
			name: "*string",
			args: args{
				v: mira.StrPtr(""),
			},
			want: jen.Op("*").String(),
		},
		{
			name: "bool",
			args: args{
				v: false,
			},
			want: jen.Bool(),
		},
		{
			name: "*bool",
			args: args{
				v: mira.BoolPtr(false),
			},
			want: jen.Op("*").Bool(),
		},
		{
			name: "big.Int",
			args: args{
				v: big.Int{},
			},
			want: jen.Qual("math/big", "Int"),
		},
		{
			name: "*big.Int",
			args: args{
				v: big.NewInt(0),
			},
			want: jen.Op("*").Qual("math/big", "Int"),
		},
		{
			name: "[]string",
			args: args{
				v: []string{},
			},
			want: jen.Index().String(),
		},
		{
			name: "[]*string",
			args: args{
				v: []*string{},
			},
			want: jen.Index().Op("*").String(),
		},
		{
			name: "[]int64",
			args: args{
				v: []int64{},
			},
			want: jen.Index().Int64(),
		},
		{
			name: "[]*int64",
			args: args{
				v: []*int64{},
			},
			want: jen.Index().Op("*").Int64(),
		},
		{
			name: "map[string]string",
			args: args{
				v: map[string]string{},
			},
			want: jen.Map(jen.String()).String(),
		},
		{
			name: "map[string]*string",
			args: args{
				v: map[string]*string{},
			},
			want: jen.Map(jen.String()).Op("*").String(),
		},
		{
			name: "map[int64]*big.Int",
			args: args{
				v: map[int64]*big.Int{},
			},
			want: jen.Map(jen.Int64()).Op("*").Qual("math/big", "Int"),
		},
		{
			name: "map[int64]big.Int",
			args: args{
				v: map[int64]big.Int{},
			},
			want: jen.Map(jen.Int64()).Qual("math/big", "Int"),
		},
		{
			name: "[16]int64",
			args: args{
				v: [16]int64{},
			},
			want: jen.Index(jen.Lit(16)).Int64(),
		},
		{
			name: "[16]*int64",
			args: args{
				v: [16]*int64{},
			},
			want: jen.Index(jen.Lit(16)).Op("*").Int64(),
		},
		{
			name: "uuid.UUID",
			args: args{
				v: uuid.UUID{},
			},
			want: jen.Qual("github.com/google/uuid", "UUID"),
		},
		{
			name: "*uuid.UUID",
			args: args{
				v: &uuid.UUID{},
			},
			want: jen.Op("*").Qual("github.com/google/uuid", "UUID"),
		},
		{
			name: "[]uuid.UUID",
			args: args{
				v: []uuid.UUID{},
			},
			want: jen.Index().Qual("github.com/google/uuid", "UUID"),
		},
		{
			name: "[]*uuid.UUID",
			args: args{
				v: []*uuid.UUID{},
			},
			want: jen.Index().Op("*").Qual("github.com/google/uuid", "UUID"),
		},
		{
			name: "[16]uuid.UUID",
			args: args{
				v: [16]uuid.UUID{},
			},
			want: jen.Index(jen.Lit(16)).Qual("github.com/google/uuid", "UUID"),
		},
		{
			name: "[16]*uuid.UUID",
			args: args{
				v: [16]*uuid.UUID{},
			},
			want: jen.Index(jen.Lit(16)).Op("*").Qual("github.com/google/uuid", "UUID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Type(tt.args.v)
			assert.Equal(t, fmt.Sprintf("%#v", tt.want),
				fmt.Sprintf("%#v", got))
		})
	}
}
