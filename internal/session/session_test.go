package session

import "testing"

func TestSession_SetTube(t *testing.T) {
	type fields struct {
		Tube string
	}
	type args struct {
		tube string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "check if the tube is setting properly or not",
			fields: fields{Tube: ""},
			args:   args{tube: "asdf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				Tube: tt.fields.Tube,
			}
			s.SetTube(tt.args.tube)
			if s.Tube != tt.args.tube {
				t.Errorf("tube not matching")
			}
		})
	}
}
