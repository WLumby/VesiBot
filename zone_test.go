package main

import "testing"

func Test_zoneCommand(t *testing.T) {
	type args struct {
		zoneName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Should return details about Hillsbrad", args{"Hillsbrad Foothills"}, "A level 20-31 contested zone. In the  Zones category. Added in Classic World of Warcraft. Always up to date with the latest patch (1.13.2).\n" +
			"https://wow.zamimg.com/uploads/screenshots/normal/80823-hillsbrad-foothills.jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zoneCommand(tt.args.zoneName); got != tt.want {
				t.Errorf("zoneCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}