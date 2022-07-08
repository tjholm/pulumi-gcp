// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RuntimeMetric {
    private final @Nullable Map<String,String> systemMetrics;

    @CustomType.Constructor
    private RuntimeMetric(@CustomType.Parameter("systemMetrics") @Nullable Map<String,String> systemMetrics) {
        this.systemMetrics = systemMetrics;
    }

    public Map<String,String> systemMetrics() {
        return this.systemMetrics == null ? Map.of() : this.systemMetrics;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuntimeMetric defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Map<String,String> systemMetrics;

        public Builder() {
    	      // Empty
        }

        public Builder(RuntimeMetric defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.systemMetrics = defaults.systemMetrics;
        }

        public Builder systemMetrics(@Nullable Map<String,String> systemMetrics) {
            this.systemMetrics = systemMetrics;
            return this;
        }        public RuntimeMetric build() {
            return new RuntimeMetric(systemMetrics);
        }
    }
}
