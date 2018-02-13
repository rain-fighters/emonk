# emonk <img src="./doc/emonk.png" align="right" height="64" width="64" />
A *Discord* bot written in *Go (Golang)*

## Main (planned) features
* open source
* people may build and run their own instance and register their own Discord App and Bot User
* also available as a service (my instance running on a virtual server at google, hetzner, ...)
* needs to be invited/added by a server/guild owner using the bot's OAUTH2_URL
* will refuse to act when given administrative privileges on server/guild
* has a command interface using mentions (@emonk ...)
	* for configuration by an admin (server/guild owner or role with administrtaive privileges on the server/guild))
	* for moderation tasks by an operator (role)
	* for service tasks by a user (role or @everyone)

## Maybe features
* Sync messages from selected discord channel(s) to *YouTube Gaming Chat*
