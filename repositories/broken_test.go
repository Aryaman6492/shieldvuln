package repositories

import (
	"context"
	"testing"

	"github.com/Aryaman6492/shieldvuln/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestBrokenStore_GetCVE(t *testing.T) {
	b := NewBrokenStorage()
	_, err := b.GetCVE(context.TODO(), "", "", "", "")
	assert.Error(t, err)
}

func TestBrokenStore_GetSBOM(t *testing.T) {
	b := NewBrokenStorage()
	_, err := b.GetSBOM(context.TODO(), "", "")
	assert.Error(t, err)
}

func TestBrokenStore_StoreCVE(t *testing.T) {
	b := NewBrokenStorage()
	err := b.StoreCVE(context.TODO(), domain.CVEManifest{}, false)
	assert.Error(t, err)
}

func TestBrokenStore_StoreSBOM(t *testing.T) {
	b := NewBrokenStorage()
	err := b.StoreSBOM(context.TODO(), domain.SBOM{}, false)
	assert.Error(t, err)
}
