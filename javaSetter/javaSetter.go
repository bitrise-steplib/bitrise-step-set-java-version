package javaSetter

import (
	"runtime"

	"github.com/bitrise-io/go-utils/log"
)

type JavaVersion string

const (
	JavaVersion_8  = JavaVersion("Java 8")
	JavaVersion_11 = JavaVersion("Java 11")
)

type Platform string

const (
	MacOS  = Platform("MacOS")
	Ubuntu = Platform("Ubuntu")
)

func (j JavaSetter) platform() Platform {
	if runtime.GOOS == "linux" {
		j.logger.Printf("Platform: Ubuntu")
		return Ubuntu
	}
	j.logger.Printf("Platform: MacOS")
	return MacOS
}

type JavaSetter struct {
	logger  log.Logger
	version JavaVersion
}

func New(logger log.Logger, version JavaVersion) *JavaSetter {
	return &JavaSetter{logger: logger, version: version}
}

func (j JavaSetter) SetJava() (JavaVersion, error) {
	switch j.platform() {
	case MacOS:
		return j.setJavaMac()
	default:
		return j.setJavaUbuntu()
	}
}

func (j JavaSetter) setJavaMac() (JavaVersion, error) {
	j.logger.Printf("jenv global 1.8...")
	return JavaVersion_8, nil
}

func (j JavaSetter) setJavaUbuntu() (JavaVersion, error) {
	j.logger.Printf("sudo update-alternatives --set javac /usr/lib/jvm/java-8-openjdk-amd64/bin/javac...")
	return JavaVersion_8, nil
}