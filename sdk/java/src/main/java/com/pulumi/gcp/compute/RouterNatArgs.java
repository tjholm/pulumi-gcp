// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.RouterNatLogConfigArgs;
import com.pulumi.gcp.compute.inputs.RouterNatRuleArgs;
import com.pulumi.gcp.compute.inputs.RouterNatSubnetworkArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouterNatArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouterNatArgs Empty = new RouterNatArgs();

    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     * 
     */
    @Import(name="drainNatIps")
    private @Nullable Output<List<String>> drainNatIps;

    /**
     * @return A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     * 
     */
    public Optional<Output<List<String>>> drainNatIps() {
        return Optional.ofNullable(this.drainNatIps);
    }

    /**
     * Enable Dynamic Port Allocation.
     * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
     * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
     * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
     * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
     * Mutually exclusive with enableEndpointIndependentMapping.
     * 
     */
    @Import(name="enableDynamicPortAllocation")
    private @Nullable Output<Boolean> enableDynamicPortAllocation;

    /**
     * @return Enable Dynamic Port Allocation.
     * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
     * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
     * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
     * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
     * Mutually exclusive with enableEndpointIndependentMapping.
     * 
     */
    public Optional<Output<Boolean>> enableDynamicPortAllocation() {
        return Optional.ofNullable(this.enableDynamicPortAllocation);
    }

    /**
     * Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
     * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
     * 
     */
    @Import(name="enableEndpointIndependentMapping")
    private @Nullable Output<Boolean> enableEndpointIndependentMapping;

    /**
     * @return Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
     * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
     * 
     */
    public Optional<Output<Boolean>> enableEndpointIndependentMapping() {
        return Optional.ofNullable(this.enableEndpointIndependentMapping);
    }

    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     * 
     */
    @Import(name="icmpIdleTimeoutSec")
    private @Nullable Output<Integer> icmpIdleTimeoutSec;

    /**
     * @return Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     * 
     */
    public Optional<Output<Integer>> icmpIdleTimeoutSec() {
        return Optional.ofNullable(this.icmpIdleTimeoutSec);
    }

    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     * 
     */
    @Import(name="logConfig")
    private @Nullable Output<RouterNatLogConfigArgs> logConfig;

    /**
     * @return Configuration for logging on NAT
     * Structure is documented below.
     * 
     */
    public Optional<Output<RouterNatLogConfigArgs>> logConfig() {
        return Optional.ofNullable(this.logConfig);
    }

    /**
     * Maximum number of ports allocated to a VM from this NAT.
     * This field can only be set when enableDynamicPortAllocation is enabled.
     * 
     */
    @Import(name="maxPortsPerVm")
    private @Nullable Output<Integer> maxPortsPerVm;

    /**
     * @return Maximum number of ports allocated to a VM from this NAT.
     * This field can only be set when enableDynamicPortAllocation is enabled.
     * 
     */
    public Optional<Output<Integer>> maxPortsPerVm() {
        return Optional.ofNullable(this.maxPortsPerVm);
    }

    /**
     * Minimum number of ports allocated to a VM from this NAT.
     * 
     */
    @Import(name="minPortsPerVm")
    private @Nullable Output<Integer> minPortsPerVm;

    /**
     * @return Minimum number of ports allocated to a VM from this NAT.
     * 
     */
    public Optional<Output<Integer>> minPortsPerVm() {
        return Optional.ofNullable(this.minPortsPerVm);
    }

    /**
     * Name of the NAT service. The name must be 1-63 characters long and
     * comply with RFC1035.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the NAT service. The name must be 1-63 characters long and
     * comply with RFC1035.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
     * 
     */
    @Import(name="natIpAllocateOption", required=true)
    private Output<String> natIpAllocateOption;

    /**
     * @return How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
     * 
     */
    public Output<String> natIpAllocateOption() {
        return this.natIpAllocateOption;
    }

    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     * 
     */
    @Import(name="natIps")
    private @Nullable Output<List<String>> natIps;

    /**
     * @return Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     * 
     */
    public Optional<Output<List<String>>> natIps() {
        return Optional.ofNullable(this.natIps);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Region where the router and NAT reside.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Region where the router and NAT reside.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The name of the Cloud Router in which this NAT will be configured.
     * 
     */
    @Import(name="router", required=true)
    private Output<String> router;

    /**
     * @return The name of the Cloud Router in which this NAT will be configured.
     * 
     */
    public Output<String> router() {
        return this.router;
    }

    /**
     * A list of rules associated with this NAT.
     * Structure is documented below.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<RouterNatRuleArgs>> rules;

    /**
     * @return A list of rules associated with this NAT.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<RouterNatRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
     * 
     */
    @Import(name="sourceSubnetworkIpRangesToNat", required=true)
    private Output<String> sourceSubnetworkIpRangesToNat;

    /**
     * @return How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
     * 
     */
    public Output<String> sourceSubnetworkIpRangesToNat() {
        return this.sourceSubnetworkIpRangesToNat;
    }

    /**
     * One or more subnetwork NAT configurations. Only used if
     * `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     * 
     */
    @Import(name="subnetworks")
    private @Nullable Output<List<RouterNatSubnetworkArgs>> subnetworks;

    /**
     * @return One or more subnetwork NAT configurations. Only used if
     * `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<RouterNatSubnetworkArgs>>> subnetworks() {
        return Optional.ofNullable(this.subnetworks);
    }

    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     * 
     */
    @Import(name="tcpEstablishedIdleTimeoutSec")
    private @Nullable Output<Integer> tcpEstablishedIdleTimeoutSec;

    /**
     * @return Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     * 
     */
    public Optional<Output<Integer>> tcpEstablishedIdleTimeoutSec() {
        return Optional.ofNullable(this.tcpEstablishedIdleTimeoutSec);
    }

    /**
     * Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
     * Defaults to 120s if not set.
     * 
     */
    @Import(name="tcpTimeWaitTimeoutSec")
    private @Nullable Output<Integer> tcpTimeWaitTimeoutSec;

    /**
     * @return Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
     * Defaults to 120s if not set.
     * 
     */
    public Optional<Output<Integer>> tcpTimeWaitTimeoutSec() {
        return Optional.ofNullable(this.tcpTimeWaitTimeoutSec);
    }

    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     * 
     */
    @Import(name="tcpTransitoryIdleTimeoutSec")
    private @Nullable Output<Integer> tcpTransitoryIdleTimeoutSec;

    /**
     * @return Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     * 
     */
    public Optional<Output<Integer>> tcpTransitoryIdleTimeoutSec() {
        return Optional.ofNullable(this.tcpTransitoryIdleTimeoutSec);
    }

    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     * 
     */
    @Import(name="udpIdleTimeoutSec")
    private @Nullable Output<Integer> udpIdleTimeoutSec;

    /**
     * @return Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     * 
     */
    public Optional<Output<Integer>> udpIdleTimeoutSec() {
        return Optional.ofNullable(this.udpIdleTimeoutSec);
    }

    private RouterNatArgs() {}

    private RouterNatArgs(RouterNatArgs $) {
        this.drainNatIps = $.drainNatIps;
        this.enableDynamicPortAllocation = $.enableDynamicPortAllocation;
        this.enableEndpointIndependentMapping = $.enableEndpointIndependentMapping;
        this.icmpIdleTimeoutSec = $.icmpIdleTimeoutSec;
        this.logConfig = $.logConfig;
        this.maxPortsPerVm = $.maxPortsPerVm;
        this.minPortsPerVm = $.minPortsPerVm;
        this.name = $.name;
        this.natIpAllocateOption = $.natIpAllocateOption;
        this.natIps = $.natIps;
        this.project = $.project;
        this.region = $.region;
        this.router = $.router;
        this.rules = $.rules;
        this.sourceSubnetworkIpRangesToNat = $.sourceSubnetworkIpRangesToNat;
        this.subnetworks = $.subnetworks;
        this.tcpEstablishedIdleTimeoutSec = $.tcpEstablishedIdleTimeoutSec;
        this.tcpTimeWaitTimeoutSec = $.tcpTimeWaitTimeoutSec;
        this.tcpTransitoryIdleTimeoutSec = $.tcpTransitoryIdleTimeoutSec;
        this.udpIdleTimeoutSec = $.udpIdleTimeoutSec;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouterNatArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouterNatArgs $;

        public Builder() {
            $ = new RouterNatArgs();
        }

        public Builder(RouterNatArgs defaults) {
            $ = new RouterNatArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param drainNatIps A list of URLs of the IP resources to be drained. These IPs must be
         * valid static external IPs that have been assigned to the NAT.
         * 
         * @return builder
         * 
         */
        public Builder drainNatIps(@Nullable Output<List<String>> drainNatIps) {
            $.drainNatIps = drainNatIps;
            return this;
        }

        /**
         * @param drainNatIps A list of URLs of the IP resources to be drained. These IPs must be
         * valid static external IPs that have been assigned to the NAT.
         * 
         * @return builder
         * 
         */
        public Builder drainNatIps(List<String> drainNatIps) {
            return drainNatIps(Output.of(drainNatIps));
        }

        /**
         * @param drainNatIps A list of URLs of the IP resources to be drained. These IPs must be
         * valid static external IPs that have been assigned to the NAT.
         * 
         * @return builder
         * 
         */
        public Builder drainNatIps(String... drainNatIps) {
            return drainNatIps(List.of(drainNatIps));
        }

        /**
         * @param enableDynamicPortAllocation Enable Dynamic Port Allocation.
         * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
         * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
         * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
         * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
         * Mutually exclusive with enableEndpointIndependentMapping.
         * 
         * @return builder
         * 
         */
        public Builder enableDynamicPortAllocation(@Nullable Output<Boolean> enableDynamicPortAllocation) {
            $.enableDynamicPortAllocation = enableDynamicPortAllocation;
            return this;
        }

        /**
         * @param enableDynamicPortAllocation Enable Dynamic Port Allocation.
         * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
         * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
         * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
         * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
         * Mutually exclusive with enableEndpointIndependentMapping.
         * 
         * @return builder
         * 
         */
        public Builder enableDynamicPortAllocation(Boolean enableDynamicPortAllocation) {
            return enableDynamicPortAllocation(Output.of(enableDynamicPortAllocation));
        }

        /**
         * @param enableEndpointIndependentMapping Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
         * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
         * 
         * @return builder
         * 
         */
        public Builder enableEndpointIndependentMapping(@Nullable Output<Boolean> enableEndpointIndependentMapping) {
            $.enableEndpointIndependentMapping = enableEndpointIndependentMapping;
            return this;
        }

        /**
         * @param enableEndpointIndependentMapping Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
         * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
         * 
         * @return builder
         * 
         */
        public Builder enableEndpointIndependentMapping(Boolean enableEndpointIndependentMapping) {
            return enableEndpointIndependentMapping(Output.of(enableEndpointIndependentMapping));
        }

        /**
         * @param icmpIdleTimeoutSec Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder icmpIdleTimeoutSec(@Nullable Output<Integer> icmpIdleTimeoutSec) {
            $.icmpIdleTimeoutSec = icmpIdleTimeoutSec;
            return this;
        }

        /**
         * @param icmpIdleTimeoutSec Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder icmpIdleTimeoutSec(Integer icmpIdleTimeoutSec) {
            return icmpIdleTimeoutSec(Output.of(icmpIdleTimeoutSec));
        }

        /**
         * @param logConfig Configuration for logging on NAT
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder logConfig(@Nullable Output<RouterNatLogConfigArgs> logConfig) {
            $.logConfig = logConfig;
            return this;
        }

        /**
         * @param logConfig Configuration for logging on NAT
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder logConfig(RouterNatLogConfigArgs logConfig) {
            return logConfig(Output.of(logConfig));
        }

        /**
         * @param maxPortsPerVm Maximum number of ports allocated to a VM from this NAT.
         * This field can only be set when enableDynamicPortAllocation is enabled.
         * 
         * @return builder
         * 
         */
        public Builder maxPortsPerVm(@Nullable Output<Integer> maxPortsPerVm) {
            $.maxPortsPerVm = maxPortsPerVm;
            return this;
        }

        /**
         * @param maxPortsPerVm Maximum number of ports allocated to a VM from this NAT.
         * This field can only be set when enableDynamicPortAllocation is enabled.
         * 
         * @return builder
         * 
         */
        public Builder maxPortsPerVm(Integer maxPortsPerVm) {
            return maxPortsPerVm(Output.of(maxPortsPerVm));
        }

        /**
         * @param minPortsPerVm Minimum number of ports allocated to a VM from this NAT.
         * 
         * @return builder
         * 
         */
        public Builder minPortsPerVm(@Nullable Output<Integer> minPortsPerVm) {
            $.minPortsPerVm = minPortsPerVm;
            return this;
        }

        /**
         * @param minPortsPerVm Minimum number of ports allocated to a VM from this NAT.
         * 
         * @return builder
         * 
         */
        public Builder minPortsPerVm(Integer minPortsPerVm) {
            return minPortsPerVm(Output.of(minPortsPerVm));
        }

        /**
         * @param name Name of the NAT service. The name must be 1-63 characters long and
         * comply with RFC1035.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the NAT service. The name must be 1-63 characters long and
         * comply with RFC1035.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param natIpAllocateOption How external IPs should be allocated for this NAT. Valid values are
         * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
         * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
         * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder natIpAllocateOption(Output<String> natIpAllocateOption) {
            $.natIpAllocateOption = natIpAllocateOption;
            return this;
        }

        /**
         * @param natIpAllocateOption How external IPs should be allocated for this NAT. Valid values are
         * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
         * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
         * Possible values are `MANUAL_ONLY` and `AUTO_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder natIpAllocateOption(String natIpAllocateOption) {
            return natIpAllocateOption(Output.of(natIpAllocateOption));
        }

        /**
         * @param natIps Self-links of NAT IPs. Only valid if natIpAllocateOption
         * is set to MANUAL_ONLY.
         * 
         * @return builder
         * 
         */
        public Builder natIps(@Nullable Output<List<String>> natIps) {
            $.natIps = natIps;
            return this;
        }

        /**
         * @param natIps Self-links of NAT IPs. Only valid if natIpAllocateOption
         * is set to MANUAL_ONLY.
         * 
         * @return builder
         * 
         */
        public Builder natIps(List<String> natIps) {
            return natIps(Output.of(natIps));
        }

        /**
         * @param natIps Self-links of NAT IPs. Only valid if natIpAllocateOption
         * is set to MANUAL_ONLY.
         * 
         * @return builder
         * 
         */
        public Builder natIps(String... natIps) {
            return natIps(List.of(natIps));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param region Region where the router and NAT reside.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Region where the router and NAT reside.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param router The name of the Cloud Router in which this NAT will be configured.
         * 
         * @return builder
         * 
         */
        public Builder router(Output<String> router) {
            $.router = router;
            return this;
        }

        /**
         * @param router The name of the Cloud Router in which this NAT will be configured.
         * 
         * @return builder
         * 
         */
        public Builder router(String router) {
            return router(Output.of(router));
        }

        /**
         * @param rules A list of rules associated with this NAT.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<RouterNatRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules A list of rules associated with this NAT.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<RouterNatRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules A list of rules associated with this NAT.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(RouterNatRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param sourceSubnetworkIpRangesToNat How NAT should be configured per Subnetwork.
         * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
         * IP ranges in every Subnetwork are allowed to Nat.
         * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
         * ranges in every Subnetwork are allowed to Nat.
         * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
         * (specified in the field subnetwork below). Note that if this field
         * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
         * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
         * other RouterNat section in any Router for this network in this region.
         * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
         * 
         * @return builder
         * 
         */
        public Builder sourceSubnetworkIpRangesToNat(Output<String> sourceSubnetworkIpRangesToNat) {
            $.sourceSubnetworkIpRangesToNat = sourceSubnetworkIpRangesToNat;
            return this;
        }

        /**
         * @param sourceSubnetworkIpRangesToNat How NAT should be configured per Subnetwork.
         * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
         * IP ranges in every Subnetwork are allowed to Nat.
         * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
         * ranges in every Subnetwork are allowed to Nat.
         * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
         * (specified in the field subnetwork below). Note that if this field
         * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
         * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
         * other RouterNat section in any Router for this network in this region.
         * Possible values are `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, and `LIST_OF_SUBNETWORKS`.
         * 
         * @return builder
         * 
         */
        public Builder sourceSubnetworkIpRangesToNat(String sourceSubnetworkIpRangesToNat) {
            return sourceSubnetworkIpRangesToNat(Output.of(sourceSubnetworkIpRangesToNat));
        }

        /**
         * @param subnetworks One or more subnetwork NAT configurations. Only used if
         * `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder subnetworks(@Nullable Output<List<RouterNatSubnetworkArgs>> subnetworks) {
            $.subnetworks = subnetworks;
            return this;
        }

        /**
         * @param subnetworks One or more subnetwork NAT configurations. Only used if
         * `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder subnetworks(List<RouterNatSubnetworkArgs> subnetworks) {
            return subnetworks(Output.of(subnetworks));
        }

        /**
         * @param subnetworks One or more subnetwork NAT configurations. Only used if
         * `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder subnetworks(RouterNatSubnetworkArgs... subnetworks) {
            return subnetworks(List.of(subnetworks));
        }

        /**
         * @param tcpEstablishedIdleTimeoutSec Timeout (in seconds) for TCP established connections.
         * Defaults to 1200s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpEstablishedIdleTimeoutSec(@Nullable Output<Integer> tcpEstablishedIdleTimeoutSec) {
            $.tcpEstablishedIdleTimeoutSec = tcpEstablishedIdleTimeoutSec;
            return this;
        }

        /**
         * @param tcpEstablishedIdleTimeoutSec Timeout (in seconds) for TCP established connections.
         * Defaults to 1200s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpEstablishedIdleTimeoutSec(Integer tcpEstablishedIdleTimeoutSec) {
            return tcpEstablishedIdleTimeoutSec(Output.of(tcpEstablishedIdleTimeoutSec));
        }

        /**
         * @param tcpTimeWaitTimeoutSec Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
         * Defaults to 120s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpTimeWaitTimeoutSec(@Nullable Output<Integer> tcpTimeWaitTimeoutSec) {
            $.tcpTimeWaitTimeoutSec = tcpTimeWaitTimeoutSec;
            return this;
        }

        /**
         * @param tcpTimeWaitTimeoutSec Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
         * Defaults to 120s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpTimeWaitTimeoutSec(Integer tcpTimeWaitTimeoutSec) {
            return tcpTimeWaitTimeoutSec(Output.of(tcpTimeWaitTimeoutSec));
        }

        /**
         * @param tcpTransitoryIdleTimeoutSec Timeout (in seconds) for TCP transitory connections.
         * Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpTransitoryIdleTimeoutSec(@Nullable Output<Integer> tcpTransitoryIdleTimeoutSec) {
            $.tcpTransitoryIdleTimeoutSec = tcpTransitoryIdleTimeoutSec;
            return this;
        }

        /**
         * @param tcpTransitoryIdleTimeoutSec Timeout (in seconds) for TCP transitory connections.
         * Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder tcpTransitoryIdleTimeoutSec(Integer tcpTransitoryIdleTimeoutSec) {
            return tcpTransitoryIdleTimeoutSec(Output.of(tcpTransitoryIdleTimeoutSec));
        }

        /**
         * @param udpIdleTimeoutSec Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder udpIdleTimeoutSec(@Nullable Output<Integer> udpIdleTimeoutSec) {
            $.udpIdleTimeoutSec = udpIdleTimeoutSec;
            return this;
        }

        /**
         * @param udpIdleTimeoutSec Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
         * 
         * @return builder
         * 
         */
        public Builder udpIdleTimeoutSec(Integer udpIdleTimeoutSec) {
            return udpIdleTimeoutSec(Output.of(udpIdleTimeoutSec));
        }

        public RouterNatArgs build() {
            $.natIpAllocateOption = Objects.requireNonNull($.natIpAllocateOption, "expected parameter 'natIpAllocateOption' to be non-null");
            $.router = Objects.requireNonNull($.router, "expected parameter 'router' to be non-null");
            $.sourceSubnetworkIpRangesToNat = Objects.requireNonNull($.sourceSubnetworkIpRangesToNat, "expected parameter 'sourceSubnetworkIpRangesToNat' to be non-null");
            return $;
        }
    }

}
