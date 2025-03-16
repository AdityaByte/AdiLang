package interpreter

import "fmt"

type Environment struct {
	variables map[string]interface{}
	parent *Environment
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		variables: make(map[string]interface{}),
		parent: parent,
	}
}

func (e *Environment) Set(name string, value interface{}) {
	e.variables[name] = value
}

func (e *Environment) Get(name string) (interface{}, error) {
	// The exists is the lookup function which was managed by the map of the environment so
	// in this when the variable exists in the current scope it return the true value
	value, exists := e.variables[name]

	// Here we have added the thing that if the variable does not exists in the current scope then it will 
	// check for the parent scope if the parent scope is not nil and the variable exists in that scope then it will return that variable value
	if !exists && e.parent != nil {
		return e.parent.Get(name)
	}

	if !exists {
		fmt.Printf("Variable is not defined in the scope %s", name)
		return nil, fmt.Errorf("undefined variable: %s",name)
	}

	return value, nil
}