package exprutils

import (
	"reflect"
	"sync"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"

	"exusiai.dev/roguestats-backend/internal/model"
)

type ExprRunner struct {
	methodEnv map[string]interface{}
	v         vm.VM
}

var exprRunnerInstance *ExprRunner
var exprRunnerOnce sync.Once

func GetExprRunner() *ExprRunner {
	exprRunnerOnce.Do(func() {
		methods := getMethods(reflect.TypeOf(ExprFunction{}))
		methodEnv := make(map[string]interface{})
		exprFunction := ExprFunction{}
		for _, method := range methods {
			methodEnv[method] = reflect.ValueOf(exprFunction).MethodByName(method).Interface()
		}
		exprRunnerInstance = &ExprRunner{
			methodEnv: methodEnv,
			v:         vm.VM{},
		}
	})
	return exprRunnerInstance
}

func (e ExprRunner) PrepareEnv(event *model.Event) map[string]interface{} {
	env := map[string]interface{}{
		"content": event.Content,
	}
	for k, v := range e.methodEnv {
		env[k] = v
	}
	return env
}

func (e ExprRunner) RunCode(code string, env map[string]interface{}) (interface{}, error) {
	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		return nil, err
	}
	output, err := e.v.Run(program, env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func getMethods(t reflect.Type) []string {
	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.PkgPath == "" {
			methods = append(methods, method.Name)
		}
	}
	return methods
}
