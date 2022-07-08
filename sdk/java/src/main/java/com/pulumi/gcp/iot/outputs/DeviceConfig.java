// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iot.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DeviceConfig {
    private final @Nullable String binaryData;
    private final @Nullable String cloudUpdateTime;
    private final @Nullable String deviceAckTime;
    private final @Nullable String version;

    @CustomType.Constructor
    private DeviceConfig(
        @CustomType.Parameter("binaryData") @Nullable String binaryData,
        @CustomType.Parameter("cloudUpdateTime") @Nullable String cloudUpdateTime,
        @CustomType.Parameter("deviceAckTime") @Nullable String deviceAckTime,
        @CustomType.Parameter("version") @Nullable String version) {
        this.binaryData = binaryData;
        this.cloudUpdateTime = cloudUpdateTime;
        this.deviceAckTime = deviceAckTime;
        this.version = version;
    }

    public Optional<String> binaryData() {
        return Optional.ofNullable(this.binaryData);
    }
    public Optional<String> cloudUpdateTime() {
        return Optional.ofNullable(this.cloudUpdateTime);
    }
    public Optional<String> deviceAckTime() {
        return Optional.ofNullable(this.deviceAckTime);
    }
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DeviceConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String binaryData;
        private @Nullable String cloudUpdateTime;
        private @Nullable String deviceAckTime;
        private @Nullable String version;

        public Builder() {
    	      // Empty
        }

        public Builder(DeviceConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.binaryData = defaults.binaryData;
    	      this.cloudUpdateTime = defaults.cloudUpdateTime;
    	      this.deviceAckTime = defaults.deviceAckTime;
    	      this.version = defaults.version;
        }

        public Builder binaryData(@Nullable String binaryData) {
            this.binaryData = binaryData;
            return this;
        }
        public Builder cloudUpdateTime(@Nullable String cloudUpdateTime) {
            this.cloudUpdateTime = cloudUpdateTime;
            return this;
        }
        public Builder deviceAckTime(@Nullable String deviceAckTime) {
            this.deviceAckTime = deviceAckTime;
            return this;
        }
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }        public DeviceConfig build() {
            return new DeviceConfig(binaryData, cloudUpdateTime, deviceAckTime, version);
        }
    }
}
