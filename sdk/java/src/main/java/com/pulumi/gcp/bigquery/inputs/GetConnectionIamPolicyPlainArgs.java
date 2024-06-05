// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetConnectionIamPolicyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetConnectionIamPolicyPlainArgs Empty = new GetConnectionIamPolicyPlainArgs();

    /**
     * Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="connectionId", required=true)
    private String connectionId;

    /**
     * @return Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public String connectionId() {
        return this.connectionId;
    }

    /**
     * The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable String location;

    /**
     * @return The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<String> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    private GetConnectionIamPolicyPlainArgs() {}

    private GetConnectionIamPolicyPlainArgs(GetConnectionIamPolicyPlainArgs $) {
        this.connectionId = $.connectionId;
        this.location = $.location;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetConnectionIamPolicyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetConnectionIamPolicyPlainArgs $;

        public Builder() {
            $ = new GetConnectionIamPolicyPlainArgs();
        }

        public Builder(GetConnectionIamPolicyPlainArgs defaults) {
            $ = new GetConnectionIamPolicyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionId Optional connection id that should be assigned to the created connection.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder connectionId(String connectionId) {
            $.connectionId = connectionId;
            return this;
        }

        /**
         * @param location The geographic location where the connection should reside.
         * Cloud SQL instance must be in the same location as the connection
         * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
         * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
         * Spanner Connections same as spanner region
         * AWS allowed regions are aws-us-east-1
         * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable String location) {
            $.location = location;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        public GetConnectionIamPolicyPlainArgs build() {
            if ($.connectionId == null) {
                throw new MissingRequiredPropertyException("GetConnectionIamPolicyPlainArgs", "connectionId");
            }
            return $;
        }
    }

}
