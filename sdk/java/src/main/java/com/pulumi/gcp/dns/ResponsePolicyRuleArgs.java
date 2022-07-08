// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dns.inputs.ResponsePolicyRuleLocalDataArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResponsePolicyRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResponsePolicyRuleArgs Empty = new ResponsePolicyRuleArgs();

    /**
     * The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     * 
     */
    @Import(name="dnsName", required=true)
    private Output<String> dnsName;

    /**
     * @return The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }

    /**
     * Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
     * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     * Structure is documented below.
     * 
     */
    @Import(name="localData")
    private @Nullable Output<ResponsePolicyRuleLocalDataArgs> localData;

    /**
     * @return Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
     * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ResponsePolicyRuleLocalDataArgs>> localData() {
        return Optional.ofNullable(this.localData);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Identifies the response policy addressed by this request.
     * 
     */
    @Import(name="responsePolicy", required=true)
    private Output<String> responsePolicy;

    /**
     * @return Identifies the response policy addressed by this request.
     * 
     */
    public Output<String> responsePolicy() {
        return this.responsePolicy;
    }

    /**
     * An identifier for this rule. Must be unique with the ResponsePolicy.
     * 
     */
    @Import(name="ruleName", required=true)
    private Output<String> ruleName;

    /**
     * @return An identifier for this rule. Must be unique with the ResponsePolicy.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }

    private ResponsePolicyRuleArgs() {}

    private ResponsePolicyRuleArgs(ResponsePolicyRuleArgs $) {
        this.dnsName = $.dnsName;
        this.localData = $.localData;
        this.project = $.project;
        this.responsePolicy = $.responsePolicy;
        this.ruleName = $.ruleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResponsePolicyRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResponsePolicyRuleArgs $;

        public Builder() {
            $ = new ResponsePolicyRuleArgs();
        }

        public Builder(ResponsePolicyRuleArgs defaults) {
            $ = new ResponsePolicyRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dnsName The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(Output<String> dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        /**
         * @param dnsName The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(String dnsName) {
            return dnsName(Output.of(dnsName));
        }

        /**
         * @param localData Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
         * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder localData(@Nullable Output<ResponsePolicyRuleLocalDataArgs> localData) {
            $.localData = localData;
            return this;
        }

        /**
         * @param localData Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
         * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder localData(ResponsePolicyRuleLocalDataArgs localData) {
            return localData(Output.of(localData));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param responsePolicy Identifies the response policy addressed by this request.
         * 
         * @return builder
         * 
         */
        public Builder responsePolicy(Output<String> responsePolicy) {
            $.responsePolicy = responsePolicy;
            return this;
        }

        /**
         * @param responsePolicy Identifies the response policy addressed by this request.
         * 
         * @return builder
         * 
         */
        public Builder responsePolicy(String responsePolicy) {
            return responsePolicy(Output.of(responsePolicy));
        }

        /**
         * @param ruleName An identifier for this rule. Must be unique with the ResponsePolicy.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        /**
         * @param ruleName An identifier for this rule. Must be unique with the ResponsePolicy.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        public ResponsePolicyRuleArgs build() {
            $.dnsName = Objects.requireNonNull($.dnsName, "expected parameter 'dnsName' to be non-null");
            $.responsePolicy = Objects.requireNonNull($.responsePolicy, "expected parameter 'responsePolicy' to be non-null");
            $.ruleName = Objects.requireNonNull($.ruleName, "expected parameter 'ruleName' to be non-null");
            return $;
        }
    }

}
