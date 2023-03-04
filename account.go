package transferkit

type AccountCheck struct {
	DeliveryMethod DeliveryMethod `json:"DeliveryMethod"`
}

type AccountFetch struct {
	DeliveryMethod DeliveryMethod `json:"DeliveryMethod"`
}
