package main

import "testing"

func Test_bislCommand(t *testing.T) {
	type args struct {
		className string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test BIS list for warrior", args{"warrior dps"},
			"Helm --->  Lionheart Helm\n" +
			"Neck --->  Mark of Fordring\n" +
			"Shoulder --->  Truestrike Shoulders\n" +
			"Cloak --->  Cape of the Black Baron\n" +
			"Chest --->  Savage Gladiator Chain\n" +
			"Wrists --->  Battleborn Armbraces\n" +
			"Gloves --->  Edgemaster's Handguards\n" +
			"Belt --->  Brigam Girdle\n" +
			"Legs --->  Cloudkeeper Legplates\n" +
			"Boots --->  Bloodmail Boots\n" +
			"Rings --->  Blackstone Ring\n" +
			"Trinkets --->  Hand of Justice\n" +
			"Main-Hand --->  Ironfoe\n" +
			"Off-Hand --->  Felstriker\n" +
			"Two-Hand --->  Blackblade of Shahram\n" +
			"Ranged --->  Riphook\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bislCommand(tt.args.className); got != tt.want {
				t.Errorf("bislCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}