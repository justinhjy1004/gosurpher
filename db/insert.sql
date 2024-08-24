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
"Denying Delusions",
"Justin Ho", 
"The occasional clarity that I have as I see sunlight peer through the leaves. I can see them glistening and momentarily waving at me as I walk towards them. I don't have the luxury of enjoying such clear vision, especially at night. To correct myself, I don't have such a privilege whilst maintaining my desire to not wear my thick corrective lenses. I have no problems with glasses, just not the weary ones I have.
Sitting outdoors at the SEAS building, I could not help but wonder if engineers really care about the greeneries that surround them. I don't see them often outside and that makes sense. I don't see many people in the summer in that respect. What a pity, I think, where many in lower Allston could enjoy such tranquility that Harvard could afford, yet to them without an access card it seems like an impenetrable forest. A forest that even I, a cardholder, sometimes hesitate for the sole reason that I do not belong. The irony of it all is that Harvard has an open campus, the degree of which is obviously debatable.
Momentary clarity helps me as I have come to accept the many beautiful delusions I enjoy having. They are fundamentally delusions, yet I have to cling onto them. In darkness, they creep and they wrap me around with the warm touch of what I think this world is about. They comfort me until they don't. But they are parts of me that give me hope and despair, if not an illusory imagination that I am contorted with.
Maybe I should not see this world that clearly and thank God for my astigmatism. It helps to just have this distorted lens that tells you that the world is conspiring for you and humbling you that not all you see is real.",
"generic/personal/denying-delusions",
"personal");

