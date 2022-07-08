// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EntryGroupIamPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final EntryGroupIamPolicyArgs Empty = new EntryGroupIamPolicyArgs();

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="entryGroup", required=true)
    private Output<String> entryGroup;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> entryGroup() {
        return this.entryGroup;
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

    @Import(name="region")
    private @Nullable Output<String> region;

    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private EntryGroupIamPolicyArgs() {}

    private EntryGroupIamPolicyArgs(EntryGroupIamPolicyArgs $) {
        this.entryGroup = $.entryGroup;
        this.policyData = $.policyData;
        this.project = $.project;
        this.region = $.region;
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

        /**
         * @param entryGroup Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder entryGroup(Output<String> entryGroup) {
            $.entryGroup = entryGroup;
            return this;
        }

        /**
         * @param entryGroup Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder entryGroup(String entryGroup) {
            return entryGroup(Output.of(entryGroup));
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

        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        public EntryGroupIamPolicyArgs build() {
            $.entryGroup = Objects.requireNonNull($.entryGroup, "expected parameter 'entryGroup' to be non-null");
            $.policyData = Objects.requireNonNull($.policyData, "expected parameter 'policyData' to be non-null");
            return $;
        }
    }

}
