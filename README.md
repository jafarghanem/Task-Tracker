
# 🧾 Task Tracker CLI (Golang)

📁 GitHub Repository: [https://github.com/jafarghanem/Task-Tracker](https://github.com/jafarghanem/Task-Tracker)

This is a Task Tracker CLI Project from the roadmap.sh website : https://roadmap.sh/projects/task-tracker
This is a simple, interactive **Command-Line Interface (CLI)** application written in Go for managing tasks locally using a JSON file. It allows users to add, update, delete, and track the progress of tasks in a structured and user-friendly way.

---

## 🚀 Features

- ✅ Add tasks with descriptions
- 🔄 Update task descriptions
- ❌ Delete tasks
- 📋 List all tasks or filter by status (`todo`, `in-progress`, `done`)
- 🔁 Mark tasks as **in-progress**
- ✅ Mark tasks as **done**
- 🗂️ All data is persisted in a `task.json` file

---

## 📦 Requirements

- [Go 1.18+](https://golang.org/dl/) installed

---

## 🛠 Usage

### 👇 Run the application:

```bash
go run main.go
```

You will enter an interactive CLI prompt:

```bash
This is a simple task-tracker.
You can add tasks, track their status, and mark them as done.
--------------------------------------------------
```

### 📥 Available Commands

| Command                           | Description                                     |
|----------------------------------|-------------------------------------------------|
| `add "description"`              | Add a new task with a description               |
| `update <id> "new description"`  | Update an existing task’s description           |
| `delete <id>`                    | Delete the task with the given ID               |
| `list`                           | List all tasks                                  |
| `list todo`                      | List only tasks with `todo` status              |
| `list in-progress`              | List only tasks marked as `in-progress`         |
| `list done`                      | List only completed tasks                       |
| `mark-in-progress <id>`         | Mark the task as `in-progress`                  |
| `mark-done <id>`                | Mark the task as `done`                         |
| `exit`                           | Exit the application                            |

---

## 🧠 Example

```bash
> add "Buy groceries"
Task added successfully (ID: 1)

> list
ID: 1, Description: Buy groceries, Status: todo, Created At: ..., Updated At: ...

> mark-in-progress 1
Task marked as in-progress.

> update 1 "Buy groceries and cook"
Task with ID 1 updated successfully.

> mark-done 1
Task marked as done.

> list done
ID: 1, Description: Buy groceries and cook, Status: done, ...
```

---

## 💾 Data Storage

All tasks are saved in a local file called `task.json` in the same directory. Each task includes:

- `ID`: Auto-incremented unique ID
- `Description`: Task text
- `Status`: `todo`, `in-progress`, or `done`
- `CreatedAt`: Timestamp when the task was added
- `UpdatedAt`: Timestamp of the last update

---

## 📂 Project Structure

```
.
├── main.go           # Entry point and CLI loop
├── task.json         # (Created automatically) stores tasks persistently
└── README.md         # You’re reading it!
```

---

## 🤝 Contributing

This is a minimal tool built for learning Go CLI design and file persistence. Feel free to fork, open issues, or suggest improvements!

---

## 📃 License

This project is licensed under the MIT License.
