// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigtable.inputs.TableColumnFamilyArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableState extends com.pulumi.resources.ResourceArgs {

    public static final TableState Empty = new TableState();

    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     * 
     */
    @Import(name="columnFamilies")
    private @Nullable Output<List<TableColumnFamilyArgs>> columnFamilies;

    /**
     * @return A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     * 
     */
    public Optional<Output<List<TableColumnFamilyArgs>>> columnFamilies() {
        return Optional.ofNullable(this.columnFamilies);
    }

    /**
     * The name of the Bigtable instance.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return The name of the Bigtable instance.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * The name of the table.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the table.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * A list of predefined keys to split the table on.
     * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     * 
     */
    @Import(name="splitKeys")
    private @Nullable Output<List<String>> splitKeys;

    /**
     * @return A list of predefined keys to split the table on.
     * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     * 
     */
    public Optional<Output<List<String>>> splitKeys() {
        return Optional.ofNullable(this.splitKeys);
    }

    private TableState() {}

    private TableState(TableState $) {
        this.columnFamilies = $.columnFamilies;
        this.instanceName = $.instanceName;
        this.name = $.name;
        this.project = $.project;
        this.splitKeys = $.splitKeys;
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
         * @param columnFamilies A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder columnFamilies(@Nullable Output<List<TableColumnFamilyArgs>> columnFamilies) {
            $.columnFamilies = columnFamilies;
            return this;
        }

        /**
         * @param columnFamilies A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder columnFamilies(List<TableColumnFamilyArgs> columnFamilies) {
            return columnFamilies(Output.of(columnFamilies));
        }

        /**
         * @param columnFamilies A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder columnFamilies(TableColumnFamilyArgs... columnFamilies) {
            return columnFamilies(List.of(columnFamilies));
        }

        /**
         * @param instanceName The name of the Bigtable instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the Bigtable instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param name The name of the table.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the table.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param splitKeys A list of predefined keys to split the table on.
         * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
         * to delete/recreate the entire `gcp.bigtable.Table` resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(@Nullable Output<List<String>> splitKeys) {
            $.splitKeys = splitKeys;
            return this;
        }

        /**
         * @param splitKeys A list of predefined keys to split the table on.
         * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
         * to delete/recreate the entire `gcp.bigtable.Table` resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(List<String> splitKeys) {
            return splitKeys(Output.of(splitKeys));
        }

        /**
         * @param splitKeys A list of predefined keys to split the table on.
         * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
         * to delete/recreate the entire `gcp.bigtable.Table` resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(String... splitKeys) {
            return splitKeys(List.of(splitKeys));
        }

        public TableState build() {
            return $;
        }
    }

}
