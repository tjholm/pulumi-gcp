// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RegionBackendServiceSubsetting {
    /**
     * @return The algorithm used for subsetting.
     * Possible values are: `CONSISTENT_HASH_SUBSETTING`.
     * 
     */
    private String policy;

    private RegionBackendServiceSubsetting() {}
    /**
     * @return The algorithm used for subsetting.
     * Possible values are: `CONSISTENT_HASH_SUBSETTING`.
     * 
     */
    public String policy() {
        return this.policy;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionBackendServiceSubsetting defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String policy;
        public Builder() {}
        public Builder(RegionBackendServiceSubsetting defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.policy = defaults.policy;
        }

        @CustomType.Setter
        public Builder policy(String policy) {
            if (policy == null) {
              throw new MissingRequiredPropertyException("RegionBackendServiceSubsetting", "policy");
            }
            this.policy = policy;
            return this;
        }
        public RegionBackendServiceSubsetting build() {
            final var _resultValue = new RegionBackendServiceSubsetting();
            _resultValue.policy = policy;
            return _resultValue;
        }
    }
}
