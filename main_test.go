package main

import (
	"errors"
	"io"
	"reflect"
	"testing"
)

type errReader struct{}

func (reader errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("bad reader")
}

type nilReader struct{}

func (reader nilReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func TestGenerateKeys(t *testing.T) {
	type args struct {
		rand io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantPk  []byte
		wantSk  []byte
		wantErr bool
	}{
		{
			name:    "err",
			args:    args{errReader{}},
			wantErr: true,
		},
		{
			name: "normal",
			args: args{nilReader{}},
			wantPk: []byte{
				47, 229, 125, 163, 71, 205, 98, 67, 21, 40, 218, 172, 95, 187, 41, 7,
				48, 255, 246, 132, 175, 196, 207, 194, 237, 144, 153, 95, 88, 203, 59, 116,
			},
			wantSk: make([]byte, KeySize),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPk, gotSk, err := GenerateKeys(tt.args.rand)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPk, tt.wantPk) {
				t.Errorf("GenerateKeys() Public Key\ngot = %v\nwant  %v", gotPk, tt.wantPk)
			}
			if !reflect.DeepEqual(gotSk, tt.wantSk) {
				t.Errorf("GenerateKeys() Private Key\ngot = %v\nwant  %v", gotSk, tt.wantSk)
			}
		})
	}
}
