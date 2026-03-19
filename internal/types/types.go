package types

type Gender string

const (
	Male         Gender = "Masculino"
	Female       Gender = "Feminino"
	NotSpecified Gender = "Não especificado"
)

func GetValidGender() []Gender {
	return []Gender{Male, Female, NotSpecified}
}
