// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceFromMachineImageBootDiskInitializeParams {
    private final @Nullable String image;
    private final @Nullable Map<String,Object> labels;
    private final @Nullable Integer size;
    private final @Nullable String type;

    @CustomType.Constructor
    private InstanceFromMachineImageBootDiskInitializeParams(
        @CustomType.Parameter("image") @Nullable String image,
        @CustomType.Parameter("labels") @Nullable Map<String,Object> labels,
        @CustomType.Parameter("size") @Nullable Integer size,
        @CustomType.Parameter("type") @Nullable String type) {
        this.image = image;
        this.labels = labels;
        this.size = size;
        this.type = type;
    }

    public Optional<String> image() {
        return Optional.ofNullable(this.image);
    }
    public Map<String,Object> labels() {
        return this.labels == null ? Map.of() : this.labels;
    }
    public Optional<Integer> size() {
        return Optional.ofNullable(this.size);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceFromMachineImageBootDiskInitializeParams defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String image;
        private @Nullable Map<String,Object> labels;
        private @Nullable Integer size;
        private @Nullable String type;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceFromMachineImageBootDiskInitializeParams defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.image = defaults.image;
    	      this.labels = defaults.labels;
    	      this.size = defaults.size;
    	      this.type = defaults.type;
        }

        public Builder image(@Nullable String image) {
            this.image = image;
            return this;
        }
        public Builder labels(@Nullable Map<String,Object> labels) {
            this.labels = labels;
            return this;
        }
        public Builder size(@Nullable Integer size) {
            this.size = size;
            return this;
        }
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }        public InstanceFromMachineImageBootDiskInitializeParams build() {
            return new InstanceFromMachineImageBootDiskInitializeParams(image, labels, size, type);
        }
    }
}
