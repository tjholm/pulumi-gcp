// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkconnectivity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HubRoutingVpc {
    private final @Nullable String uri;

    @CustomType.Constructor
    private HubRoutingVpc(@CustomType.Parameter("uri") @Nullable String uri) {
        this.uri = uri;
    }

    public Optional<String> uri() {
        return Optional.ofNullable(this.uri);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HubRoutingVpc defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String uri;

        public Builder() {
    	      // Empty
        }

        public Builder(HubRoutingVpc defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.uri = defaults.uri;
        }

        public Builder uri(@Nullable String uri) {
            this.uri = uri;
            return this;
        }        public HubRoutingVpc build() {
            return new HubRoutingVpc(uri);
        }
    }
}
