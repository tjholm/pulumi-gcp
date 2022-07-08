// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessAuthorizedDatasetArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessViewArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatasetAccessState extends com.pulumi.resources.ResourceArgs {

    public static final DatasetAccessState Empty = new DatasetAccessState();

    /**
     * If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
     * stored in state as a different member type
     * 
     */
    @Import(name="apiUpdatedMember")
    private @Nullable Output<Boolean> apiUpdatedMember;

    /**
     * @return If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
     * stored in state as a different member type
     * 
     */
    public Optional<Output<Boolean>> apiUpdatedMember() {
        return Optional.ofNullable(this.apiUpdatedMember);
    }

    /**
     * The dataset this entry applies to
     * Structure is documented below.
     * 
     */
    @Import(name="authorizedDataset")
    private @Nullable Output<DatasetAccessAuthorizedDatasetArgs> authorizedDataset;

    /**
     * @return The dataset this entry applies to
     * Structure is documented below.
     * 
     */
    public Optional<Output<DatasetAccessAuthorizedDatasetArgs>> authorizedDataset() {
        return Optional.ofNullable(this.authorizedDataset);
    }

    /**
     * The ID of the dataset containing this table.
     * 
     */
    @Import(name="datasetId")
    private @Nullable Output<String> datasetId;

    /**
     * @return The ID of the dataset containing this table.
     * 
     */
    public Optional<Output<String>> datasetId() {
        return Optional.ofNullable(this.datasetId);
    }

    /**
     * A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * An email address of a Google Group to grant access to.
     * 
     */
    @Import(name="groupByEmail")
    private @Nullable Output<String> groupByEmail;

    /**
     * @return An email address of a Google Group to grant access to.
     * 
     */
    public Optional<Output<String>> groupByEmail() {
        return Optional.ofNullable(this.groupByEmail);
    }

    /**
     * Some other type of member that appears in the IAM Policy but isn&#39;t a user,
     * group, domain, or special group. For example: `allUsers`
     * 
     */
    @Import(name="iamMember")
    private @Nullable Output<String> iamMember;

    /**
     * @return Some other type of member that appears in the IAM Policy but isn&#39;t a user,
     * group, domain, or special group. For example: `allUsers`
     * 
     */
    public Optional<Output<String>> iamMember() {
        return Optional.ofNullable(this.iamMember);
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
     * Describes the rights granted to the user specified by the other
     * member of the access object. Basic, predefined, and custom roles are
     * supported. Predefined roles that have equivalent basic roles are
     * swapped by the API to their basic counterparts, and will show a diff
     * post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Describes the rights granted to the user specified by the other
     * member of the access object. Basic, predefined, and custom roles are
     * supported. Predefined roles that have equivalent basic roles are
     * swapped by the API to their basic counterparts, and will show a diff
     * post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * A special group to grant access to. Possible values include:
     * 
     */
    @Import(name="specialGroup")
    private @Nullable Output<String> specialGroup;

    /**
     * @return A special group to grant access to. Possible values include:
     * 
     */
    public Optional<Output<String>> specialGroup() {
        return Optional.ofNullable(this.specialGroup);
    }

    /**
     * An email address of a user to grant access to. For example:
     * fred@example.com
     * 
     */
    @Import(name="userByEmail")
    private @Nullable Output<String> userByEmail;

    /**
     * @return An email address of a user to grant access to. For example:
     * fred@example.com
     * 
     */
    public Optional<Output<String>> userByEmail() {
        return Optional.ofNullable(this.userByEmail);
    }

    /**
     * A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    @Import(name="view")
    private @Nullable Output<DatasetAccessViewArgs> view;

    /**
     * @return A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    public Optional<Output<DatasetAccessViewArgs>> view() {
        return Optional.ofNullable(this.view);
    }

    private DatasetAccessState() {}

    private DatasetAccessState(DatasetAccessState $) {
        this.apiUpdatedMember = $.apiUpdatedMember;
        this.authorizedDataset = $.authorizedDataset;
        this.datasetId = $.datasetId;
        this.domain = $.domain;
        this.groupByEmail = $.groupByEmail;
        this.iamMember = $.iamMember;
        this.project = $.project;
        this.role = $.role;
        this.specialGroup = $.specialGroup;
        this.userByEmail = $.userByEmail;
        this.view = $.view;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatasetAccessState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatasetAccessState $;

        public Builder() {
            $ = new DatasetAccessState();
        }

        public Builder(DatasetAccessState defaults) {
            $ = new DatasetAccessState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiUpdatedMember If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
         * stored in state as a different member type
         * 
         * @return builder
         * 
         */
        public Builder apiUpdatedMember(@Nullable Output<Boolean> apiUpdatedMember) {
            $.apiUpdatedMember = apiUpdatedMember;
            return this;
        }

        /**
         * @param apiUpdatedMember If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
         * stored in state as a different member type
         * 
         * @return builder
         * 
         */
        public Builder apiUpdatedMember(Boolean apiUpdatedMember) {
            return apiUpdatedMember(Output.of(apiUpdatedMember));
        }

        /**
         * @param authorizedDataset The dataset this entry applies to
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder authorizedDataset(@Nullable Output<DatasetAccessAuthorizedDatasetArgs> authorizedDataset) {
            $.authorizedDataset = authorizedDataset;
            return this;
        }

        /**
         * @param authorizedDataset The dataset this entry applies to
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder authorizedDataset(DatasetAccessAuthorizedDatasetArgs authorizedDataset) {
            return authorizedDataset(Output.of(authorizedDataset));
        }

        /**
         * @param datasetId The ID of the dataset containing this table.
         * 
         * @return builder
         * 
         */
        public Builder datasetId(@Nullable Output<String> datasetId) {
            $.datasetId = datasetId;
            return this;
        }

        /**
         * @param datasetId The ID of the dataset containing this table.
         * 
         * @return builder
         * 
         */
        public Builder datasetId(String datasetId) {
            return datasetId(Output.of(datasetId));
        }

        /**
         * @param domain A domain to grant access to. Any users signed in with the
         * domain specified will be granted the specified access
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain A domain to grant access to. Any users signed in with the
         * domain specified will be granted the specified access
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param groupByEmail An email address of a Google Group to grant access to.
         * 
         * @return builder
         * 
         */
        public Builder groupByEmail(@Nullable Output<String> groupByEmail) {
            $.groupByEmail = groupByEmail;
            return this;
        }

        /**
         * @param groupByEmail An email address of a Google Group to grant access to.
         * 
         * @return builder
         * 
         */
        public Builder groupByEmail(String groupByEmail) {
            return groupByEmail(Output.of(groupByEmail));
        }

        /**
         * @param iamMember Some other type of member that appears in the IAM Policy but isn&#39;t a user,
         * group, domain, or special group. For example: `allUsers`
         * 
         * @return builder
         * 
         */
        public Builder iamMember(@Nullable Output<String> iamMember) {
            $.iamMember = iamMember;
            return this;
        }

        /**
         * @param iamMember Some other type of member that appears in the IAM Policy but isn&#39;t a user,
         * group, domain, or special group. For example: `allUsers`
         * 
         * @return builder
         * 
         */
        public Builder iamMember(String iamMember) {
            return iamMember(Output.of(iamMember));
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
         * @param role Describes the rights granted to the user specified by the other
         * member of the access object. Basic, predefined, and custom roles are
         * supported. Predefined roles that have equivalent basic roles are
         * swapped by the API to their basic counterparts, and will show a diff
         * post-create. See
         * [official docs](https://cloud.google.com/bigquery/docs/access-control).
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Describes the rights granted to the user specified by the other
         * member of the access object. Basic, predefined, and custom roles are
         * supported. Predefined roles that have equivalent basic roles are
         * swapped by the API to their basic counterparts, and will show a diff
         * post-create. See
         * [official docs](https://cloud.google.com/bigquery/docs/access-control).
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param specialGroup A special group to grant access to. Possible values include:
         * 
         * @return builder
         * 
         */
        public Builder specialGroup(@Nullable Output<String> specialGroup) {
            $.specialGroup = specialGroup;
            return this;
        }

        /**
         * @param specialGroup A special group to grant access to. Possible values include:
         * 
         * @return builder
         * 
         */
        public Builder specialGroup(String specialGroup) {
            return specialGroup(Output.of(specialGroup));
        }

        /**
         * @param userByEmail An email address of a user to grant access to. For example:
         * fred@example.com
         * 
         * @return builder
         * 
         */
        public Builder userByEmail(@Nullable Output<String> userByEmail) {
            $.userByEmail = userByEmail;
            return this;
        }

        /**
         * @param userByEmail An email address of a user to grant access to. For example:
         * fred@example.com
         * 
         * @return builder
         * 
         */
        public Builder userByEmail(String userByEmail) {
            return userByEmail(Output.of(userByEmail));
        }

        /**
         * @param view A view from a different dataset to grant access to. Queries
         * executed against that view will have read access to tables in
         * this dataset. The role field is not required when this field is
         * set. If that view is updated by any user, access to the view
         * needs to be granted again via an update operation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder view(@Nullable Output<DatasetAccessViewArgs> view) {
            $.view = view;
            return this;
        }

        /**
         * @param view A view from a different dataset to grant access to. Queries
         * executed against that view will have read access to tables in
         * this dataset. The role field is not required when this field is
         * set. If that view is updated by any user, access to the view
         * needs to be granted again via an update operation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder view(DatasetAccessViewArgs view) {
            return view(Output.of(view));
        }

        public DatasetAccessState build() {
            return $;
        }
    }

}
