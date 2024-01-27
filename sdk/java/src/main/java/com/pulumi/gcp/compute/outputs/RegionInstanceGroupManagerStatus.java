// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.RegionInstanceGroupManagerStatusAllInstancesConfig;
import com.pulumi.gcp.compute.outputs.RegionInstanceGroupManagerStatusStateful;
import com.pulumi.gcp.compute.outputs.RegionInstanceGroupManagerStatusVersionTarget;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegionInstanceGroupManagerStatus {
    /**
     * @return Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    private @Nullable List<RegionInstanceGroupManagerStatusAllInstancesConfig> allInstancesConfigs;
    /**
     * @return A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
     * 
     */
    private @Nullable Boolean isStable;
    /**
     * @return Stateful status of the given Instance Group Manager.
     * 
     */
    private @Nullable List<RegionInstanceGroupManagerStatusStateful> statefuls;
    /**
     * @return A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
     * 
     */
    private @Nullable List<RegionInstanceGroupManagerStatusVersionTarget> versionTargets;

    private RegionInstanceGroupManagerStatus() {}
    /**
     * @return Properties to set on all instances in the group. After setting
     * allInstancesConfig on the group, you must update the group&#39;s instances to
     * apply the configuration.
     * 
     */
    public List<RegionInstanceGroupManagerStatusAllInstancesConfig> allInstancesConfigs() {
        return this.allInstancesConfigs == null ? List.of() : this.allInstancesConfigs;
    }
    /**
     * @return A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
     * 
     */
    public Optional<Boolean> isStable() {
        return Optional.ofNullable(this.isStable);
    }
    /**
     * @return Stateful status of the given Instance Group Manager.
     * 
     */
    public List<RegionInstanceGroupManagerStatusStateful> statefuls() {
        return this.statefuls == null ? List.of() : this.statefuls;
    }
    /**
     * @return A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances&#39; target version are specified by version field on Instance Group Manager.
     * 
     */
    public List<RegionInstanceGroupManagerStatusVersionTarget> versionTargets() {
        return this.versionTargets == null ? List.of() : this.versionTargets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionInstanceGroupManagerStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<RegionInstanceGroupManagerStatusAllInstancesConfig> allInstancesConfigs;
        private @Nullable Boolean isStable;
        private @Nullable List<RegionInstanceGroupManagerStatusStateful> statefuls;
        private @Nullable List<RegionInstanceGroupManagerStatusVersionTarget> versionTargets;
        public Builder() {}
        public Builder(RegionInstanceGroupManagerStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allInstancesConfigs = defaults.allInstancesConfigs;
    	      this.isStable = defaults.isStable;
    	      this.statefuls = defaults.statefuls;
    	      this.versionTargets = defaults.versionTargets;
        }

        @CustomType.Setter
        public Builder allInstancesConfigs(@Nullable List<RegionInstanceGroupManagerStatusAllInstancesConfig> allInstancesConfigs) {

            this.allInstancesConfigs = allInstancesConfigs;
            return this;
        }
        public Builder allInstancesConfigs(RegionInstanceGroupManagerStatusAllInstancesConfig... allInstancesConfigs) {
            return allInstancesConfigs(List.of(allInstancesConfigs));
        }
        @CustomType.Setter
        public Builder isStable(@Nullable Boolean isStable) {

            this.isStable = isStable;
            return this;
        }
        @CustomType.Setter
        public Builder statefuls(@Nullable List<RegionInstanceGroupManagerStatusStateful> statefuls) {

            this.statefuls = statefuls;
            return this;
        }
        public Builder statefuls(RegionInstanceGroupManagerStatusStateful... statefuls) {
            return statefuls(List.of(statefuls));
        }
        @CustomType.Setter
        public Builder versionTargets(@Nullable List<RegionInstanceGroupManagerStatusVersionTarget> versionTargets) {

            this.versionTargets = versionTargets;
            return this;
        }
        public Builder versionTargets(RegionInstanceGroupManagerStatusVersionTarget... versionTargets) {
            return versionTargets(List.of(versionTargets));
        }
        public RegionInstanceGroupManagerStatus build() {
            final var _resultValue = new RegionInstanceGroupManagerStatus();
            _resultValue.allInstancesConfigs = allInstancesConfigs;
            _resultValue.isStable = isStable;
            _resultValue.statefuls = statefuls;
            _resultValue.versionTargets = versionTargets;
            return _resultValue;
        }
    }
}
