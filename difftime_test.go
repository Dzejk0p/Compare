package compare

import (
	"reflect"
	"testing"
	"time"
)

func TestDiffer_diffTime(t *testing.T) {
	type args struct {
		a    reflect.Value
		b    reflect.Value
		name string
	}
	tests := []struct {
		name          string
		d             *Differ
		args          args
		wantErr       bool
		want          Zmiany
		wantCzyZmiana bool
	}{
		{
			name: "zwykle rozne",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(time.Now()),
				b:    reflect.ValueOf(time.Now().Add(1 * time.Hour)),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: time.Now(),
					Jest: time.Now().Add(1 * time.Hour),
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "takie same",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(time.Now()),
				b:    reflect.ValueOf(time.Now()),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: nil,
					Jest: time.Now(),
				},
			},
			wantCzyZmiana: false,
		},
		{
			name: "a nil",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(nil),
				b:    reflect.ValueOf(time.Now()),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: nil,
					Jest: time.Now(),
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "b nil",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(time.Now()),
				b:    reflect.ValueOf(nil),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: time.Now(),
					Jest: nil,
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "rozne typy",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf("a"),
				b:    reflect.ValueOf(time.Now()),
				name: "test",
			},
			wantErr:       true,
			want:          Zmiany{},
			wantCzyZmiana: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.d.diffTime(tt.args.a, tt.args.b, tt.args.name, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Differ.diffTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !tt.wantErr {
				if !reflect.DeepEqual(tt.want, tt.d.Zmiany) {
					t.Errorf("Differ.diffTime() błędny wynik.\nChce:\n%v\nMam:\n%v\n", tt.want, tt.d.Zmiany)
				}

				if tt.wantCzyZmiana != tt.d.CzyZmiana {
					t.Error("Differ.diffTime() błąd flagi zmiany.")
				}
			}
		})
	}
}
