// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class ClusterNodePoolUpgradeSettings {
    private final Integer maxSurge;
    private final Integer maxUnavailable;

    @CustomType.Constructor
    private ClusterNodePoolUpgradeSettings(
        @CustomType.Parameter("maxSurge") Integer maxSurge,
        @CustomType.Parameter("maxUnavailable") Integer maxUnavailable) {
        this.maxSurge = maxSurge;
        this.maxUnavailable = maxUnavailable;
    }

    public Integer maxSurge() {
        return this.maxSurge;
    }
    public Integer maxUnavailable() {
        return this.maxUnavailable;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterNodePoolUpgradeSettings defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Integer maxSurge;
        private Integer maxUnavailable;

        public Builder() {
    	      // Empty
        }

        public Builder(ClusterNodePoolUpgradeSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxSurge = defaults.maxSurge;
    	      this.maxUnavailable = defaults.maxUnavailable;
        }

        public Builder maxSurge(Integer maxSurge) {
            this.maxSurge = Objects.requireNonNull(maxSurge);
            return this;
        }
        public Builder maxUnavailable(Integer maxUnavailable) {
            this.maxUnavailable = Objects.requireNonNull(maxUnavailable);
            return this;
        }        public ClusterNodePoolUpgradeSettings build() {
            return new ClusterNodePoolUpgradeSettings(maxSurge, maxUnavailable);
        }
    }
}
