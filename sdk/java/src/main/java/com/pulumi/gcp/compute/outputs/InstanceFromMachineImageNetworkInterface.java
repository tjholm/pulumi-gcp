// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.InstanceFromMachineImageNetworkInterfaceAccessConfig;
import com.pulumi.gcp.compute.outputs.InstanceFromMachineImageNetworkInterfaceAliasIpRange;
import com.pulumi.gcp.compute.outputs.InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceFromMachineImageNetworkInterface {
    private final @Nullable List<InstanceFromMachineImageNetworkInterfaceAccessConfig> accessConfigs;
    private final @Nullable List<InstanceFromMachineImageNetworkInterfaceAliasIpRange> aliasIpRanges;
    private final @Nullable List<InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig> ipv6AccessConfigs;
    private final @Nullable String ipv6AccessType;
    /**
     * @return A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     * 
     */
    private final @Nullable String name;
    private final @Nullable String network;
    private final @Nullable String networkIp;
    private final @Nullable String nicType;
    private final @Nullable Integer queueCount;
    private final @Nullable String stackType;
    private final @Nullable String subnetwork;
    private final @Nullable String subnetworkProject;

    @CustomType.Constructor
    private InstanceFromMachineImageNetworkInterface(
        @CustomType.Parameter("accessConfigs") @Nullable List<InstanceFromMachineImageNetworkInterfaceAccessConfig> accessConfigs,
        @CustomType.Parameter("aliasIpRanges") @Nullable List<InstanceFromMachineImageNetworkInterfaceAliasIpRange> aliasIpRanges,
        @CustomType.Parameter("ipv6AccessConfigs") @Nullable List<InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig> ipv6AccessConfigs,
        @CustomType.Parameter("ipv6AccessType") @Nullable String ipv6AccessType,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("network") @Nullable String network,
        @CustomType.Parameter("networkIp") @Nullable String networkIp,
        @CustomType.Parameter("nicType") @Nullable String nicType,
        @CustomType.Parameter("queueCount") @Nullable Integer queueCount,
        @CustomType.Parameter("stackType") @Nullable String stackType,
        @CustomType.Parameter("subnetwork") @Nullable String subnetwork,
        @CustomType.Parameter("subnetworkProject") @Nullable String subnetworkProject) {
        this.accessConfigs = accessConfigs;
        this.aliasIpRanges = aliasIpRanges;
        this.ipv6AccessConfigs = ipv6AccessConfigs;
        this.ipv6AccessType = ipv6AccessType;
        this.name = name;
        this.network = network;
        this.networkIp = networkIp;
        this.nicType = nicType;
        this.queueCount = queueCount;
        this.stackType = stackType;
        this.subnetwork = subnetwork;
        this.subnetworkProject = subnetworkProject;
    }

    public List<InstanceFromMachineImageNetworkInterfaceAccessConfig> accessConfigs() {
        return this.accessConfigs == null ? List.of() : this.accessConfigs;
    }
    public List<InstanceFromMachineImageNetworkInterfaceAliasIpRange> aliasIpRanges() {
        return this.aliasIpRanges == null ? List.of() : this.aliasIpRanges;
    }
    public List<InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig> ipv6AccessConfigs() {
        return this.ipv6AccessConfigs == null ? List.of() : this.ipv6AccessConfigs;
    }
    public Optional<String> ipv6AccessType() {
        return Optional.ofNullable(this.ipv6AccessType);
    }
    /**
     * @return A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> network() {
        return Optional.ofNullable(this.network);
    }
    public Optional<String> networkIp() {
        return Optional.ofNullable(this.networkIp);
    }
    public Optional<String> nicType() {
        return Optional.ofNullable(this.nicType);
    }
    public Optional<Integer> queueCount() {
        return Optional.ofNullable(this.queueCount);
    }
    public Optional<String> stackType() {
        return Optional.ofNullable(this.stackType);
    }
    public Optional<String> subnetwork() {
        return Optional.ofNullable(this.subnetwork);
    }
    public Optional<String> subnetworkProject() {
        return Optional.ofNullable(this.subnetworkProject);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceFromMachineImageNetworkInterface defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<InstanceFromMachineImageNetworkInterfaceAccessConfig> accessConfigs;
        private @Nullable List<InstanceFromMachineImageNetworkInterfaceAliasIpRange> aliasIpRanges;
        private @Nullable List<InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig> ipv6AccessConfigs;
        private @Nullable String ipv6AccessType;
        private @Nullable String name;
        private @Nullable String network;
        private @Nullable String networkIp;
        private @Nullable String nicType;
        private @Nullable Integer queueCount;
        private @Nullable String stackType;
        private @Nullable String subnetwork;
        private @Nullable String subnetworkProject;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceFromMachineImageNetworkInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessConfigs = defaults.accessConfigs;
    	      this.aliasIpRanges = defaults.aliasIpRanges;
    	      this.ipv6AccessConfigs = defaults.ipv6AccessConfigs;
    	      this.ipv6AccessType = defaults.ipv6AccessType;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.networkIp = defaults.networkIp;
    	      this.nicType = defaults.nicType;
    	      this.queueCount = defaults.queueCount;
    	      this.stackType = defaults.stackType;
    	      this.subnetwork = defaults.subnetwork;
    	      this.subnetworkProject = defaults.subnetworkProject;
        }

        public Builder accessConfigs(@Nullable List<InstanceFromMachineImageNetworkInterfaceAccessConfig> accessConfigs) {
            this.accessConfigs = accessConfigs;
            return this;
        }
        public Builder accessConfigs(InstanceFromMachineImageNetworkInterfaceAccessConfig... accessConfigs) {
            return accessConfigs(List.of(accessConfigs));
        }
        public Builder aliasIpRanges(@Nullable List<InstanceFromMachineImageNetworkInterfaceAliasIpRange> aliasIpRanges) {
            this.aliasIpRanges = aliasIpRanges;
            return this;
        }
        public Builder aliasIpRanges(InstanceFromMachineImageNetworkInterfaceAliasIpRange... aliasIpRanges) {
            return aliasIpRanges(List.of(aliasIpRanges));
        }
        public Builder ipv6AccessConfigs(@Nullable List<InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig> ipv6AccessConfigs) {
            this.ipv6AccessConfigs = ipv6AccessConfigs;
            return this;
        }
        public Builder ipv6AccessConfigs(InstanceFromMachineImageNetworkInterfaceIpv6AccessConfig... ipv6AccessConfigs) {
            return ipv6AccessConfigs(List.of(ipv6AccessConfigs));
        }
        public Builder ipv6AccessType(@Nullable String ipv6AccessType) {
            this.ipv6AccessType = ipv6AccessType;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder network(@Nullable String network) {
            this.network = network;
            return this;
        }
        public Builder networkIp(@Nullable String networkIp) {
            this.networkIp = networkIp;
            return this;
        }
        public Builder nicType(@Nullable String nicType) {
            this.nicType = nicType;
            return this;
        }
        public Builder queueCount(@Nullable Integer queueCount) {
            this.queueCount = queueCount;
            return this;
        }
        public Builder stackType(@Nullable String stackType) {
            this.stackType = stackType;
            return this;
        }
        public Builder subnetwork(@Nullable String subnetwork) {
            this.subnetwork = subnetwork;
            return this;
        }
        public Builder subnetworkProject(@Nullable String subnetworkProject) {
            this.subnetworkProject = subnetworkProject;
            return this;
        }        public InstanceFromMachineImageNetworkInterface build() {
            return new InstanceFromMachineImageNetworkInterface(accessConfigs, aliasIpRanges, ipv6AccessConfigs, ipv6AccessType, name, network, networkIp, nicType, queueCount, stackType, subnetwork, subnetworkProject);
        }
    }
}
