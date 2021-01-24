# noCloudStream
#### But thanks anyway :)


Hello there!  
You might wonder what this crappy code is all about; Well, allow me to explain :D  

## What is this?

At the time of writing, this is a program that connects to the Twitch chat and
(only) allows a streamer to highlight comments on the chat with a nice animation.
That's also a thing some nice cloud service does, but I don't like cloud stuff.
(Yes, I have problems, I am well aware of them.)
Also does not need an API key from Twitch (yet).

## Why it exists?

Because I liked the feature very much, I haven't found any open source program (or
plugin) to do it and don't like like cloud platforms. Also I have no idea (nor
willpower) to learn the OBS APIs (or C++) and make a plugin (if that's possible).

## And how does this look?

Very bad indeed. At least the "backend". And the code structure. And the way it
works. But I find the comment animation to look good enough to use it in
a science stream I direct. So yeah, good enough.  
Screen:
![]

## How does it work under the hood?

I'm so sorry you asked this. If you didn't, that's even better.
Feel free to skip the section. Or stop here. Or not, I'm a file, not a cop.
<content>
Literally a fork of the Gorilla's chat example (it totally fit my usecase)
</content>


## Todo (I'm probably never gonna do lol)
[ ] Queue messages on the OBS page side 
[ ] Allow multiple animations tipe
[ ] Listen for events from twitch (like sub/whatever happens on twitch) (I __totally__ follow twitch streams)