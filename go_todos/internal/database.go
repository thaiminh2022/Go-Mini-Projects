package internal

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Database struct {
	Todos []TodoData `json:"todo"`
}

type TodoData struct {
	Complete bool   `json:"complete"`
	Value    string `json:"value"`
}

func (p *Database) AddTodo(todo TodoData) {
	p.Todos = append(p.Todos, todo)
}
func (p *Database) DeleteTodoAt(index int) {
	p.Todos = append(p.Todos[:index], p.Todos[index+1:]...)
}

func (p *Database) ClearTodo() {
	p.Todos = []TodoData{}
}
func (p *Database) ToggleTodoComplete(index int) {
	theTodo := p.Todos[index]
	theTodo.Complete = !theTodo.Complete
	p.Todos[index] = theTodo
}

func get_executable_dir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
func (p *Database) LoadDatabase() error {
	var err error
	data, err := os.ReadFile(get_executable_dir() + "/database.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	return nil

}
func (p *Database) NewDatabase() error {
	p.Todos = []TodoData{}
	dbJson, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = os.WriteFile(get_executable_dir()+"/database.json", dbJson, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (p *Database) SaveDatabase() error {
	dbJson, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = os.WriteFile(get_executable_dir()+"/database.json", dbJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
