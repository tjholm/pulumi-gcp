// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.gkeonprem.outputs.VMwareClusterNetworkConfigStaticIpConfigIpBlockIp;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class VMwareClusterNetworkConfigStaticIpConfigIpBlock {
    /**
     * @return The network gateway used by the VMware User Cluster.
     * 
     */
    private String gateway;
    /**
     * @return The node&#39;s network configurations used by the VMware User Cluster.
     * Structure is documented below.
     * 
     */
    private List<VMwareClusterNetworkConfigStaticIpConfigIpBlockIp> ips;
    /**
     * @return The netmask used by the VMware User Cluster.
     * 
     */
    private String netmask;

    private VMwareClusterNetworkConfigStaticIpConfigIpBlock() {}
    /**
     * @return The network gateway used by the VMware User Cluster.
     * 
     */
    public String gateway() {
        return this.gateway;
    }
    /**
     * @return The node&#39;s network configurations used by the VMware User Cluster.
     * Structure is documented below.
     * 
     */
    public List<VMwareClusterNetworkConfigStaticIpConfigIpBlockIp> ips() {
        return this.ips;
    }
    /**
     * @return The netmask used by the VMware User Cluster.
     * 
     */
    public String netmask() {
        return this.netmask;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VMwareClusterNetworkConfigStaticIpConfigIpBlock defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String gateway;
        private List<VMwareClusterNetworkConfigStaticIpConfigIpBlockIp> ips;
        private String netmask;
        public Builder() {}
        public Builder(VMwareClusterNetworkConfigStaticIpConfigIpBlock defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.gateway = defaults.gateway;
    	      this.ips = defaults.ips;
    	      this.netmask = defaults.netmask;
        }

        @CustomType.Setter
        public Builder gateway(String gateway) {
            this.gateway = Objects.requireNonNull(gateway);
            return this;
        }
        @CustomType.Setter
        public Builder ips(List<VMwareClusterNetworkConfigStaticIpConfigIpBlockIp> ips) {
            this.ips = Objects.requireNonNull(ips);
            return this;
        }
        public Builder ips(VMwareClusterNetworkConfigStaticIpConfigIpBlockIp... ips) {
            return ips(List.of(ips));
        }
        @CustomType.Setter
        public Builder netmask(String netmask) {
            this.netmask = Objects.requireNonNull(netmask);
            return this;
        }
        public VMwareClusterNetworkConfigStaticIpConfigIpBlock build() {
            final var o = new VMwareClusterNetworkConfigStaticIpConfigIpBlock();
            o.gateway = gateway;
            o.ips = ips;
            o.netmask = netmask;
            return o;
        }
    }
}
