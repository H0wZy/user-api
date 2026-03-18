package types

type Gender string

const (
	Masculino Gender = "Masculino"
	Feminino  Gender = "Feminino"
)

func GetValidGender() []Gender {
	return []Gender{Masculino, Feminino}
}
