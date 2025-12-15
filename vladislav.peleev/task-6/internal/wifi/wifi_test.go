package wifi_test

import (
	"net"
	"testing"

	"github.com/VlasfimosY/task-6/internal/wifi"
	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface(nil), assert.AnError)

	wifiService := wifi.New(mockHandle)

	_, err := wifiService.GetAddresses()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "getting interfaces:")
}

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	t.Parallel()

	addr1, _ := net.ParseMAC("00:11:22:33:44:55")
	addr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: addr1, Name: "wlan0"},
		{HardwareAddr: addr2, Name: "wlan1"},
	}, nil)

	wifiService := wifi.New(mockHandle)

	addrs, err := wifiService.GetAddresses()
	require.NoError(t, err)
	assert.Equal(t, []net.HardwareAddr{addr1, addr2}, addrs)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface(nil), assert.AnError)

	wifiService := wifi.New(mockHandle)

	_, err := wifiService.GetNames()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "getting interfaces:")
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	t.Parallel()

	addr1, _ := net.ParseMAC("00:11:22:33:44:55")

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: addr1, Name: "wlan0"},
		{HardwareAddr: nil, Name: "lo"},
	}, nil)

	wifiService := wifi.New(mockHandle)

	names, err := wifiService.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"wlan0", "lo"}, names)
}
