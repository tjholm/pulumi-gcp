// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.InstanceGroupManagerStatusAllInstancesConfigArgs;
import com.pulumi.gcp.compute.inputs.InstanceGroupManagerStatusStatefulArgs;
import com.pulumi.gcp.compute.inputs.InstanceGroupManagerStatusVersionTargetArgs;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceGroupManagerStatusArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceGroupManagerStatusArgs Empty = new InstanceGroupManagerStatusArgs();

    /**
     * Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    @Import(name="allInstancesConfigs")
    private @Nullable Output<List<InstanceGroupManagerStatusAllInstancesConfigArgs>> allInstancesConfigs;

    /**
     * @return Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    public Optional<Output<List<InstanceGroupManagerStatusAllInstancesConfigArgs>>> allInstancesConfigs() {
        return Optional.ofNullable(this.allInstancesConfigs);
    }

    /**
     * A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
     * 
     */
    @Import(name="isStable")
    private @Nullable Output<Boolean> isStable;

    /**
     * @return A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
     * 
     */
    public Optional<Output<Boolean>> isStable() {
        return Optional.ofNullable(this.isStable);
    }

    /**
     * Stateful status of the given Instance Group Manager.
     * 
     */
    @Import(name="statefuls")
    private @Nullable Output<List<InstanceGroupManagerStatusStatefulArgs>> statefuls;

    /**
     * @return Stateful status of the given Instance Group Manager.
     * 
     */
    public Optional<Output<List<InstanceGroupManagerStatusStatefulArgs>>> statefuls() {
        return Optional.ofNullable(this.statefuls);
    }

    /**
     * A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
     * 
     */
    @Import(name="versionTargets")
    private @Nullable Output<List<InstanceGroupManagerStatusVersionTargetArgs>> versionTargets;

    /**
     * @return A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
     * 
     */
    public Optional<Output<List<InstanceGroupManagerStatusVersionTargetArgs>>> versionTargets() {
        return Optional.ofNullable(this.versionTargets);
    }

    private InstanceGroupManagerStatusArgs() {}

    private InstanceGroupManagerStatusArgs(InstanceGroupManagerStatusArgs $) {
        this.allInstancesConfigs = $.allInstancesConfigs;
        this.isStable = $.isStable;
        this.statefuls = $.statefuls;
        this.versionTargets = $.versionTargets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceGroupManagerStatusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceGroupManagerStatusArgs $;

        public Builder() {
            $ = new InstanceGroupManagerStatusArgs();
        }

        public Builder(InstanceGroupManagerStatusArgs defaults) {
            $ = new InstanceGroupManagerStatusArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allInstancesConfigs Properties to set on all instances in the group. After setting
         * allInstancesConfig on the group, you must update the group&#39;s instances to
         * apply the configuration.
         * 
         * @return builder
         * 
         */
        public Builder allInstancesConfigs(@Nullable Output<List<InstanceGroupManagerStatusAllInstancesConfigArgs>> allInstancesConfigs) {
            $.allInstancesConfigs = allInstancesConfigs;
            return this;
        }

        /**
         * @param allInstancesConfigs Properties to set on all instances in the group. After setting
         * allInstancesConfig on the group, you must update the group&#39;s instances to
         * apply the configuration.
         * 
         * @return builder
         * 
         */
        public Builder allInstancesConfigs(List<InstanceGroupManagerStatusAllInstancesConfigArgs> allInstancesConfigs) {
            return allInstancesConfigs(Output.of(allInstancesConfigs));
        }

        /**
         * @param allInstancesConfigs Properties to set on all instances in the group. After setting
         * allInstancesConfig on the group, you must update the group&#39;s instances to
         * apply the configuration.
         * 
         * @return builder
         * 
         */
        public Builder allInstancesConfigs(InstanceGroupManagerStatusAllInstancesConfigArgs... allInstancesConfigs) {
            return allInstancesConfigs(List.of(allInstancesConfigs));
        }

        /**
         * @param isStable A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
         * 
         * @return builder
         * 
         */
        public Builder isStable(@Nullable Output<Boolean> isStable) {
            $.isStable = isStable;
            return this;
        }

        /**
         * @param isStable A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
         * 
         * @return builder
         * 
         */
        public Builder isStable(Boolean isStable) {
            return isStable(Output.of(isStable));
        }

        /**
         * @param statefuls Stateful status of the given Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder statefuls(@Nullable Output<List<InstanceGroupManagerStatusStatefulArgs>> statefuls) {
            $.statefuls = statefuls;
            return this;
        }

        /**
         * @param statefuls Stateful status of the given Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder statefuls(List<InstanceGroupManagerStatusStatefulArgs> statefuls) {
            return statefuls(Output.of(statefuls));
        }

        /**
         * @param statefuls Stateful status of the given Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder statefuls(InstanceGroupManagerStatusStatefulArgs... statefuls) {
            return statefuls(List.of(statefuls));
        }

        /**
         * @param versionTargets A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder versionTargets(@Nullable Output<List<InstanceGroupManagerStatusVersionTargetArgs>> versionTargets) {
            $.versionTargets = versionTargets;
            return this;
        }

        /**
         * @param versionTargets A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder versionTargets(List<InstanceGroupManagerStatusVersionTargetArgs> versionTargets) {
            return versionTargets(Output.of(versionTargets));
        }

        /**
         * @param versionTargets A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
         * 
         * @return builder
         * 
         */
        public Builder versionTargets(InstanceGroupManagerStatusVersionTargetArgs... versionTargets) {
            return versionTargets(List.of(versionTargets));
        }

        public InstanceGroupManagerStatusArgs build() {
            return $;
        }
    }

}
