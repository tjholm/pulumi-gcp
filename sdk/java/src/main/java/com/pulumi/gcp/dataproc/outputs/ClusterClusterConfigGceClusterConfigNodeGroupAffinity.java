// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
    /**
     * @return The URI of a sole-tenant node group resource that the cluster will be created on.
     * 
     */
    private String nodeGroupUri;

    private ClusterClusterConfigGceClusterConfigNodeGroupAffinity() {}
    /**
     * @return The URI of a sole-tenant node group resource that the cluster will be created on.
     * 
     */
    public String nodeGroupUri() {
        return this.nodeGroupUri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterClusterConfigGceClusterConfigNodeGroupAffinity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String nodeGroupUri;
        public Builder() {}
        public Builder(ClusterClusterConfigGceClusterConfigNodeGroupAffinity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.nodeGroupUri = defaults.nodeGroupUri;
        }

        @CustomType.Setter
        public Builder nodeGroupUri(String nodeGroupUri) {
            this.nodeGroupUri = Objects.requireNonNull(nodeGroupUri);
            return this;
        }
        public ClusterClusterConfigGceClusterConfigNodeGroupAffinity build() {
            final var o = new ClusterClusterConfigGceClusterConfigNodeGroupAffinity();
            o.nodeGroupUri = nodeGroupUri;
            return o;
        }
    }
}
