// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SecurityPolicyRuleMatchExpr {
    /**
     * @return Textual representation of an expression in Common Expression Language syntax.
     * The application context of the containing message determines which well-known feature set of CEL is supported.
     * 
     */
    private final String expression;

    @CustomType.Constructor
    private SecurityPolicyRuleMatchExpr(@CustomType.Parameter("expression") String expression) {
        this.expression = expression;
    }

    /**
     * @return Textual representation of an expression in Common Expression Language syntax.
     * The application context of the containing message determines which well-known feature set of CEL is supported.
     * 
     */
    public String expression() {
        return this.expression;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecurityPolicyRuleMatchExpr defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String expression;

        public Builder() {
    	      // Empty
        }

        public Builder(SecurityPolicyRuleMatchExpr defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expression = defaults.expression;
        }

        public Builder expression(String expression) {
            this.expression = Objects.requireNonNull(expression);
            return this;
        }        public SecurityPolicyRuleMatchExpr build() {
            return new SecurityPolicyRuleMatchExpr(expression);
        }
    }
}
