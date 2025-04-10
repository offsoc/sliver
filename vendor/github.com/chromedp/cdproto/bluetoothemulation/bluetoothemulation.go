// Package bluetoothemulation provides the Chrome DevTools Protocol
// commands, types, and events for the BluetoothEmulation domain.
//
// This domain allows configuring virtual Bluetooth devices to test the
// web-bluetooth API.
//
// Generated by the cdproto-gen command.
package bluetoothemulation

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// EnableParams enable the BluetoothEmulation domain.
type EnableParams struct {
	State       CentralState `json:"state"`       // State of the simulated central.
	LeSupported bool         `json:"leSupported"` // If the simulated central supports low-energy.
}

// Enable enable the BluetoothEmulation domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BluetoothEmulation#method-enable
//
// parameters:
//
//	state - State of the simulated central.
//	leSupported - If the simulated central supports low-energy.
func Enable(state CentralState, leSupported bool) *EnableParams {
	return &EnableParams{
		State:       state,
		LeSupported: leSupported,
	}
}

// Do executes BluetoothEmulation.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, p, nil)
}

// SetSimulatedCentralStateParams set the state of the simulated central.
type SetSimulatedCentralStateParams struct {
	State CentralState `json:"state"` // State of the simulated central.
}

// SetSimulatedCentralState set the state of the simulated central.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BluetoothEmulation#method-setSimulatedCentralState
//
// parameters:
//
//	state - State of the simulated central.
func SetSimulatedCentralState(state CentralState) *SetSimulatedCentralStateParams {
	return &SetSimulatedCentralStateParams{
		State: state,
	}
}

// Do executes BluetoothEmulation.setSimulatedCentralState against the provided context.
func (p *SetSimulatedCentralStateParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetSimulatedCentralState, p, nil)
}

// DisableParams disable the BluetoothEmulation domain.
type DisableParams struct{}

// Disable disable the BluetoothEmulation domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BluetoothEmulation#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes BluetoothEmulation.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// SimulatePreconnectedPeripheralParams simulates a peripheral with
// |address|, |name| and |knownServiceUuids| that has already been connected to
// the system.
type SimulatePreconnectedPeripheralParams struct {
	Address           string              `json:"address"`
	Name              string              `json:"name"`
	ManufacturerData  []*ManufacturerData `json:"manufacturerData"`
	KnownServiceUUIDs []string            `json:"knownServiceUuids"`
}

// SimulatePreconnectedPeripheral simulates a peripheral with |address|,
// |name| and |knownServiceUuids| that has already been connected to the system.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BluetoothEmulation#method-simulatePreconnectedPeripheral
//
// parameters:
//
//	address
//	name
//	manufacturerData
//	knownServiceUUIDs
func SimulatePreconnectedPeripheral(address string, name string, manufacturerData []*ManufacturerData, knownServiceUUIDs []string) *SimulatePreconnectedPeripheralParams {
	return &SimulatePreconnectedPeripheralParams{
		Address:           address,
		Name:              name,
		ManufacturerData:  manufacturerData,
		KnownServiceUUIDs: knownServiceUUIDs,
	}
}

// Do executes BluetoothEmulation.simulatePreconnectedPeripheral against the provided context.
func (p *SimulatePreconnectedPeripheralParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSimulatePreconnectedPeripheral, p, nil)
}

// SimulateAdvertisementParams simulates an advertisement packet described in
// |entry| being received by the central.
type SimulateAdvertisementParams struct {
	Entry *ScanEntry `json:"entry"`
}

// SimulateAdvertisement simulates an advertisement packet described in
// |entry| being received by the central.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BluetoothEmulation#method-simulateAdvertisement
//
// parameters:
//
//	entry
func SimulateAdvertisement(entry *ScanEntry) *SimulateAdvertisementParams {
	return &SimulateAdvertisementParams{
		Entry: entry,
	}
}

// Do executes BluetoothEmulation.simulateAdvertisement against the provided context.
func (p *SimulateAdvertisementParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSimulateAdvertisement, p, nil)
}

// Command names.
const (
	CommandEnable                         = "BluetoothEmulation.enable"
	CommandSetSimulatedCentralState       = "BluetoothEmulation.setSimulatedCentralState"
	CommandDisable                        = "BluetoothEmulation.disable"
	CommandSimulatePreconnectedPeripheral = "BluetoothEmulation.simulatePreconnectedPeripheral"
	CommandSimulateAdvertisement          = "BluetoothEmulation.simulateAdvertisement"
)
