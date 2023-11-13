// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.databasemigrationservice.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileAlloydbArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileCloudsqlArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileErrorArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileMysqlArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfilePostgresqlArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionProfileState extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionProfileState Empty = new ConnectionProfileState();

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
     * ***
     * 
     */
    @Import(name="connectionProfileId")
    private @Nullable Output<String> connectionProfileId;

    /**
     * @return The ID of the connection profile.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> connectionProfileId() {
        return Optional.ofNullable(this.connectionProfileId);
    }

    /**
     * Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC &#39;Zulu&#39; format, accurate to nanoseconds. Example: &#39;2014-10-02T15:01:23.045123456Z&#39;.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC &#39;Zulu&#39; format, accurate to nanoseconds. Example: &#39;2014-10-02T15:01:23.045123456Z&#39;.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The database provider.
     * 
     */
    @Import(name="dbprovider")
    private @Nullable Output<String> dbprovider;

    /**
     * @return The database provider.
     * 
     */
    public Optional<Output<String>> dbprovider() {
        return Optional.ofNullable(this.dbprovider);
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
     * Output only. The error details in case of state FAILED.
     * Structure is documented below.
     * 
     */
    @Import(name="errors")
    private @Nullable Output<List<ConnectionProfileErrorArgs>> errors;

    /**
     * @return Output only. The error details in case of state FAILED.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<ConnectionProfileErrorArgs>>> errors() {
        return Optional.ofNullable(this.errors);
    }

    /**
     * The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
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
     * The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * The current connection profile state.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The current connection profile state.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private ConnectionProfileState() {}

    private ConnectionProfileState(ConnectionProfileState $) {
        this.alloydb = $.alloydb;
        this.cloudsql = $.cloudsql;
        this.connectionProfileId = $.connectionProfileId;
        this.createTime = $.createTime;
        this.dbprovider = $.dbprovider;
        this.displayName = $.displayName;
        this.effectiveLabels = $.effectiveLabels;
        this.errors = $.errors;
        this.labels = $.labels;
        this.location = $.location;
        this.mysql = $.mysql;
        this.name = $.name;
        this.postgresql = $.postgresql;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionProfileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionProfileState $;

        public Builder() {
            $ = new ConnectionProfileState();
        }

        public Builder(ConnectionProfileState defaults) {
            $ = new ConnectionProfileState(Objects.requireNonNull(defaults));
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
         * ***
         * 
         * @return builder
         * 
         */
        public Builder connectionProfileId(@Nullable Output<String> connectionProfileId) {
            $.connectionProfileId = connectionProfileId;
            return this;
        }

        /**
         * @param connectionProfileId The ID of the connection profile.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder connectionProfileId(String connectionProfileId) {
            return connectionProfileId(Output.of(connectionProfileId));
        }

        /**
         * @param createTime Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC &#39;Zulu&#39; format, accurate to nanoseconds. Example: &#39;2014-10-02T15:01:23.045123456Z&#39;.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC &#39;Zulu&#39; format, accurate to nanoseconds. Example: &#39;2014-10-02T15:01:23.045123456Z&#39;.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param dbprovider The database provider.
         * 
         * @return builder
         * 
         */
        public Builder dbprovider(@Nullable Output<String> dbprovider) {
            $.dbprovider = dbprovider;
            return this;
        }

        /**
         * @param dbprovider The database provider.
         * 
         * @return builder
         * 
         */
        public Builder dbprovider(String dbprovider) {
            return dbprovider(Output.of(dbprovider));
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
         * @param errors Output only. The error details in case of state FAILED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder errors(@Nullable Output<List<ConnectionProfileErrorArgs>> errors) {
            $.errors = errors;
            return this;
        }

        /**
         * @param errors Output only. The error details in case of state FAILED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder errors(List<ConnectionProfileErrorArgs> errors) {
            return errors(Output.of(errors));
        }

        /**
         * @param errors Output only. The error details in case of state FAILED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder errors(ConnectionProfileErrorArgs... errors) {
            return errors(List.of(errors));
        }

        /**
         * @param labels The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
         * 
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
         * @param labels The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
         * 
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
         * @param name The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param state The current connection profile state.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The current connection profile state.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public ConnectionProfileState build() {
            return $;
        }
    }

}
