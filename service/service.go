package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAT   time.Time
	CompletedAt time.Time
}

type List []Item

func HandlePanic() {
	if a := recover(); a != nil {
		fmt.Printf("recovered from the panic :%v\n", a)
	}
}

func (l *List) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAT:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, todo)
}

func (l *List) Complete(i int) error {
	list := *l

	if i <= 0 || i > len(list) {
		return fmt.Errorf("item does not exist %v", i)
	}

	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Save(fileName string) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, jsonData, 0644)
}

func (l *List) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
