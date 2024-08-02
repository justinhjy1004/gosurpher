/* The schema of Blog table
CREATE TABLE Blog (
Id int primary key,
Date datetime,
Title text,
Autor text, 
Content text, 
Route text,
Type text); */


-- Template for creating a new Blog entry
INSERT INTO Blog (Id, Date, Title, Autor, Content, Route, Type) VALUES
((SELECT max(Id) FROM Blog) + 1, CURRENT_TIMESTAMP, "The only Hacking that I am Doing is p-hacking", "Justin Ho", 
	"Working in academia is so interesting. Perhaps a little disappointing, but hey, most people find joy in the tiny things in life, and I think I have found my partial calling. Note that I am in no way an academic, not in the true sense. I aspired to be one, when I was naively hopeful that academia meant the highest honour of discovering knowledge and understanding difficult questions that is worth answering. I do still currently have that dream of course. But a dream is just a dream, it is fleeting, and most probably not coming true.
	Not too long ago, I started becoming interested in cybersecurity. It has always been in the back of my mind, and truth to be told, I was simply not able to understand most concepts. With time and experience working in the technology space, my interests (and fears) in cybersecurity grew. Of course, I do not think that I am really that good in it. I could only vaguely understand when Jonathan was describing a 'libc attack' to me. I was confused. It apparently is a type of bufferoverflow attack, but more like a mutated version of it because some pesky cybersecurity experts found a good way to stop code injections by having unexecutable stacks. I nodded along of course, but it is a hot and humid summer, and I had enough of outdoors for the day.
	This blog has nothing to do with cybersecurity though. This is more of an observation (rant) of the many peculiarities of academia as an institution. I do not work in any respectable form of academia, ie. I work in the social sciences, and I like it. I am annoyed though by the absurd (yet understandable) obsession with the p-value. Most of my 'best work' in this reputable institution I work in (its initials literally has 'BS' in it) involved a lot of p-hacking. For those who are unaware of what p-hacking is, p-hacking is the act of data 'cleaning/manipulation' that 'explores/exploits' statistical methods to find significant effects when there is none. Since how we statistically find effects from data involves rejecting/accepting a null hypothesis",
	"generic/statisticallyspeaking/p-hacking", "Statistically Speaking");

