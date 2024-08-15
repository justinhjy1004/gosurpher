/* The schema of Blog table
CREATE TABLE Blog (
Id int primary key,
Date datetime,
Title text,
Autor text, 
Content text, 
Route text,
Type text); */

-- Note: For links, use the format (text to be rendered, https://----)
-- Template for creating a new Blog entry

INSERT INTO Blog (Id, Date, Title, Autor, Content, Route, Type) VALUES
((SELECT max(Id) FROM Blog) + 1, 
CURRENT_TIMESTAMP, 
"The Only Hacking I am Doing is P-Hacking", 
"Justin Ho", 
"There often comes a time when you realize that the work that you are doing is bullshit. Mine came when my dear professor asked me to 'Take a deeper look into it!' when her main experiment's hypothesis never quite panned out. Of course, nobody knew beforehand what would have happened in the experiment. I would never profess to have such an ability to predict the unknown. All I heard from that command is of course to p-hack the shit out of the dataset to find a gem, whatever it is. 
For any reader that is unaware of what p-hacking is, it is basically the act of sifting through the data, slicing, cleaning, manipulating it, and conduct various statistical tests in hopes of a few of them being statistically significant. A primer can be found here from the (Embassy of Good Science,https://embassy.science/wiki/Theme:6b584d4e-2c9d-4e27-b370-5fbdb983ab46). You might say, isn't that normal to do so? And if statistics is worth anything, it is a tool to prove your hypotheses, right? You my friend, should consider being a business school professor, as you are teeming with the right kind of optimism in such a place. Andrew Gelman wrote a great paper talking about p-hacking in social science in this (article, http://www.stat.columbia.edu/~gelman/research/unpublished/p_hacking.pdf) for those who are interested in the misleading nature of research driven by p-hacking.
I have just gotten interested in cybersecurity. Learning about lower level programming which I used to abhor. Yet now the only hacking that I get to do is p-hacking. I would not call such a practice fraud. It is not. I think many well meaning people find themselves rethinking their hypotheses from initial studies. But coming from professors who should know better, it is deeply troubling, it is deeply unsettling. The (Center of Open Science, https://www.cos.io/initiatives/prereg) as well as the (Wharton Credibility Lab, https://aspredicted.org/) set up ways to combat such practices called (pre-registrations, https://en.wikipedia.org/wiki/Preregistration_(science)). Yet I see multiple professors constantly flouting such rules in the 'top' research institutions in the United States. A few examples of such practices are not publishing pre-registrations that 'failed' to pan out, pre-registring bullshit with very vague hypotheses and goals to only later claim that they have followed the pre-registered analyses, and some even just straight up send folks to dead links. The Data Colada team does provide an article describing how to (pre-register your studies properly, https://datacolada.org/64). 
In any sense, I wish that this practice will wither and die out. We need to really change the culture of research. I am honestly coming from a place of selfishness. I see how these practices crushed my soul (and still crushing it as an RA), but I don't yet see a way out. If anything, things might get worse before it gets better",
"generic/statistically-speaking/only-p-hacking",
 "Statistically Speaking");

