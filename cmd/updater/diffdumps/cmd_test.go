package diffdumps

import (
	"testing"

	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stackrox/scanner/database"
	"github.com/stackrox/scanner/ext/versionfmt"
	"github.com/stretchr/testify/assert"
)

func getVuln(namespace string, fixInVersions ...string) database.Vulnerability {
	fixedIn := make([]database.FeatureVersion, 0, len(fixInVersions))
	for _, fixInVersion := range fixInVersions {
		fixedIn = append(fixedIn, database.FeatureVersion{Feature: database.Feature{Name: uuid.NewV4().String()}, Version: fixInVersion})
	}
	return database.Vulnerability{
		Namespace: database.Namespace{Name: namespace},
		Name:      uuid.NewV4().String(),
		FixedIn:   fixedIn,
	}
}

func TestFilterFixableCentOSVulns(t *testing.T) {
	nonCentOSVulnWithNonFixable := getVuln("debian:8", versionfmt.MaxVersion)
	nonCentOSVulnWithFixable := getVuln("debian:8", "1.2.3")
	nonCentOSVulnWithFixableAndNonFixable := getVuln("debian:8", "1.2.3", versionfmt.MaxVersion)
	centOSVulnWithNonFixable := getVuln("centos:8", versionfmt.MaxVersion)
	centOSVulnWithFixable := getVuln("centos:8", "1.2.3")
	centOSVulnWithFixableAndNonFixable := getVuln("centos:8", "1.2.3", versionfmt.MaxVersion)
	out := filterFixableCentOSVulns([]database.Vulnerability{
		nonCentOSVulnWithNonFixable, nonCentOSVulnWithFixable, nonCentOSVulnWithFixableAndNonFixable,
		centOSVulnWithNonFixable, centOSVulnWithFixable, centOSVulnWithFixableAndNonFixable,
	})
	// Remove the non-fixable feature.
	centOSVulnWithFixableAndNonFixable.FixedIn = centOSVulnWithFixableAndNonFixable.FixedIn[:1]
	assert.Equal(t, []database.Vulnerability{
		nonCentOSVulnWithNonFixable, nonCentOSVulnWithFixable, nonCentOSVulnWithFixableAndNonFixable,
		centOSVulnWithFixable, centOSVulnWithFixableAndNonFixable,
	}, out)
}