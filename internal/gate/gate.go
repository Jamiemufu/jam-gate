// Package gate controls the physical gate.
package gate

import "jam-gate/internal/status"

// Gate exposes gate operations without tying callers to the current hardware.
// During development, the status lights simulate the future motor-controller
// signals. The implementation can later be replaced without changing main.
type Gate struct {
	simulator *status.Status
}

// NewSimulator creates a gate that represents movement using status lights.
func NewSimulator(simulator *status.Status) *Gate {
	return &Gate{simulator: simulator}
}

// Open signals the gate to open.
func (g *Gate) Open() error {
	println("Opening gate")
	return g.simulator.Unlock()
}

// Close signals the gate to close.
func (g *Gate) Close() error {
	println("Closing gate")
	return g.simulator.Lock()
}

// Stop signals the gate to stop moving.
func (g *Gate) Stop() error {
	println("Stopping gate")
	return g.simulator.Stop()
}
