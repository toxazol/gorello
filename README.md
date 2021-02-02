## To Run locally:

```
docker-compose up
```


## User is able to manage (create, read, update, delete) Projects:

- Query `/project` with method POST to CreateProject
    + json payload should contain name and description fields. id will be auto generated
- Query `/projects` with method GET to GetProjects
- Query `/project` with method POST to UpdateProject
    + json payload should contain project object with new values in fields that are to be updated. id should stay the same
- Project object example:
    ```json
    {
        "id": 1,
        "name": "docker",
        "description": "manual test empty db"
    }
    ```
- Query `/project/:project_id` with method DELETE to DeleteProject

- Projects are listed by name
- Project contains at least one column: 
    + the first column created by default when a Project created (default name: New Column)


## User is able to manage (create, read, update, delete) Columns:

- Query `/column` with method POST to CreateColumn
    + json payload should contain name and project_id
- Query `/columns/:project_id` with method GET to GetColumns
- Query `/column/:column_id` with method GET to GetColumn
- Query `/column` with method POST to UpdateColumn (User is able to rename a Column)
    + name and project_id can be updated
- Query `/move/column/:column_id/?new_pos=int` with method PUT to MoveColumn to new position (User is able to move a Column left or right)
    + position is a positive integer starting from 0 (uppermost)
- Column object example:
    ```json
    {
        "id": 1,
        "name": "New Column",
        "project_id": 1
    }
    ```
- Query `/column/:column_id` with method DELETE to DeleteColumn

- Columns are listed by their position specified by User
- Column name is unique
- When a Column is deleted its tasks are moved to the Column to the left of the current
- The last column cannot be deleted

## User is able to manage (create, read, update, delete) Tasks:
- Query `/task` with method POST to CreateTask
    + json payload should contain name and description and column_id
- Query `/tasks/:column_id` with method GET to GetTasks
- Query `/task/:task_id` with method GET to GetTask
- Query `/task` with method POST to UpdateTask
    + name, description (User can update the name and the description of the Task) and column_id (User is able to move a Task across the Columns to change its status) can be updated
- Query `/move/task/:task_id` with method PUT to MoveTask to new position (User is able to move a Task within the Column (up and down) to prioritize it)
    + position is a positive integer starting from 0 (uppermost)
- Task object example:
    ```json
    {
        "id": 11,
        "name": "new task",
        "description": "created in postman",
        "column_id": 18
    }
    ```

- Query `/task/:task_id` with method DELETE to DeleteTask with all Comments related to this Task
- Task can be created only within a Column
- User can view Tasks in all Columns of a Project


## User is able to manage (create, read, update, delete) Comments:

- Query `/comment` with method POST to CreateComment
    + json payload should contain text and task_id

- Query `/comments/:task_id` with method GET to GetComments (User can view Comments related to a Task)
- Query `/comment/:comment_id` with method GET to GetComment

- Query `/comment` with method POST to UpdateComment
    + User can update the Comment text
- Comment object example:
    ```json
    {
        "id": 3,
        "text": "new",
        "task_id": 6,
        "time_stamp": "2021-02-02 03:04:31"
    }
    ```

- Query `/comment/:comment_id` with method DELETE to DeleteComment
- Comment can be created only within a Task
- Comments in a list are sorted by their creation date (from newest to oldest)

