package tools


map[string]float64

type Word struct {
	text   string
	cipher string
	score int
}

type WordSorter []Word

func (c WordSorter) Len() int           { return len(c) }
func (c WordSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c WordSorter) Less(i, j int) bool { return c[i].score < c[j].score }


// CheckFrequency checks frequency of etaoin shrdlu.
// The higher the counter, the most like the phrase
// is an English phrase. Very unsophisticated
func CheckFrequency(data string) int {
	mostFrequent := "etaoin shrdlu"
	var counter int
	data = strings.ToLower(data)
	for _, l := range data {
	Loop:
		for _, f := range mostFrequent {
			if l == f {
				counter++
				break Loop
			}
		}
	}

	return counter
}
