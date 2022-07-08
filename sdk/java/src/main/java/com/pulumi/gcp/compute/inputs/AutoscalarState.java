// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.AutoscalarAutoscalingPolicyArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AutoscalarState extends com.pulumi.resources.ResourceArgs {

    public static final AutoscalarState Empty = new AutoscalarState();

    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     * 
     */
    @Import(name="autoscalingPolicy")
    private @Nullable Output<AutoscalarAutoscalingPolicyArgs> autoscalingPolicy;

    /**
     * @return The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AutoscalarAutoscalingPolicyArgs>> autoscalingPolicy() {
        return Optional.ofNullable(this.autoscalingPolicy);
    }

    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Import(name="creationTimestamp")
    private @Nullable Output<String> creationTimestamp;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Optional<Output<String>> creationTimestamp() {
        return Optional.ofNullable(this.creationTimestamp);
    }

    /**
     * An optional description of this resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The identifier for this object. Format specified above.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The identifier for this object. Format specified above.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * The URI of the created resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * URL of the zone where the instance group resides.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return URL of the zone where the instance group resides.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private AutoscalarState() {}

    private AutoscalarState(AutoscalarState $) {
        this.autoscalingPolicy = $.autoscalingPolicy;
        this.creationTimestamp = $.creationTimestamp;
        this.description = $.description;
        this.name = $.name;
        this.project = $.project;
        this.selfLink = $.selfLink;
        this.target = $.target;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AutoscalarState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AutoscalarState $;

        public Builder() {
            $ = new AutoscalarState();
        }

        public Builder(AutoscalarState defaults) {
            $ = new AutoscalarState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoscalingPolicy The configuration parameters for the autoscaling algorithm. You can
         * define one or more of the policies for an autoscaler: cpuUtilization,
         * customMetricUtilizations, and loadBalancingUtilization.
         * If none of these are specified, the default will be to autoscale based
         * on cpuUtilization to 0.6 or 60%.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder autoscalingPolicy(@Nullable Output<AutoscalarAutoscalingPolicyArgs> autoscalingPolicy) {
            $.autoscalingPolicy = autoscalingPolicy;
            return this;
        }

        /**
         * @param autoscalingPolicy The configuration parameters for the autoscaling algorithm. You can
         * define one or more of the policies for an autoscaler: cpuUtilization,
         * customMetricUtilizations, and loadBalancingUtilization.
         * If none of these are specified, the default will be to autoscale based
         * on cpuUtilization to 0.6 or 60%.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder autoscalingPolicy(AutoscalarAutoscalingPolicyArgs autoscalingPolicy) {
            return autoscalingPolicy(Output.of(autoscalingPolicy));
        }

        /**
         * @param creationTimestamp Creation timestamp in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder creationTimestamp(@Nullable Output<String> creationTimestamp) {
            $.creationTimestamp = creationTimestamp;
            return this;
        }

        /**
         * @param creationTimestamp Creation timestamp in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder creationTimestamp(String creationTimestamp) {
            return creationTimestamp(Output.of(creationTimestamp));
        }

        /**
         * @param description An optional description of this resource.
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
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The identifier for this object. Format specified above.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The identifier for this object. Format specified above.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param target Fraction of backend capacity utilization (set in HTTP(s) load
         * balancing configuration) that autoscaler should maintain. Must
         * be a positive float value. If not defined, the default is 0.8.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Fraction of backend capacity utilization (set in HTTP(s) load
         * balancing configuration) that autoscaler should maintain. Must
         * be a positive float value. If not defined, the default is 0.8.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        /**
         * @param zone URL of the zone where the instance group resides.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone URL of the zone where the instance group resides.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public AutoscalarState build() {
            return $;
        }
    }

}
