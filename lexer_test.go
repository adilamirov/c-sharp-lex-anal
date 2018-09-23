package main

import "testing"

func TestLexer_LoadText(t *testing.T) {
	tests := []struct {
		name string
		text []byte
	}{
		{"empty code", []byte{}},
		{"some code", []byte("bajsfklasdiofhasdf")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{}
			l.LoadText(tt.text)
		})
	}
}

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		text   []byte
		cursor int
		token  string
		tipe   TokenType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"end of code",
			fields{
				[]byte("   		  \n\n	   "),
				5,
				"",
				0},
			false,
		},
		{
			"next is operator",
			fields{
				[]byte("class Program{static void Main(){int test = 2;float item = 5.5f;char unit = 'e';string basic = \"c#\";}}"),
				29,
				"",
				0},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				text:   tt.fields.text,
				cursor: tt.fields.cursor,
				token:  tt.fields.token,
				tipe:   tt.fields.tipe,
			}
			if got := l.NextToken(); got != tt.want {
				t.Errorf("Lexer.NextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
