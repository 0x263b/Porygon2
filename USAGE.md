### Functions

* [8ball](#8ball)
* [Bing](#bing)
* [Lastfm](#lastfm)
* [Random](#random)
* [Translate](#translate)
* [TVMaze](#tvmaze)
* [Twitter](#twitter)
* [Urban Dictionary](#urban-dictionary)
* [URL Parser](#url-parser)
* [User Profiles](#user-profiles)
* [Weather](#weather)
* [WolframAlpha](#wolframalpha)
* [Youtube](#youtube)
* [Admin functions](#admin-functions)

***

### 8Ball
Gives and 8ball style answer to a *question*

**-8ball** *question*

	-8ball Am I going to score with this one girl I just finished talking to?
	My sources say no


### Bing
Gets the first result from [Bing](https://www.bing.com/) for *search query*

**-b/-bing** *search query*

	-bing Richard Stallman
	Bing | Richard Stallman's Personal Page | http://stallman.org/

Gets the first result from [Bing image search](https://www.bing.com/images) for *search query*

**-img** *search query*

	-img Richard Stallman
	Bing | Richard Stallman → image/jpeg 257 kB | http://www.straferight.com/photopost/data/500/richard-stallman.jpg


### Last.fm
Associates your current irc nick with *user*.
Other lastfm functions will default to this nick if no user is provided.

**-set lastfm** *user*
	
	<joebloggs> -set lastfm JosefBloggs
	<Porygon> joebloggs: last.fm user updated to: JosefBloggs
 

Weekly stats for *user*

**-charts** *user*

	-charts Cbbleh
	Last.fm | Top 5 Weekly artists for Cbbleh | Slayer (26), Iced Earth (25), Jean-Féry Rebel (23), Morbid Saint (15), Judas Priest (14)


Returns the currently playing/last scrobbled track for *user* and top artist tags

**-np** *user*
	
	-np Cbbleh
	Last.fm | cbbleh is playing: "Super X-9" by Daikaiju from Daikaiju | Surf, surf rock, instrumental, instrumental surf rock


### Random
Randomly picks an option from an array separated by |

**-r/-rand** `one | two | three`

	-r do work | don't do work
	don't do work


### Translate
Translates *text* using [Bing translate](http://www.bing.com/translator) and provides a link to [Google Translate](http://translate.google.com/)

**-tr/-translate** *from to text*

	-translate en fr pig disgusting
	Translate | http://mnn.im/uoo4u | en=>fr | "porc écoeurant"

| code | Language           | code | Language           | code | Language |
| ---: | :----------------- | ---: | :----------------- | ---: | :------- |
| ar   | Arabic					| ht   | Haitian Creole		| fa   | Persian (Farsi) 
| bg   | Bulgarian				| he   | Hebrew				| pl   | Polish	
| ca   | Catalan				| hi   | Hindi				| ro   | Romanian
| cn   | Chinese Simplified		| mww  | Himong Daw			| ru   | Russian
| tw   | Chinese Traditional 	| hu   | Hungarian			| sk   | Slovak
| cs   | Czech					| id   | Indonesian			| sl   | Slovenian
| da   | Danish					| it   | Italian			| es   | Spanish
| nl   | Dutch					| ja   | Japanese			| sv   | Swedish
| en   | English				| ko   | Korean				| th   | Thai
| et   | Estonian				| lv   | Latvian			| tr   | Turkish
| fi   | Finnish				| lt   | Lithuanian			| uk   | Ukrainian
| fr   | French					| ms   | Malay				| ur   | Urdu
| de   | German					| mt   | Maltese			| vi   | Vietnamese
| el   | Greek					| no   | Norwegian


### TVMaze
Info for *tv show* with episode airtime if available **-tv** *tv show*

	-tv Better call saul
	TVmaze | Better Call Saul | Airtime: Monday 22:00 on AMC | Status: Running | Next Ep: S2E6 at 22:00 2016-03-21
	
	-tv Mr Robot
	TVmaze | Mr. Robot | Airtime: Wednesday 22:00 on USA Network | Status: Running


### Twitter
Latest tweet for *user* **-tw/-twitter** *user*

	-twitter Guardian
	Twitter | The Guardian (@guardian) | Aston Villa target Rémi Garde after sacking Tim Sherwood https://t.co/cqcgUpiEOJ via @guardian_sport | 31 seconds ago


### Urban Dictionary
Gets the first definition of *query* at [UrbanDictionary](http://www.urbandictionary.com/)

**-u/-ur/-urban** *query*

	-urban 4chan
	Urban Dictionary | 4chan | http://mnn.im/upucr | you have just entered the very heart, soul, and life force of the internet. this is a place beyond sanity, wild and untamed. there is nothing new here. "new" content on 4chan is not found; it is created from old material. every interesting, offensive, shoc…


Gets the *n*th definition for *query* (only works for definitions 1-7)

**-u/-ur/-urban** *n* *query*

	-urban 3 4chan
	UrbanDictionary | 4chan | 4chan.org is the absolute hell hole of the internet, but still amusing. Entering this website requires you leave your humanity behind before entering. WARNING: You will see things on /b/ that you wish you had never seen in your life.


### URL Parser
Returns the title of a page and the host for html URLs.
Returns the type, size, and (sometimes) filename of a file URL.

	https://news.ycombinator.com/
	Title | Hacker News | news.ycombinator.com

	https://41.media.tumblr.com/bca28cbcbba3718cd67fd20062df19b9/tumblr_nl8gekhnLU1tdhimpo1_1280.png
	File | image/png 272kB | 41.media.tumblr.com


### User Profiles
Returns the set variables for a *user*

	-whois qb
	qb | Twitter: @abnormcore | URL: https://dribbble.com/qb
	
Variables are set using **-set url** *url* or **-set twitter** *handle*

	-set twitter someone
	twitter updated to: someone

	-set url http://www.something.com/
	url updated to: http://www.something.com/


### Weather
[Yahoo Weather](http://weather.yahoo.com/) for *location*
**-w/-we/-weather** *location*

	-weather Washington, DC
	Weather | Washington | Cloudy 15°C. Wind chill: 15°C. Humidity: 72%

[Yahoo Weather Forecast](http://weather.yahoo.com/) for *location*
**-f/-fo/-forecast** *location*

	-forecast Washington, DC
	Forecast | Washington | Sun: Clouds Early/Clearing Late 16°C/10°C | Mon: Mostly Sunny 19°C/8°C | Tue: Mostly Sunny 23°C/11°C | Wed: Partly Cloudy 24°C/11°C
	
Associates your current irc nick with *location*.
Other weather functions will default to this location if none is provided.

**-set location** *location* 

	<joebloggs> -set location Washington, DC
	<Porygon> joebloggs: location updated to: Washington, DC


### WolframAlpha
Finds the answer of *question* using [WolfarmAlpha](http://www.wolframalpha.com/)

**-wa** *question*

	-wa time in Bosnia
	Wolfram | current time in Bosnia and Herzegovina >>> 12:55:38 pm CEST | Tuesday, October 6, 2015


### Youtube
Gets the first result from [Youtube](https://www.youtube.com) for *search query* 

**-yt/-youtube** *search query*

	-yt Richard Stallman interject
	YouTube | I'd just like to interject... | 3m1s | https://youtu.be/QlD9UBTcSW4

***

### Admin functions
These functions are limited to bot admins and can only be used in a private message.
	
Ignore a user

**-set ignore** *nick*

	-set ignore Cbbleh
	<Porygon2> I never liked him anyway
	
Unignore a user

**-set unignore** *nick*

	-set unignore Cbbleh
	<Porygon2> Sorry about that
	
Toggles the URL parser for the channel

**-set urls on/off** *channel*

	-set urls on #lobby
	<Porygon2> Now reacting to URLs in #lobby
	
Toggles the file URL parser for the channel

**-set files on/off** *channel*

	-set files on #lobby
	<Porygon2> No longer displaying file info in #lobby	
Joins a channel and adds it to auto join

**-join** *channel*

	-join #foobar
	* Porygon has joined #foobar
	
Parts a channel and removes it from auto join

**-part** *channel*

	-part #foobar
	* Porygon has left the channel

