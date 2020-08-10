package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamel(t *testing.T) {
	tests := []struct {
		input,
		want string
	}{
		{
			input: "Hello camel",
			want:  "HelloCamel",
		},
		{
			input: "Hello-camel",
			want:  "HelloCamel",
		},
		{
			input: "Hello_under_camel",
			want:  "HelloUnderCamel",
		},
		{
			input: "Asd__--  _sep _every_wheres asdf___",
			want:  "AsdSepEveryWheresAsdf",
		},
		{
			input: "hello",
			want:  "Hello",
		},
		{
			input: "id",
			want:  "Id",
		},
		{
			input: "HttpServer",
			want:  "HttpServer",
		},
		{
			input: "Id",
			want:  "Id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ToCamel(tt.input)
			assert.Equal(t, tt.want, got, fmt.Sprintf("input: %q, expect: %q actual: %q", tt.input, tt.want, got))
		})
	}
}

var resultCamel string

func BenchmarkToCamel(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = ToLowerCamel("_--34-asd__--  _sep _every_wherea asdf___")
	}
	resultCamel = s
}

func TestToLowerCamel(t *testing.T) {
	tests := []struct {
		input,
		want string
	}{
		{
			input: "Hello world of 123camel",
			want:  "helloWorldOf123Camel",
		},
		{
			input: "Hello-camel",
			want:  "helloCamel",
		},
		{
			input: "Hello_under_camel",
			want:  "helloUnderCamel",
		},
		{
			input: "Asd__--  _sep_3 _every_wherea asdf___",
			want:  "asdSep3EveryWhereaAsdf",
		},
		{
			input: "Hello",
			want:  "hello",
		},
		{
			input: "Http-Server",
			want:  "httpServer",
		},
		{
			input: "Http_Test",
			want:  "httpTest",
		},
		{
			input: "id",
			want:  "id",
		},
		{
			input: "mY Id",
			want:  "mYId",
		},
		{
			input: "ID",
			want:  "id",
		},
		{
			input: "HTTP",
			want:  "http",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ToLowerCamel(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}

}

var resultLowerCamel string

func BenchmarkToLowerCamel(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = ToLowerCamel("_--34-asd__--  _sep _every_wherea asdf___")
	}
	resultLowerCamel = s
}
