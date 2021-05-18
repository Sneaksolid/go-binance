package binance

import (
	"context"
	"encoding/json"
)

type TradeFeeService struct {
	c *Client
}

func (s TradeFeeService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeFee, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/asset/tradeFee",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*TradeFee, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type TradeFee struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
