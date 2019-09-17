package main

import "testing"

func Test_findCommand(t *testing.T) {
	type args struct {
		npcName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Should return correct info for Ragnaros", args{"Ragnaros"},
			"Ragnaros is a boss that can be found in Molten Core. The location of this NPC is unknown. In the NPCs category. Added in Classic World of Warcraft.\n" +
			"https://wow.zamimg.com/uploads/screenshots/normal/1061-ragnaros-ragnaros-after-killing-us.jpg"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommand(tt.args.npcName); got != tt.want {
				t.Errorf("findCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}