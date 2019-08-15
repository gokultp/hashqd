package commands

import (
	"io"
	"testing"
)

func TestPut_Decode(t *testing.T) {
	type fields struct {
		data   []byte
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Put{
				data:   tt.fields.data,
				reader: tt.fields.reader,
			}
			if err := c.Decode(); (err != nil) != tt.wantErr {
				t.Errorf("Put.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
