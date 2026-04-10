package types

type Gender string

const (
	Male         Gender = "Masculino"
	Female       Gender = "Feminino"
	NotSpecified Gender = "Não especificado"
)

func GetValidGenders() []Gender {
	return []Gender{Male, Female, NotSpecified}
}
