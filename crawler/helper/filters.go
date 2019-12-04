package helper

func StringSliceFilter(o []string, f func(str string)bool) []string{
	var filteredSlice []string
	for _,item := range o {
		if f(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}