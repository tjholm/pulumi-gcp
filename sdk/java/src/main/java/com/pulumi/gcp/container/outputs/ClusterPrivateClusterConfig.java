// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.container.outputs.ClusterPrivateClusterConfigMasterGlobalAccessConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterPrivateClusterConfig {
    /**
     * @return When `true`, the cluster&#39;s private
     * endpoint is used as the cluster endpoint and access through the public endpoint
     * is disabled. When `false`, either endpoint can be used. This field only applies
     * to private clusters, when `enable_private_nodes` is `true`.
     * 
     */
    private @Nullable Boolean enablePrivateEndpoint;
    /**
     * @return Enables the private cluster feature,
     * creating a private endpoint on the cluster. In a private cluster, nodes only
     * have RFC 1918 private addresses and communicate with the master&#39;s private
     * endpoint via private networking.
     * 
     */
    private @Nullable Boolean enablePrivateNodes;
    private @Nullable ClusterPrivateClusterConfigMasterGlobalAccessConfig masterGlobalAccessConfig;
    /**
     * @return The IP range in CIDR notation to use for
     * the hosted master network. This range will be used for assigning private IP
     * addresses to the cluster master(s) and the ILB VIP. This range must not overlap
     * with any other ranges in use within the cluster&#39;s network, and it must be a /28
     * subnet. See [Private Cluster Limitations](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters#req_res_lim)
     * for more details. This field only applies to private clusters, when
     * `enable_private_nodes` is `true`.
     * 
     */
    private @Nullable String masterIpv4CidrBlock;
    /**
     * @return The name of the peering between this cluster and the Google owned VPC.
     * 
     */
    private @Nullable String peeringName;
    /**
     * @return The internal IP address of this cluster&#39;s master endpoint.
     * 
     */
    private @Nullable String privateEndpoint;
    /**
     * @return Subnetwork in cluster&#39;s network where master&#39;s endpoint will be provisioned.
     * 
     */
    private @Nullable String privateEndpointSubnetwork;
    /**
     * @return The external IP address of this cluster&#39;s master endpoint.
     * 
     */
    private @Nullable String publicEndpoint;

    private ClusterPrivateClusterConfig() {}
    /**
     * @return When `true`, the cluster&#39;s private
     * endpoint is used as the cluster endpoint and access through the public endpoint
     * is disabled. When `false`, either endpoint can be used. This field only applies
     * to private clusters, when `enable_private_nodes` is `true`.
     * 
     */
    public Optional<Boolean> enablePrivateEndpoint() {
        return Optional.ofNullable(this.enablePrivateEndpoint);
    }
    /**
     * @return Enables the private cluster feature,
     * creating a private endpoint on the cluster. In a private cluster, nodes only
     * have RFC 1918 private addresses and communicate with the master&#39;s private
     * endpoint via private networking.
     * 
     */
    public Optional<Boolean> enablePrivateNodes() {
        return Optional.ofNullable(this.enablePrivateNodes);
    }
    public Optional<ClusterPrivateClusterConfigMasterGlobalAccessConfig> masterGlobalAccessConfig() {
        return Optional.ofNullable(this.masterGlobalAccessConfig);
    }
    /**
     * @return The IP range in CIDR notation to use for
     * the hosted master network. This range will be used for assigning private IP
     * addresses to the cluster master(s) and the ILB VIP. This range must not overlap
     * with any other ranges in use within the cluster&#39;s network, and it must be a /28
     * subnet. See [Private Cluster Limitations](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters#req_res_lim)
     * for more details. This field only applies to private clusters, when
     * `enable_private_nodes` is `true`.
     * 
     */
    public Optional<String> masterIpv4CidrBlock() {
        return Optional.ofNullable(this.masterIpv4CidrBlock);
    }
    /**
     * @return The name of the peering between this cluster and the Google owned VPC.
     * 
     */
    public Optional<String> peeringName() {
        return Optional.ofNullable(this.peeringName);
    }
    /**
     * @return The internal IP address of this cluster&#39;s master endpoint.
     * 
     */
    public Optional<String> privateEndpoint() {
        return Optional.ofNullable(this.privateEndpoint);
    }
    /**
     * @return Subnetwork in cluster&#39;s network where master&#39;s endpoint will be provisioned.
     * 
     */
    public Optional<String> privateEndpointSubnetwork() {
        return Optional.ofNullable(this.privateEndpointSubnetwork);
    }
    /**
     * @return The external IP address of this cluster&#39;s master endpoint.
     * 
     */
    public Optional<String> publicEndpoint() {
        return Optional.ofNullable(this.publicEndpoint);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterPrivateClusterConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enablePrivateEndpoint;
        private @Nullable Boolean enablePrivateNodes;
        private @Nullable ClusterPrivateClusterConfigMasterGlobalAccessConfig masterGlobalAccessConfig;
        private @Nullable String masterIpv4CidrBlock;
        private @Nullable String peeringName;
        private @Nullable String privateEndpoint;
        private @Nullable String privateEndpointSubnetwork;
        private @Nullable String publicEndpoint;
        public Builder() {}
        public Builder(ClusterPrivateClusterConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enablePrivateEndpoint = defaults.enablePrivateEndpoint;
    	      this.enablePrivateNodes = defaults.enablePrivateNodes;
    	      this.masterGlobalAccessConfig = defaults.masterGlobalAccessConfig;
    	      this.masterIpv4CidrBlock = defaults.masterIpv4CidrBlock;
    	      this.peeringName = defaults.peeringName;
    	      this.privateEndpoint = defaults.privateEndpoint;
    	      this.privateEndpointSubnetwork = defaults.privateEndpointSubnetwork;
    	      this.publicEndpoint = defaults.publicEndpoint;
        }

        @CustomType.Setter
        public Builder enablePrivateEndpoint(@Nullable Boolean enablePrivateEndpoint) {
            this.enablePrivateEndpoint = enablePrivateEndpoint;
            return this;
        }
        @CustomType.Setter
        public Builder enablePrivateNodes(@Nullable Boolean enablePrivateNodes) {
            this.enablePrivateNodes = enablePrivateNodes;
            return this;
        }
        @CustomType.Setter
        public Builder masterGlobalAccessConfig(@Nullable ClusterPrivateClusterConfigMasterGlobalAccessConfig masterGlobalAccessConfig) {
            this.masterGlobalAccessConfig = masterGlobalAccessConfig;
            return this;
        }
        @CustomType.Setter
        public Builder masterIpv4CidrBlock(@Nullable String masterIpv4CidrBlock) {
            this.masterIpv4CidrBlock = masterIpv4CidrBlock;
            return this;
        }
        @CustomType.Setter
        public Builder peeringName(@Nullable String peeringName) {
            this.peeringName = peeringName;
            return this;
        }
        @CustomType.Setter
        public Builder privateEndpoint(@Nullable String privateEndpoint) {
            this.privateEndpoint = privateEndpoint;
            return this;
        }
        @CustomType.Setter
        public Builder privateEndpointSubnetwork(@Nullable String privateEndpointSubnetwork) {
            this.privateEndpointSubnetwork = privateEndpointSubnetwork;
            return this;
        }
        @CustomType.Setter
        public Builder publicEndpoint(@Nullable String publicEndpoint) {
            this.publicEndpoint = publicEndpoint;
            return this;
        }
        public ClusterPrivateClusterConfig build() {
            final var o = new ClusterPrivateClusterConfig();
            o.enablePrivateEndpoint = enablePrivateEndpoint;
            o.enablePrivateNodes = enablePrivateNodes;
            o.masterGlobalAccessConfig = masterGlobalAccessConfig;
            o.masterIpv4CidrBlock = masterIpv4CidrBlock;
            o.peeringName = peeringName;
            o.privateEndpoint = privateEndpoint;
            o.privateEndpointSubnetwork = privateEndpointSubnetwork;
            o.publicEndpoint = publicEndpoint;
            return o;
        }
    }
}
