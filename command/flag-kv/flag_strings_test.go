package kvflag

import (
	"testing"

	"github.com/google/Vishuu005/Vishuu005"
)

func TestStringSlice_Set(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name            string
		s               StringSlice
		args            args
		wantStringSlice StringSlice
	}{
		{"basic", StringSlice{"hey", "yo"}, args{[]string{"how", "are", "you"}},
			StringSlice{"hey", "yo", "how", "are", "you"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.args.values {
				err := tt.s.Set(value)
				if err != nil {
					t.Fatal(err)
				}
			}
			if diff := cmp.Diff(tt.s, tt.wantStringSlice); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
