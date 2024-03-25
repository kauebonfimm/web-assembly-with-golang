package passwordengine_test

import (
	"strings"
	"testing"

	. "github.com/kauebonfimm/web-assembly-with-golang/internal/core/password_engine"
)

func TestGeneratePassword(t *testing.T) {
	type args struct {
		length     uint16
		hasLetters bool
		hasDigits  bool
		hasSymbos  bool
		removing   string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Succeess to generate password",
			args: args{
				length:     10,
				hasLetters: true,
				hasDigits:  true,
				hasSymbos:  true,
				removing:   "",
			},
			wantErr: false,
		},
		{
			name: "Succeess to generate password without letters",
			args: args{
				length:     10,
				hasLetters: false,
				hasDigits:  true,
				hasSymbos:  true,
				removing:   "",
			},
			wantErr: false,
		},
		{
			name: "Succeess to generate password without digits",
			args: args{
				length:     10,
				hasLetters: true,
				hasDigits:  false,
				hasSymbos:  true,
				removing:   "",
			},

			wantErr: false,
		},
		{
			name: "Succeess to generate password without symbols",
			args: args{
				length:     10,
				hasLetters: true,
				hasDigits:  true,
				hasSymbos:  false,
				removing:   "",
			},
			wantErr: false,
		},
		{
			name: "Succeess to generate password removing characters",
			args: args{
				length:     10,
				hasLetters: true,
				hasDigits:  true,
				hasSymbos:  true,
				removing:   "g",
			},
			wantErr: false,
		},
		{
			name: "Fail to generate password without any parameters",
			args: args{
				length:     10,
				hasLetters: false,
				hasDigits:  false,
				hasSymbos:  false,
				removing:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := GeneratePassword(tt.args.length, tt.args.hasLetters, tt.args.hasDigits, tt.args.hasSymbos, tt.args.removing)

			t.Log(password, err)
			if tt.wantErr && err == nil {

				t.Errorf("GeneratePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(password) != int(tt.args.length) && !tt.wantErr {
				t.Errorf("GeneratePassword() = %v, want %v", len(password), tt.args.length)
				return
			}

			if tt.args.hasLetters {
				if !strings.ContainsAny(password, LETTERS) {
					t.Errorf("GeneratePassword() = %v, want %v", password, LETTERS)
					return
				}
			}

			if tt.args.hasDigits {
				if !strings.ContainsAny(password, DIGITS) {
					t.Errorf("GeneratePassword() = %v, want %v", password, DIGITS)
					return
				}
			}

			if tt.args.hasSymbos {
				if !strings.ContainsAny(password, SYMBOLS) {
					t.Errorf("GeneratePassword() = %v, want %v", password, SYMBOLS)
					return
				}
			}

			if tt.args.removing != "" {
				for _, r := range tt.args.removing {
					if strings.ContainsAny(password, string(r)) {
						t.Errorf("GeneratePassword() = %v, want %v", password, tt.args.removing)
						return
					}
				}
			}

		})
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(10, true, true, true, "")
	}
}

func BenchmarkGeneratePasswordWithoutLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(10, false, true, true, "")
	}
}

func BenchmarkGeneratePasswordWithoutDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(10, true, false, true, "")
	}
}

func BenchmarkGeneratePasswordWithoutSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(10, true, true, false, "")
	}
}

func BenchmarkGeneratePasswordRemovingCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(10, true, true, true, "g")
	}
}
