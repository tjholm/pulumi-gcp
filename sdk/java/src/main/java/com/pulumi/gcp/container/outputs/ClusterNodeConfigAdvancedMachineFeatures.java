// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class ClusterNodeConfigAdvancedMachineFeatures {
    /**
     * @return The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
     * 
     */
    private Integer threadsPerCore;

    private ClusterNodeConfigAdvancedMachineFeatures() {}
    /**
     * @return The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
     * 
     */
    public Integer threadsPerCore() {
        return this.threadsPerCore;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterNodeConfigAdvancedMachineFeatures defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer threadsPerCore;
        public Builder() {}
        public Builder(ClusterNodeConfigAdvancedMachineFeatures defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.threadsPerCore = defaults.threadsPerCore;
        }

        @CustomType.Setter
        public Builder threadsPerCore(Integer threadsPerCore) {
            this.threadsPerCore = Objects.requireNonNull(threadsPerCore);
            return this;
        }
        public ClusterNodeConfigAdvancedMachineFeatures build() {
            final var o = new ClusterNodeConfigAdvancedMachineFeatures();
            o.threadsPerCore = threadsPerCore;
            return o;
        }
    }
}
