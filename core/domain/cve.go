package domain

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/Aryaman6492/storage/pkg/apis/softwarecomposition/v1beta1"
)

const (
	CriticalSeverity   = "Critical"
	HighSeverity       = "High"
	MediumSeverity     = "Medium"
	LowSeverity        = "Low"
	NegligibleSeverity = "Negligible"
	UnknownSeverity    = "Unknown"
)

type CVEExceptions []armotypes.VulnerabilityExceptionPolicy

// CVEManifest contains a JSON CVE report manifest with some metadata
type CVEManifest struct {
	Content            *v1beta1.GrypeDocument
	Annotations        map[string]string
	Labels             map[string]string
	Name               string
	Wlid               string
	SBOMCreatorName    string
	SBOMCreatorVersion string
	CVEScannerName     string
	CVEScannerVersion  string
	CVEDBVersion       string
}
