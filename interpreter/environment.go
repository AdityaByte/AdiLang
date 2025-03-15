package interpreter

import "fmt"

type Environment struct {
	variables map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		variables: make(map[string]interface{}),
	}
}

func (e *Environment) Set(name string, value interface{}) {
	e.variables[name] = value
}

func (e *Environment) Get(name string) (interface{}, error) {
	value, exists := e.variables[name]
	if !exists {
		return nil, fmt.Errorf("undefined variable: %s",name)
	}
	return value, nil
}