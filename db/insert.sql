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
"The Gamma Distribution", 
"Justin Ho", 
"The Gamma distribution has always fascinated me, mostly because it is a less commonly used Greek alphabet, and the PDF somehow has a Gamma function nested within it, which makes it terrifying. This is true with the Beta as well, but it just lacks that mysterious vibe. It is also extremely important since it is related a number of important distributions, but I'd like to focus on its connections to the Exponential and the Poisson.",
"gamma",
"Stories of Distribution");

