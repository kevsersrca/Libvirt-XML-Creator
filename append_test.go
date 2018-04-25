package bin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	test := New("test")

	assert.Equal(t, test.LibvirtDomainConfigurationPath, "/etc/libvirt/qemu/test.xml")
	require.Equal(t, test.LibvirtDomainConfigurationPath, "/etc/libvirt/qemu/test.xml")
}
