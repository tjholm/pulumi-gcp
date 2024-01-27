// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.InstanceSchedulingLocalSsdRecoveryTimeout;
import com.pulumi.gcp.compute.outputs.InstanceSchedulingMaxRunDuration;
import com.pulumi.gcp.compute.outputs.InstanceSchedulingNodeAffinity;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceScheduling {
    /**
     * @return Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * Defaults to true.
     * 
     */
    private @Nullable Boolean automaticRestart;
    /**
     * @return Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    private @Nullable String instanceTerminationAction;
    private @Nullable InstanceSchedulingLocalSsdRecoveryTimeout localSsdRecoveryTimeout;
    /**
     * @return Beta Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
     * &lt;a name=&#34;nested_guest_accelerator&#34;&gt;&lt;/a&gt;The `guest_accelerator` block supports:
     * 
     */
    private @Nullable String maintenanceInterval;
    /**
     * @return Beta The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    private @Nullable InstanceSchedulingMaxRunDuration maxRunDuration;
    /**
     * @return The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
     * 
     */
    private @Nullable Integer minNodeCpus;
    /**
     * @return Specifies node affinities or anti-affinities
     * to determine which sole-tenant nodes your instances and managed instance
     * groups will use as host systems. Read more on sole-tenant node creation
     * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
     * Structure documented below.
     * 
     */
    private @Nullable List<InstanceSchedulingNodeAffinity> nodeAffinities;
    /**
     * @return Describes maintenance behavior for the
     * instance. Can be MIGRATE or TERMINATE, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
     * 
     */
    private @Nullable String onHostMaintenance;
    /**
     * @return Specifies if the instance is preemptible.
     * If this field is set to true, then `automatic_restart` must be
     * set to false.  Defaults to false.
     * 
     */
    private @Nullable Boolean preemptible;
    /**
     * @return Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
     * `preemptible` should be `true` and `automatic_restart` should be
     * `false`. For more info about
     * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
     * 
     */
    private @Nullable String provisioningModel;

    private InstanceScheduling() {}
    /**
     * @return Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * Defaults to true.
     * 
     */
    public Optional<Boolean> automaticRestart() {
        return Optional.ofNullable(this.automaticRestart);
    }
    /**
     * @return Describe the type of termination action for VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    public Optional<String> instanceTerminationAction() {
        return Optional.ofNullable(this.instanceTerminationAction);
    }
    public Optional<InstanceSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeout() {
        return Optional.ofNullable(this.localSsdRecoveryTimeout);
    }
    /**
     * @return Beta Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
     * &lt;a name=&#34;nested_guest_accelerator&#34;&gt;&lt;/a&gt;The `guest_accelerator` block supports:
     * 
     */
    public Optional<String> maintenanceInterval() {
        return Optional.ofNullable(this.maintenanceInterval);
    }
    /**
     * @return Beta The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    public Optional<InstanceSchedulingMaxRunDuration> maxRunDuration() {
        return Optional.ofNullable(this.maxRunDuration);
    }
    /**
     * @return The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
     * 
     */
    public Optional<Integer> minNodeCpus() {
        return Optional.ofNullable(this.minNodeCpus);
    }
    /**
     * @return Specifies node affinities or anti-affinities
     * to determine which sole-tenant nodes your instances and managed instance
     * groups will use as host systems. Read more on sole-tenant node creation
     * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
     * Structure documented below.
     * 
     */
    public List<InstanceSchedulingNodeAffinity> nodeAffinities() {
        return this.nodeAffinities == null ? List.of() : this.nodeAffinities;
    }
    /**
     * @return Describes maintenance behavior for the
     * instance. Can be MIGRATE or TERMINATE, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
     * 
     */
    public Optional<String> onHostMaintenance() {
        return Optional.ofNullable(this.onHostMaintenance);
    }
    /**
     * @return Specifies if the instance is preemptible.
     * If this field is set to true, then `automatic_restart` must be
     * set to false.  Defaults to false.
     * 
     */
    public Optional<Boolean> preemptible() {
        return Optional.ofNullable(this.preemptible);
    }
    /**
     * @return Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
     * `preemptible` should be `true` and `automatic_restart` should be
     * `false`. For more info about
     * `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
     * 
     */
    public Optional<String> provisioningModel() {
        return Optional.ofNullable(this.provisioningModel);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceScheduling defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean automaticRestart;
        private @Nullable String instanceTerminationAction;
        private @Nullable InstanceSchedulingLocalSsdRecoveryTimeout localSsdRecoveryTimeout;
        private @Nullable String maintenanceInterval;
        private @Nullable InstanceSchedulingMaxRunDuration maxRunDuration;
        private @Nullable Integer minNodeCpus;
        private @Nullable List<InstanceSchedulingNodeAffinity> nodeAffinities;
        private @Nullable String onHostMaintenance;
        private @Nullable Boolean preemptible;
        private @Nullable String provisioningModel;
        public Builder() {}
        public Builder(InstanceScheduling defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.automaticRestart = defaults.automaticRestart;
    	      this.instanceTerminationAction = defaults.instanceTerminationAction;
    	      this.localSsdRecoveryTimeout = defaults.localSsdRecoveryTimeout;
    	      this.maintenanceInterval = defaults.maintenanceInterval;
    	      this.maxRunDuration = defaults.maxRunDuration;
    	      this.minNodeCpus = defaults.minNodeCpus;
    	      this.nodeAffinities = defaults.nodeAffinities;
    	      this.onHostMaintenance = defaults.onHostMaintenance;
    	      this.preemptible = defaults.preemptible;
    	      this.provisioningModel = defaults.provisioningModel;
        }

        @CustomType.Setter
        public Builder automaticRestart(@Nullable Boolean automaticRestart) {

            this.automaticRestart = automaticRestart;
            return this;
        }
        @CustomType.Setter
        public Builder instanceTerminationAction(@Nullable String instanceTerminationAction) {

            this.instanceTerminationAction = instanceTerminationAction;
            return this;
        }
        @CustomType.Setter
        public Builder localSsdRecoveryTimeout(@Nullable InstanceSchedulingLocalSsdRecoveryTimeout localSsdRecoveryTimeout) {

            this.localSsdRecoveryTimeout = localSsdRecoveryTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceInterval(@Nullable String maintenanceInterval) {

            this.maintenanceInterval = maintenanceInterval;
            return this;
        }
        @CustomType.Setter
        public Builder maxRunDuration(@Nullable InstanceSchedulingMaxRunDuration maxRunDuration) {

            this.maxRunDuration = maxRunDuration;
            return this;
        }
        @CustomType.Setter
        public Builder minNodeCpus(@Nullable Integer minNodeCpus) {

            this.minNodeCpus = minNodeCpus;
            return this;
        }
        @CustomType.Setter
        public Builder nodeAffinities(@Nullable List<InstanceSchedulingNodeAffinity> nodeAffinities) {

            this.nodeAffinities = nodeAffinities;
            return this;
        }
        public Builder nodeAffinities(InstanceSchedulingNodeAffinity... nodeAffinities) {
            return nodeAffinities(List.of(nodeAffinities));
        }
        @CustomType.Setter
        public Builder onHostMaintenance(@Nullable String onHostMaintenance) {

            this.onHostMaintenance = onHostMaintenance;
            return this;
        }
        @CustomType.Setter
        public Builder preemptible(@Nullable Boolean preemptible) {

            this.preemptible = preemptible;
            return this;
        }
        @CustomType.Setter
        public Builder provisioningModel(@Nullable String provisioningModel) {

            this.provisioningModel = provisioningModel;
            return this;
        }
        public InstanceScheduling build() {
            final var _resultValue = new InstanceScheduling();
            _resultValue.automaticRestart = automaticRestart;
            _resultValue.instanceTerminationAction = instanceTerminationAction;
            _resultValue.localSsdRecoveryTimeout = localSsdRecoveryTimeout;
            _resultValue.maintenanceInterval = maintenanceInterval;
            _resultValue.maxRunDuration = maxRunDuration;
            _resultValue.minNodeCpus = minNodeCpus;
            _resultValue.nodeAffinities = nodeAffinities;
            _resultValue.onHostMaintenance = onHostMaintenance;
            _resultValue.preemptible = preemptible;
            _resultValue.provisioningModel = provisioningModel;
            return _resultValue;
        }
    }
}
