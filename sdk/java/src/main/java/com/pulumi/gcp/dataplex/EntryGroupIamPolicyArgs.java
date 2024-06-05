// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EntryGroupIamPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final EntryGroupIamPolicyArgs Empty = new EntryGroupIamPolicyArgs();

    @Import(name="entryGroupId", required=true)
    private Output<String> entryGroupId;

    public Output<String> entryGroupId() {
        return this.entryGroupId;
    }

    /**
     * The location where entry group will be created in.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location where entry group will be created in.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Import(name="policyData", required=true)
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private EntryGroupIamPolicyArgs() {}

    private EntryGroupIamPolicyArgs(EntryGroupIamPolicyArgs $) {
        this.entryGroupId = $.entryGroupId;
        this.location = $.location;
        this.policyData = $.policyData;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EntryGroupIamPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EntryGroupIamPolicyArgs $;

        public Builder() {
            $ = new EntryGroupIamPolicyArgs();
        }

        public Builder(EntryGroupIamPolicyArgs defaults) {
            $ = new EntryGroupIamPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder entryGroupId(Output<String> entryGroupId) {
            $.entryGroupId = entryGroupId;
            return this;
        }

        public Builder entryGroupId(String entryGroupId) {
            return entryGroupId(Output.of(entryGroupId));
        }

        /**
         * @param location The location where entry group will be created in.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location where entry group will be created in.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(Output<String> policyData) {
            $.policyData = policyData;
            return this;
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(String policyData) {
            return policyData(Output.of(policyData));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
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
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public EntryGroupIamPolicyArgs build() {
            if ($.entryGroupId == null) {
                throw new MissingRequiredPropertyException("EntryGroupIamPolicyArgs", "entryGroupId");
            }
            if ($.policyData == null) {
                throw new MissingRequiredPropertyException("EntryGroupIamPolicyArgs", "policyData");
            }
            return $;
        }
    }

}
