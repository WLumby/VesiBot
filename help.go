package main

const helpText = "List of commands:\n" +
	"!vesi help - Displays list of commands\n" +
	"!vesi sorc - Displays source code for bot\n" +
	"!vesi find <npc name> - Find the zone an NPC is located in\n" +
	"!vesi item <item name> - Find where an item is from\n" +
	"!vesi zone <zone name> - Find details about a zone\n" +
	"!vesi dung <dungeon name> - Find strategy guide for a dungeon\n" +
	"!vesi bisl <class and role> (E.g. !vesi bisl priest healer)\n" +
	"!vesi qust <quest name / related npc>\n" +
	"!vesi strs - Til the day I ******* die.\n"

func HelpCommand() string {
	return helpText
}
