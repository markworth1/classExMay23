package repo

//Repo is an internal memory bank
type Repo struct {
	data []string
}

func (r *Repo) Insert(s string) {
	r.data = append(r.data, s)
}

func (r *Repo) GetAllData() []string {
	return r.data
}