package main

import "testing"

func Test_itemCommand(t *testing.T) {
	type args struct {
		itemName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Should find correct information for Ashkandi", args{"Ashkandi"}, "This epic two-handed sword has an item level of 81. It is looted from Nefarian. In the Two-Handed Swords category. Added in Classic World of Warcraft.\n" +
			"https://wow.zamimg.com/uploads/screenshots/normal/1321-ashkandi-greatsword-of-the-brotherhood.jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itemCommand(tt.args.itemName); got != tt.want {
				t.Errorf("itemCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}