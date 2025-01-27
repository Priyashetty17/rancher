//go:build (validation || infra.any || cluster.any || extended) && !sanity && !stress

package globalroles

import (
	"testing"

	"github.com/rancher/rancher/tests/v2/actions/rbac"
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/clusters"
	"github.com/rancher/shepherd/pkg/session"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GlobalRolesTestSuite struct {
	suite.Suite
	client  *rancher.Client
	session *session.Session
	cluster *management.Cluster
}

func (gr *GlobalRolesTestSuite) TearDownSuite() {
	gr.session.Cleanup()
}

func (gr *GlobalRolesTestSuite) SetupSuite() {
	gr.session = session.NewSession()

	client, err := rancher.NewClient("", gr.session)
	require.NoError(gr.T(), err)
	gr.client = client

	log.Info("Getting cluster name from the config file and append cluster details in rbos")
	clusterName := client.RancherConfig.ClusterName
	require.NotEmptyf(gr.T(), clusterName, "Cluster name to install should be set")
	clusterID, err := clusters.GetClusterIDByName(gr.client, clusterName)
	require.NoError(gr.T(), err, "Error getting cluster ID")
	gr.cluster, err = gr.client.Management.Cluster.ByID(clusterID)
	assert.NoError(gr.T(), err)
}

func (gr *GlobalRolesTestSuite) TestGlobalRoleCustom() {
	subSession := gr.session.NewSession()
	defer subSession.Cleanup()

	log.Info("Create a user with custom global role.")
	createdGlobalRole, createdUser, err := createGlobalRoleAndUser(gr.client)
	require.NoError(gr.T(), err)

	log.Info("Verify that the global role binding is created for the user.")
	grb, err := rbac.GetGlobalRoleBindingByUserAndRole(gr.client, createdUser.ID, createdGlobalRole.Name)
	require.NoError(gr.T(), err)
	require.NotEmpty(gr.T(), grb, "Global Role Binding not found for the user")

	log.Info(".")

}

func (gr *GlobalRolesTestSuite) TestBuiltinGlobalRole() {
	subSession := gr.session.NewSession()
	defer subSession.Cleanup()

}

func (gr *GlobalRolesTestSuite) TestUpdateBuiltinGlobalRoleFails() {
	subSession := gr.session.NewSession()
	defer subSession.Cleanup()

}

func (gr *GlobalRolesTestSuite) TestDeleteBuiltinGlobalRoleFails() {
	subSession := gr.session.NewSession()
	defer subSession.Cleanup()

}

func (gr *GlobalRolesTestSuite) TestConvertCustomGlobalRoleToBuiltinFails() {
	subSession := gr.session.NewSession()
	defer subSession.Cleanup()

}

func TestGlobalRolesTestSuite(t *testing.T) {
	suite.Run(t, new(GlobalRolesTestSuite))
}
