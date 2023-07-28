// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vpcaccess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.vpcaccess.inputs.ConnectorSubnetArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectorState extends com.pulumi.resources.ResourceArgs {

    public static final ConnectorState Empty = new ConnectorState();

    /**
     * List of projects using the connector.
     * 
     */
    @Import(name="connectedProjects")
    private @Nullable Output<List<String>> connectedProjects;

    /**
     * @return List of projects using the connector.
     * 
     */
    public Optional<Output<List<String>>> connectedProjects() {
        return Optional.ofNullable(this.connectedProjects);
    }

    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     * 
     */
    @Import(name="ipCidrRange")
    private @Nullable Output<String> ipCidrRange;

    /**
     * @return The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     * 
     */
    public Optional<Output<String>> ipCidrRange() {
        return Optional.ofNullable(this.ipCidrRange);
    }

    /**
     * Machine type of VM Instance underlying connector. Default is e2-micro
     * 
     */
    @Import(name="machineType")
    private @Nullable Output<String> machineType;

    /**
     * @return Machine type of VM Instance underlying connector. Default is e2-micro
     * 
     */
    public Optional<Output<String>> machineType() {
        return Optional.ofNullable(this.machineType);
    }

    /**
     * Maximum value of instances in autoscaling group underlying the connector.
     * 
     */
    @Import(name="maxInstances")
    private @Nullable Output<Integer> maxInstances;

    /**
     * @return Maximum value of instances in autoscaling group underlying the connector.
     * 
     */
    public Optional<Output<Integer>> maxInstances() {
        return Optional.ofNullable(this.maxInstances);
    }

    /**
     * Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300.
     * 
     */
    @Import(name="maxThroughput")
    private @Nullable Output<Integer> maxThroughput;

    /**
     * @return Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300.
     * 
     */
    public Optional<Output<Integer>> maxThroughput() {
        return Optional.ofNullable(this.maxThroughput);
    }

    /**
     * Minimum value of instances in autoscaling group underlying the connector.
     * 
     */
    @Import(name="minInstances")
    private @Nullable Output<Integer> minInstances;

    /**
     * @return Minimum value of instances in autoscaling group underlying the connector.
     * 
     */
    public Optional<Output<Integer>> minInstances() {
        return Optional.ofNullable(this.minInstances);
    }

    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200.
     * 
     */
    @Import(name="minThroughput")
    private @Nullable Output<Integer> minThroughput;

    /**
     * @return Minimum throughput of the connector in Mbps. Default and min is 200.
     * 
     */
    public Optional<Output<Integer>> minThroughput() {
        return Optional.ofNullable(this.minThroughput);
    }

    /**
     * The name of the resource (Max 25 characters).
     * 
     * ***
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the resource (Max 25 characters).
     * 
     * ***
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
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
     * Region where the VPC Access connector resides. If it is not provided, the provider region is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Region where the VPC Access connector resides. If it is not provided, the provider region is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The fully qualified name of this VPC connector
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return The fully qualified name of this VPC connector
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * State of the VPC access connector.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return State of the VPC access connector.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * The subnet in which to house the connector
     * Structure is documented below.
     * 
     */
    @Import(name="subnet")
    private @Nullable Output<ConnectorSubnetArgs> subnet;

    /**
     * @return The subnet in which to house the connector
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectorSubnetArgs>> subnet() {
        return Optional.ofNullable(this.subnet);
    }

    private ConnectorState() {}

    private ConnectorState(ConnectorState $) {
        this.connectedProjects = $.connectedProjects;
        this.ipCidrRange = $.ipCidrRange;
        this.machineType = $.machineType;
        this.maxInstances = $.maxInstances;
        this.maxThroughput = $.maxThroughput;
        this.minInstances = $.minInstances;
        this.minThroughput = $.minThroughput;
        this.name = $.name;
        this.network = $.network;
        this.project = $.project;
        this.region = $.region;
        this.selfLink = $.selfLink;
        this.state = $.state;
        this.subnet = $.subnet;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectorState $;

        public Builder() {
            $ = new ConnectorState();
        }

        public Builder(ConnectorState defaults) {
            $ = new ConnectorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectedProjects List of projects using the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectedProjects(@Nullable Output<List<String>> connectedProjects) {
            $.connectedProjects = connectedProjects;
            return this;
        }

        /**
         * @param connectedProjects List of projects using the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectedProjects(List<String> connectedProjects) {
            return connectedProjects(Output.of(connectedProjects));
        }

        /**
         * @param connectedProjects List of projects using the connector.
         * 
         * @return builder
         * 
         */
        public Builder connectedProjects(String... connectedProjects) {
            return connectedProjects(List.of(connectedProjects));
        }

        /**
         * @param ipCidrRange The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
         * 
         * @return builder
         * 
         */
        public Builder ipCidrRange(@Nullable Output<String> ipCidrRange) {
            $.ipCidrRange = ipCidrRange;
            return this;
        }

        /**
         * @param ipCidrRange The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
         * 
         * @return builder
         * 
         */
        public Builder ipCidrRange(String ipCidrRange) {
            return ipCidrRange(Output.of(ipCidrRange));
        }

        /**
         * @param machineType Machine type of VM Instance underlying connector. Default is e2-micro
         * 
         * @return builder
         * 
         */
        public Builder machineType(@Nullable Output<String> machineType) {
            $.machineType = machineType;
            return this;
        }

        /**
         * @param machineType Machine type of VM Instance underlying connector. Default is e2-micro
         * 
         * @return builder
         * 
         */
        public Builder machineType(String machineType) {
            return machineType(Output.of(machineType));
        }

        /**
         * @param maxInstances Maximum value of instances in autoscaling group underlying the connector.
         * 
         * @return builder
         * 
         */
        public Builder maxInstances(@Nullable Output<Integer> maxInstances) {
            $.maxInstances = maxInstances;
            return this;
        }

        /**
         * @param maxInstances Maximum value of instances in autoscaling group underlying the connector.
         * 
         * @return builder
         * 
         */
        public Builder maxInstances(Integer maxInstances) {
            return maxInstances(Output.of(maxInstances));
        }

        /**
         * @param maxThroughput Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300.
         * 
         * @return builder
         * 
         */
        public Builder maxThroughput(@Nullable Output<Integer> maxThroughput) {
            $.maxThroughput = maxThroughput;
            return this;
        }

        /**
         * @param maxThroughput Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300.
         * 
         * @return builder
         * 
         */
        public Builder maxThroughput(Integer maxThroughput) {
            return maxThroughput(Output.of(maxThroughput));
        }

        /**
         * @param minInstances Minimum value of instances in autoscaling group underlying the connector.
         * 
         * @return builder
         * 
         */
        public Builder minInstances(@Nullable Output<Integer> minInstances) {
            $.minInstances = minInstances;
            return this;
        }

        /**
         * @param minInstances Minimum value of instances in autoscaling group underlying the connector.
         * 
         * @return builder
         * 
         */
        public Builder minInstances(Integer minInstances) {
            return minInstances(Output.of(minInstances));
        }

        /**
         * @param minThroughput Minimum throughput of the connector in Mbps. Default and min is 200.
         * 
         * @return builder
         * 
         */
        public Builder minThroughput(@Nullable Output<Integer> minThroughput) {
            $.minThroughput = minThroughput;
            return this;
        }

        /**
         * @param minThroughput Minimum throughput of the connector in Mbps. Default and min is 200.
         * 
         * @return builder
         * 
         */
        public Builder minThroughput(Integer minThroughput) {
            return minThroughput(Output.of(minThroughput));
        }

        /**
         * @param name The name of the resource (Max 25 characters).
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the resource (Max 25 characters).
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param network Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
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
         * @param region Region where the VPC Access connector resides. If it is not provided, the provider region is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Region where the VPC Access connector resides. If it is not provided, the provider region is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param selfLink The fully qualified name of this VPC connector
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink The fully qualified name of this VPC connector
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param state State of the VPC access connector.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state State of the VPC access connector.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param subnet The subnet in which to house the connector
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder subnet(@Nullable Output<ConnectorSubnetArgs> subnet) {
            $.subnet = subnet;
            return this;
        }

        /**
         * @param subnet The subnet in which to house the connector
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder subnet(ConnectorSubnetArgs subnet) {
            return subnet(Output.of(subnet));
        }

        public ConnectorState build() {
            return $;
        }
    }

}
