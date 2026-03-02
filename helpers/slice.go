package helpers

// Contains нь slice дотор элемент байгаа эсэхийг шалгана.
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Map нь slice-ийн элемент бүр дээр function ажиллуулж шинэ slice буцаана.
func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter нь нөхцөл хангасан элементүүдийг шүүнэ.
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce нь slice-ийг нэг утга руу нэгтгэнэ.
func Reduce[T any, R any](slice []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Chunk нь slice-ийг заасан хэмжээтэй хэсгүүдэд хуваана.
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Unique нь slice-аас давхардлыг арилгана.
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// PaginateSlice нь slice-ийг хуудаслана.
func PaginateSlice[T any](slice []T, page, pageSize int) []T {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	start := (page - 1) * pageSize
	if start >= len(slice) {
		return []T{}
	}

	end := start + pageSize
	if end > len(slice) {
		end = len(slice)
	}

	return slice[start:end]
}
