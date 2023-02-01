// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.TableEncryptionConfigurationArgs;
import com.pulumi.gcp.bigquery.inputs.TableExternalDataConfigurationArgs;
import com.pulumi.gcp.bigquery.inputs.TableMaterializedViewArgs;
import com.pulumi.gcp.bigquery.inputs.TableRangePartitioningArgs;
import com.pulumi.gcp.bigquery.inputs.TableTimePartitioningArgs;
import com.pulumi.gcp.bigquery.inputs.TableViewArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableState extends com.pulumi.resources.ResourceArgs {

    public static final TableState Empty = new TableState();

    /**
     * Specifies column names to use for data clustering.
     * Up to four top-level columns are allowed, and should be specified in
     * descending priority order.
     * 
     */
    @Import(name="clusterings")
    private @Nullable Output<List<String>> clusterings;

    /**
     * @return Specifies column names to use for data clustering.
     * Up to four top-level columns are allowed, and should be specified in
     * descending priority order.
     * 
     */
    public Optional<Output<List<String>>> clusterings() {
        return Optional.ofNullable(this.clusterings);
    }

    /**
     * The time when this table was created, in milliseconds since the epoch.
     * 
     */
    @Import(name="creationTime")
    private @Nullable Output<Integer> creationTime;

    /**
     * @return The time when this table was created, in milliseconds since the epoch.
     * 
     */
    public Optional<Output<Integer>> creationTime() {
        return Optional.ofNullable(this.creationTime);
    }

    /**
     * The dataset ID to create the table in.
     * Changing this forces a new resource to be created.
     * 
     */
    @Import(name="datasetId")
    private @Nullable Output<String> datasetId;

    /**
     * @return The dataset ID to create the table in.
     * Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> datasetId() {
        return Optional.ofNullable(this.datasetId);
    }

    /**
     * Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a
     * terraform destroy or terraform apply that would delete the instance will fail.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a
     * terraform destroy or terraform apply that would delete the instance will fail.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * The field description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The field description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies how the table should be encrypted.
     * If left blank, the table will be encrypted with a Google-managed key; that process
     * is transparent to the user.  Structure is documented below.
     * 
     */
    @Import(name="encryptionConfiguration")
    private @Nullable Output<TableEncryptionConfigurationArgs> encryptionConfiguration;

    /**
     * @return Specifies how the table should be encrypted.
     * If left blank, the table will be encrypted with a Google-managed key; that process
     * is transparent to the user.  Structure is documented below.
     * 
     */
    public Optional<Output<TableEncryptionConfigurationArgs>> encryptionConfiguration() {
        return Optional.ofNullable(this.encryptionConfiguration);
    }

    /**
     * A hash of the resource.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return A hash of the resource.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The time when this table expires, in
     * milliseconds since the epoch. If not present, the table will persist
     * indefinitely. Expired tables will be deleted and their storage
     * reclaimed.
     * 
     */
    @Import(name="expirationTime")
    private @Nullable Output<Integer> expirationTime;

    /**
     * @return The time when this table expires, in
     * milliseconds since the epoch. If not present, the table will persist
     * indefinitely. Expired tables will be deleted and their storage
     * reclaimed.
     * 
     */
    public Optional<Output<Integer>> expirationTime() {
        return Optional.ofNullable(this.expirationTime);
    }

    /**
     * Describes the data format,
     * location, and other properties of a table stored outside of BigQuery.
     * By defining these properties, the data source can then be queried as
     * if it were a standard BigQuery table. Structure is documented below.
     * 
     */
    @Import(name="externalDataConfiguration")
    private @Nullable Output<TableExternalDataConfigurationArgs> externalDataConfiguration;

    /**
     * @return Describes the data format,
     * location, and other properties of a table stored outside of BigQuery.
     * By defining these properties, the data source can then be queried as
     * if it were a standard BigQuery table. Structure is documented below.
     * 
     */
    public Optional<Output<TableExternalDataConfigurationArgs>> externalDataConfiguration() {
        return Optional.ofNullable(this.externalDataConfiguration);
    }

    /**
     * A descriptive name for the table.
     * 
     */
    @Import(name="friendlyName")
    private @Nullable Output<String> friendlyName;

    /**
     * @return A descriptive name for the table.
     * 
     */
    public Optional<Output<String>> friendlyName() {
        return Optional.ofNullable(this.friendlyName);
    }

    /**
     * A mapping of labels to assign to the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A mapping of labels to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The time when this table was last modified, in milliseconds since the epoch.
     * 
     */
    @Import(name="lastModifiedTime")
    private @Nullable Output<Integer> lastModifiedTime;

    /**
     * @return The time when this table was last modified, in milliseconds since the epoch.
     * 
     */
    public Optional<Output<Integer>> lastModifiedTime() {
        return Optional.ofNullable(this.lastModifiedTime);
    }

    /**
     * The geographic location where the table resides. This value is inherited from the dataset.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The geographic location where the table resides. This value is inherited from the dataset.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * If specified, configures this table as a materialized view.
     * Structure is documented below.
     * 
     */
    @Import(name="materializedView")
    private @Nullable Output<TableMaterializedViewArgs> materializedView;

    /**
     * @return If specified, configures this table as a materialized view.
     * Structure is documented below.
     * 
     */
    public Optional<Output<TableMaterializedViewArgs>> materializedView() {
        return Optional.ofNullable(this.materializedView);
    }

    /**
     * The size of this table in bytes, excluding any data in the streaming buffer.
     * 
     */
    @Import(name="numBytes")
    private @Nullable Output<Integer> numBytes;

    /**
     * @return The size of this table in bytes, excluding any data in the streaming buffer.
     * 
     */
    public Optional<Output<Integer>> numBytes() {
        return Optional.ofNullable(this.numBytes);
    }

    /**
     * The number of bytes in the table that are considered &#34;long-term storage&#34;.
     * 
     */
    @Import(name="numLongTermBytes")
    private @Nullable Output<Integer> numLongTermBytes;

    /**
     * @return The number of bytes in the table that are considered &#34;long-term storage&#34;.
     * 
     */
    public Optional<Output<Integer>> numLongTermBytes() {
        return Optional.ofNullable(this.numLongTermBytes);
    }

    /**
     * The number of rows of data in this table, excluding any data in the streaming buffer.
     * 
     */
    @Import(name="numRows")
    private @Nullable Output<Integer> numRows;

    /**
     * @return The number of rows of data in this table, excluding any data in the streaming buffer.
     * 
     */
    public Optional<Output<Integer>> numRows() {
        return Optional.ofNullable(this.numRows);
    }

    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * If specified, configures range-based
     * partitioning for this table. Structure is documented below.
     * 
     */
    @Import(name="rangePartitioning")
    private @Nullable Output<TableRangePartitioningArgs> rangePartitioning;

    /**
     * @return If specified, configures range-based
     * partitioning for this table. Structure is documented below.
     * 
     */
    public Optional<Output<TableRangePartitioningArgs>> rangePartitioning() {
        return Optional.ofNullable(this.rangePartitioning);
    }

    /**
     * A JSON schema for the table.
     * 
     */
    @Import(name="schema")
    private @Nullable Output<String> schema;

    /**
     * @return A JSON schema for the table.
     * 
     */
    public Optional<Output<String>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * The URI of the created resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * A unique ID for the resource.
     * Changing this forces a new resource to be created.
     * 
     */
    @Import(name="tableId")
    private @Nullable Output<String> tableId;

    /**
     * @return A unique ID for the resource.
     * Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> tableId() {
        return Optional.ofNullable(this.tableId);
    }

    /**
     * If specified, configures time-based
     * partitioning for this table. Structure is documented below.
     * 
     */
    @Import(name="timePartitioning")
    private @Nullable Output<TableTimePartitioningArgs> timePartitioning;

    /**
     * @return If specified, configures time-based
     * partitioning for this table. Structure is documented below.
     * 
     */
    public Optional<Output<TableTimePartitioningArgs>> timePartitioning() {
        return Optional.ofNullable(this.timePartitioning);
    }

    /**
     * The supported types are DAY, HOUR, MONTH, and YEAR,
     * which will generate one partition per day, hour, month, and year, respectively.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The supported types are DAY, HOUR, MONTH, and YEAR,
     * which will generate one partition per day, hour, month, and year, respectively.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * If specified, configures this table as a view.
     * Structure is documented below.
     * 
     */
    @Import(name="view")
    private @Nullable Output<TableViewArgs> view;

    /**
     * @return If specified, configures this table as a view.
     * Structure is documented below.
     * 
     */
    public Optional<Output<TableViewArgs>> view() {
        return Optional.ofNullable(this.view);
    }

    private TableState() {}

    private TableState(TableState $) {
        this.clusterings = $.clusterings;
        this.creationTime = $.creationTime;
        this.datasetId = $.datasetId;
        this.deletionProtection = $.deletionProtection;
        this.description = $.description;
        this.encryptionConfiguration = $.encryptionConfiguration;
        this.etag = $.etag;
        this.expirationTime = $.expirationTime;
        this.externalDataConfiguration = $.externalDataConfiguration;
        this.friendlyName = $.friendlyName;
        this.labels = $.labels;
        this.lastModifiedTime = $.lastModifiedTime;
        this.location = $.location;
        this.materializedView = $.materializedView;
        this.numBytes = $.numBytes;
        this.numLongTermBytes = $.numLongTermBytes;
        this.numRows = $.numRows;
        this.project = $.project;
        this.rangePartitioning = $.rangePartitioning;
        this.schema = $.schema;
        this.selfLink = $.selfLink;
        this.tableId = $.tableId;
        this.timePartitioning = $.timePartitioning;
        this.type = $.type;
        this.view = $.view;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableState $;

        public Builder() {
            $ = new TableState();
        }

        public Builder(TableState defaults) {
            $ = new TableState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterings Specifies column names to use for data clustering.
         * Up to four top-level columns are allowed, and should be specified in
         * descending priority order.
         * 
         * @return builder
         * 
         */
        public Builder clusterings(@Nullable Output<List<String>> clusterings) {
            $.clusterings = clusterings;
            return this;
        }

        /**
         * @param clusterings Specifies column names to use for data clustering.
         * Up to four top-level columns are allowed, and should be specified in
         * descending priority order.
         * 
         * @return builder
         * 
         */
        public Builder clusterings(List<String> clusterings) {
            return clusterings(Output.of(clusterings));
        }

        /**
         * @param clusterings Specifies column names to use for data clustering.
         * Up to four top-level columns are allowed, and should be specified in
         * descending priority order.
         * 
         * @return builder
         * 
         */
        public Builder clusterings(String... clusterings) {
            return clusterings(List.of(clusterings));
        }

        /**
         * @param creationTime The time when this table was created, in milliseconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder creationTime(@Nullable Output<Integer> creationTime) {
            $.creationTime = creationTime;
            return this;
        }

        /**
         * @param creationTime The time when this table was created, in milliseconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder creationTime(Integer creationTime) {
            return creationTime(Output.of(creationTime));
        }

        /**
         * @param datasetId The dataset ID to create the table in.
         * Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder datasetId(@Nullable Output<String> datasetId) {
            $.datasetId = datasetId;
            return this;
        }

        /**
         * @param datasetId The dataset ID to create the table in.
         * Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder datasetId(String datasetId) {
            return datasetId(Output.of(datasetId));
        }

        /**
         * @param deletionProtection Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a
         * terraform destroy or terraform apply that would delete the instance will fail.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a
         * terraform destroy or terraform apply that would delete the instance will fail.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param description The field description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The field description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param encryptionConfiguration Specifies how the table should be encrypted.
         * If left blank, the table will be encrypted with a Google-managed key; that process
         * is transparent to the user.  Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfiguration(@Nullable Output<TableEncryptionConfigurationArgs> encryptionConfiguration) {
            $.encryptionConfiguration = encryptionConfiguration;
            return this;
        }

        /**
         * @param encryptionConfiguration Specifies how the table should be encrypted.
         * If left blank, the table will be encrypted with a Google-managed key; that process
         * is transparent to the user.  Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfiguration(TableEncryptionConfigurationArgs encryptionConfiguration) {
            return encryptionConfiguration(Output.of(encryptionConfiguration));
        }

        /**
         * @param etag A hash of the resource.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag A hash of the resource.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param expirationTime The time when this table expires, in
         * milliseconds since the epoch. If not present, the table will persist
         * indefinitely. Expired tables will be deleted and their storage
         * reclaimed.
         * 
         * @return builder
         * 
         */
        public Builder expirationTime(@Nullable Output<Integer> expirationTime) {
            $.expirationTime = expirationTime;
            return this;
        }

        /**
         * @param expirationTime The time when this table expires, in
         * milliseconds since the epoch. If not present, the table will persist
         * indefinitely. Expired tables will be deleted and their storage
         * reclaimed.
         * 
         * @return builder
         * 
         */
        public Builder expirationTime(Integer expirationTime) {
            return expirationTime(Output.of(expirationTime));
        }

        /**
         * @param externalDataConfiguration Describes the data format,
         * location, and other properties of a table stored outside of BigQuery.
         * By defining these properties, the data source can then be queried as
         * if it were a standard BigQuery table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder externalDataConfiguration(@Nullable Output<TableExternalDataConfigurationArgs> externalDataConfiguration) {
            $.externalDataConfiguration = externalDataConfiguration;
            return this;
        }

        /**
         * @param externalDataConfiguration Describes the data format,
         * location, and other properties of a table stored outside of BigQuery.
         * By defining these properties, the data source can then be queried as
         * if it were a standard BigQuery table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder externalDataConfiguration(TableExternalDataConfigurationArgs externalDataConfiguration) {
            return externalDataConfiguration(Output.of(externalDataConfiguration));
        }

        /**
         * @param friendlyName A descriptive name for the table.
         * 
         * @return builder
         * 
         */
        public Builder friendlyName(@Nullable Output<String> friendlyName) {
            $.friendlyName = friendlyName;
            return this;
        }

        /**
         * @param friendlyName A descriptive name for the table.
         * 
         * @return builder
         * 
         */
        public Builder friendlyName(String friendlyName) {
            return friendlyName(Output.of(friendlyName));
        }

        /**
         * @param labels A mapping of labels to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels A mapping of labels to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param lastModifiedTime The time when this table was last modified, in milliseconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder lastModifiedTime(@Nullable Output<Integer> lastModifiedTime) {
            $.lastModifiedTime = lastModifiedTime;
            return this;
        }

        /**
         * @param lastModifiedTime The time when this table was last modified, in milliseconds since the epoch.
         * 
         * @return builder
         * 
         */
        public Builder lastModifiedTime(Integer lastModifiedTime) {
            return lastModifiedTime(Output.of(lastModifiedTime));
        }

        /**
         * @param location The geographic location where the table resides. This value is inherited from the dataset.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The geographic location where the table resides. This value is inherited from the dataset.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param materializedView If specified, configures this table as a materialized view.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder materializedView(@Nullable Output<TableMaterializedViewArgs> materializedView) {
            $.materializedView = materializedView;
            return this;
        }

        /**
         * @param materializedView If specified, configures this table as a materialized view.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder materializedView(TableMaterializedViewArgs materializedView) {
            return materializedView(Output.of(materializedView));
        }

        /**
         * @param numBytes The size of this table in bytes, excluding any data in the streaming buffer.
         * 
         * @return builder
         * 
         */
        public Builder numBytes(@Nullable Output<Integer> numBytes) {
            $.numBytes = numBytes;
            return this;
        }

        /**
         * @param numBytes The size of this table in bytes, excluding any data in the streaming buffer.
         * 
         * @return builder
         * 
         */
        public Builder numBytes(Integer numBytes) {
            return numBytes(Output.of(numBytes));
        }

        /**
         * @param numLongTermBytes The number of bytes in the table that are considered &#34;long-term storage&#34;.
         * 
         * @return builder
         * 
         */
        public Builder numLongTermBytes(@Nullable Output<Integer> numLongTermBytes) {
            $.numLongTermBytes = numLongTermBytes;
            return this;
        }

        /**
         * @param numLongTermBytes The number of bytes in the table that are considered &#34;long-term storage&#34;.
         * 
         * @return builder
         * 
         */
        public Builder numLongTermBytes(Integer numLongTermBytes) {
            return numLongTermBytes(Output.of(numLongTermBytes));
        }

        /**
         * @param numRows The number of rows of data in this table, excluding any data in the streaming buffer.
         * 
         * @return builder
         * 
         */
        public Builder numRows(@Nullable Output<Integer> numRows) {
            $.numRows = numRows;
            return this;
        }

        /**
         * @param numRows The number of rows of data in this table, excluding any data in the streaming buffer.
         * 
         * @return builder
         * 
         */
        public Builder numRows(Integer numRows) {
            return numRows(Output.of(numRows));
        }

        /**
         * @param project The ID of the project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param rangePartitioning If specified, configures range-based
         * partitioning for this table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rangePartitioning(@Nullable Output<TableRangePartitioningArgs> rangePartitioning) {
            $.rangePartitioning = rangePartitioning;
            return this;
        }

        /**
         * @param rangePartitioning If specified, configures range-based
         * partitioning for this table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rangePartitioning(TableRangePartitioningArgs rangePartitioning) {
            return rangePartitioning(Output.of(rangePartitioning));
        }

        /**
         * @param schema A JSON schema for the table.
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<String> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema A JSON schema for the table.
         * 
         * @return builder
         * 
         */
        public Builder schema(String schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param tableId A unique ID for the resource.
         * Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder tableId(@Nullable Output<String> tableId) {
            $.tableId = tableId;
            return this;
        }

        /**
         * @param tableId A unique ID for the resource.
         * Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder tableId(String tableId) {
            return tableId(Output.of(tableId));
        }

        /**
         * @param timePartitioning If specified, configures time-based
         * partitioning for this table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder timePartitioning(@Nullable Output<TableTimePartitioningArgs> timePartitioning) {
            $.timePartitioning = timePartitioning;
            return this;
        }

        /**
         * @param timePartitioning If specified, configures time-based
         * partitioning for this table. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder timePartitioning(TableTimePartitioningArgs timePartitioning) {
            return timePartitioning(Output.of(timePartitioning));
        }

        /**
         * @param type The supported types are DAY, HOUR, MONTH, and YEAR,
         * which will generate one partition per day, hour, month, and year, respectively.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The supported types are DAY, HOUR, MONTH, and YEAR,
         * which will generate one partition per day, hour, month, and year, respectively.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param view If specified, configures this table as a view.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder view(@Nullable Output<TableViewArgs> view) {
            $.view = view;
            return this;
        }

        /**
         * @param view If specified, configures this table as a view.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder view(TableViewArgs view) {
            return view(Output.of(view));
        }

        public TableState build() {
            return $;
        }
    }

}
