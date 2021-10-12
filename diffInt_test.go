package diff

import (
	"reflect"
	"testing"
)

func TestDiffer_diffInt(t *testing.T) {
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
				a:    reflect.ValueOf(1),
				b:    reflect.ValueOf(2),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: 1,
					Jest: 2,
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "takie same",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(2),
				b:    reflect.ValueOf(2),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: nil,
					Jest: 2,
				},
			},
			wantCzyZmiana: false,
		},
		{
			name: "a nil",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(nil),
				b:    reflect.ValueOf(2),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: nil,
					Jest: 2,
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "b nil",
			d:    NewDiffer(),
			args: args{
				a:    reflect.ValueOf(2),
				b:    reflect.ValueOf(nil),
				name: "test",
			},
			wantErr: false,
			want: Zmiany{
				"test": Pole{
					Bylo: 2,
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
				b:    reflect.ValueOf(2),
				name: "test",
			},
			wantErr:       true,
			want:          Zmiany{},
			wantCzyZmiana: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.d.diffInt(tt.args.a, tt.args.b, tt.args.name, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Differ.diffInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !tt.wantErr {
				if !reflect.DeepEqual(tt.want, tt.d.Zmiany) {
					t.Errorf("Differ.diffInt() błędny wynik.\nChce:\n%v\nMam:\n%v\n", tt.want, tt.d.Zmiany)
				}

				if tt.wantCzyZmiana != tt.d.CzyZmiana {
					t.Error("Differ.diffInt() błąd flagi zmiany.")
				}
			}
		})
	}
}
