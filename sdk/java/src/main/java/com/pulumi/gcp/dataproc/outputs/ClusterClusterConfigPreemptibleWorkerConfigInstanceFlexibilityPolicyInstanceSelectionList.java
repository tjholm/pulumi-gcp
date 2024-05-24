// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList {
    /**
     * @return Full machine-type names, e.g. `&#34;n1-standard-16&#34;`.
     * 
     */
    private @Nullable List<String> machineTypes;
    /**
     * @return Preference of this instance selection. A lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
     * 
     * ***
     * 
     */
    private @Nullable Integer rank;

    private ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList() {}
    /**
     * @return Full machine-type names, e.g. `&#34;n1-standard-16&#34;`.
     * 
     */
    public List<String> machineTypes() {
        return this.machineTypes == null ? List.of() : this.machineTypes;
    }
    /**
     * @return Preference of this instance selection. A lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
     * 
     * ***
     * 
     */
    public Optional<Integer> rank() {
        return Optional.ofNullable(this.rank);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> machineTypes;
        private @Nullable Integer rank;
        public Builder() {}
        public Builder(ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.machineTypes = defaults.machineTypes;
    	      this.rank = defaults.rank;
        }

        @CustomType.Setter
        public Builder machineTypes(@Nullable List<String> machineTypes) {

            this.machineTypes = machineTypes;
            return this;
        }
        public Builder machineTypes(String... machineTypes) {
            return machineTypes(List.of(machineTypes));
        }
        @CustomType.Setter
        public Builder rank(@Nullable Integer rank) {

            this.rank = rank;
            return this;
        }
        public ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList build() {
            final var _resultValue = new ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList();
            _resultValue.machineTypes = machineTypes;
            _resultValue.rank = rank;
            return _resultValue;
        }
    }
}
