package service

func IsValidCity(cidy string) bool {
	validCities := map[string]bool{
		"Москва":          true,
		"Санкт-Петербург": true,
		"Казань":          true,
	}
	_, ok := validCities[cidy]
	return ok
}
