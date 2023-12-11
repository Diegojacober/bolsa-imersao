package entity

type Asset struct {
	ID			string
	Name		string
	MarketVolme	int
}

func NewAsset(id string, name string, marketVolme int) *Asset {
	return &Asset{
		ID: 			id,
		Name:			name,
		MarketVolme:	marketVolme,
	}
}