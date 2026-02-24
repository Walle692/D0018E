package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils/user"
)

func Test_CreateUser_InsertsRow(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	username, password, role := "test_username", "test_password", "buyer"
	err := user.CreateUser(username, password, role)
	require.NoError(t, err)

	var gotRole string
	var gotUserID int
	err = pool.QueryRow(ctx, `SELECT role, user_id FROM myschema.users WHERE username=$1`, username).Scan(&gotRole, &gotUserID)
	require.NoError(t, err)
	require.Equal(t, role, gotRole)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.basket WHERE basket_user_id=$1`, gotUserID); e != nil {
			require.NoError(t, e)
			t.Logf("clean up user: %v", e)
		}
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.users WHERE username=$1`, username); e != nil {
			require.NoError(t, e)
			t.Logf("clean up user: %v", e)
		}
	})
}
