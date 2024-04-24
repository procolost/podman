package define

import "os"

const UserCertsTargetPath = "/etc/containers/certs.d"
const DefaultIdentityName = "machine"

var (
	DefaultFilePerm os.FileMode = 0644
)

type CreateVMOpts struct {
	Name               string
	Dirs               *MachineDirs
	ReExec             bool
	UserModeNetworking bool
}

type MachineDirs struct {
	ConfigDir     *VMFile
	DataDir       *VMFile
	ImageCacheDir *VMFile
	RuntimeDir    *VMFile
}
