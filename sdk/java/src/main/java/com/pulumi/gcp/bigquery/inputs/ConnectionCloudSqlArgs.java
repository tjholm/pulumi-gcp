// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.ConnectionCloudSqlCredentialArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionCloudSqlArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionCloudSqlArgs Empty = new ConnectionCloudSqlArgs();

    /**
     * Cloud SQL properties.
     * Structure is documented below.
     * 
     */
    @Import(name="credential", required=true)
    private Output<ConnectionCloudSqlCredentialArgs> credential;

    /**
     * @return Cloud SQL properties.
     * Structure is documented below.
     * 
     */
    public Output<ConnectionCloudSqlCredentialArgs> credential() {
        return this.credential;
    }

    /**
     * Database name.
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return Database name.
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * Cloud SQL instance ID in the form project:location:instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return Cloud SQL instance ID in the form project:location:instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
     * 
     */
    @Import(name="serviceAccountId")
    private @Nullable Output<String> serviceAccountId;

    /**
     * @return When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
     * 
     */
    public Optional<Output<String>> serviceAccountId() {
        return Optional.ofNullable(this.serviceAccountId);
    }

    /**
     * Type of the Cloud SQL database.
     * Possible values are `DATABASE_TYPE_UNSPECIFIED`, `POSTGRES`, and `MYSQL`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of the Cloud SQL database.
     * Possible values are `DATABASE_TYPE_UNSPECIFIED`, `POSTGRES`, and `MYSQL`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ConnectionCloudSqlArgs() {}

    private ConnectionCloudSqlArgs(ConnectionCloudSqlArgs $) {
        this.credential = $.credential;
        this.database = $.database;
        this.instanceId = $.instanceId;
        this.serviceAccountId = $.serviceAccountId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionCloudSqlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionCloudSqlArgs $;

        public Builder() {
            $ = new ConnectionCloudSqlArgs();
        }

        public Builder(ConnectionCloudSqlArgs defaults) {
            $ = new ConnectionCloudSqlArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param credential Cloud SQL properties.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder credential(Output<ConnectionCloudSqlCredentialArgs> credential) {
            $.credential = credential;
            return this;
        }

        /**
         * @param credential Cloud SQL properties.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder credential(ConnectionCloudSqlCredentialArgs credential) {
            return credential(Output.of(credential));
        }

        /**
         * @param database Database name.
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database Database name.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param instanceId Cloud SQL instance ID in the form project:location:instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Cloud SQL instance ID in the form project:location:instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param serviceAccountId When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(@Nullable Output<String> serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        /**
         * @param serviceAccountId When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            return serviceAccountId(Output.of(serviceAccountId));
        }

        /**
         * @param type Type of the Cloud SQL database.
         * Possible values are `DATABASE_TYPE_UNSPECIFIED`, `POSTGRES`, and `MYSQL`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the Cloud SQL database.
         * Possible values are `DATABASE_TYPE_UNSPECIFIED`, `POSTGRES`, and `MYSQL`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ConnectionCloudSqlArgs build() {
            $.credential = Objects.requireNonNull($.credential, "expected parameter 'credential' to be non-null");
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
