package drop

import (
	"os"
	"os/user"
	"strconv"
	"testing"
)

func TestDropToDaemon(t *testing.T) {
	usr, err := user.Lookup("daemon")
	if err != nil {
		t.Error(err)
	}

	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		t.Error(err)
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		t.Error(err)
	}

	if err = DropPrivileges("daemon"); err != nil {
		t.Error(err)
	}

	if os.Getgid() != gid {
		t.Error("Failed to drop gid")
	}

	if os.Getuid() != uid {
		t.Error("Failed to drop uid")
	}
}
