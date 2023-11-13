// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networksecurity.inputs.AuthorizationPolicyRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthorizationPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationPolicyState Empty = new AuthorizationPolicyState();

    /**
     * The action to take when a rule match is found. Possible values are &#34;ALLOW&#34; or &#34;DENY&#34;.
     * Possible values are: `ALLOW`, `DENY`.
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return The action to take when a rule match is found. Possible values are &#34;ALLOW&#34; or &#34;DENY&#34;.
     * Possible values are: `ALLOW`, `DENY`.
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * Time the AuthorizationPolicy was created in UTC.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Time the AuthorizationPolicy was created in UTC.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * A free-text description of the resource. Max length 1024 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A free-text description of the resource. Max length 1024 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * Set of label tags associated with the AuthorizationPolicy resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the AuthorizationPolicy resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location of the authorization policy.
     * The default value is `global`.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the authorization policy.
     * The default value is `global`.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Name of the AuthorizationPolicy resource.
     * 
     * ***
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the AuthorizationPolicy resource.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * List of rules to match. Note that at least one of the rules must match in order for the action specified in the &#39;action&#39; field to be taken.
     * A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
     * Structure is documented below.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<AuthorizationPolicyRuleArgs>> rules;

    /**
     * @return List of rules to match. Note that at least one of the rules must match in order for the action specified in the &#39;action&#39; field to be taken.
     * A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<AuthorizationPolicyRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Time the AuthorizationPolicy was updated in UTC.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Time the AuthorizationPolicy was updated in UTC.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private AuthorizationPolicyState() {}

    private AuthorizationPolicyState(AuthorizationPolicyState $) {
        this.action = $.action;
        this.createTime = $.createTime;
        this.description = $.description;
        this.effectiveLabels = $.effectiveLabels;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.rules = $.rules;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationPolicyState $;

        public Builder() {
            $ = new AuthorizationPolicyState();
        }

        public Builder(AuthorizationPolicyState defaults) {
            $ = new AuthorizationPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The action to take when a rule match is found. Possible values are &#34;ALLOW&#34; or &#34;DENY&#34;.
         * Possible values are: `ALLOW`, `DENY`.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The action to take when a rule match is found. Possible values are &#34;ALLOW&#34; or &#34;DENY&#34;.
         * Possible values are: `ALLOW`, `DENY`.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param createTime Time the AuthorizationPolicy was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Time the AuthorizationPolicy was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param labels Set of label tags associated with the AuthorizationPolicy resource.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Set of label tags associated with the AuthorizationPolicy resource.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location of the authorization policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the authorization policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name Name of the AuthorizationPolicy resource.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the AuthorizationPolicy resource.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param rules List of rules to match. Note that at least one of the rules must match in order for the action specified in the &#39;action&#39; field to be taken.
         * A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<AuthorizationPolicyRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules List of rules to match. Note that at least one of the rules must match in order for the action specified in the &#39;action&#39; field to be taken.
         * A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<AuthorizationPolicyRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules List of rules to match. Note that at least one of the rules must match in order for the action specified in the &#39;action&#39; field to be taken.
         * A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(AuthorizationPolicyRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param updateTime Time the AuthorizationPolicy was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Time the AuthorizationPolicy was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public AuthorizationPolicyState build() {
            return $;
        }
    }

}
