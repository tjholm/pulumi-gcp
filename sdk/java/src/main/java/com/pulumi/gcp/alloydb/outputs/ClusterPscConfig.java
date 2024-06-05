// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterPscConfig {
    /**
     * @return Create an instance that allows connections from Private Service Connect endpoints to the instance.
     * 
     */
    private @Nullable Boolean pscEnabled;

    private ClusterPscConfig() {}
    /**
     * @return Create an instance that allows connections from Private Service Connect endpoints to the instance.
     * 
     */
    public Optional<Boolean> pscEnabled() {
        return Optional.ofNullable(this.pscEnabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterPscConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean pscEnabled;
        public Builder() {}
        public Builder(ClusterPscConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pscEnabled = defaults.pscEnabled;
        }

        @CustomType.Setter
        public Builder pscEnabled(@Nullable Boolean pscEnabled) {

            this.pscEnabled = pscEnabled;
            return this;
        }
        public ClusterPscConfig build() {
            final var _resultValue = new ClusterPscConfig();
            _resultValue.pscEnabled = pscEnabled;
            return _resultValue;
        }
    }
}
