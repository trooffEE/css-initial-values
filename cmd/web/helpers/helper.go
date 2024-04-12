package helpers

// TODO Вынести в отдельный пакет
func Join(separator string, path ...string) string {
	result := ""
	for i := range path {
		result += path[i]
		if i != len(path)-1 {
			result += separator
		}
	}
	return result
}
