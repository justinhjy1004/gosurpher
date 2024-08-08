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
((SELECT max(Id) FROM Blog) + 1, CURRENT_TIMESTAMP, "My Thoughts on AI (Because Of Course!)", "Justin Ho", 
	"I am probably too late for the AI dunking party since the stock market is treating tech companies pretty poorly as I am writing, but I do think that even at this moment, there are so many people who are still on the hype train. With a knack of being bad in predictions, I think that the AI hype will end soon enough (and God do I hope it does!) but I am just sitting here, listening to this random guy behind me talking about (prompt engineering, https://www.cartoonstock.com/cartoon?searchID=CS587847) and how cool it is. And it is! It is truly marvelous that us nerds have found a way to create something that is willing to hold a conversations with us! But I do not think that it is the miracle that some claim it to be.
	In all honesty, my negative views of AI probably came from work. Constantly being surrounded by AI hypers who do not know a single thing about AI, Machine Learning, or basic technology, it is difficult to not feel like I am drowning. I have been told 'Have you tried ChatGPT?' enough times that I might do what (Lucidity, https://ludic.mataroa.blog/blog/i-will-fucking-piledrive-you-if-you-mention-ai-again/) would do to you if you mention AI again. However, I am just not quite ready to be homicidal.
	I think fundamentally, AI would not yet be the amorphous 'AGI'. I do not think that we have reached a point where we have discovered how to algorithmically replicate human intelligence. In fact, my position is not that controversial. Meta's AI Chief Scientist (Yann LeCun, https://www.wired.com/story/artificial-intelligence-meta-yann-lecun-interview/) agrees with me on this point. Though I do think that we have created a very different and novel kind of intelligence with the recent advancements of LLMs. I do want, in all honesty and sincerity, that this hype will die down soon, and we can see AI for what it truly is. Appreciate its wonders but recognize its failures. Try not to be too much like an Asian parent and expect groundbreaking things from AI, and that is okay. And maybe with all of that, I can start enjoying AI again.",
	"generic/twocents/my-thoughts-on-ai", "2 Cents");

