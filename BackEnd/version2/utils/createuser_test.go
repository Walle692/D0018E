package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/walle692/D0018E/BackEnd/version2/test_setup"
	"github.com/walle692/D0018E/BackEnd/version2/utils"
)

func Test_CreateUser_InsertsRow(t *testing.T) {
	ctx, pool := test_setup.SetUpDB(t)

	username, password, role := "test_username", "test_password", "buyer"
	err := utils.CreateUser(username, password, role)
	require.NoError(t, err)

	var gotRole string
	err = pool.QueryRow(ctx, `SELECT role FROM myschema.users WHERE username=$1`, username).Scan(&gotRole)
	require.NoError(t, err)
	require.Equal(t, role, gotRole)

	t.Cleanup(func() {
		if _, e := pool.Exec(ctx, `DELETE FROM myschema.users WHERE username=$1`, username); e != nil {
			t.Logf("clean up user: %v", e)
		}
	})
}
