// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigtable.inputs.TableColumnFamilyArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableArgs extends com.pulumi.resources.ResourceArgs {

    public static final TableArgs Empty = new TableArgs();

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
     * A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<String> deletionProtection;

    /**
     * @return A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     * 
     */
    public Optional<Output<String>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * The name of the Bigtable instance.
     * 
     */
    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    /**
     * @return The name of the Bigtable instance.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }

    /**
     * The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
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
     * A list of predefined keys to split the table on. !&gt; Warning: Modifying the split_keys of an existing table will cause
     * Terraform to delete/recreate the entire google_bigtable_table resource.
     * 
     */
    @Import(name="splitKeys")
    private @Nullable Output<List<String>> splitKeys;

    /**
     * @return A list of predefined keys to split the table on. !&gt; Warning: Modifying the split_keys of an existing table will cause
     * Terraform to delete/recreate the entire google_bigtable_table resource.
     * 
     */
    public Optional<Output<List<String>>> splitKeys() {
        return Optional.ofNullable(this.splitKeys);
    }

    private TableArgs() {}

    private TableArgs(TableArgs $) {
        this.columnFamilies = $.columnFamilies;
        this.deletionProtection = $.deletionProtection;
        this.instanceName = $.instanceName;
        this.name = $.name;
        this.project = $.project;
        this.splitKeys = $.splitKeys;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableArgs $;

        public Builder() {
            $ = new TableArgs();
        }

        public Builder(TableArgs defaults) {
            $ = new TableArgs(Objects.requireNonNull(defaults));
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
         * @param deletionProtection A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<String> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(String deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param instanceName The name of the Bigtable instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(Output<String> instanceName) {
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
         * @param name The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
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
         * @param splitKeys A list of predefined keys to split the table on. !&gt; Warning: Modifying the split_keys of an existing table will cause
         * Terraform to delete/recreate the entire google_bigtable_table resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(@Nullable Output<List<String>> splitKeys) {
            $.splitKeys = splitKeys;
            return this;
        }

        /**
         * @param splitKeys A list of predefined keys to split the table on. !&gt; Warning: Modifying the split_keys of an existing table will cause
         * Terraform to delete/recreate the entire google_bigtable_table resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(List<String> splitKeys) {
            return splitKeys(Output.of(splitKeys));
        }

        /**
         * @param splitKeys A list of predefined keys to split the table on. !&gt; Warning: Modifying the split_keys of an existing table will cause
         * Terraform to delete/recreate the entire google_bigtable_table resource.
         * 
         * @return builder
         * 
         */
        public Builder splitKeys(String... splitKeys) {
            return splitKeys(List.of(splitKeys));
        }

        public TableArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            return $;
        }
    }

}
