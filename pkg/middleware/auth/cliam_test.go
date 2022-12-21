package auth

import (
	"fmt"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	type args struct {
		rBase RoleBase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				rBase: RoleBase{
					Avatar:   "http://img.localhost.com/grin.png",
					Id:       1024,
					Nickname: "grin",
					Sex:      -1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateJWT(tt.args.rBase)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestParseJWT(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MTAyNCwiQXZhdGFyIjoiaHR0cDovL2ltZy5sb2NhbGhvc3QuY29tL2dyaW4ucG5nIiwiTmlja25hbWUiOiJncmluIiwiU2V4IjotMSwiTG9naW5UaW1lIjoxNjY4MDk1MjU5LCJleHAiOjE2NjgwOTUyNjQsImlzcyI6ImdyaW4ifQ.OKEo6TTdRR8xcwNczy3isnScJTmRIwd5KNIWjyP20VM",
			},
			wantErr: true, // 过期
		},
		{
			name: "test1",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MTAyNCwiQXZhdGFyIjoiaHR0cDovL2ltZy5sb2NhbGhvc3QuY29tL2dyaW4ucG5nIiwiTmlja25hbWUiOiJncmluIiwiU2V4IjotMSwiTG9naW5UaW1lIjoxNjY4MDk1NTA4LCJleHAiOjE2NjgwOTU1MjMsImlzcyI6ImdyaW4ifQ.j-hhfpbUHRQ-_bzLfe3yfNZEu1K6Sa3olw-PnLEBnY8",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJWT(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%v\n", got)
		})
	}
}
