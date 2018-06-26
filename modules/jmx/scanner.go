// Package jmx provides a zgrab2 module that scans for jmx.
// TODO: Describe module, the flags, the probe, the output, etc.
package jmx

import (
	log "github.com/sirupsen/logrus"
	"github.com/zmap/zgrab2"
    "bytes"
)

// ScanResults instances are returned by the module's Scan function.
type ScanResults struct {
    ret bool `long:"ret" description:"show if the port is a rmi port"`
	// Protocols that support TLS should include
	// TLSLog      *zgrab2.TLSLog `json:"tls,omitempty"`
}

// Flags holds the command-line configuration for the jmx scan module.
// Populated by the framework.
type Flags struct {
	zgrab2.BaseFlags
	// Protocols that support TLS should include zgrab2.TLSFlags

	Verbose bool `long:"verbose" description:"More verbose logging, include debug fields in the scan results"`
}

// Module implements the zgrab2.Module interface.
type Module struct {
	// TODO: Add any module-global state
}

// Scanner implements the zgrab2.Scanner interface.
type Scanner struct {
	config *Flags
	// TODO: Add scan state
}

// RegisterModule registers the zgrab2 module.
func RegisterModule() {
	var module Module
	_, err := zgrab2.AddCommand("jmx", "jmx", "Probe for jmx", 1099, &module)
	if err != nil {
		log.Fatal(err)
	}
}

// NewFlags returns a default Flags object.
func (module *Module) NewFlags() interface{} {
	return new(Flags)
}

// NewScanner returns a new Scanner instance.
func (module *Module) NewScanner() zgrab2.Scanner {
	return new(Scanner)
}

// Validate checks that the flags are valid.
// On success, returns nil.
// On failure, returns an error instance describing the error.
func (flags *Flags) Validate(args []string) error {
	return nil
}

// Help returns the module's help string.
func (flags *Flags) Help() string {
	return ""
}

// Init initializes the Scanner.
func (scanner *Scanner) Init(flags zgrab2.ScanFlags) error {
	f, _ := flags.(*Flags)
	scanner.config = f
	return nil
}

// InitPerSender initializes the scanner for a given sender.
func (scanner *Scanner) InitPerSender(senderID int) error {
	return nil
}

// GetName returns the Scanner name defined in the Flags.
func (scanner *Scanner) GetName() string {
	return scanner.config.Name
}

// Protocol returns the protocol identifier of the scan.
func (scanner *Scanner) Protocol() string {
	return "jmx"
}

// GetPort returns the port being scanned.
func (scanner *Scanner) GetPort() uint {
	return scanner.config.Port
}

// Scan TODO: describe what is scanned
func (scanner *Scanner) Scan(target zgrab2.ScanTarget) (zgrab2.ScanStatus, interface{}, error) {
	conn, err := target.Open(&scanner.config.BaseFlags)
	if err != nil {
		return zgrab2.TryGetScanStatus(err), nil, err
	}
	defer conn.Close()
	// TODO: implement

    result := &ScanResults{}
	if _, err := conn.Write([]byte("\x4a\x52\x4d\x49\x00\x02\x4b")); err != nil {
		return zgrab2.SCAN_CONNECTION_CLOSED, "", err
	}
	ret := make([]byte, 1)
    _, err = conn.Read(ret)
    if err != nil {
		return zgrab2.SCAN_CONNECTION_CLOSED, "", err
    }
    if bytes.Equal([]byte("\x4e"), ret) {
        result.ret = true
    } else {
		return zgrab2.SCAN_CONNECTION_CLOSED, "", err
    }

    return zgrab2.SCAN_SUCCESS, result, nil
}
