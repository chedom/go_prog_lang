package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotofiesUser(t *testing.T) {
	saved := notifyUser
	defer func() { notifyUser = saved}()

	// Install the test's fake notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}

	// ...simulate a 980MB-used condition...
	const user = "jeo@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}

	const wantSubstring = "98% of you quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, " +
			"want substring %q", notifiedMsg, wantSubstring)
	}
}