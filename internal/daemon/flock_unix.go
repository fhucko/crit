//go:build !windows

package daemon

import (
	"errors"
	"os"
	"syscall"
)

// flockExclusive acquires an exclusive advisory lock on f, blocking until the
// lock is available. Released automatically when the process dies.
func FlockExclusive(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
}

// flockExclusiveNB acquires an exclusive advisory lock on f without blocking.
// Returns an error immediately if another process holds the lock.
func flockExclusiveNB(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

// lockHeld reports whether a flockExclusiveNB error means another process
// holds the lock, as opposed to locking being unavailable.
func lockHeld(err error) bool {
	return errors.Is(err, syscall.EWOULDBLOCK)
}

// Funlock releases an advisory lock previously acquired on f.
func Funlock(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
