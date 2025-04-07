package ports

import (
	"context"

	"github.com/Aryaman6492/shieldvuln/core/domain"
	"github.com/Aryaman6492/storage/pkg/apis/softwarecomposition/v1beta1"
)

type ApplicationProfileRepository interface {
	GetApplicationProfile(ctx context.Context, namespace string, name string) (v1beta1.ApplicationProfile, error)
}

// CVERepository is the port implemented by adapters to be used in ScanService to store CVE manifests
type CVERepository interface {
	GetCVE(ctx context.Context, name, SBOMCreatorVersion, CVEScannerVersion, CVEDBVersion string) (domain.CVEManifest, error)
	GetCVESummary(ctx context.Context) (*v1beta1.VulnerabilityManifestSummary, error)
	StoreCVE(ctx context.Context, cve domain.CVEManifest, withRelevancy bool) error
	StoreCVESummary(ctx context.Context, cve domain.CVEManifest, cvep domain.CVEManifest, withRelevancy bool) error
	StoreVEX(ctx context.Context, cve domain.CVEManifest, cvep domain.CVEManifest, withRelevancy bool) error
}

// SBOMRepository is the port implemented by adapters to be used in ScanService to store SBOMs
type SBOMRepository interface {
	GetSBOM(ctx context.Context, name, SBOMCreatorVersion string) (domain.SBOM, error)
	StoreSBOM(ctx context.Context, sbom domain.SBOM, isFiltered bool) error
}
