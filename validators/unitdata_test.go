package validators

import "testing"

func TestNickname(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := Key(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("Nickname() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
