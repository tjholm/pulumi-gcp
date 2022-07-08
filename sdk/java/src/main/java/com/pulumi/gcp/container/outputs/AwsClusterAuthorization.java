// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.container.outputs.AwsClusterAuthorizationAdminUser;
import java.util.List;
import java.util.Objects;

@CustomType
public final class AwsClusterAuthorization {
    /**
     * @return Users to perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
     * 
     */
    private final List<AwsClusterAuthorizationAdminUser> adminUsers;

    @CustomType.Constructor
    private AwsClusterAuthorization(@CustomType.Parameter("adminUsers") List<AwsClusterAuthorizationAdminUser> adminUsers) {
        this.adminUsers = adminUsers;
    }

    /**
     * @return Users to perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
     * 
     */
    public List<AwsClusterAuthorizationAdminUser> adminUsers() {
        return this.adminUsers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AwsClusterAuthorization defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<AwsClusterAuthorizationAdminUser> adminUsers;

        public Builder() {
    	      // Empty
        }

        public Builder(AwsClusterAuthorization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminUsers = defaults.adminUsers;
        }

        public Builder adminUsers(List<AwsClusterAuthorizationAdminUser> adminUsers) {
            this.adminUsers = Objects.requireNonNull(adminUsers);
            return this;
        }
        public Builder adminUsers(AwsClusterAuthorizationAdminUser... adminUsers) {
            return adminUsers(List.of(adminUsers));
        }        public AwsClusterAuthorization build() {
            return new AwsClusterAuthorization(adminUsers);
        }
    }
}
