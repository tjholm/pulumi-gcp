// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.ConnectionIamMemberConditionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionIamMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionIamMemberArgs Empty = new ConnectionIamMemberArgs();

    @Import(name="condition")
    private @Nullable Output<ConnectionIamMemberConditionArgs> condition;

    public Optional<Output<ConnectionIamMemberConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="connectionId", required=true)
    private Output<String> connectionId;

    /**
     * @return Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> connectionId() {
        return this.connectionId;
    }

    /**
     * The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    @Import(name="member", required=true)
    private Output<String> member;

    public Output<String> member() {
        return this.member;
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

    /**
     * The role that should be applied. Only one
     * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private ConnectionIamMemberArgs() {}

    private ConnectionIamMemberArgs(ConnectionIamMemberArgs $) {
        this.condition = $.condition;
        this.connectionId = $.connectionId;
        this.location = $.location;
        this.member = $.member;
        this.project = $.project;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionIamMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionIamMemberArgs $;

        public Builder() {
            $ = new ConnectionIamMemberArgs();
        }

        public Builder(ConnectionIamMemberArgs defaults) {
            $ = new ConnectionIamMemberArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<ConnectionIamMemberConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(ConnectionIamMemberConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param connectionId Optional connection id that should be assigned to the created connection.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder connectionId(Output<String> connectionId) {
            $.connectionId = connectionId;
            return this;
        }

        /**
         * @param connectionId Optional connection id that should be assigned to the created connection.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder connectionId(String connectionId) {
            return connectionId(Output.of(connectionId));
        }

        /**
         * @param location The geographic location where the connection should reside.
         * Cloud SQL instance must be in the same location as the connection
         * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
         * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
         * Spanner Connections same as spanner region
         * AWS allowed regions are aws-us-east-1
         * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The geographic location where the connection should reside.
         * Cloud SQL instance must be in the same location as the connection
         * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
         * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
         * Spanner Connections same as spanner region
         * AWS allowed regions are aws-us-east-1
         * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        public Builder member(Output<String> member) {
            $.member = member;
            return this;
        }

        public Builder member(String member) {
            return member(Output.of(member));
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

        /**
         * @param role The role that should be applied. Only one
         * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public ConnectionIamMemberArgs build() {
            $.connectionId = Objects.requireNonNull($.connectionId, "expected parameter 'connectionId' to be non-null");
            $.member = Objects.requireNonNull($.member, "expected parameter 'member' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
