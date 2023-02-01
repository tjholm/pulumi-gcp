// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs Empty = new ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs();

    /**
     * The URI of a sole-tenant node group resource that the cluster will be created on.
     * 
     */
    @Import(name="nodeGroupUri", required=true)
    private Output<String> nodeGroupUri;

    /**
     * @return The URI of a sole-tenant node group resource that the cluster will be created on.
     * 
     */
    public Output<String> nodeGroupUri() {
        return this.nodeGroupUri;
    }

    private ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs() {}

    private ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs(ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs $) {
        this.nodeGroupUri = $.nodeGroupUri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs $;

        public Builder() {
            $ = new ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs();
        }

        public Builder(ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs defaults) {
            $ = new ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param nodeGroupUri The URI of a sole-tenant node group resource that the cluster will be created on.
         * 
         * @return builder
         * 
         */
        public Builder nodeGroupUri(Output<String> nodeGroupUri) {
            $.nodeGroupUri = nodeGroupUri;
            return this;
        }

        /**
         * @param nodeGroupUri The URI of a sole-tenant node group resource that the cluster will be created on.
         * 
         * @return builder
         * 
         */
        public Builder nodeGroupUri(String nodeGroupUri) {
            return nodeGroupUri(Output.of(nodeGroupUri));
        }

        public ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs build() {
            $.nodeGroupUri = Objects.requireNonNull($.nodeGroupUri, "expected parameter 'nodeGroupUri' to be non-null");
            return $;
        }
    }

}
