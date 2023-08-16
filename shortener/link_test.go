package shortener

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	type args struct {
		shortURL string
		longURL  string
	}
	tests := []struct {
		name    string
		args    args
		want    Link
		wantErr error
	}{
		{
			name: "test should fail because shortURL is empty",
			args: args{
				shortURL: "",
			},
			want:    Link{},
			wantErr: errors.New("shorturl: is empty"),
		},
		{
			name: "test should fail because longURL is empty",
			args: args{
				shortURL: "asfd",
				longURL:  "",
			},
			want:    Link{},
			wantErr: errors.New("longurl: is empty"),
		},
		{
			name: "test should fail because its not a valid URL",
			args: args{
				shortURL: "sdf",
				longURL:  "wrongurl",
			},
			want:    Link{},
			wantErr: errors.New("not a valid URL"),
		},
		{
			name: "test should pass",
			args: args{
				shortURL: "twt",
				longURL:  "https://twitter.com/urusername",
			},
			want: Link{
				ShortURL: "twt",
				LongURL:  "https://twitter.com/urusername",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLink(tt.args.shortURL, tt.args.longURL)
			tt.want.ID = got.ID
			tt.want.CreatedAt = got.CreatedAt
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
