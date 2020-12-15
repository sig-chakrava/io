package execute

import (
	"sig-gitlab.internal.synopsys.com/swip/io/adapters/binary"
	"sig-gitlab.internal.synopsys.com/swip/io/adapters/tool/SAST/coverity"
	"sig-gitlab.internal.synopsys.com/swip/io/interfaces/adapterInterfaces"
	"sig-gitlab.internal.synopsys.com/swip/io/models/adapterModels"
)

func GetBinaryAdapter(adapterBase *adapterModels.AdapterBase) adapterInterfaces.Adapter {
	binaryState := &binary.State{}
	binaryConfig := &adapterModels.BinaryAdapterConfig{}

	binary := binary.New(adapterBase, binaryState, binaryConfig)
	return binary
}

func GetCoverityAdapter(adapterBase *adapterModels.AdapterBase) adapterInterfaces.Adapter {
	coverityState := &coverity.State{}

	coverity := coverity.New(adapterBase, coverityState)
	return coverity
}
