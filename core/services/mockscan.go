package services

import (
	"context"

	"github.com/Aryaman6492/shieldvuln/core/domain"
	"github.com/Aryaman6492/shieldvuln/core/ports"
)

type MockScanService struct {
	happy bool
}

var _ ports.ScanService = (*MockScanService)(nil)

func NewMockScanService(happy bool) *MockScanService {
	return &MockScanService{happy: happy}
}

func (m MockScanService) GenerateSBOM(context.Context) error {
	if m.happy {
		return nil
	}
	return domain.ErrMockError
}

func (m MockScanService) Ready(context.Context) bool {
	return m.happy
}

func (m MockScanService) ScanAP(context.Context) error {
	if m.happy {
		return nil
	}
	return domain.ErrMockError
}

func (m MockScanService) ScanCVE(context.Context) error {
	if m.happy {
		return nil
	}
	return domain.ErrMockError
}

func (m MockScanService) ScanRegistry(context.Context) error {
	if m.happy {
		return nil
	}
	return domain.ErrMockError
}

func (m MockScanService) ValidateGenerateSBOM(ctx context.Context, _ domain.ScanCommand) (context.Context, error) {
	if m.happy {
		return ctx, nil
	}
	return ctx, domain.ErrMockError
}

func (m MockScanService) ValidateScanAP(ctx context.Context, _ domain.ScanCommand) (context.Context, error) {
	if m.happy {
		return ctx, nil
	}
	return ctx, domain.ErrMockError
}

func (m MockScanService) ValidateScanCVE(ctx context.Context, _ domain.ScanCommand) (context.Context, error) {
	if m.happy {
		return ctx, nil
	}
	return ctx, domain.ErrMockError
}

func (m MockScanService) ValidateScanRegistry(ctx context.Context, _ domain.ScanCommand) (context.Context, error) {
	if m.happy {
		return ctx, nil
	}
	return ctx, domain.ErrMockError
}
