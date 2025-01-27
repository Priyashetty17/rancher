package globalroles

import (
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/rancher/tests/v2/actions/rbac"

	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/users"
	namegen "github.com/rancher/shepherd/pkg/namegenerator"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	customGlobalRole = v3.GlobalRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "",
		},
		Rules: []rbacv1.PolicyRule{
			{
				Verbs:     []string{"get", "list", "watch"},
				APIGroups: []string{"*"},
				Resources: []string{"*"},
			},
		},
	}

	globalRoleBinding = &v3.GlobalRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "",
		},
		GlobalRoleName: "",
		UserName:       "",
	}
)

func createGlobalRole(client *rancher.Client) (*v3.GlobalRole, error) {
	customGlobalRole.Name = namegen.AppendRandomString("testgr")
	createdGlobalRole, err := client.WranglerContext.Mgmt.GlobalRole().Create(&customGlobalRole)
	if err != nil {
		return nil, err
	}

	return createdGlobalRole, err
}

func createGlobalRoleAndUser(client *rancher.Client) (*v3.GlobalRole, *management.User, error) {
	customGlobalRole.Name = namegen.AppendRandomString("testgr")
	createdGlobalRole, err := client.WranglerContext.Mgmt.GlobalRole().Create(&customGlobalRole)
	if err != nil {
		return nil, nil, err
	}

	createdUser, err := users.CreateUserWithRole(client, users.UserConfig(), rbac.StandardUser.String(), customGlobalRole.Name)
	if err != nil {
		return nil, nil, err
	}

	return createdGlobalRole, createdUser, err
}
