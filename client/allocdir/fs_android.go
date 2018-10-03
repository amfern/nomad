package allocdir

import (
	"os"
	"path/filepath"
)

var (
	// SharedAllocContainerPath is the path inside container for mounted
	// directory shared across tasks in a task group.
	SharedAllocContainerPath = filepath.Join("/", SharedAllocName)

	// TaskLocalContainer is the path inside a container for mounted directory
	// for local storage.
	TaskLocalContainerPath = filepath.Join("/", TaskLocal)

	// TaskSecretsContainerPath is the path inside a container for mounted
	// secrets directory
	TaskSecretsContainerPath = filepath.Join("/", TaskSecrets)
)

// linkOrCopy attempts to hardlink dst to src and fallsback to copying if the
// hardlink fails.
func linkOrCopy(src, dst string, uid, gid int, perm os.FileMode) error {
	// Avoid link/copy if the file already exists in the chroot
	// TODO 0.6 clean this up. This was needed because chroot creation fails
	// when a process restarts.
	if fileInfo, _ := os.Stat(dst); fileInfo != nil {
		return nil
	}
	// Attempt to hardlink.
	if err := os.Link(src, dst); err == nil {
		return nil
	}

	return fileCopy(src, dst, uid, gid, perm)
}

// The windows version does nothing currently.
func linkDir(src, dst string) error {
	return nil
}

// The windows version does nothing currently.
func unlinkDir(dir string) error {
	return nil
}

// createSecretDir creates the secrets dir folder at the given path
func createSecretDir(dir string) error {
	return os.MkdirAll(dir, 0777)
}

// removeSecretDir removes the secrets dir folder
func removeSecretDir(dir string) error {
	return os.RemoveAll(dir)
}

// The android version does nothing currently.
func dropDirPermissions(path string, desired os.FileMode) error {
	return nil
}

// MountSpecialDirs mounts the dev and proc file system on the chroot of the
// task. It's a no-op on android.
func MountSpecialDirs(taskDir string) error {
	return nil
}

// unmountSpecialDirs unmounts the dev and proc file system from the chroot
func unmountSpecialDirs(taskDir string) error {
	return nil
}

// getOwner is irrelevant in android as users
func getOwner(os.FileInfo) (int, int) {
	return idUnsupported, idUnsupported
}

// func getOwner(fi os.FileInfo) (int, int) {
// 	stat, ok := fi.Sys().(*syscall.Stat_t)
// 	if !ok {
// 		return -1, -1
// 	}
// 	return int(stat.Uid), int(stat.Gid)
// }
