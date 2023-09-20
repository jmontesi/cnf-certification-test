package configuration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	configuration "github.com/test-network-function/cnf-certification-test/pkg/configuration"
)

const (
	filePath     = "testdata/tnf_test_config.yml"
	nsLength     = 2
	ns1          = "tnf"
	ns2          = "test2"
	labels       = 2
	label1Prefix = "targetPod1.com"
	label1Name   = "name1"
	label1Value  = "value1"
	label2Prefix = "targetPod2.com"
	label2Name   = "name2"
	label2Value  = "value2"
	crds         = 2
	crdSuffix1   = "group1.test.com"
	crdSuffix2   = "group2.test.com"
	containers   = 1
)

func TestLoadConfiguration(t *testing.T) {
	env, err := configuration.LoadConfiguration(filePath)
	assert.Nil(t, err)
	// check if targetNameSpaces section is parsed properly
	assert.Equal(t, nsLength, len(env.TargetNameSpaces))
	ns := configuration.Namespace{Name: ns1}
	assert.Contains(t, env.TargetNameSpaces, ns)
	ns.Name = ns2
	assert.Contains(t, env.TargetNameSpaces, ns)
	// check if targetPodlabels section is parsed properly
	assert.Equal(t, labels, len(env.TargetPodLabels))
	podlabel1 := configuration.Label{Prefix: label1Prefix, Name: label1Name, Value: label1Value}
	assert.Contains(t, env.TargetPodLabels, podlabel1)
	podlabel2 := configuration.Label{Prefix: label2Prefix, Name: label2Name, Value: label2Value}
	assert.Contains(t, env.TargetPodLabels, podlabel2)
	// check if targetCrdFilters section is parsed properly
	assert.Equal(t, crds, len(env.CrdFilters))
	crd1 := configuration.CrdFilter{NameSuffix: crdSuffix1}
	assert.Contains(t, env.CrdFilters, crd1)
	crd2 := configuration.CrdFilter{NameSuffix: crdSuffix2}
	assert.Contains(t, env.CrdFilters, crd2)
}
