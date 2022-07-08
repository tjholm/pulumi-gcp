// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.clouddeploy.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DeliveryPipelineConditionTargetsPresentCondition {
    private final @Nullable List<String> missingTargets;
    private final @Nullable Boolean status;
    private final @Nullable String updateTime;

    @CustomType.Constructor
    private DeliveryPipelineConditionTargetsPresentCondition(
        @CustomType.Parameter("missingTargets") @Nullable List<String> missingTargets,
        @CustomType.Parameter("status") @Nullable Boolean status,
        @CustomType.Parameter("updateTime") @Nullable String updateTime) {
        this.missingTargets = missingTargets;
        this.status = status;
        this.updateTime = updateTime;
    }

    public List<String> missingTargets() {
        return this.missingTargets == null ? List.of() : this.missingTargets;
    }
    public Optional<Boolean> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DeliveryPipelineConditionTargetsPresentCondition defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> missingTargets;
        private @Nullable Boolean status;
        private @Nullable String updateTime;

        public Builder() {
    	      // Empty
        }

        public Builder(DeliveryPipelineConditionTargetsPresentCondition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.missingTargets = defaults.missingTargets;
    	      this.status = defaults.status;
    	      this.updateTime = defaults.updateTime;
        }

        public Builder missingTargets(@Nullable List<String> missingTargets) {
            this.missingTargets = missingTargets;
            return this;
        }
        public Builder missingTargets(String... missingTargets) {
            return missingTargets(List.of(missingTargets));
        }
        public Builder status(@Nullable Boolean status) {
            this.status = status;
            return this;
        }
        public Builder updateTime(@Nullable String updateTime) {
            this.updateTime = updateTime;
            return this;
        }        public DeliveryPipelineConditionTargetsPresentCondition build() {
            return new DeliveryPipelineConditionTargetsPresentCondition(missingTargets, status, updateTime);
        }
    }
}
