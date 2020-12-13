package hello

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{
			name: "world",
			args: args{
				name: "world",
			},
			wantWriter: "Hello, world!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			Hello(writer, tt.args.name)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Hello() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
