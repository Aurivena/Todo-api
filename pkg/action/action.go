package action

import "Todo/pkg/domain"

type Action struct {
	domains *domain.Domain
}

func NewAction(domain *domain.Domain) *Action {
	return &Action{domains: domain}
}
