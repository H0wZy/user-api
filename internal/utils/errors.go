package utils

import "errors"

var UserNotFound = errors.New("usuário não encontrado")
var EmailAlreadyExists = errors.New("e-mail já cadastrado")
var UsernameAlreadyExists = errors.New("nome de usuário já cadastrado")
var InvalidGender = errors.New("gênero inválido")
var EmptyGender = errors.New("gender não pode estar vazio")
var EmptyGenderString = errors.New("gênero não pode ser um texto vazio")
