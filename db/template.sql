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
((SELECT max(Id) FROM Blog) + 1, CURRENT_TIMESTAMP, "Agency and our Choice to Live", "Justin Ho", 
	"Not too long ago, I watched a YouTube video of this neurosurgeon who decided to quit shortly after his residency, to assume a tonne of debt that he could never afford to pay without his high-paying job, and currently live a contented life. One that perhaps was way better than his life, performing precise operations that do not truly benefit anyone in the long run, except to line the pockets of the medical industry. This broke him. It was not drama. It was not personal vendetta. It was just cold-hearted cash grabs that destroyed his soul.
	Jonathan sent me that video soon after I left Jackie's team. However, a part of me still feels stuck. I do know that quitting does not equal to a resolution, that I am aware of. What I am saddened is the partial clarity of the pointlessness of my work. Rebecca (my counselor) summarized it up so well, noting that this is a crisis, an existential one. I exercised my agency to walk away once. I think I can do it again. Though I recognize the importance of this transition of my life, I recognize that I cannot do my PhD with good conscience in an institution like Harvard Business School. 
	That is in my control. All of that concern is in my control. To pursue my PhD in an institution, and do research in a field that I like, that is in my control. Of course, having a fulfilling worklife is indeed a privilege, and I recognize that. But since this is a choice thrusted upon me, I do have to ask, should I take the route of vestige of prestige? Probably not. My agency is truly tested when I am given this difficult choice, a difficult choice to choose between the 'natural' path or otherwise. All of the ideas of success instilled in us since childhood appear so fickle now in my head, and maybe what Brent said is true. It is indeed liberating to come to that realization.
	As for working with Julian, time would tell. The only thing that I really have to answer to is my love for Nathan. I am of course not putting this responsibility on you Nathan, I don't. Noting that this is just a transitory period is what I need to keep in mind. Noting that my decisions are shocking to folks that think it is crazy to walk away makes me happy, just a little, in the inside. Noting that life is totally out of control, but it is fine to admit that, and only assume the parts that are within your agency simultaneously makes you feel small and weak, but you are on firm ground, Justin. 
	Good luck!", 
	"generic/personal/agency-and-our-choice-to-live", "Personal");

