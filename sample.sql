insert into gorello.columns (name) values ('col1'), ('col2');

insert into gorello.projects (name, description) values ('gorello', 'trello in go'), ('doulike', 'do u like?');


testProject := NewProject("golang", "")
	testProject.AddColumn("TO_DO").AddTask("gorello", "").AddComment("gambare!")
	testProject.Columns[1].AddTask("NOT important task", "")
	testProject.AddColumn("Done")
	importantTask := testProject.Columns[1].AddTask("very important task", "")
	// printObj(testProject)
	importantTask.ChangePosition(0)
	// printObj(testProject)
