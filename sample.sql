insert into gorello.projects (name, description) values ('gorello', 'trello in go'), ('doulike', 'do u like?');

insert into gorello.columns (name, project_id) values ('TO_DO', 1), ('DONE', 1), ('default', 2);

insert into gorello.tasks (name, description, priority, column_id) values
('finish me', 'finish gorello', 1, 3),
('NOT important task', 'trash', 2, 3),
('most important task', 'first priority', 3, 3);

insert into gorello.comments (text, task_id) values ('gambare!', 4);

testProject := NewProject("golang", "")
	testProject.AddColumn("TO_DO").AddTask("gorello", "").AddComment("gambare!")
	testProject.Columns[1].AddTask("NOT important task", "")
	testProject.AddColumn("Done")
	importantTask := testProject.Columns[1].AddTask("very important task", "")
	// printObj(testProject)
	importantTask.ChangePosition(0)
	// printObj(testProject)	
