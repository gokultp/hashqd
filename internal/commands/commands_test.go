package commands

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestGetCommand(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should get put cmd object if command is put",
			args: args{
				r: bytes.NewReader([]byte("put 4\r\ntest\r\n")),
			},
			want:    "put",
			wantErr: false,
		},
		{
			name: "should get error if command is invalid",
			args: args{
				r: bytes.NewReader([]byte("puts 4\r\ntest\r\n")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommand(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got.Command(), tt.want) {
				t.Errorf("GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCommand(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "should return the command", args: args{r: strings.NewReader("put 4\r\ntest")}, want: "put", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCommand(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
