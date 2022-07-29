// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class InstanceGroupManagerAllInstancesConfig {
    /**
     * @return ), The label key-value pairs that you want to patch onto the instance.
     * 
     */
    private final @Nullable Map<String,String> labels;
    /**
     * @return ), The metadata key-value pairs that you want to patch onto the instance. For more information, see [Project and instance metadata](https://cloud.google.com/compute/docs/metadata#project_and_instance_metadata).
     * 
     */
    private final @Nullable Map<String,String> metadata;

    @CustomType.Constructor
    private InstanceGroupManagerAllInstancesConfig(
        @CustomType.Parameter("labels") @Nullable Map<String,String> labels,
        @CustomType.Parameter("metadata") @Nullable Map<String,String> metadata) {
        this.labels = labels;
        this.metadata = metadata;
    }

    /**
     * @return ), The label key-value pairs that you want to patch onto the instance.
     * 
     */
    public Map<String,String> labels() {
        return this.labels == null ? Map.of() : this.labels;
    }
    /**
     * @return ), The metadata key-value pairs that you want to patch onto the instance. For more information, see [Project and instance metadata](https://cloud.google.com/compute/docs/metadata#project_and_instance_metadata).
     * 
     */
    public Map<String,String> metadata() {
        return this.metadata == null ? Map.of() : this.metadata;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceGroupManagerAllInstancesConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Map<String,String> labels;
        private @Nullable Map<String,String> metadata;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceGroupManagerAllInstancesConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labels = defaults.labels;
    	      this.metadata = defaults.metadata;
        }

        public Builder labels(@Nullable Map<String,String> labels) {
            this.labels = labels;
            return this;
        }
        public Builder metadata(@Nullable Map<String,String> metadata) {
            this.metadata = metadata;
            return this;
        }        public InstanceGroupManagerAllInstancesConfig build() {
            return new InstanceGroupManagerAllInstancesConfig(labels, metadata);
        }
    }
}
