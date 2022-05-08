package homework

import "sort"

type KeyValueArr []KeyValue

type KeyValue struct {
	key   int
	value string
}

func (a KeyValueArr) Len() int           { return len(a) }
func (a KeyValueArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyValueArr) Less(i, j int) bool { return a[i].key < a[j].key }

func sortMapValues(input map[int]string) (result []string) {
	var kvA KeyValueArr

	for k, v := range input {
		kvA = append(kvA, KeyValue{
			k,
			v,
		})
	}

	sort.Sort(kvA)

	for _, value := range kvA {
		result = append(result, value.value)
	}

	return result
}
