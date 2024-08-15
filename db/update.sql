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
SET Route = "generic/personal/my-thoughts-on-ai", Type = "Personal"
WHERE Type = "2 Cents";