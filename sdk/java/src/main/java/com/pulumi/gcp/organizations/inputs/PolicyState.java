// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.organizations.inputs.PolicyBooleanPolicyArgs;
import com.pulumi.gcp.organizations.inputs.PolicyListPolicyArgs;
import com.pulumi.gcp.organizations.inputs.PolicyRestorePolicyArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyState extends com.pulumi.resources.ResourceArgs {

    public static final PolicyState Empty = new PolicyState();

    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented
     * below.
     * 
     */
    @Import(name="booleanPolicy")
    private @Nullable Output<PolicyBooleanPolicyArgs> booleanPolicy;

    /**
     * @return A boolean policy is a constraint that is either enforced or not. Structure is documented
     * below.
     * 
     */
    public Optional<Output<PolicyBooleanPolicyArgs>> booleanPolicy() {
        return Optional.ofNullable(this.booleanPolicy);
    }

    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     * 
     */
    @Import(name="constraint")
    private @Nullable Output<String> constraint;

    /**
     * @return The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     * 
     */
    public Optional<Output<String>> constraint() {
        return Optional.ofNullable(this.constraint);
    }

    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
     * 
     */
    @Import(name="listPolicy")
    private @Nullable Output<PolicyListPolicyArgs> listPolicy;

    /**
     * @return A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
     * 
     */
    public Optional<Output<PolicyListPolicyArgs>> listPolicy() {
        return Optional.ofNullable(this.listPolicy);
    }

    /**
     * The numeric ID of the organization to set the policy for.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return The numeric ID of the organization to set the policy for.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * A restore policy is a constraint to restore the default policy. Structure is documented below.
     * 
     */
    @Import(name="restorePolicy")
    private @Nullable Output<PolicyRestorePolicyArgs> restorePolicy;

    /**
     * @return A restore policy is a constraint to restore the default policy. Structure is documented below.
     * 
     */
    public Optional<Output<PolicyRestorePolicyArgs>> restorePolicy() {
        return Optional.ofNullable(this.restorePolicy);
    }

    /**
     * (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    /**
     * Version of the Policy. Default version is 0.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return Version of the Policy. Default version is 0.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private PolicyState() {}

    private PolicyState(PolicyState $) {
        this.booleanPolicy = $.booleanPolicy;
        this.constraint = $.constraint;
        this.etag = $.etag;
        this.listPolicy = $.listPolicy;
        this.orgId = $.orgId;
        this.restorePolicy = $.restorePolicy;
        this.updateTime = $.updateTime;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyState $;

        public Builder() {
            $ = new PolicyState();
        }

        public Builder(PolicyState defaults) {
            $ = new PolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param booleanPolicy A boolean policy is a constraint that is either enforced or not. Structure is documented
         * below.
         * 
         * @return builder
         * 
         */
        public Builder booleanPolicy(@Nullable Output<PolicyBooleanPolicyArgs> booleanPolicy) {
            $.booleanPolicy = booleanPolicy;
            return this;
        }

        /**
         * @param booleanPolicy A boolean policy is a constraint that is either enforced or not. Structure is documented
         * below.
         * 
         * @return builder
         * 
         */
        public Builder booleanPolicy(PolicyBooleanPolicyArgs booleanPolicy) {
            return booleanPolicy(Output.of(booleanPolicy));
        }

        /**
         * @param constraint The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
         * 
         * @return builder
         * 
         */
        public Builder constraint(@Nullable Output<String> constraint) {
            $.constraint = constraint;
            return this;
        }

        /**
         * @param constraint The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
         * 
         * @return builder
         * 
         */
        public Builder constraint(String constraint) {
            return constraint(Output.of(constraint));
        }

        /**
         * @param etag (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param listPolicy A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder listPolicy(@Nullable Output<PolicyListPolicyArgs> listPolicy) {
            $.listPolicy = listPolicy;
            return this;
        }

        /**
         * @param listPolicy A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder listPolicy(PolicyListPolicyArgs listPolicy) {
            return listPolicy(Output.of(listPolicy));
        }

        /**
         * @param orgId The numeric ID of the organization to set the policy for.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId The numeric ID of the organization to set the policy for.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param restorePolicy A restore policy is a constraint to restore the default policy. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restorePolicy(@Nullable Output<PolicyRestorePolicyArgs> restorePolicy) {
            $.restorePolicy = restorePolicy;
            return this;
        }

        /**
         * @param restorePolicy A restore policy is a constraint to restore the default policy. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restorePolicy(PolicyRestorePolicyArgs restorePolicy) {
            return restorePolicy(Output.of(restorePolicy));
        }

        /**
         * @param updateTime (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        /**
         * @param version Version of the Policy. Default version is 0.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the Policy. Default version is 0.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public PolicyState build() {
            return $;
        }
    }

}
