// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.composer.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EnvironmentConfigWorkloadsConfigWorker {
    private final @Nullable Double cpu;
    private final @Nullable Integer maxCount;
    private final @Nullable Double memoryGb;
    private final @Nullable Integer minCount;
    private final @Nullable Double storageGb;

    @CustomType.Constructor
    private EnvironmentConfigWorkloadsConfigWorker(
        @CustomType.Parameter("cpu") @Nullable Double cpu,
        @CustomType.Parameter("maxCount") @Nullable Integer maxCount,
        @CustomType.Parameter("memoryGb") @Nullable Double memoryGb,
        @CustomType.Parameter("minCount") @Nullable Integer minCount,
        @CustomType.Parameter("storageGb") @Nullable Double storageGb) {
        this.cpu = cpu;
        this.maxCount = maxCount;
        this.memoryGb = memoryGb;
        this.minCount = minCount;
        this.storageGb = storageGb;
    }

    public Optional<Double> cpu() {
        return Optional.ofNullable(this.cpu);
    }
    public Optional<Integer> maxCount() {
        return Optional.ofNullable(this.maxCount);
    }
    public Optional<Double> memoryGb() {
        return Optional.ofNullable(this.memoryGb);
    }
    public Optional<Integer> minCount() {
        return Optional.ofNullable(this.minCount);
    }
    public Optional<Double> storageGb() {
        return Optional.ofNullable(this.storageGb);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EnvironmentConfigWorkloadsConfigWorker defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Double cpu;
        private @Nullable Integer maxCount;
        private @Nullable Double memoryGb;
        private @Nullable Integer minCount;
        private @Nullable Double storageGb;

        public Builder() {
    	      // Empty
        }

        public Builder(EnvironmentConfigWorkloadsConfigWorker defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cpu = defaults.cpu;
    	      this.maxCount = defaults.maxCount;
    	      this.memoryGb = defaults.memoryGb;
    	      this.minCount = defaults.minCount;
    	      this.storageGb = defaults.storageGb;
        }

        public Builder cpu(@Nullable Double cpu) {
            this.cpu = cpu;
            return this;
        }
        public Builder maxCount(@Nullable Integer maxCount) {
            this.maxCount = maxCount;
            return this;
        }
        public Builder memoryGb(@Nullable Double memoryGb) {
            this.memoryGb = memoryGb;
            return this;
        }
        public Builder minCount(@Nullable Integer minCount) {
            this.minCount = minCount;
            return this;
        }
        public Builder storageGb(@Nullable Double storageGb) {
            this.storageGb = storageGb;
            return this;
        }        public EnvironmentConfigWorkloadsConfigWorker build() {
            return new EnvironmentConfigWorkloadsConfigWorker(cpu, maxCount, memoryGb, minCount, storageGb);
        }
    }
}
