package utils

import "errors"

var UserNotFound = errors.New("usuário não encontrado")
var EmailAlreadyExists = errors.New("e-mail já cadastrado")
var UsernameAlreadyExists = errors.New("nome de usuário já cadastrado")
