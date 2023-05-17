// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkeonprem.inputs.VMwareNodePoolStatusConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VMwareNodePoolStatusArgs extends com.pulumi.resources.ResourceArgs {

    public static final VMwareNodePoolStatusArgs Empty = new VMwareNodePoolStatusArgs();

    /**
     * (Output)
     * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
     * Structure is documented below.
     * 
     */
    @Import(name="conditions")
    private @Nullable Output<List<VMwareNodePoolStatusConditionArgs>> conditions;

    /**
     * @return (Output)
     * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<VMwareNodePoolStatusConditionArgs>>> conditions() {
        return Optional.ofNullable(this.conditions);
    }

    /**
     * (Output)
     * Human-friendly representation of the error message from the user cluster
     * controller. The error message can be temporary as the user cluster
     * controller creates a cluster or node pool. If the error message persists
     * for a longer period of time, it can be used to surface error message to
     * indicate real problems requiring user intervention.
     * 
     */
    @Import(name="errorMessage")
    private @Nullable Output<String> errorMessage;

    /**
     * @return (Output)
     * Human-friendly representation of the error message from the user cluster
     * controller. The error message can be temporary as the user cluster
     * controller creates a cluster or node pool. If the error message persists
     * for a longer period of time, it can be used to surface error message to
     * indicate real problems requiring user intervention.
     * 
     */
    public Optional<Output<String>> errorMessage() {
        return Optional.ofNullable(this.errorMessage);
    }

    private VMwareNodePoolStatusArgs() {}

    private VMwareNodePoolStatusArgs(VMwareNodePoolStatusArgs $) {
        this.conditions = $.conditions;
        this.errorMessage = $.errorMessage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VMwareNodePoolStatusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VMwareNodePoolStatusArgs $;

        public Builder() {
            $ = new VMwareNodePoolStatusArgs();
        }

        public Builder(VMwareNodePoolStatusArgs defaults) {
            $ = new VMwareNodePoolStatusArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param conditions (Output)
         * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder conditions(@Nullable Output<List<VMwareNodePoolStatusConditionArgs>> conditions) {
            $.conditions = conditions;
            return this;
        }

        /**
         * @param conditions (Output)
         * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder conditions(List<VMwareNodePoolStatusConditionArgs> conditions) {
            return conditions(Output.of(conditions));
        }

        /**
         * @param conditions (Output)
         * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder conditions(VMwareNodePoolStatusConditionArgs... conditions) {
            return conditions(List.of(conditions));
        }

        /**
         * @param errorMessage (Output)
         * Human-friendly representation of the error message from the user cluster
         * controller. The error message can be temporary as the user cluster
         * controller creates a cluster or node pool. If the error message persists
         * for a longer period of time, it can be used to surface error message to
         * indicate real problems requiring user intervention.
         * 
         * @return builder
         * 
         */
        public Builder errorMessage(@Nullable Output<String> errorMessage) {
            $.errorMessage = errorMessage;
            return this;
        }

        /**
         * @param errorMessage (Output)
         * Human-friendly representation of the error message from the user cluster
         * controller. The error message can be temporary as the user cluster
         * controller creates a cluster or node pool. If the error message persists
         * for a longer period of time, it can be used to surface error message to
         * indicate real problems requiring user intervention.
         * 
         * @return builder
         * 
         */
        public Builder errorMessage(String errorMessage) {
            return errorMessage(Output.of(errorMessage));
        }

        public VMwareNodePoolStatusArgs build() {
            return $;
        }
    }

}
