package domain

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func Test_cutRange(t *testing.T) {
	type args struct {
		r     string
		slice []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				r:     "1-4",
				slice: []byte("1234567890"),
			},
			want:    []byte("1234"),
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				r:     "2-4",
				slice: []byte("1234567890"),
			},
			want:    []byte("234"),
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				r:     "1-4,7",
				slice: []byte("1234567890"),
			},
			want:    []byte("12347"),
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				r:     "-4,7",
				slice: []byte("1234567890"),
			},
			want:    []byte("12347"),
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				r:     "-4,7-",
				slice: []byte("1234567890"),
			},
			want:    []byte("12347890"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cutRange(tt.args.r, tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("cutRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRange(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name    string
		args    args
		want    []parsedRange
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				r: "1-4",
			},
			want:    []parsedRange{{1, 4, false}},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				r: "1-4,5-7",
			},
			want:    []parsedRange{{1, 4, false}, {5, 7, false}},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				r: "1-4,5-7,9",
			},
			want:    []parsedRange{{1, 4, false}, {5, 7, false}, {9, 0, true}},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				r: "1-4,5-7,9-",
			},
			want:    []parsedRange{{1, 4, false}, {5, 7, false}, {9, 0, false}},
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				r: "-4,5-7,9-",
			},
			want:    []parsedRange{{0, 4, false}, {5, 7, false}, {9, 0, false}},
			wantErr: false,
		},
		{
			name: "test6",
			args: args{
				r: "1,-4,5-7,9-,12",
			},
			want:    []parsedRange{{1, 0, true}, {0, 4, false}, {5, 7, false}, {9, 0, false}, {12, 0, true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRange(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomain_Cut(t *testing.T) {
	type fields struct {
		Log *zap.Logger
	}
	type args struct {
		in    io.Reader
		flags Flags
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Log: zap.NewNop(),
			},
			args: args{
				in: strings.NewReader("1234567890"),
				flags: Flags{
					B: "1-4,7-",
				},
			},
			want:    []string{"12347890"},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Log: zap.NewNop(),
			},
			args: args{
				in: strings.NewReader("Hello, world!"),
				flags: Flags{
					B: "1-4,7-",
				},
			},
			want:    []string{"Hell world!"},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Log: zap.NewNop(),
			},
			args: args{
				in: strings.NewReader("Hello, world!"),
				flags: Flags{
					B: "-4,7-",
				},
			},
			want:    []string{"Hell world!"},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Log: zap.NewNop(),
			},
			args: args{
				in: strings.NewReader("Hello, world!"),
				flags: Flags{
					C: "1-4,7-",
				},
			},
			want:    []string{"Hell world!"},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Log: zap.NewNop(),
			},
			args: args{
				in: strings.NewReader("Hello, world!"),
				flags: Flags{
					F: "1",
					D: ",",
				},
			},
			want:    []string{"Hello"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Domain{
				Log: tt.fields.Log,
			}
			got, err := d.Cut(tt.args.in, tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("Domain.Cut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Domain.Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}
