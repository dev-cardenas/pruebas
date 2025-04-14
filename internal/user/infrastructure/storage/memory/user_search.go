package memory_user

import (
	domain "pruebas/internal/user/domain"
	"strings"
)

// Search busca usuarios por nombre y apellido en el repositorio en memoria.
func (r *InMemoryUserRepository) Search(name, lastname string, offset, limit int) ([]*domain.User, error) {
	var result []*domain.User

	// Recorremos los usuarios en memoria y los filtramos por nombre y apellido.
	for _, u := range r.users {
		// Compara si el nombre y apellido contienen los términos de búsqueda (no distingue mayúsculas/minúsculas).
		if strings.Contains(strings.ToLower(u.Name), strings.ToLower(name)) &&
			strings.Contains(strings.ToLower(u.LastName), strings.ToLower(lastname)) {
			result = append(result, u)
		}
	}

	// Aplicar paginación
	if offset < len(result) {
		end := offset + limit
		if end > len(result) {
			end = len(result)
		}
		result = result[offset:end]
	} else {
		result = nil // Si el offset es mayor que la cantidad de usuarios, se devuelve nil.
	}

	return result, nil
}
