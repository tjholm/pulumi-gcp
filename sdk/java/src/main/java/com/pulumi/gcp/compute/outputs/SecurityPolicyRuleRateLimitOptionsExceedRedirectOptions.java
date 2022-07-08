// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions {
    /**
     * @return External redirection target when &#34;EXTERNAL_302&#34; is set in &#39;type&#39;.
     * 
     */
    private final @Nullable String target;
    /**
     * @return Type of redirect action.
     * 
     */
    private final String type;

    @CustomType.Constructor
    private SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions(
        @CustomType.Parameter("target") @Nullable String target,
        @CustomType.Parameter("type") String type) {
        this.target = target;
        this.type = type;
    }

    /**
     * @return External redirection target when &#34;EXTERNAL_302&#34; is set in &#39;type&#39;.
     * 
     */
    public Optional<String> target() {
        return Optional.ofNullable(this.target);
    }
    /**
     * @return Type of redirect action.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String target;
        private String type;

        public Builder() {
    	      // Empty
        }

        public Builder(SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.target = defaults.target;
    	      this.type = defaults.type;
        }

        public Builder target(@Nullable String target) {
            this.target = target;
            return this;
        }
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }        public SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions build() {
            return new SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions(target, type);
        }
    }
}
