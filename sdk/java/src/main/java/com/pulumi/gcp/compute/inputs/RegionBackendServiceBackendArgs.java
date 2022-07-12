// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionBackendServiceBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionBackendServiceBackendArgs Empty = new RegionBackendServiceBackendArgs();

    /**
     * Specifies the balancing mode for this backend.
     * See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
     * for an explanation of load balancing modes.
     * Default value is `CONNECTION`.
     * Possible values are `UTILIZATION`, `RATE`, and `CONNECTION`.
     * 
     */
    @Import(name="balancingMode")
    private @Nullable Output<String> balancingMode;

    /**
     * @return Specifies the balancing mode for this backend.
     * See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
     * for an explanation of load balancing modes.
     * Default value is `CONNECTION`.
     * Possible values are `UTILIZATION`, `RATE`, and `CONNECTION`.
     * 
     */
    public Optional<Output<String>> balancingMode() {
        return Optional.ofNullable(this.balancingMode);
    }

    /**
     * A multiplier applied to the group&#39;s maximum servicing capacity
     * (based on UTILIZATION, RATE or CONNECTION).
     * ~&gt;**NOTE**: This field cannot be set for
     * INTERNAL region backend services (default loadBalancingScheme),
     * but is required for non-INTERNAL backend service. The total
     * capacity_scaler for all backends must be non-zero.
     * A setting of 0 means the group is completely drained, offering
     * 0% of its available Capacity. Valid range is [0.0,1.0].
     * 
     */
    @Import(name="capacityScaler")
    private @Nullable Output<Double> capacityScaler;

    /**
     * @return A multiplier applied to the group&#39;s maximum servicing capacity
     * (based on UTILIZATION, RATE or CONNECTION).
     * ~&gt;**NOTE**: This field cannot be set for
     * INTERNAL region backend services (default loadBalancingScheme),
     * but is required for non-INTERNAL backend service. The total
     * capacity_scaler for all backends must be non-zero.
     * A setting of 0 means the group is completely drained, offering
     * 0% of its available Capacity. Valid range is [0.0,1.0].
     * 
     */
    public Optional<Output<Double>> capacityScaler() {
        return Optional.ofNullable(this.capacityScaler);
    }

    /**
     * An optional description of this resource.
     * Provide this property when you create the resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description of this resource.
     * Provide this property when you create the resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * This field designates whether this is a failover backend. More
     * than one failover backend can be configured for a given RegionBackendService.
     * 
     */
    @Import(name="failover")
    private @Nullable Output<Boolean> failover;

    /**
     * @return This field designates whether this is a failover backend. More
     * than one failover backend can be configured for a given RegionBackendService.
     * 
     */
    public Optional<Output<Boolean>> failover() {
        return Optional.ofNullable(this.failover);
    }

    /**
     * The fully-qualified URL of an Instance Group or Network Endpoint
     * Group resource. In case of instance group this defines the list
     * of instances that serve traffic. Member virtual machine
     * instances from each instance group must live in the same zone as
     * the instance group itself. No two backends in a backend service
     * are allowed to use same Instance Group resource.
     * For Network Endpoint Groups this defines list of endpoints. All
     * endpoints of Network Endpoint Group must be hosted on instances
     * located in the same zone as the Network Endpoint Group.
     * Backend services cannot mix Instance Group and
     * Network Endpoint Group backends.
     * When the `load_balancing_scheme` is INTERNAL, only instance groups
     * are supported.
     * Note that you must specify an Instance Group or Network Endpoint
     * Group resource using the fully-qualified URL, rather than a
     * partial URL.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The fully-qualified URL of an Instance Group or Network Endpoint
     * Group resource. In case of instance group this defines the list
     * of instances that serve traffic. Member virtual machine
     * instances from each instance group must live in the same zone as
     * the instance group itself. No two backends in a backend service
     * are allowed to use same Instance Group resource.
     * For Network Endpoint Groups this defines list of endpoints. All
     * endpoints of Network Endpoint Group must be hosted on instances
     * located in the same zone as the Network Endpoint Group.
     * Backend services cannot mix Instance Group and
     * Network Endpoint Group backends.
     * When the `load_balancing_scheme` is INTERNAL, only instance groups
     * are supported.
     * Note that you must specify an Instance Group or Network Endpoint
     * Group resource using the fully-qualified URL, rather than a
     * partial URL.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The maximum number of connections to the backend cluster.
     * Defaults to 1024.
     * 
     */
    @Import(name="maxConnections")
    private @Nullable Output<Integer> maxConnections;

    /**
     * @return The maximum number of connections to the backend cluster.
     * Defaults to 1024.
     * 
     */
    public Optional<Output<Integer>> maxConnections() {
        return Optional.ofNullable(this.maxConnections);
    }

    /**
     * The max number of simultaneous connections that a single backend
     * network endpoint can handle. Cannot be set
     * for INTERNAL backend services.
     * This is used to calculate the capacity of the group. Can be
     * used in either CONNECTION or UTILIZATION balancing modes. For
     * CONNECTION mode, either maxConnections or
     * maxConnectionsPerEndpoint must be set.
     * 
     */
    @Import(name="maxConnectionsPerEndpoint")
    private @Nullable Output<Integer> maxConnectionsPerEndpoint;

    /**
     * @return The max number of simultaneous connections that a single backend
     * network endpoint can handle. Cannot be set
     * for INTERNAL backend services.
     * This is used to calculate the capacity of the group. Can be
     * used in either CONNECTION or UTILIZATION balancing modes. For
     * CONNECTION mode, either maxConnections or
     * maxConnectionsPerEndpoint must be set.
     * 
     */
    public Optional<Output<Integer>> maxConnectionsPerEndpoint() {
        return Optional.ofNullable(this.maxConnectionsPerEndpoint);
    }

    /**
     * The max number of simultaneous connections that a single
     * backend instance can handle. Cannot be set for INTERNAL backend
     * services.
     * This is used to calculate the capacity of the group.
     * Can be used in either CONNECTION or UTILIZATION balancing modes.
     * For CONNECTION mode, either maxConnections or
     * maxConnectionsPerInstance must be set.
     * 
     */
    @Import(name="maxConnectionsPerInstance")
    private @Nullable Output<Integer> maxConnectionsPerInstance;

    /**
     * @return The max number of simultaneous connections that a single
     * backend instance can handle. Cannot be set for INTERNAL backend
     * services.
     * This is used to calculate the capacity of the group.
     * Can be used in either CONNECTION or UTILIZATION balancing modes.
     * For CONNECTION mode, either maxConnections or
     * maxConnectionsPerInstance must be set.
     * 
     */
    public Optional<Output<Integer>> maxConnectionsPerInstance() {
        return Optional.ofNullable(this.maxConnectionsPerInstance);
    }

    /**
     * The max requests per second (RPS) of the group. Cannot be set
     * for INTERNAL backend services.
     * Can be used with either RATE or UTILIZATION balancing modes,
     * but required if RATE mode. Either maxRate or one
     * of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
     * group type, must be set.
     * 
     */
    @Import(name="maxRate")
    private @Nullable Output<Integer> maxRate;

    /**
     * @return The max requests per second (RPS) of the group. Cannot be set
     * for INTERNAL backend services.
     * Can be used with either RATE or UTILIZATION balancing modes,
     * but required if RATE mode. Either maxRate or one
     * of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
     * group type, must be set.
     * 
     */
    public Optional<Output<Integer>> maxRate() {
        return Optional.ofNullable(this.maxRate);
    }

    /**
     * The max requests per second (RPS) that a single backend network
     * endpoint can handle. This is used to calculate the capacity of
     * the group. Can be used in either balancing mode. For RATE mode,
     * either maxRate or maxRatePerEndpoint must be set. Cannot be set
     * for INTERNAL backend services.
     * 
     */
    @Import(name="maxRatePerEndpoint")
    private @Nullable Output<Double> maxRatePerEndpoint;

    /**
     * @return The max requests per second (RPS) that a single backend network
     * endpoint can handle. This is used to calculate the capacity of
     * the group. Can be used in either balancing mode. For RATE mode,
     * either maxRate or maxRatePerEndpoint must be set. Cannot be set
     * for INTERNAL backend services.
     * 
     */
    public Optional<Output<Double>> maxRatePerEndpoint() {
        return Optional.ofNullable(this.maxRatePerEndpoint);
    }

    /**
     * The max requests per second (RPS) that a single backend
     * instance can handle. This is used to calculate the capacity of
     * the group. Can be used in either balancing mode. For RATE mode,
     * either maxRate or maxRatePerInstance must be set. Cannot be set
     * for INTERNAL backend services.
     * 
     */
    @Import(name="maxRatePerInstance")
    private @Nullable Output<Double> maxRatePerInstance;

    /**
     * @return The max requests per second (RPS) that a single backend
     * instance can handle. This is used to calculate the capacity of
     * the group. Can be used in either balancing mode. For RATE mode,
     * either maxRate or maxRatePerInstance must be set. Cannot be set
     * for INTERNAL backend services.
     * 
     */
    public Optional<Output<Double>> maxRatePerInstance() {
        return Optional.ofNullable(this.maxRatePerInstance);
    }

    /**
     * Used when balancingMode is UTILIZATION. This ratio defines the
     * CPU utilization target for the group. Valid range is [0.0, 1.0].
     * Cannot be set for INTERNAL backend services.
     * 
     */
    @Import(name="maxUtilization")
    private @Nullable Output<Double> maxUtilization;

    /**
     * @return Used when balancingMode is UTILIZATION. This ratio defines the
     * CPU utilization target for the group. Valid range is [0.0, 1.0].
     * Cannot be set for INTERNAL backend services.
     * 
     */
    public Optional<Output<Double>> maxUtilization() {
        return Optional.ofNullable(this.maxUtilization);
    }

    private RegionBackendServiceBackendArgs() {}

    private RegionBackendServiceBackendArgs(RegionBackendServiceBackendArgs $) {
        this.balancingMode = $.balancingMode;
        this.capacityScaler = $.capacityScaler;
        this.description = $.description;
        this.failover = $.failover;
        this.group = $.group;
        this.maxConnections = $.maxConnections;
        this.maxConnectionsPerEndpoint = $.maxConnectionsPerEndpoint;
        this.maxConnectionsPerInstance = $.maxConnectionsPerInstance;
        this.maxRate = $.maxRate;
        this.maxRatePerEndpoint = $.maxRatePerEndpoint;
        this.maxRatePerInstance = $.maxRatePerInstance;
        this.maxUtilization = $.maxUtilization;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionBackendServiceBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionBackendServiceBackendArgs $;

        public Builder() {
            $ = new RegionBackendServiceBackendArgs();
        }

        public Builder(RegionBackendServiceBackendArgs defaults) {
            $ = new RegionBackendServiceBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param balancingMode Specifies the balancing mode for this backend.
         * See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
         * for an explanation of load balancing modes.
         * Default value is `CONNECTION`.
         * Possible values are `UTILIZATION`, `RATE`, and `CONNECTION`.
         * 
         * @return builder
         * 
         */
        public Builder balancingMode(@Nullable Output<String> balancingMode) {
            $.balancingMode = balancingMode;
            return this;
        }

        /**
         * @param balancingMode Specifies the balancing mode for this backend.
         * See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
         * for an explanation of load balancing modes.
         * Default value is `CONNECTION`.
         * Possible values are `UTILIZATION`, `RATE`, and `CONNECTION`.
         * 
         * @return builder
         * 
         */
        public Builder balancingMode(String balancingMode) {
            return balancingMode(Output.of(balancingMode));
        }

        /**
         * @param capacityScaler A multiplier applied to the group&#39;s maximum servicing capacity
         * (based on UTILIZATION, RATE or CONNECTION).
         * ~&gt;**NOTE**: This field cannot be set for
         * INTERNAL region backend services (default loadBalancingScheme),
         * but is required for non-INTERNAL backend service. The total
         * capacity_scaler for all backends must be non-zero.
         * A setting of 0 means the group is completely drained, offering
         * 0% of its available Capacity. Valid range is [0.0,1.0].
         * 
         * @return builder
         * 
         */
        public Builder capacityScaler(@Nullable Output<Double> capacityScaler) {
            $.capacityScaler = capacityScaler;
            return this;
        }

        /**
         * @param capacityScaler A multiplier applied to the group&#39;s maximum servicing capacity
         * (based on UTILIZATION, RATE or CONNECTION).
         * ~&gt;**NOTE**: This field cannot be set for
         * INTERNAL region backend services (default loadBalancingScheme),
         * but is required for non-INTERNAL backend service. The total
         * capacity_scaler for all backends must be non-zero.
         * A setting of 0 means the group is completely drained, offering
         * 0% of its available Capacity. Valid range is [0.0,1.0].
         * 
         * @return builder
         * 
         */
        public Builder capacityScaler(Double capacityScaler) {
            return capacityScaler(Output.of(capacityScaler));
        }

        /**
         * @param description An optional description of this resource.
         * Provide this property when you create the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description of this resource.
         * Provide this property when you create the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param failover This field designates whether this is a failover backend. More
         * than one failover backend can be configured for a given RegionBackendService.
         * 
         * @return builder
         * 
         */
        public Builder failover(@Nullable Output<Boolean> failover) {
            $.failover = failover;
            return this;
        }

        /**
         * @param failover This field designates whether this is a failover backend. More
         * than one failover backend can be configured for a given RegionBackendService.
         * 
         * @return builder
         * 
         */
        public Builder failover(Boolean failover) {
            return failover(Output.of(failover));
        }

        /**
         * @param group The fully-qualified URL of an Instance Group or Network Endpoint
         * Group resource. In case of instance group this defines the list
         * of instances that serve traffic. Member virtual machine
         * instances from each instance group must live in the same zone as
         * the instance group itself. No two backends in a backend service
         * are allowed to use same Instance Group resource.
         * For Network Endpoint Groups this defines list of endpoints. All
         * endpoints of Network Endpoint Group must be hosted on instances
         * located in the same zone as the Network Endpoint Group.
         * Backend services cannot mix Instance Group and
         * Network Endpoint Group backends.
         * When the `load_balancing_scheme` is INTERNAL, only instance groups
         * are supported.
         * Note that you must specify an Instance Group or Network Endpoint
         * Group resource using the fully-qualified URL, rather than a
         * partial URL.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The fully-qualified URL of an Instance Group or Network Endpoint
         * Group resource. In case of instance group this defines the list
         * of instances that serve traffic. Member virtual machine
         * instances from each instance group must live in the same zone as
         * the instance group itself. No two backends in a backend service
         * are allowed to use same Instance Group resource.
         * For Network Endpoint Groups this defines list of endpoints. All
         * endpoints of Network Endpoint Group must be hosted on instances
         * located in the same zone as the Network Endpoint Group.
         * Backend services cannot mix Instance Group and
         * Network Endpoint Group backends.
         * When the `load_balancing_scheme` is INTERNAL, only instance groups
         * are supported.
         * Note that you must specify an Instance Group or Network Endpoint
         * Group resource using the fully-qualified URL, rather than a
         * partial URL.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param maxConnections The maximum number of connections to the backend cluster.
         * Defaults to 1024.
         * 
         * @return builder
         * 
         */
        public Builder maxConnections(@Nullable Output<Integer> maxConnections) {
            $.maxConnections = maxConnections;
            return this;
        }

        /**
         * @param maxConnections The maximum number of connections to the backend cluster.
         * Defaults to 1024.
         * 
         * @return builder
         * 
         */
        public Builder maxConnections(Integer maxConnections) {
            return maxConnections(Output.of(maxConnections));
        }

        /**
         * @param maxConnectionsPerEndpoint The max number of simultaneous connections that a single backend
         * network endpoint can handle. Cannot be set
         * for INTERNAL backend services.
         * This is used to calculate the capacity of the group. Can be
         * used in either CONNECTION or UTILIZATION balancing modes. For
         * CONNECTION mode, either maxConnections or
         * maxConnectionsPerEndpoint must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionsPerEndpoint(@Nullable Output<Integer> maxConnectionsPerEndpoint) {
            $.maxConnectionsPerEndpoint = maxConnectionsPerEndpoint;
            return this;
        }

        /**
         * @param maxConnectionsPerEndpoint The max number of simultaneous connections that a single backend
         * network endpoint can handle. Cannot be set
         * for INTERNAL backend services.
         * This is used to calculate the capacity of the group. Can be
         * used in either CONNECTION or UTILIZATION balancing modes. For
         * CONNECTION mode, either maxConnections or
         * maxConnectionsPerEndpoint must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionsPerEndpoint(Integer maxConnectionsPerEndpoint) {
            return maxConnectionsPerEndpoint(Output.of(maxConnectionsPerEndpoint));
        }

        /**
         * @param maxConnectionsPerInstance The max number of simultaneous connections that a single
         * backend instance can handle. Cannot be set for INTERNAL backend
         * services.
         * This is used to calculate the capacity of the group.
         * Can be used in either CONNECTION or UTILIZATION balancing modes.
         * For CONNECTION mode, either maxConnections or
         * maxConnectionsPerInstance must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionsPerInstance(@Nullable Output<Integer> maxConnectionsPerInstance) {
            $.maxConnectionsPerInstance = maxConnectionsPerInstance;
            return this;
        }

        /**
         * @param maxConnectionsPerInstance The max number of simultaneous connections that a single
         * backend instance can handle. Cannot be set for INTERNAL backend
         * services.
         * This is used to calculate the capacity of the group.
         * Can be used in either CONNECTION or UTILIZATION balancing modes.
         * For CONNECTION mode, either maxConnections or
         * maxConnectionsPerInstance must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxConnectionsPerInstance(Integer maxConnectionsPerInstance) {
            return maxConnectionsPerInstance(Output.of(maxConnectionsPerInstance));
        }

        /**
         * @param maxRate The max requests per second (RPS) of the group. Cannot be set
         * for INTERNAL backend services.
         * Can be used with either RATE or UTILIZATION balancing modes,
         * but required if RATE mode. Either maxRate or one
         * of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
         * group type, must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxRate(@Nullable Output<Integer> maxRate) {
            $.maxRate = maxRate;
            return this;
        }

        /**
         * @param maxRate The max requests per second (RPS) of the group. Cannot be set
         * for INTERNAL backend services.
         * Can be used with either RATE or UTILIZATION balancing modes,
         * but required if RATE mode. Either maxRate or one
         * of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
         * group type, must be set.
         * 
         * @return builder
         * 
         */
        public Builder maxRate(Integer maxRate) {
            return maxRate(Output.of(maxRate));
        }

        /**
         * @param maxRatePerEndpoint The max requests per second (RPS) that a single backend network
         * endpoint can handle. This is used to calculate the capacity of
         * the group. Can be used in either balancing mode. For RATE mode,
         * either maxRate or maxRatePerEndpoint must be set. Cannot be set
         * for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxRatePerEndpoint(@Nullable Output<Double> maxRatePerEndpoint) {
            $.maxRatePerEndpoint = maxRatePerEndpoint;
            return this;
        }

        /**
         * @param maxRatePerEndpoint The max requests per second (RPS) that a single backend network
         * endpoint can handle. This is used to calculate the capacity of
         * the group. Can be used in either balancing mode. For RATE mode,
         * either maxRate or maxRatePerEndpoint must be set. Cannot be set
         * for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxRatePerEndpoint(Double maxRatePerEndpoint) {
            return maxRatePerEndpoint(Output.of(maxRatePerEndpoint));
        }

        /**
         * @param maxRatePerInstance The max requests per second (RPS) that a single backend
         * instance can handle. This is used to calculate the capacity of
         * the group. Can be used in either balancing mode. For RATE mode,
         * either maxRate or maxRatePerInstance must be set. Cannot be set
         * for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxRatePerInstance(@Nullable Output<Double> maxRatePerInstance) {
            $.maxRatePerInstance = maxRatePerInstance;
            return this;
        }

        /**
         * @param maxRatePerInstance The max requests per second (RPS) that a single backend
         * instance can handle. This is used to calculate the capacity of
         * the group. Can be used in either balancing mode. For RATE mode,
         * either maxRate or maxRatePerInstance must be set. Cannot be set
         * for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxRatePerInstance(Double maxRatePerInstance) {
            return maxRatePerInstance(Output.of(maxRatePerInstance));
        }

        /**
         * @param maxUtilization Used when balancingMode is UTILIZATION. This ratio defines the
         * CPU utilization target for the group. Valid range is [0.0, 1.0].
         * Cannot be set for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxUtilization(@Nullable Output<Double> maxUtilization) {
            $.maxUtilization = maxUtilization;
            return this;
        }

        /**
         * @param maxUtilization Used when balancingMode is UTILIZATION. This ratio defines the
         * CPU utilization target for the group. Valid range is [0.0, 1.0].
         * Cannot be set for INTERNAL backend services.
         * 
         * @return builder
         * 
         */
        public Builder maxUtilization(Double maxUtilization) {
            return maxUtilization(Output.of(maxUtilization));
        }

        public RegionBackendServiceBackendArgs build() {
            $.group = Objects.requireNonNull($.group, "expected parameter 'group' to be non-null");
            return $;
        }
    }

}
