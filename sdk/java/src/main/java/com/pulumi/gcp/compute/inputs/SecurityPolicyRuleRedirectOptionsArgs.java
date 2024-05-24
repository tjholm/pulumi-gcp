// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityPolicyRuleRedirectOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityPolicyRuleRedirectOptionsArgs Empty = new SecurityPolicyRuleRedirectOptionsArgs();

    /**
     * External redirection target when `EXTERNAL_302` is set in `type`.
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return External redirection target when `EXTERNAL_302` is set in `type`.
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * Type of redirect action.
     * 
     * * `EXTERNAL_302`: Redirect to an external address, configured in `target`.
     * * `GOOGLE_RECAPTCHA`: Redirect to Google reCAPTCHA.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of redirect action.
     * 
     * * `EXTERNAL_302`: Redirect to an external address, configured in `target`.
     * * `GOOGLE_RECAPTCHA`: Redirect to Google reCAPTCHA.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private SecurityPolicyRuleRedirectOptionsArgs() {}

    private SecurityPolicyRuleRedirectOptionsArgs(SecurityPolicyRuleRedirectOptionsArgs $) {
        this.target = $.target;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityPolicyRuleRedirectOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityPolicyRuleRedirectOptionsArgs $;

        public Builder() {
            $ = new SecurityPolicyRuleRedirectOptionsArgs();
        }

        public Builder(SecurityPolicyRuleRedirectOptionsArgs defaults) {
            $ = new SecurityPolicyRuleRedirectOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param target External redirection target when `EXTERNAL_302` is set in `type`.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target External redirection target when `EXTERNAL_302` is set in `type`.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        /**
         * @param type Type of redirect action.
         * 
         * * `EXTERNAL_302`: Redirect to an external address, configured in `target`.
         * * `GOOGLE_RECAPTCHA`: Redirect to Google reCAPTCHA.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of redirect action.
         * 
         * * `EXTERNAL_302`: Redirect to an external address, configured in `target`.
         * * `GOOGLE_RECAPTCHA`: Redirect to Google reCAPTCHA.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SecurityPolicyRuleRedirectOptionsArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("SecurityPolicyRuleRedirectOptionsArgs", "type");
            }
            return $;
        }
    }

}
