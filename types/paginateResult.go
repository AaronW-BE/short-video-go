package types

type PaginateResult struct {
	List  []interface{} `json:"list"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Total int           `json:"total"`
}

func (p *PaginateResult) ToJson() {

}
