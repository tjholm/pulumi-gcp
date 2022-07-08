// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.orgpolicy.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class PolicySpecRuleValues {
    /**
     * @return List of values allowed at this resource.
     * 
     */
    private final @Nullable List<String> allowedValues;
    /**
     * @return List of values denied at this resource.
     * 
     */
    private final @Nullable List<String> deniedValues;

    @CustomType.Constructor
    private PolicySpecRuleValues(
        @CustomType.Parameter("allowedValues") @Nullable List<String> allowedValues,
        @CustomType.Parameter("deniedValues") @Nullable List<String> deniedValues) {
        this.allowedValues = allowedValues;
        this.deniedValues = deniedValues;
    }

    /**
     * @return List of values allowed at this resource.
     * 
     */
    public List<String> allowedValues() {
        return this.allowedValues == null ? List.of() : this.allowedValues;
    }
    /**
     * @return List of values denied at this resource.
     * 
     */
    public List<String> deniedValues() {
        return this.deniedValues == null ? List.of() : this.deniedValues;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicySpecRuleValues defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> allowedValues;
        private @Nullable List<String> deniedValues;

        public Builder() {
    	      // Empty
        }

        public Builder(PolicySpecRuleValues defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedValues = defaults.allowedValues;
    	      this.deniedValues = defaults.deniedValues;
        }

        public Builder allowedValues(@Nullable List<String> allowedValues) {
            this.allowedValues = allowedValues;
            return this;
        }
        public Builder allowedValues(String... allowedValues) {
            return allowedValues(List.of(allowedValues));
        }
        public Builder deniedValues(@Nullable List<String> deniedValues) {
            this.deniedValues = deniedValues;
            return this;
        }
        public Builder deniedValues(String... deniedValues) {
            return deniedValues(List.of(deniedValues));
        }        public PolicySpecRuleValues build() {
            return new PolicySpecRuleValues(allowedValues, deniedValues);
        }
    }
}
