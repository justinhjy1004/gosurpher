/* The schema of Blog table
CREATE TABLE Blog (
Id int primary key,
Date datetime,
Title text,
Autor text, 
Content text, 
Route text,
Type text); */

UPDATE Blog
SET Content = "Inspired by our favourite Statistics teacher Blitztein, I was inspired to write stories about distributions that relate to events in my life as a way to journal, and also remember the things that happened, and the underlying probability distributions that govern them. In this spirit, I begin with the Poisson Distribuion."
WHERE Route = "poisson";
