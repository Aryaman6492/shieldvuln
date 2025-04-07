package adapters

import (
	"context"

	"github.com/Aryaman6492/shieldvuln/core/ports"
)

type MockRelevancyAdapter struct {
}

var _ ports.Relevancy = (*MockRelevancyAdapter)(nil)

func NewMockRelevancyAdapter() *MockRelevancyAdapter {
	return &MockRelevancyAdapter{}
}

func (m MockRelevancyAdapter) GetContainerRelevancyScans(_ context.Context, _, _ string) ([]ports.ContainerRelevancyScan, error) {
	return []ports.ContainerRelevancyScan{}, nil
}
