// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebaserules.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RulesetMetadata {
    private final @Nullable List<String> services;

    @CustomType.Constructor
    private RulesetMetadata(@CustomType.Parameter("services") @Nullable List<String> services) {
        this.services = services;
    }

    public List<String> services() {
        return this.services == null ? List.of() : this.services;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RulesetMetadata defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> services;

        public Builder() {
    	      // Empty
        }

        public Builder(RulesetMetadata defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.services = defaults.services;
        }

        public Builder services(@Nullable List<String> services) {
            this.services = services;
            return this;
        }
        public Builder services(String... services) {
            return services(List.of(services));
        }        public RulesetMetadata build() {
            return new RulesetMetadata(services);
        }
    }
}
