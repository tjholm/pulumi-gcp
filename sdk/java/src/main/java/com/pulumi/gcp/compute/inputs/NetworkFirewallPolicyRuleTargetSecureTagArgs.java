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


public final class NetworkFirewallPolicyRuleTargetSecureTagArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkFirewallPolicyRuleTargetSecureTagArgs Empty = new NetworkFirewallPolicyRuleTargetSecureTagArgs();

    /**
     * Name of the secure tag, created with TagManager&#39;s TagValue API. {@literal @}pattern tagValues/[0-9]+
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the secure tag, created with TagManager&#39;s TagValue API. {@literal @}pattern tagValues/[0-9]+
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private NetworkFirewallPolicyRuleTargetSecureTagArgs() {}

    private NetworkFirewallPolicyRuleTargetSecureTagArgs(NetworkFirewallPolicyRuleTargetSecureTagArgs $) {
        this.name = $.name;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkFirewallPolicyRuleTargetSecureTagArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkFirewallPolicyRuleTargetSecureTagArgs $;

        public Builder() {
            $ = new NetworkFirewallPolicyRuleTargetSecureTagArgs();
        }

        public Builder(NetworkFirewallPolicyRuleTargetSecureTagArgs defaults) {
            $ = new NetworkFirewallPolicyRuleTargetSecureTagArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the secure tag, created with TagManager&#39;s TagValue API. {@literal @}pattern tagValues/[0-9]+
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the secure tag, created with TagManager&#39;s TagValue API. {@literal @}pattern tagValues/[0-9]+
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param state [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public NetworkFirewallPolicyRuleTargetSecureTagArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("NetworkFirewallPolicyRuleTargetSecureTagArgs", "name");
            }
            return $;
        }
    }

}
