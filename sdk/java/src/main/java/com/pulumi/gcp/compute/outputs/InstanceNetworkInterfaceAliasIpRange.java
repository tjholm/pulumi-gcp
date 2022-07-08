// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetworkInterfaceAliasIpRange {
    /**
     * @return The IP CIDR range represented by this alias IP range. This IP CIDR range
     * must belong to the specified subnetwork and cannot contain IP addresses reserved by
     * system or used by other network interfaces. This range may be a single IP address
     * (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format string (e.g. 10.1.2.0/24).
     * 
     */
    private final String ipCidrRange;
    /**
     * @return The subnetwork secondary range name specifying
     * the secondary range from which to allocate the IP CIDR range for this alias IP
     * range. If left unspecified, the primary range of the subnetwork will be used.
     * 
     */
    private final @Nullable String subnetworkRangeName;

    @CustomType.Constructor
    private InstanceNetworkInterfaceAliasIpRange(
        @CustomType.Parameter("ipCidrRange") String ipCidrRange,
        @CustomType.Parameter("subnetworkRangeName") @Nullable String subnetworkRangeName) {
        this.ipCidrRange = ipCidrRange;
        this.subnetworkRangeName = subnetworkRangeName;
    }

    /**
     * @return The IP CIDR range represented by this alias IP range. This IP CIDR range
     * must belong to the specified subnetwork and cannot contain IP addresses reserved by
     * system or used by other network interfaces. This range may be a single IP address
     * (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format string (e.g. 10.1.2.0/24).
     * 
     */
    public String ipCidrRange() {
        return this.ipCidrRange;
    }
    /**
     * @return The subnetwork secondary range name specifying
     * the secondary range from which to allocate the IP CIDR range for this alias IP
     * range. If left unspecified, the primary range of the subnetwork will be used.
     * 
     */
    public Optional<String> subnetworkRangeName() {
        return Optional.ofNullable(this.subnetworkRangeName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceNetworkInterfaceAliasIpRange defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String ipCidrRange;
        private @Nullable String subnetworkRangeName;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceNetworkInterfaceAliasIpRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipCidrRange = defaults.ipCidrRange;
    	      this.subnetworkRangeName = defaults.subnetworkRangeName;
        }

        public Builder ipCidrRange(String ipCidrRange) {
            this.ipCidrRange = Objects.requireNonNull(ipCidrRange);
            return this;
        }
        public Builder subnetworkRangeName(@Nullable String subnetworkRangeName) {
            this.subnetworkRangeName = subnetworkRangeName;
            return this;
        }        public InstanceNetworkInterfaceAliasIpRange build() {
            return new InstanceNetworkInterfaceAliasIpRange(ipCidrRange, subnetworkRangeName);
        }
    }
}
