//полтупление потока задач, решение этих затач независимо от види и отправка результатов

package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type Data struct {
	Name    string
	Email   string
	Message string
}

func (d *Data) Process() {
	fmt.Printf("Сообщение для %v, от %v\n", d.Message, d.Name)
	time.Sleep(10 * time.Millisecond)
}

type Image struct {
	Id    int
	Image string
}

func (im *Image) Process() {
	fmt.Printf("Task for %v in completed, ID %v\n", im.Image, im.Id)
	time.Sleep(100 * time.Millisecond)
}

type Video struct {
	Id    int
	Video string
}

func (v *Video) Process() {
	fmt.Printf("Vidio for %v in completed, ID %v\n", v.Video, v.Id)
	time.Sleep(1 * time.Second)
}

type TaskEngine struct {
	Tasks      []Task
	Concurence int
	TaskChan   chan Task
	wg         sync.WaitGroup
}

func (te *TaskEngine) CompletedTask() {
	for task := range te.TaskChan {
		task.Process()
		te.wg.Done()
	}
}

func (te *TaskEngine) RunTasksGoroutine() {
	te.TaskChan = make(chan Task, len(te.Tasks))

	te.wg.Add(len(te.Tasks))
	for i := 0; i < te.Concurence; i++ {
		go te.CompletedTask()
	}

	for _, t := range te.Tasks {
		te.TaskChan <- t
	}

	close(te.TaskChan)
	te.wg.Wait()

}

func main() {

	tasks := []Task{
		&Data{"RT", "@mail.ru", "Hello World"},
		&Data{"ART", "@mail.ru", "Hello Two World"},
		&Image{1, "Image1"},
		&Image{2, "Image2"},
		&Image{3, "Image3"},
		&Video{1, "Video1"},
		&Video{2, "Video2"},
		&Video{3, "Video3"},
		&Video{4, "Video4"},
	}

	taslEng := TaskEngine{
		Tasks:      tasks,
		Concurence: 3,
	}

	taslEng.RunTasksGoroutine()
}
