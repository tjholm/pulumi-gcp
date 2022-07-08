// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AlertPolicyCreationRecord {
    private final @Nullable String mutateTime;
    private final @Nullable String mutatedBy;

    @CustomType.Constructor
    private AlertPolicyCreationRecord(
        @CustomType.Parameter("mutateTime") @Nullable String mutateTime,
        @CustomType.Parameter("mutatedBy") @Nullable String mutatedBy) {
        this.mutateTime = mutateTime;
        this.mutatedBy = mutatedBy;
    }

    public Optional<String> mutateTime() {
        return Optional.ofNullable(this.mutateTime);
    }
    public Optional<String> mutatedBy() {
        return Optional.ofNullable(this.mutatedBy);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AlertPolicyCreationRecord defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String mutateTime;
        private @Nullable String mutatedBy;

        public Builder() {
    	      // Empty
        }

        public Builder(AlertPolicyCreationRecord defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.mutateTime = defaults.mutateTime;
    	      this.mutatedBy = defaults.mutatedBy;
        }

        public Builder mutateTime(@Nullable String mutateTime) {
            this.mutateTime = mutateTime;
            return this;
        }
        public Builder mutatedBy(@Nullable String mutatedBy) {
            this.mutatedBy = mutatedBy;
            return this;
        }        public AlertPolicyCreationRecord build() {
            return new AlertPolicyCreationRecord(mutateTime, mutatedBy);
        }
    }
}
