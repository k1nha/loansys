package handler

import "fmt"

func errorParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func (cRole *CreateRoleRequest) Validate() error {
	if cRole.Name == "" {
		return errorParamRequired("role", "string")
	}

	return nil
}
