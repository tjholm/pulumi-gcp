// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.gkeonprem.outputs.VMwareClusterStatusCondition;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VMwareClusterStatus {
    /**
     * @return (Output)
     * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
     * Structure is documented below.
     * 
     */
    private @Nullable List<VMwareClusterStatusCondition> conditions;
    /**
     * @return (Output)
     * Human-friendly representation of the error message from the user cluster
     * controller. The error message can be temporary as the user cluster
     * controller creates a cluster or node pool. If the error message persists
     * for a longer period of time, it can be used to surface error message to
     * indicate real problems requiring user intervention.
     * 
     */
    private @Nullable String errorMessage;

    private VMwareClusterStatus() {}
    /**
     * @return (Output)
     * ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
     * Structure is documented below.
     * 
     */
    public List<VMwareClusterStatusCondition> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    /**
     * @return (Output)
     * Human-friendly representation of the error message from the user cluster
     * controller. The error message can be temporary as the user cluster
     * controller creates a cluster or node pool. If the error message persists
     * for a longer period of time, it can be used to surface error message to
     * indicate real problems requiring user intervention.
     * 
     */
    public Optional<String> errorMessage() {
        return Optional.ofNullable(this.errorMessage);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VMwareClusterStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<VMwareClusterStatusCondition> conditions;
        private @Nullable String errorMessage;
        public Builder() {}
        public Builder(VMwareClusterStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.conditions = defaults.conditions;
    	      this.errorMessage = defaults.errorMessage;
        }

        @CustomType.Setter
        public Builder conditions(@Nullable List<VMwareClusterStatusCondition> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(VMwareClusterStatusCondition... conditions) {
            return conditions(List.of(conditions));
        }
        @CustomType.Setter
        public Builder errorMessage(@Nullable String errorMessage) {
            this.errorMessage = errorMessage;
            return this;
        }
        public VMwareClusterStatus build() {
            final var o = new VMwareClusterStatus();
            o.conditions = conditions;
            o.errorMessage = errorMessage;
            return o;
        }
    }
}
