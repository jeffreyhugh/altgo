What is this?
-----
This is a self-hosted Discord bot capable of managing multiple Minecraft alts simultaneously.
It's designed for AFK farms where multiple accounts need to be loading multiple chunks. Since
none of the accounts are rendering the game, it's incredibly lightweight.

**I offer no guarantee that accounts will not be banned either by Mojang or third-party server
owners. Use at your own risk.**

Setup
-----
In `constants/tokens/`, there are two files prefixed with `EXAMPLE-CHANGEME`. Both need to be
updated.

In `constants/tokens/EXAMPLE-CHANGEME-discord.go`:
  1. Uncomment all the lines
  1. Add your Discord token
  1. Rename the file to `constants/tokens/discord.go`
  
In `constants/tokens/EXAMPLE-CHANGEME-minecraft.go`:
  1. Uncomment all the lines
  1. Add your Minecraft accounts
     - If the account is **migrated**, you'll need to include the email address to sign in with
     - If the account is **unmigrated**, set the email address to a blank string (`""`)
     - Use the `a(ign, email, password string, migrated bool)` function.
  1. Rename the file to `constants/tokens/minecraft.go`
    
Then, `go run main.go`

Commands
-----
All commands are prefixed with `/` by default (changed in `constants/constants.go`)

  - `/connect <IGN> <server><:port>`
  - Example: `/connect Xx_Example_xX mc.hypixel.net`
  - Example: `/connect Xx_Example_xX 127.0.0.1:25577`
    - `<IGN>` : an IGN with login information in `constants/tokens/minecraft.go`
    - `<server>` : an IP address or domain name
    - `<:port>` : (optional) port, default to `25565`. 
---
  - `/disconnect <IGN>`
  - Example: `/disconnect Xx_Example_xX`
    - `<IGN>` : an IGN to disconnect
---
  - `/chat <IGN> <message>`
  - Example: `/chat Xx_Example_xX hello, world!`
  - Example: `/chat Xx_Example_xX /summon minecraft:lightning_bolt ~ ~ ~`
    - `<IGN>` : a logged-in account to chat from
    - `<message>` : a message with length \leq 256
---
  - `/follow <IGN> <target> <duration> <nocheck>`
  - Example: `/follow Xx_Example_xX target69 10`
  - Example: `/follow Xx_Example_xX target69 60 nocheck` (dangerous)
    - `<IGN>` : a logged-in account
    - `<target>` : player to follow
    - `<duration>` : how long to follow in seconds
    - `<nocheck>` : DO NOT USE! Forgo the check to make sure the account and target are in the
    same block.
---
  - `/goto <IGN> <target> <nocheck>`
  - Example: `/goto Xx_Example_xX target69`
  - Example: `/goto Xx_Example_xX target69 nocheck` (dangerous)
    - `<IGN>` : a logged-in account
    - `<target>` : player to goto
    - `<nocheck>` : DO NOT USE! Forgo the check to make sure the account and target are in the
    same block.
    
Known Bugs
-----
  - ~~Certain servers send an MoTD on join that causes an error. This is a bug according to the library
  developer, and may or may not get fixed.~~ [Fixed according to dev](https://github.com/Tnze/go-mc/commit/67806abcdb744eebeca5f3a1d8d0107a2d5cbf46)
  - Certain servers have a proxy which redirects traffic to the "real server" (e.g. `hypixel.net --> 
  mc.hypixel.net`). This causes an error.
  - Movement generally sucks because of the client's inability to keep track of player positions. Use
  with extreme caution, knowing it will likely cause a slew of "moved too fast" or "moved wrongly" 
  warnings (especially when changing altitude). 
    
TODO
-----
  - [X] Send chat message
  - [X] Follow player command
  - [ ] Anti-afk jitter
  - [ ] POST to Discord webhook if account disconnects
  - [ ] Get UUIDs from helper instead of hardcode
  
License
-----
This code is licensed under Gnu GPL3, with the additional stipulation that derivatives and
redistributions must give partial credit to the original author.

This part isn't required, but if you end up using the software, I'd love to hear about your use case!
Feel free to reach me on Discord (QueueBot#1111) or email (q@queue.bot).