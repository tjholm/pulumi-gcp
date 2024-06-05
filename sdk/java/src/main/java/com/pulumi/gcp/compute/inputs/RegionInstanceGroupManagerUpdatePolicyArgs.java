// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionInstanceGroupManagerUpdatePolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionInstanceGroupManagerUpdatePolicyArgs Empty = new RegionInstanceGroupManagerUpdatePolicyArgs();

    /**
     * The instance redistribution policy for regional managed instance groups. Valid values are: `&#34;PROACTIVE&#34;`, `&#34;NONE&#34;`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
     * 
     */
    @Import(name="instanceRedistributionType")
    private @Nullable Output<String> instanceRedistributionType;

    /**
     * @return The instance redistribution policy for regional managed instance groups. Valid values are: `&#34;PROACTIVE&#34;`, `&#34;NONE&#34;`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
     * 
     */
    public Optional<Output<String>> instanceRedistributionType() {
        return Optional.ofNullable(this.instanceRedistributionType);
    }

    /**
     * , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
     * 
     */
    @Import(name="maxSurgeFixed")
    private @Nullable Output<Integer> maxSurgeFixed;

    /**
     * @return , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
     * 
     */
    public Optional<Output<Integer>> maxSurgeFixed() {
        return Optional.ofNullable(this.maxSurgeFixed);
    }

    /**
     * , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
     * 
     */
    @Import(name="maxSurgePercent")
    private @Nullable Output<Integer> maxSurgePercent;

    /**
     * @return , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
     * 
     */
    public Optional<Output<Integer>> maxSurgePercent() {
        return Optional.ofNullable(this.maxSurgePercent);
    }

    /**
     * , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
     * 
     */
    @Import(name="maxUnavailableFixed")
    private @Nullable Output<Integer> maxUnavailableFixed;

    /**
     * @return , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
     * 
     */
    public Optional<Output<Integer>> maxUnavailableFixed() {
        return Optional.ofNullable(this.maxUnavailableFixed);
    }

    /**
     * , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
     * 
     */
    @Import(name="maxUnavailablePercent")
    private @Nullable Output<Integer> maxUnavailablePercent;

    /**
     * @return , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
     * 
     */
    public Optional<Output<Integer>> maxUnavailablePercent() {
        return Optional.ofNullable(this.maxUnavailablePercent);
    }

    /**
     * , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
     * 
     */
    @Import(name="minReadySec")
    private @Nullable Output<Integer> minReadySec;

    /**
     * @return , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
     * 
     */
    public Optional<Output<Integer>> minReadySec() {
        return Optional.ofNullable(this.minReadySec);
    }

    /**
     * Minimal action to be taken on an instance. You can specify either `NONE` to forbid any actions, `REFRESH` to update without stopping instances, `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `REFRESH`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
     * 
     */
    @Import(name="minimalAction", required=true)
    private Output<String> minimalAction;

    /**
     * @return Minimal action to be taken on an instance. You can specify either `NONE` to forbid any actions, `REFRESH` to update without stopping instances, `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `REFRESH`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
     * 
     */
    public Output<String> minimalAction() {
        return this.minimalAction;
    }

    /**
     * Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
     * 
     */
    @Import(name="mostDisruptiveAllowedAction")
    private @Nullable Output<String> mostDisruptiveAllowedAction;

    /**
     * @return Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
     * 
     */
    public Optional<Output<String>> mostDisruptiveAllowedAction() {
        return Optional.ofNullable(this.mostDisruptiveAllowedAction);
    }

    /**
     * , The instance replacement method for managed instance groups. Valid values are: &#34;RECREATE&#34;, &#34;SUBSTITUTE&#34;. If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
     * ***
     * 
     */
    @Import(name="replacementMethod")
    private @Nullable Output<String> replacementMethod;

    /**
     * @return , The instance replacement method for managed instance groups. Valid values are: &#34;RECREATE&#34;, &#34;SUBSTITUTE&#34;. If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
     * ***
     * 
     */
    public Optional<Output<String>> replacementMethod() {
        return Optional.ofNullable(this.replacementMethod);
    }

    /**
     * The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private RegionInstanceGroupManagerUpdatePolicyArgs() {}

    private RegionInstanceGroupManagerUpdatePolicyArgs(RegionInstanceGroupManagerUpdatePolicyArgs $) {
        this.instanceRedistributionType = $.instanceRedistributionType;
        this.maxSurgeFixed = $.maxSurgeFixed;
        this.maxSurgePercent = $.maxSurgePercent;
        this.maxUnavailableFixed = $.maxUnavailableFixed;
        this.maxUnavailablePercent = $.maxUnavailablePercent;
        this.minReadySec = $.minReadySec;
        this.minimalAction = $.minimalAction;
        this.mostDisruptiveAllowedAction = $.mostDisruptiveAllowedAction;
        this.replacementMethod = $.replacementMethod;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionInstanceGroupManagerUpdatePolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionInstanceGroupManagerUpdatePolicyArgs $;

        public Builder() {
            $ = new RegionInstanceGroupManagerUpdatePolicyArgs();
        }

        public Builder(RegionInstanceGroupManagerUpdatePolicyArgs defaults) {
            $ = new RegionInstanceGroupManagerUpdatePolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceRedistributionType The instance redistribution policy for regional managed instance groups. Valid values are: `&#34;PROACTIVE&#34;`, `&#34;NONE&#34;`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
         * 
         * @return builder
         * 
         */
        public Builder instanceRedistributionType(@Nullable Output<String> instanceRedistributionType) {
            $.instanceRedistributionType = instanceRedistributionType;
            return this;
        }

        /**
         * @param instanceRedistributionType The instance redistribution policy for regional managed instance groups. Valid values are: `&#34;PROACTIVE&#34;`, `&#34;NONE&#34;`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
         * 
         * @return builder
         * 
         */
        public Builder instanceRedistributionType(String instanceRedistributionType) {
            return instanceRedistributionType(Output.of(instanceRedistributionType));
        }

        /**
         * @param maxSurgeFixed , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
         * 
         * @return builder
         * 
         */
        public Builder maxSurgeFixed(@Nullable Output<Integer> maxSurgeFixed) {
            $.maxSurgeFixed = maxSurgeFixed;
            return this;
        }

        /**
         * @param maxSurgeFixed , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
         * 
         * @return builder
         * 
         */
        public Builder maxSurgeFixed(Integer maxSurgeFixed) {
            return maxSurgeFixed(Output.of(maxSurgeFixed));
        }

        /**
         * @param maxSurgePercent , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
         * 
         * @return builder
         * 
         */
        public Builder maxSurgePercent(@Nullable Output<Integer> maxSurgePercent) {
            $.maxSurgePercent = maxSurgePercent;
            return this;
        }

        /**
         * @param maxSurgePercent , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
         * 
         * @return builder
         * 
         */
        public Builder maxSurgePercent(Integer maxSurgePercent) {
            return maxSurgePercent(Output.of(maxSurgePercent));
        }

        /**
         * @param maxUnavailableFixed , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
         * 
         * @return builder
         * 
         */
        public Builder maxUnavailableFixed(@Nullable Output<Integer> maxUnavailableFixed) {
            $.maxUnavailableFixed = maxUnavailableFixed;
            return this;
        }

        /**
         * @param maxUnavailableFixed , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
         * 
         * @return builder
         * 
         */
        public Builder maxUnavailableFixed(Integer maxUnavailableFixed) {
            return maxUnavailableFixed(Output.of(maxUnavailableFixed));
        }

        /**
         * @param maxUnavailablePercent , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
         * 
         * @return builder
         * 
         */
        public Builder maxUnavailablePercent(@Nullable Output<Integer> maxUnavailablePercent) {
            $.maxUnavailablePercent = maxUnavailablePercent;
            return this;
        }

        /**
         * @param maxUnavailablePercent , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
         * 
         * @return builder
         * 
         */
        public Builder maxUnavailablePercent(Integer maxUnavailablePercent) {
            return maxUnavailablePercent(Output.of(maxUnavailablePercent));
        }

        /**
         * @param minReadySec , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
         * 
         * @return builder
         * 
         */
        public Builder minReadySec(@Nullable Output<Integer> minReadySec) {
            $.minReadySec = minReadySec;
            return this;
        }

        /**
         * @param minReadySec , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
         * 
         * @return builder
         * 
         */
        public Builder minReadySec(Integer minReadySec) {
            return minReadySec(Output.of(minReadySec));
        }

        /**
         * @param minimalAction Minimal action to be taken on an instance. You can specify either `NONE` to forbid any actions, `REFRESH` to update without stopping instances, `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `REFRESH`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
         * 
         * @return builder
         * 
         */
        public Builder minimalAction(Output<String> minimalAction) {
            $.minimalAction = minimalAction;
            return this;
        }

        /**
         * @param minimalAction Minimal action to be taken on an instance. You can specify either `NONE` to forbid any actions, `REFRESH` to update without stopping instances, `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `REFRESH`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
         * 
         * @return builder
         * 
         */
        public Builder minimalAction(String minimalAction) {
            return minimalAction(Output.of(minimalAction));
        }

        /**
         * @param mostDisruptiveAllowedAction Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
         * 
         * @return builder
         * 
         */
        public Builder mostDisruptiveAllowedAction(@Nullable Output<String> mostDisruptiveAllowedAction) {
            $.mostDisruptiveAllowedAction = mostDisruptiveAllowedAction;
            return this;
        }

        /**
         * @param mostDisruptiveAllowedAction Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
         * 
         * @return builder
         * 
         */
        public Builder mostDisruptiveAllowedAction(String mostDisruptiveAllowedAction) {
            return mostDisruptiveAllowedAction(Output.of(mostDisruptiveAllowedAction));
        }

        /**
         * @param replacementMethod , The instance replacement method for managed instance groups. Valid values are: &#34;RECREATE&#34;, &#34;SUBSTITUTE&#34;. If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
         * ***
         * 
         * @return builder
         * 
         */
        public Builder replacementMethod(@Nullable Output<String> replacementMethod) {
            $.replacementMethod = replacementMethod;
            return this;
        }

        /**
         * @param replacementMethod , The instance replacement method for managed instance groups. Valid values are: &#34;RECREATE&#34;, &#34;SUBSTITUTE&#34;. If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
         * ***
         * 
         * @return builder
         * 
         */
        public Builder replacementMethod(String replacementMethod) {
            return replacementMethod(Output.of(replacementMethod));
        }

        /**
         * @param type The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public RegionInstanceGroupManagerUpdatePolicyArgs build() {
            if ($.minimalAction == null) {
                throw new MissingRequiredPropertyException("RegionInstanceGroupManagerUpdatePolicyArgs", "minimalAction");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("RegionInstanceGroupManagerUpdatePolicyArgs", "type");
            }
            return $;
        }
    }

}
