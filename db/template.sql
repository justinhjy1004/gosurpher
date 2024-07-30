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
((SELECT max(Id) FROM Blog) + 1, CURRENT_TIMESTAMP, "The Joy of Good Work", "Justin Ho", 
		"There is an obsession of the consumer society, and no, this is not a rant about how we have transformed our way of life into a cycle of constant hedonism and endless consumption, though that warrants its own set of criticisms, what I truly want to discuss is the unfortunate neglect of something I believe to be inherent in human life, that we feel some sense of satisfaction with how we choose to conduct our life, to provide, to give, to contribute. It is an essence in us that we are so willing to forget given the optimizing goal of consuming more, which means producing more, though not in a better way. As a result we have created a society that aims to squeeze the very last bit of human effort in the name of producing more, for the sake of consumption.
		As of writing, I left one of the most painful experiences in life. A huge part of it lies in my warped perception of the value of work (or rather the lack thereof) in the situation I was in, and I didn't believe in a bit the things I produced. Harvard Business School has a knack for getting professors of the lowest quality perhaps, though as Aiqi said, it could be a selection bias issue. Afterall, who would truly go to a business school academia unless you truly do not care deep down about discovering life's most interesting questions (Sorry Julian, you are honestly one of the good ones). But I dare say that though the publication records of these professors might be great on paper, the papers they produce are abhorrently bad and sad. I shall not generalize to Harvard as an institution as a whole, since I had read many great works by Harvard researchers, though it does paint a very somber picture to me the state of research in the world.
		Of course, every where has bullshit problems. But reflecting on that experience really shaped how I feel about work, for better or worse. I no longer hold research as a sacred work that the most intelligent of us partake in. No. It is after all a form of business. However, I do recognize the good that can come from good research. The ones that are worthwhile. The ones that I know is going to matter to people living in this world, and not the constant stream of 'AI this' or 'AI that'. The farcical nature of research that one can find here makes me think that we should have a collective effort to humble the worst of the offenders. 
		Yet they sit on the ivory tower. I wish I was not as intelligent as they are, and I wish that I have learned something that matters here in this place, but I cannot say that unfortunately. What I can say is that it opened up my mind to what is worthwhile doing, work that matters. Work that changes lives for the better, and that truly comes from the passion to obtain knowledge and to know. I don't want to be a manager of sorts, getting RAs in line to do all of my painful work. I want to do that work, the pain, the headaches and all. And most importantly, research that matters to the world, that I can look back proudly and say, I did my research on that.",
		"generic/personal/joy-of-good-work", "Personal");

