// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceFromMachineImageShieldedInstanceConfig {
    private final @Nullable Boolean enableIntegrityMonitoring;
    private final @Nullable Boolean enableSecureBoot;
    private final @Nullable Boolean enableVtpm;

    @CustomType.Constructor
    private InstanceFromMachineImageShieldedInstanceConfig(
        @CustomType.Parameter("enableIntegrityMonitoring") @Nullable Boolean enableIntegrityMonitoring,
        @CustomType.Parameter("enableSecureBoot") @Nullable Boolean enableSecureBoot,
        @CustomType.Parameter("enableVtpm") @Nullable Boolean enableVtpm) {
        this.enableIntegrityMonitoring = enableIntegrityMonitoring;
        this.enableSecureBoot = enableSecureBoot;
        this.enableVtpm = enableVtpm;
    }

    public Optional<Boolean> enableIntegrityMonitoring() {
        return Optional.ofNullable(this.enableIntegrityMonitoring);
    }
    public Optional<Boolean> enableSecureBoot() {
        return Optional.ofNullable(this.enableSecureBoot);
    }
    public Optional<Boolean> enableVtpm() {
        return Optional.ofNullable(this.enableVtpm);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceFromMachineImageShieldedInstanceConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean enableIntegrityMonitoring;
        private @Nullable Boolean enableSecureBoot;
        private @Nullable Boolean enableVtpm;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceFromMachineImageShieldedInstanceConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableIntegrityMonitoring = defaults.enableIntegrityMonitoring;
    	      this.enableSecureBoot = defaults.enableSecureBoot;
    	      this.enableVtpm = defaults.enableVtpm;
        }

        public Builder enableIntegrityMonitoring(@Nullable Boolean enableIntegrityMonitoring) {
            this.enableIntegrityMonitoring = enableIntegrityMonitoring;
            return this;
        }
        public Builder enableSecureBoot(@Nullable Boolean enableSecureBoot) {
            this.enableSecureBoot = enableSecureBoot;
            return this;
        }
        public Builder enableVtpm(@Nullable Boolean enableVtpm) {
            this.enableVtpm = enableVtpm;
            return this;
        }        public InstanceFromMachineImageShieldedInstanceConfig build() {
            return new InstanceFromMachineImageShieldedInstanceConfig(enableIntegrityMonitoring, enableSecureBoot, enableVtpm);
        }
    }
}
