package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task struct （数据结构）
type Task struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}

// File 存储路径
const filePath = "tasks.json"

// 加载 tasks.json
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		// 文件不存在 → 返回空任务列表
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// 保存 tasks.json
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// 添加任务
func addTask(text string) {
	tasks, _ := loadTasks()
	newTask := Task{
		ID:   len(tasks) + 1,
		Text: text,
		Done: false,
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Println("Added:", text)
}

// 列出任务
func listTasks() {
	tasks, _ := loadTasks()
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Text)
	}
}

// 完成任务
func completeTask(idStr string) {
	id, _ := strconv.Atoi(idStr)

	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			saveTasks(tasks)
			fmt.Println("Completed:", tasks[i].Text)
			return
		}
	}

	fmt.Println("Task not found.")
}

// 主入口
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <add|list|done> [...args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <text>")
			return
		}
		addTask(os.Args[2])

	case "list":
		listTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <id>")
			return
		}
		completeTask(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}
