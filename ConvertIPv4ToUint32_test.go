package main

import (
	"testing"
)

func TestConvertIPv4ToUint32(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{name: "Valid IP", args: args{ip: "172.168.5.1"}, want: 2896692481, wantErr: false},
		{name: "Valid IP with spaces", args: args{ip: " 172. 168.5.1 "}, want: 2896692481, wantErr: false},
		{name: "All zeroes", args: args{ip: "0.0.0.0"}, want: 0, wantErr: false},
		{name: "All ones", args: args{ip: "255.255.255.255"}, want: 4294967295, wantErr: false},
		{name: "Common private IP", args: args{ip: "192.168.1.1"}, want: 3232235777, wantErr: false},
		{name: "Spaces within segment", args: args{ip: "1 72.168.5.1"}, want: 0, wantErr: true},
		{name: "Segment out of range", args: args{ip: "172.168.5.256"}, want: 0, wantErr: true},
		{name: "Empty segment", args: args{ip: "172.168..1"}, want: 0, wantErr: true},
		{name: "Trailing space in segment", args: args{ip: "172.168.5. 1"}, want: 2896692481, wantErr: false},
		{name: "All segments out of range", args: args{ip: "256.256.256.256"}, want: 0, wantErr: true},
		{name: "Non-digit character", args: args{ip: "172.a.5.1"}, want: 0, wantErr: true},
		{name: "Extra dots", args: args{ip: "172.168.5.1."}, want: 0, wantErr: true},
		{name: "Invalid IP format", args: args{ip: "172.168.5"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertIPv4ToUint32(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertIPv4ToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertIPv4ToUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}
