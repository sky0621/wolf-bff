package domain

type Title string

func (t *Title) Validate() bool {
	if t == nil {
		return false
	}
	if *t == "" {

	}
}
