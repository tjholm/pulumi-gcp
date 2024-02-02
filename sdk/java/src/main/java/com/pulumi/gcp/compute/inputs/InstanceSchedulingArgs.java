// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.InstanceSchedulingLocalSsdRecoveryTimeoutArgs;
import com.pulumi.gcp.compute.inputs.InstanceSchedulingMaxRunDurationArgs;
import com.pulumi.gcp.compute.inputs.InstanceSchedulingNodeAffinityArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceSchedulingArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceSchedulingArgs Empty = new InstanceSchedulingArgs();

    /**
     * Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * Defaults to true.
     * 
     */
    @Import(name="automaticRestart")
    private @Nullable Output<Boolean> automaticRestart;

    /**
     * @return Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> automaticRestart() {
        return Optional.ofNullable(this.automaticRestart);
    }

    /**
     * Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    @Import(name="instanceTerminationAction")
    private @Nullable Output<String> instanceTerminationAction;

    /**
     * @return Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    public Optional<Output<String>> instanceTerminationAction() {
        return Optional.ofNullable(this.instanceTerminationAction);
    }

    /**
     * Specifies the maximum amount of time a Local Ssd Vm should wait while
     *   recovery of the Local Ssd state is attempted. Its value should be in
     *   between 0 and 168 hours with hour granularity and the default value being 1
     *   hour.
     * 
     */
    @Import(name="localSsdRecoveryTimeout")
    private @Nullable Output<InstanceSchedulingLocalSsdRecoveryTimeoutArgs> localSsdRecoveryTimeout;

    /**
     * @return Specifies the maximum amount of time a Local Ssd Vm should wait while
     *   recovery of the Local Ssd state is attempted. Its value should be in
     *   between 0 and 168 hours with hour granularity and the default value being 1
     *   hour.
     * 
     */
    public Optional<Output<InstanceSchedulingLocalSsdRecoveryTimeoutArgs>> localSsdRecoveryTimeout() {
        return Optional.ofNullable(this.localSsdRecoveryTimeout);
    }

    /**
     * Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
     * 
     */
    @Import(name="maintenanceInterval")
    private @Nullable Output<String> maintenanceInterval;

    /**
     * @return Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
     * 
     */
    public Optional<Output<String>> maintenanceInterval() {
        return Optional.ofNullable(this.maintenanceInterval);
    }

    /**
     * The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    @Import(name="maxRunDuration")
    private @Nullable Output<InstanceSchedulingMaxRunDurationArgs> maxRunDuration;

    /**
     * @return The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    public Optional<Output<InstanceSchedulingMaxRunDurationArgs>> maxRunDuration() {
        return Optional.ofNullable(this.maxRunDuration);
    }

    /**
     * The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
     * 
     */
    @Import(name="minNodeCpus")
    private @Nullable Output<Integer> minNodeCpus;

    /**
     * @return The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
     * 
     */
    public Optional<Output<Integer>> minNodeCpus() {
        return Optional.ofNullable(this.minNodeCpus);
    }

    /**
     * Specifies node affinities or anti-affinities
     * to determine which sole-tenant nodes your instances and managed instance
     * groups will use as host systems. Read more on sole-tenant node creation
     * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
     * Structure documented below.
     * 
     */
    @Import(name="nodeAffinities")
    private @Nullable Output<List<InstanceSchedulingNodeAffinityArgs>> nodeAffinities;

    /**
     * @return Specifies node affinities or anti-affinities
     * to determine which sole-tenant nodes your instances and managed instance
     * groups will use as host systems. Read more on sole-tenant node creation
     * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
     * Structure documented below.
     * 
     */
    public Optional<Output<List<InstanceSchedulingNodeAffinityArgs>>> nodeAffinities() {
        return Optional.ofNullable(this.nodeAffinities);
    }

    /**
     * Describes maintenance behavior for the
     * instance. Can be MIGRATE or TERMINATE, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
     * 
     */
    @Import(name="onHostMaintenance")
    private @Nullable Output<String> onHostMaintenance;

    /**
     * @return Describes maintenance behavior for the
     * instance. Can be MIGRATE or TERMINATE, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
     * 
     */
    public Optional<Output<String>> onHostMaintenance() {
        return Optional.ofNullable(this.onHostMaintenance);
    }

    /**
     * Specifies if the instance is preemptible.
     * If this field is set to true, then `automatic_restart` must be
     * set to false.  Defaults to false.
     * 
     */
    @Import(name="preemptible")
    private @Nullable Output<Boolean> preemptible;

    /**
     * @return Specifies if the instance is preemptible.
     * If this field is set to true, then `automatic_restart` must be
     * set to false.  Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> preemptible() {
        return Optional.ofNullable(this.preemptible);
    }

    /**
     * Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
     * `preemptible` should be `true` and `automatic_restart` should be
     * `false`. For more info about
     * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
     * 
     */
    @Import(name="provisioningModel")
    private @Nullable Output<String> provisioningModel;

    /**
     * @return Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
     * `preemptible` should be `true` and `automatic_restart` should be
     * `false`. For more info about
     * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
     * 
     */
    public Optional<Output<String>> provisioningModel() {
        return Optional.ofNullable(this.provisioningModel);
    }

    private InstanceSchedulingArgs() {}

    private InstanceSchedulingArgs(InstanceSchedulingArgs $) {
        this.automaticRestart = $.automaticRestart;
        this.instanceTerminationAction = $.instanceTerminationAction;
        this.localSsdRecoveryTimeout = $.localSsdRecoveryTimeout;
        this.maintenanceInterval = $.maintenanceInterval;
        this.maxRunDuration = $.maxRunDuration;
        this.minNodeCpus = $.minNodeCpus;
        this.nodeAffinities = $.nodeAffinities;
        this.onHostMaintenance = $.onHostMaintenance;
        this.preemptible = $.preemptible;
        this.provisioningModel = $.provisioningModel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceSchedulingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceSchedulingArgs $;

        public Builder() {
            $ = new InstanceSchedulingArgs();
        }

        public Builder(InstanceSchedulingArgs defaults) {
            $ = new InstanceSchedulingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param automaticRestart Specifies if the instance should be
         * restarted if it was terminated by Compute Engine (not a user).
         * Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder automaticRestart(@Nullable Output<Boolean> automaticRestart) {
            $.automaticRestart = automaticRestart;
            return this;
        }

        /**
         * @param automaticRestart Specifies if the instance should be
         * restarted if it was terminated by Compute Engine (not a user).
         * Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder automaticRestart(Boolean automaticRestart) {
            return automaticRestart(Output.of(automaticRestart));
        }

        /**
         * @param instanceTerminationAction Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
         * 
         * @return builder
         * 
         */
        public Builder instanceTerminationAction(@Nullable Output<String> instanceTerminationAction) {
            $.instanceTerminationAction = instanceTerminationAction;
            return this;
        }

        /**
         * @param instanceTerminationAction Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
         * 
         * @return builder
         * 
         */
        public Builder instanceTerminationAction(String instanceTerminationAction) {
            return instanceTerminationAction(Output.of(instanceTerminationAction));
        }

        /**
         * @param localSsdRecoveryTimeout Specifies the maximum amount of time a Local Ssd Vm should wait while
         *   recovery of the Local Ssd state is attempted. Its value should be in
         *   between 0 and 168 hours with hour granularity and the default value being 1
         *   hour.
         * 
         * @return builder
         * 
         */
        public Builder localSsdRecoveryTimeout(@Nullable Output<InstanceSchedulingLocalSsdRecoveryTimeoutArgs> localSsdRecoveryTimeout) {
            $.localSsdRecoveryTimeout = localSsdRecoveryTimeout;
            return this;
        }

        /**
         * @param localSsdRecoveryTimeout Specifies the maximum amount of time a Local Ssd Vm should wait while
         *   recovery of the Local Ssd state is attempted. Its value should be in
         *   between 0 and 168 hours with hour granularity and the default value being 1
         *   hour.
         * 
         * @return builder
         * 
         */
        public Builder localSsdRecoveryTimeout(InstanceSchedulingLocalSsdRecoveryTimeoutArgs localSsdRecoveryTimeout) {
            return localSsdRecoveryTimeout(Output.of(localSsdRecoveryTimeout));
        }

        /**
         * @param maintenanceInterval Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceInterval(@Nullable Output<String> maintenanceInterval) {
            $.maintenanceInterval = maintenanceInterval;
            return this;
        }

        /**
         * @param maintenanceInterval Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceInterval(String maintenanceInterval) {
            return maintenanceInterval(Output.of(maintenanceInterval));
        }

        /**
         * @param maxRunDuration The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
         * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
         * 
         * @return builder
         * 
         */
        public Builder maxRunDuration(@Nullable Output<InstanceSchedulingMaxRunDurationArgs> maxRunDuration) {
            $.maxRunDuration = maxRunDuration;
            return this;
        }

        /**
         * @param maxRunDuration The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
         * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
         * 
         * @return builder
         * 
         */
        public Builder maxRunDuration(InstanceSchedulingMaxRunDurationArgs maxRunDuration) {
            return maxRunDuration(Output.of(maxRunDuration));
        }

        /**
         * @param minNodeCpus The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
         * 
         * @return builder
         * 
         */
        public Builder minNodeCpus(@Nullable Output<Integer> minNodeCpus) {
            $.minNodeCpus = minNodeCpus;
            return this;
        }

        /**
         * @param minNodeCpus The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
         * 
         * @return builder
         * 
         */
        public Builder minNodeCpus(Integer minNodeCpus) {
            return minNodeCpus(Output.of(minNodeCpus));
        }

        /**
         * @param nodeAffinities Specifies node affinities or anti-affinities
         * to determine which sole-tenant nodes your instances and managed instance
         * groups will use as host systems. Read more on sole-tenant node creation
         * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
         * Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder nodeAffinities(@Nullable Output<List<InstanceSchedulingNodeAffinityArgs>> nodeAffinities) {
            $.nodeAffinities = nodeAffinities;
            return this;
        }

        /**
         * @param nodeAffinities Specifies node affinities or anti-affinities
         * to determine which sole-tenant nodes your instances and managed instance
         * groups will use as host systems. Read more on sole-tenant node creation
         * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
         * Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder nodeAffinities(List<InstanceSchedulingNodeAffinityArgs> nodeAffinities) {
            return nodeAffinities(Output.of(nodeAffinities));
        }

        /**
         * @param nodeAffinities Specifies node affinities or anti-affinities
         * to determine which sole-tenant nodes your instances and managed instance
         * groups will use as host systems. Read more on sole-tenant node creation
         * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
         * Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder nodeAffinities(InstanceSchedulingNodeAffinityArgs... nodeAffinities) {
            return nodeAffinities(List.of(nodeAffinities));
        }

        /**
         * @param onHostMaintenance Describes maintenance behavior for the
         * instance. Can be MIGRATE or TERMINATE, for more info, read
         * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
         * 
         * @return builder
         * 
         */
        public Builder onHostMaintenance(@Nullable Output<String> onHostMaintenance) {
            $.onHostMaintenance = onHostMaintenance;
            return this;
        }

        /**
         * @param onHostMaintenance Describes maintenance behavior for the
         * instance. Can be MIGRATE or TERMINATE, for more info, read
         * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
         * 
         * @return builder
         * 
         */
        public Builder onHostMaintenance(String onHostMaintenance) {
            return onHostMaintenance(Output.of(onHostMaintenance));
        }

        /**
         * @param preemptible Specifies if the instance is preemptible.
         * If this field is set to true, then `automatic_restart` must be
         * set to false.  Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder preemptible(@Nullable Output<Boolean> preemptible) {
            $.preemptible = preemptible;
            return this;
        }

        /**
         * @param preemptible Specifies if the instance is preemptible.
         * If this field is set to true, then `automatic_restart` must be
         * set to false.  Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder preemptible(Boolean preemptible) {
            return preemptible(Output.of(preemptible));
        }

        /**
         * @param provisioningModel Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
         * `preemptible` should be `true` and `automatic_restart` should be
         * `false`. For more info about
         * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
         * 
         * @return builder
         * 
         */
        public Builder provisioningModel(@Nullable Output<String> provisioningModel) {
            $.provisioningModel = provisioningModel;
            return this;
        }

        /**
         * @param provisioningModel Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
         * `preemptible` should be `true` and `automatic_restart` should be
         * `false`. For more info about
         * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
         * 
         * @return builder
         * 
         */
        public Builder provisioningModel(String provisioningModel) {
            return provisioningModel(Output.of(provisioningModel));
        }

        public InstanceSchedulingArgs build() {
            return $;
        }
    }

}
