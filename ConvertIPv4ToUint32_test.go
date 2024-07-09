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
		{name: "正确的 ip", args: args{ip: "172.168.5.1"}, want: 2896692481, wantErr: false},
		{name: "数字前有空格", args: args{ip: "172. 168.5.1"}, want: 2896692481, wantErr: false},
		{name: "数字内有空格", args: args{ip: "1 72.168.5.1"}, want: 0, wantErr: true},
		{name: "数字前后都有空格", args: args{ip: "172.168.5. 1 "}, want: 2896692481, wantErr: false},
		{name: "多个数字前后空格同时存在", args: args{ip: "172 . 168 .    5.1"}, want: 2896692481, wantErr: false},
		{name: "0.0.0.0", args: args{ip: "0.0.0.0"}, want: 0, wantErr: false},
		{name: "255.255.255.255", args: args{ip: "255.255.255.255"}, want: 4294967295, wantErr: false},
		{name: "数字超出范围 256", args: args{ip: "172.168.5.256"}, want: 0, wantErr: true},
		{name: "数字超出范围 -1", args: args{ip: "172.168.5.-1"}, want: 0, wantErr: true},
		{name: "数字为空", args: args{ip: "172.168..1"}, want: 0, wantErr: true},
		{name: "非数字的字符", args: args{ip: "172.a.5.1"}, want: 0, wantErr: true},
		{name: "ip末尾多余.", args: args{ip: "172.168.5.1."}, want: 0, wantErr: true},
		{name: "ip首端多余.", args: args{ip: ".172.168.5.1"}, want: 0, wantErr: true},
		{name: "数字不足 4 个", args: args{ip: "172.168.5"}, want: 0, wantErr: true},
		{name: "数字超出 4 个", args: args{ip: "172.168.5.1.1"}, want: 0, wantErr: true},
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
