// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.gkeonprem.outputs.VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock {
    /**
     * @return The network gateway used by the VMware User Cluster.
     * 
     */
    private @Nullable String gateway;
    /**
     * @return The node&#39;s network configurations used by the VMware User Cluster.
     * Structure is documented below.
     * 
     */
    private @Nullable List<VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp> ips;
    /**
     * @return The netmask used by the VMware User Cluster.
     * 
     */
    private @Nullable String netmask;

    private VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock() {}
    /**
     * @return The network gateway used by the VMware User Cluster.
     * 
     */
    public Optional<String> gateway() {
        return Optional.ofNullable(this.gateway);
    }
    /**
     * @return The node&#39;s network configurations used by the VMware User Cluster.
     * Structure is documented below.
     * 
     */
    public List<VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp> ips() {
        return this.ips == null ? List.of() : this.ips;
    }
    /**
     * @return The netmask used by the VMware User Cluster.
     * 
     */
    public Optional<String> netmask() {
        return Optional.ofNullable(this.netmask);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String gateway;
        private @Nullable List<VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp> ips;
        private @Nullable String netmask;
        public Builder() {}
        public Builder(VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.gateway = defaults.gateway;
    	      this.ips = defaults.ips;
    	      this.netmask = defaults.netmask;
        }

        @CustomType.Setter
        public Builder gateway(@Nullable String gateway) {
            this.gateway = gateway;
            return this;
        }
        @CustomType.Setter
        public Builder ips(@Nullable List<VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp> ips) {
            this.ips = ips;
            return this;
        }
        public Builder ips(VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp... ips) {
            return ips(List.of(ips));
        }
        @CustomType.Setter
        public Builder netmask(@Nullable String netmask) {
            this.netmask = netmask;
            return this;
        }
        public VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock build() {
            final var o = new VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock();
            o.gateway = gateway;
            o.ips = ips;
            o.netmask = netmask;
            return o;
        }
    }
}
