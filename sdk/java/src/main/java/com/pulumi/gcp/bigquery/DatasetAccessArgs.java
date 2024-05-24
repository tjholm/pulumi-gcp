// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessAuthorizedDatasetArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessRoutineArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessViewArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatasetAccessArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatasetAccessArgs Empty = new DatasetAccessArgs();

    /**
     * Grants all resources of particular types in a particular dataset read access to the current dataset.
     * Structure is documented below.
     * 
     */
    @Import(name="authorizedDataset")
    private @Nullable Output<DatasetAccessAuthorizedDatasetArgs> authorizedDataset;

    /**
     * @return Grants all resources of particular types in a particular dataset read access to the current dataset.
     * Structure is documented below.
     * 
     */
    public Optional<Output<DatasetAccessAuthorizedDatasetArgs>> authorizedDataset() {
        return Optional.ofNullable(this.authorizedDataset);
    }

    /**
     * A unique ID for this dataset, without the project name. The ID
     * must contain only letters (a-z, A-Z), numbers (0-9), or
     * underscores (_). The maximum length is 1,024 characters.
     * 
     * ***
     * 
     */
    @Import(name="datasetId", required=true)
    private Output<String> datasetId;

    /**
     * @return A unique ID for this dataset, without the project name. The ID
     * must contain only letters (a-z, A-Z), numbers (0-9), or
     * underscores (_). The maximum length is 1,024 characters.
     * 
     * ***
     * 
     */
    public Output<String> datasetId() {
        return this.datasetId;
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
     * A routine from a different dataset to grant access to. Queries
     * executed against that routine will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that routine is updated by any user, access to the routine
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    @Import(name="routine")
    private @Nullable Output<DatasetAccessRoutineArgs> routine;

    /**
     * @return A routine from a different dataset to grant access to. Queries
     * executed against that routine will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that routine is updated by any user, access to the routine
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    public Optional<Output<DatasetAccessRoutineArgs>> routine() {
        return Optional.ofNullable(this.routine);
    }

    /**
     * A special group to grant access to. Possible values include:
     * 
     * * `projectOwners`: Owners of the enclosing project.
     * 
     * * `projectReaders`: Readers of the enclosing project.
     * 
     * * `projectWriters`: Writers of the enclosing project.
     * 
     * * `allAuthenticatedUsers`: All authenticated BigQuery users.
     * 
     */
    @Import(name="specialGroup")
    private @Nullable Output<String> specialGroup;

    /**
     * @return A special group to grant access to. Possible values include:
     * 
     * * `projectOwners`: Owners of the enclosing project.
     * 
     * * `projectReaders`: Readers of the enclosing project.
     * 
     * * `projectWriters`: Writers of the enclosing project.
     * 
     * * `allAuthenticatedUsers`: All authenticated BigQuery users.
     * 
     */
    public Optional<Output<String>> specialGroup() {
        return Optional.ofNullable(this.specialGroup);
    }

    /**
     * An email address of a user to grant access to. For example:
     * fred{@literal @}example.com
     * 
     */
    @Import(name="userByEmail")
    private @Nullable Output<String> userByEmail;

    /**
     * @return An email address of a user to grant access to. For example:
     * fred{@literal @}example.com
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

    private DatasetAccessArgs() {}

    private DatasetAccessArgs(DatasetAccessArgs $) {
        this.authorizedDataset = $.authorizedDataset;
        this.datasetId = $.datasetId;
        this.domain = $.domain;
        this.groupByEmail = $.groupByEmail;
        this.iamMember = $.iamMember;
        this.project = $.project;
        this.role = $.role;
        this.routine = $.routine;
        this.specialGroup = $.specialGroup;
        this.userByEmail = $.userByEmail;
        this.view = $.view;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatasetAccessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatasetAccessArgs $;

        public Builder() {
            $ = new DatasetAccessArgs();
        }

        public Builder(DatasetAccessArgs defaults) {
            $ = new DatasetAccessArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizedDataset Grants all resources of particular types in a particular dataset read access to the current dataset.
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
         * @param authorizedDataset Grants all resources of particular types in a particular dataset read access to the current dataset.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder authorizedDataset(DatasetAccessAuthorizedDatasetArgs authorizedDataset) {
            return authorizedDataset(Output.of(authorizedDataset));
        }

        /**
         * @param datasetId A unique ID for this dataset, without the project name. The ID
         * must contain only letters (a-z, A-Z), numbers (0-9), or
         * underscores (_). The maximum length is 1,024 characters.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder datasetId(Output<String> datasetId) {
            $.datasetId = datasetId;
            return this;
        }

        /**
         * @param datasetId A unique ID for this dataset, without the project name. The ID
         * must contain only letters (a-z, A-Z), numbers (0-9), or
         * underscores (_). The maximum length is 1,024 characters.
         * 
         * ***
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
         * @param routine A routine from a different dataset to grant access to. Queries
         * executed against that routine will have read access to tables in
         * this dataset. The role field is not required when this field is
         * set. If that routine is updated by any user, access to the routine
         * needs to be granted again via an update operation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder routine(@Nullable Output<DatasetAccessRoutineArgs> routine) {
            $.routine = routine;
            return this;
        }

        /**
         * @param routine A routine from a different dataset to grant access to. Queries
         * executed against that routine will have read access to tables in
         * this dataset. The role field is not required when this field is
         * set. If that routine is updated by any user, access to the routine
         * needs to be granted again via an update operation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder routine(DatasetAccessRoutineArgs routine) {
            return routine(Output.of(routine));
        }

        /**
         * @param specialGroup A special group to grant access to. Possible values include:
         * 
         * * `projectOwners`: Owners of the enclosing project.
         * 
         * * `projectReaders`: Readers of the enclosing project.
         * 
         * * `projectWriters`: Writers of the enclosing project.
         * 
         * * `allAuthenticatedUsers`: All authenticated BigQuery users.
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
         * * `projectOwners`: Owners of the enclosing project.
         * 
         * * `projectReaders`: Readers of the enclosing project.
         * 
         * * `projectWriters`: Writers of the enclosing project.
         * 
         * * `allAuthenticatedUsers`: All authenticated BigQuery users.
         * 
         * @return builder
         * 
         */
        public Builder specialGroup(String specialGroup) {
            return specialGroup(Output.of(specialGroup));
        }

        /**
         * @param userByEmail An email address of a user to grant access to. For example:
         * fred{@literal @}example.com
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
         * fred{@literal @}example.com
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

        public DatasetAccessArgs build() {
            if ($.datasetId == null) {
                throw new MissingRequiredPropertyException("DatasetAccessArgs", "datasetId");
            }
            return $;
        }
    }

}
