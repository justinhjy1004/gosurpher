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
((SELECT max(Id) FROM Blog) + 1, CURRENT_TIMESTAMP, "Title", "Autor", 
" Content here ",
"Route", "Type");

-- id INTEGER PRIMARY KEY AUTOINCREMENT 
