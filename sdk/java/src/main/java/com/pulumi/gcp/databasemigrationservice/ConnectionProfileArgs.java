// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.databasemigrationservice;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileAlloydbArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileCloudsqlArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileMysqlArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfilePostgresqlArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionProfileArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionProfileArgs Empty = new ConnectionProfileArgs();

    /**
     * Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
     * Structure is documented below.
     * 
     */
    @Import(name="alloydb")
    private @Nullable Output<ConnectionProfileAlloydbArgs> alloydb;

    /**
     * @return Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionProfileAlloydbArgs>> alloydb() {
        return Optional.ofNullable(this.alloydb);
    }

    /**
     * Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
     * Structure is documented below.
     * 
     */
    @Import(name="cloudsql")
    private @Nullable Output<ConnectionProfileCloudsqlArgs> cloudsql;

    /**
     * @return Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionProfileCloudsqlArgs>> cloudsql() {
        return Optional.ofNullable(this.cloudsql);
    }

    /**
     * The ID of the connection profile.
     * 
     */
    @Import(name="connectionProfileId", required=true)
    private Output<String> connectionProfileId;

    /**
     * @return The ID of the connection profile.
     * 
     */
    public Output<String> connectionProfileId() {
        return this.connectionProfileId;
    }

    /**
     * The connection profile display name.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The connection profile display name.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location where the connection profile should reside.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location where the connection profile should reside.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Specifies connection parameters required specifically for MySQL databases.
     * Structure is documented below.
     * 
     */
    @Import(name="mysql")
    private @Nullable Output<ConnectionProfileMysqlArgs> mysql;

    /**
     * @return Specifies connection parameters required specifically for MySQL databases.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionProfileMysqlArgs>> mysql() {
        return Optional.ofNullable(this.mysql);
    }

    /**
     * Specifies connection parameters required specifically for PostgreSQL databases.
     * Structure is documented below.
     * 
     */
    @Import(name="postgresql")
    private @Nullable Output<ConnectionProfilePostgresqlArgs> postgresql;

    /**
     * @return Specifies connection parameters required specifically for PostgreSQL databases.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionProfilePostgresqlArgs>> postgresql() {
        return Optional.ofNullable(this.postgresql);
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

    private ConnectionProfileArgs() {}

    private ConnectionProfileArgs(ConnectionProfileArgs $) {
        this.alloydb = $.alloydb;
        this.cloudsql = $.cloudsql;
        this.connectionProfileId = $.connectionProfileId;
        this.displayName = $.displayName;
        this.labels = $.labels;
        this.location = $.location;
        this.mysql = $.mysql;
        this.postgresql = $.postgresql;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionProfileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionProfileArgs $;

        public Builder() {
            $ = new ConnectionProfileArgs();
        }

        public Builder(ConnectionProfileArgs defaults) {
            $ = new ConnectionProfileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alloydb Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder alloydb(@Nullable Output<ConnectionProfileAlloydbArgs> alloydb) {
            $.alloydb = alloydb;
            return this;
        }

        /**
         * @param alloydb Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder alloydb(ConnectionProfileAlloydbArgs alloydb) {
            return alloydb(Output.of(alloydb));
        }

        /**
         * @param cloudsql Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudsql(@Nullable Output<ConnectionProfileCloudsqlArgs> cloudsql) {
            $.cloudsql = cloudsql;
            return this;
        }

        /**
         * @param cloudsql Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudsql(ConnectionProfileCloudsqlArgs cloudsql) {
            return cloudsql(Output.of(cloudsql));
        }

        /**
         * @param connectionProfileId The ID of the connection profile.
         * 
         * @return builder
         * 
         */
        public Builder connectionProfileId(Output<String> connectionProfileId) {
            $.connectionProfileId = connectionProfileId;
            return this;
        }

        /**
         * @param connectionProfileId The ID of the connection profile.
         * 
         * @return builder
         * 
         */
        public Builder connectionProfileId(String connectionProfileId) {
            return connectionProfileId(Output.of(connectionProfileId));
        }

        /**
         * @param displayName The connection profile display name.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The connection profile display name.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param labels The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location where the connection profile should reside.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location where the connection profile should reside.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param mysql Specifies connection parameters required specifically for MySQL databases.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder mysql(@Nullable Output<ConnectionProfileMysqlArgs> mysql) {
            $.mysql = mysql;
            return this;
        }

        /**
         * @param mysql Specifies connection parameters required specifically for MySQL databases.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder mysql(ConnectionProfileMysqlArgs mysql) {
            return mysql(Output.of(mysql));
        }

        /**
         * @param postgresql Specifies connection parameters required specifically for PostgreSQL databases.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder postgresql(@Nullable Output<ConnectionProfilePostgresqlArgs> postgresql) {
            $.postgresql = postgresql;
            return this;
        }

        /**
         * @param postgresql Specifies connection parameters required specifically for PostgreSQL databases.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder postgresql(ConnectionProfilePostgresqlArgs postgresql) {
            return postgresql(Output.of(postgresql));
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

        public ConnectionProfileArgs build() {
            $.connectionProfileId = Objects.requireNonNull($.connectionProfileId, "expected parameter 'connectionProfileId' to be non-null");
            return $;
        }
    }

}
