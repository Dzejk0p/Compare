package compare

import (
	"reflect"
	"testing"
	"time"
)

func TestDiffer_diff(t *testing.T) {

	type testStruct struct {
		Int    int       `diff:"int"`
		String string    `diff:"string"`
		Date   time.Time `diff:"date"`
	}

	type testStruct2 struct {
		Int int               `diff:"int"`
		Map map[string]string `diff:"map"`
	}

	data1 := time.Now()
	data2 := time.Now().Add(1 * time.Hour)

	type args struct {
		a reflect.Value
		b reflect.Value
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
				a: reflect.ValueOf(testStruct{1, "a", data1}),
				b: reflect.ValueOf(testStruct{2, "b", data2}),
			},
			wantErr: false,
			want: Zmiany{
				"int": Pole{
					Bylo: 1,
					Jest: 2,
				},
				"string": Pole{
					Bylo: "a",
					Jest: "b",
				},
				"date": Pole{
					Bylo: data1,
					Jest: data2,
				},
			},
			wantCzyZmiana: true,
		},
		{
			name: "takie same",
			d:    NewDiffer(),
			args: args{
				a: reflect.ValueOf(testStruct{1, "a", data1}),
				b: reflect.ValueOf(testStruct{1, "a", data1}),
			},
			wantErr: false,
			want: Zmiany{
				"int": Pole{
					Bylo: nil,
					Jest: 1,
				},
				"string": Pole{
					Bylo: nil,
					Jest: "a",
				},
				"date": Pole{
					Bylo: nil,
					Jest: data1,
				},
			},
			wantCzyZmiana: false,
		},
		{
			name: "nieobslugiwany typ",
			d:    NewDiffer(),
			args: args{
				a: reflect.ValueOf(testStruct2{1, map[string]string{}}),
				b: reflect.ValueOf(testStruct2{1, map[string]string{}}),
			},
			wantErr:       true,
			want:          Zmiany{},
			wantCzyZmiana: false,
		},
		{
			name: "rozne typy",
			d:    NewDiffer(),
			args: args{
				a: reflect.ValueOf(testStruct2{1, map[string]string{}}),
				b: reflect.ValueOf(testStruct{1, "a", time.Now()}),
			},
			wantErr:       true,
			want:          Zmiany{},
			wantCzyZmiana: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.d.diff(tt.args.a, tt.args.b, "", nil)
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
