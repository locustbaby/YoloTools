package main

import (
	"fmt"
	"os"
	"os/exec"
)

type WorkflowStep struct {
	Name     string
	Command  string
	Args     []string
	Required bool
}

type Workflow struct {
	Steps []WorkflowStep
}

func (w *Workflow) AddStep(step WorkflowStep) {
	w.Steps = append(w.Steps, step)
}

func (w *Workflow) Execute() error {
	for _, step := range w.Steps {
		fmt.Printf("Executing step: %s\n", step.Name)

		cmd := exec.Command(step.Command, step.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Step execution failed: %s\n", step.Name)
			if step.Required {
				return err
			}
		}
	}

	return nil
}

func main() {
	// 定义工作流
	workflow := Workflow{}

	// 添加工作流步骤
	workflow.AddStep(WorkflowStep{
		Name:     "Step 1",
		Command:  "echo",
		Args:     []string{"Hello", "World!"},
		Required: true,
	})

	workflow.AddStep(WorkflowStep{
		Name:     "Step 2",
		Command:  "ls",
		Args:     []string{"-l"},
		Required: true,
	})

	workflow.AddStep(WorkflowStep{
		Name:     "Step 3",
		Command:  "not-existing-command",
		Args:     []string{},
		Required: false,
	})

	// 执行工作流
	err := workflow.Execute()
	if err != nil {
		fmt.Printf("Workflow execution failed: %v\n", err)
	} else {
		fmt.Println("Workflow executed successfully.")
	}
}
