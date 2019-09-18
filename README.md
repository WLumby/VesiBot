# VesiBot

#### The basic WoW Classic Helper bot.

VesiBot is a basic discord bot written in Go that can quickly grab scraped information
from sources about WoW Classic; including NPC locations, quest information, and BiS lists
for different classes. See below for usage information.

### Setup

Place your bot token (https://discordapp.com/developers/) in a text file named
"token.txt" in the main /VesiBot/ directory. This file will be used to connect to your
discord developer API.

Navigate to /bin/ in your CLI and type:

```
./run.sh
```

The shell script should then test your version and start the bot, connecting to
your developer API.

You should then see in your CLI:

```
VesiBot has started on # servers
```

This means the bot has started successfully.

To use the bot in your server, you will need to add it and give it
the requested permissions.
https://github.com/jagrosh/MusicBot/wiki/Adding-Your-Bot-To-Your-Server

### Usage

VesiBot has a continually updating list of official commands:

| Command   | Description |
| --------- | ----------- |
|!vesi help | Displays list of commands
|!vesi find npc name | Find the zone an NPC is located in
|!vesi item item name | Find where an item is from
|!vesi zone zone name | Find details about a zone
|!vesi dung dungeon name | Find strategy guide for a dungeon
|!vesi bisl class and role | Find BIS list for a class/role
|!vesi qust quest name / related npc | Find a quest
|!vesi strs | Til the day I ******* die. |

These commands can also be seen in discord by using !vesi help