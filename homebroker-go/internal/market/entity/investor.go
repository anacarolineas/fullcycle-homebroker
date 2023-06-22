package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition //ponteiro
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) GetAssetPosition(assetId string, qtdShares int) {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition == nil {
			i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetId, qtdShares))
		} else {
			assetPosition.Shares += qtdShares
		}
	}
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetId,
		Shares:  shares,
	}
}
