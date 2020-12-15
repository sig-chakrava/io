package executeAdapters

import (
	"sig-gitlab.internal.synopsys.com/swip/io/adapters/binary"
	"sig-gitlab.internal.synopsys.com/swip/io/adapters/toolAdapters/SAST/coverity"
	"sig-gitlab.internal.synopsys.com/swip/io/config"
	"sig-gitlab.internal.synopsys.com/swip/io/interfaces/adapterInterfaces"
	"sig-gitlab.internal.synopsys.com/swip/io/models/adapterModels"
)

type ExecuteAdapters struct {
}

func (e *ExecuteAdapters) GetBinaryAdapter(adapterBase *adapterModels.AdapterBase, config *config.Config) adapterInterfaces.Adapter {
	binaryState := &binary.State{}
	binaryConfig := &adapterModels.BinaryAdapterConfig{
		Name:           "polaris",
		BinaryFilePath: "polaris.exe",
	}

	binary := binary.New(adapterBase, binaryState, binaryConfig)
	return binary
}

func (e *ExecuteAdapters) GetCoverityAdapter(adapterBase *adapterModels.AdapterBase, config *config.Config) adapterInterfaces.Adapter {
	coverityState := &coverity.State{
		Name:                "polaris",
		DownloadURLTemplate: "http://artifactory.internal.synopsys.com/artifactory/clops-local/clops.sig.synopsys.com/polaris_cli/%[1]s/polaris_cli-%[1]s.zip",
		Version:             "1.11.210",
		ConfigFilePath:      "",
		Args:                []string{"analyze", "-w"},
	}

	coverity := coverity.New(adapterBase, coverityState)
	return coverity
}
