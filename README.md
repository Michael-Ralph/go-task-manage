# go-task-manage
A to-do list cli program to practice golang and how to use and make an api

Run application:

```
go run main.go
```

Create task(POST):

```
curl -X POST -H "Content-Type: application/json" -d '{"title": "Buy groceries", "description": "Milk, eggs, bread"}' http://localhost:9000/tasks
```

Get all tasks(GET):

```
curl http://localhost:9000/tasks
```
Get task by ID(GET):

```
curl http://localhost:9000/tasks/YOUR_TASK_ID
```

Update task(PUT):

```
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Buy all groceries", "description": "Milk, eggs, bread, cheese", "completed": true}' http://localhost:9000/tasks/YOUR_TASK_ID
```

Delete task(DELETE):

```
curl -X DELETE http://localhost:9000/tasks/YOUR_TASK_ID
```