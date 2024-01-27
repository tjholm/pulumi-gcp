// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.InstanceTemplateSchedulingLocalSsdRecoveryTimeout;
import com.pulumi.gcp.compute.outputs.InstanceTemplateSchedulingMaxRunDuration;
import com.pulumi.gcp.compute.outputs.InstanceTemplateSchedulingNodeAffinity;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceTemplateScheduling {
    /**
     * @return Specifies whether the instance should be
     * automatically restarted if it is terminated by Compute Engine (not
     * terminated by a user). This defaults to true.
     * 
     */
    private @Nullable Boolean automaticRestart;
    /**
     * @return Describe the type of termination action for `SPOT` VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    private @Nullable String instanceTerminationAction;
    private @Nullable List<InstanceTemplateSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeouts;
    /**
     * @return Beta Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
     * &lt;a name=&#34;nested_guest_accelerator&#34;&gt;&lt;/a&gt;The `guest_accelerator` block supports:
     * 
     */
    private @Nullable String maintenanceInterval;
    /**
     * @return Beta - The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    private @Nullable InstanceTemplateSchedulingMaxRunDuration maxRunDuration;
    private @Nullable Integer minNodeCpus;
    /**
     * @return Specifies node affinities or anti-affinities
     * to determine which sole-tenant nodes your instances and managed instance
     * groups will use as host systems. Read more on sole-tenant node creation
     * [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
     * Structure documented below.
     * 
     */
    private @Nullable List<InstanceTemplateSchedulingNodeAffinity> nodeAffinities;
    /**
     * @return Defines the maintenance behavior for this
     * instance.
     * 
     */
    private @Nullable String onHostMaintenance;
    /**
     * @return Allows instance to be preempted. This defaults to
     * false. Read more on this
     * [here](https://cloud.google.com/compute/docs/instances/preemptible).
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

    private InstanceTemplateScheduling() {}
    /**
     * @return Specifies whether the instance should be
     * automatically restarted if it is terminated by Compute Engine (not
     * terminated by a user). This defaults to true.
     * 
     */
    public Optional<Boolean> automaticRestart() {
        return Optional.ofNullable(this.automaticRestart);
    }
    /**
     * @return Describe the type of termination action for `SPOT` VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
     * 
     */
    public Optional<String> instanceTerminationAction() {
        return Optional.ofNullable(this.instanceTerminationAction);
    }
    public List<InstanceTemplateSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeouts() {
        return this.localSsdRecoveryTimeouts == null ? List.of() : this.localSsdRecoveryTimeouts;
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
     * @return Beta - The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
     * &lt;a name=&#34;nested_max_run_duration&#34;&gt;&lt;/a&gt;The `max_run_duration` block supports:
     * 
     */
    public Optional<InstanceTemplateSchedulingMaxRunDuration> maxRunDuration() {
        return Optional.ofNullable(this.maxRunDuration);
    }
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
    public List<InstanceTemplateSchedulingNodeAffinity> nodeAffinities() {
        return this.nodeAffinities == null ? List.of() : this.nodeAffinities;
    }
    /**
     * @return Defines the maintenance behavior for this
     * instance.
     * 
     */
    public Optional<String> onHostMaintenance() {
        return Optional.ofNullable(this.onHostMaintenance);
    }
    /**
     * @return Allows instance to be preempted. This defaults to
     * false. Read more on this
     * [here](https://cloud.google.com/compute/docs/instances/preemptible).
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

    public static Builder builder(InstanceTemplateScheduling defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean automaticRestart;
        private @Nullable String instanceTerminationAction;
        private @Nullable List<InstanceTemplateSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeouts;
        private @Nullable String maintenanceInterval;
        private @Nullable InstanceTemplateSchedulingMaxRunDuration maxRunDuration;
        private @Nullable Integer minNodeCpus;
        private @Nullable List<InstanceTemplateSchedulingNodeAffinity> nodeAffinities;
        private @Nullable String onHostMaintenance;
        private @Nullable Boolean preemptible;
        private @Nullable String provisioningModel;
        public Builder() {}
        public Builder(InstanceTemplateScheduling defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.automaticRestart = defaults.automaticRestart;
    	      this.instanceTerminationAction = defaults.instanceTerminationAction;
    	      this.localSsdRecoveryTimeouts = defaults.localSsdRecoveryTimeouts;
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
        public Builder localSsdRecoveryTimeouts(@Nullable List<InstanceTemplateSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeouts) {

            this.localSsdRecoveryTimeouts = localSsdRecoveryTimeouts;
            return this;
        }
        public Builder localSsdRecoveryTimeouts(InstanceTemplateSchedulingLocalSsdRecoveryTimeout... localSsdRecoveryTimeouts) {
            return localSsdRecoveryTimeouts(List.of(localSsdRecoveryTimeouts));
        }
        @CustomType.Setter
        public Builder maintenanceInterval(@Nullable String maintenanceInterval) {

            this.maintenanceInterval = maintenanceInterval;
            return this;
        }
        @CustomType.Setter
        public Builder maxRunDuration(@Nullable InstanceTemplateSchedulingMaxRunDuration maxRunDuration) {

            this.maxRunDuration = maxRunDuration;
            return this;
        }
        @CustomType.Setter
        public Builder minNodeCpus(@Nullable Integer minNodeCpus) {

            this.minNodeCpus = minNodeCpus;
            return this;
        }
        @CustomType.Setter
        public Builder nodeAffinities(@Nullable List<InstanceTemplateSchedulingNodeAffinity> nodeAffinities) {

            this.nodeAffinities = nodeAffinities;
            return this;
        }
        public Builder nodeAffinities(InstanceTemplateSchedulingNodeAffinity... nodeAffinities) {
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
        public InstanceTemplateScheduling build() {
            final var _resultValue = new InstanceTemplateScheduling();
            _resultValue.automaticRestart = automaticRestart;
            _resultValue.instanceTerminationAction = instanceTerminationAction;
            _resultValue.localSsdRecoveryTimeouts = localSsdRecoveryTimeouts;
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
