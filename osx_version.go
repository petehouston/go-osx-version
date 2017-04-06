package osxversion

import (
	"os/exec"
)

const CMD_NAME = "/usr/bin/sw_vers"
const CMD_OPTION_NAME = "-productName"
const CMD_OPTION_VERSION = "-productVersion"
const CMD_OPTION_BUILDVERSION = "-buildVersion"

type OSXVersion struct {
	Name string
	Version string
	BuildVersion string
}

func callCommand(cmd string, opt string) (string, error)  {
	out, err := exec.Command(cmd, opt).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func getOSXVersion() (string, error) {
	return callCommand(CMD_NAME, CMD_OPTION_VERSION)
}

func getOSXName() (string, error) {
	return callCommand(CMD_NAME, CMD_OPTION_NAME)
}

func getOSXBuildVersion() (string, error) {
	return callCommand(CMD_NAME, CMD_OPTION_BUILDVERSION)
}

func (v *OSXVersion) Query() error {
	version, err := getOSXVersion()
	if err != nil {
		return err
	}
	v.Version = version

	name, err := getOSXName()
	if err != nil {
		return err
	}
	v.Name = name

	buildVersion, err := getOSXBuildVersion()
	if err != nil {
		return err
	}
	v.BuildVersion = buildVersion

	return nil
}