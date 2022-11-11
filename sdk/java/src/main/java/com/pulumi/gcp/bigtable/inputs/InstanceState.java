// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigtable.inputs.InstanceClusterArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceState Empty = new InstanceState();

    /**
     * A block of cluster configuration options. This can be specified at least once, and up
     * to as many as possible within 8 cloud regions. See structure below.
     * 
     */
    @Import(name="clusters")
    private @Nullable Output<List<InstanceClusterArgs>> clusters;

    /**
     * @return A block of cluster configuration options. This can be specified at least once, and up
     * to as many as possible within 8 cloud regions. See structure below.
     * 
     */
    public Optional<Output<List<InstanceClusterArgs>>> clusters() {
        return Optional.ofNullable(this.clusters);
    }

    /**
     * Whether or not to allow this provider to destroy the instance. Unless this field is set to false
     * in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return Whether or not to allow this provider to destroy the instance. Unless this field is set to false
     * in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The instance type to create. One of `&#34;DEVELOPMENT&#34;` or `&#34;PRODUCTION&#34;`. Defaults to `&#34;PRODUCTION&#34;`.
     * It is recommended to leave this field unspecified since the distinction between `&#34;DEVELOPMENT&#34;` and `&#34;PRODUCTION&#34;` instances is going away,
     * and all instances will become `&#34;PRODUCTION&#34;` instances. This means that new and existing `&#34;DEVELOPMENT&#34;` instances will be converted to
     * `&#34;PRODUCTION&#34;` instances. It is recommended for users to use `&#34;PRODUCTION&#34;` instances in any case, since a 1-node `&#34;PRODUCTION&#34;` instance
     * is functionally identical to a `&#34;DEVELOPMENT&#34;` instance, but without the accompanying restrictions.
     * 
     * @deprecated
     * It is recommended to leave this field unspecified since the distinction between &#34;DEVELOPMENT&#34; and &#34;PRODUCTION&#34; instances is going away, and all instances will become &#34;PRODUCTION&#34; instances. This means that new and existing &#34;DEVELOPMENT&#34; instances will be converted to &#34;PRODUCTION&#34; instances. It is recommended for users to use &#34;PRODUCTION&#34; instances in any case, since a 1-node &#34;PRODUCTION&#34; instance is functionally identical to a &#34;DEVELOPMENT&#34; instance, but without the accompanying restrictions.
     * 
     */
    @Deprecated /* It is recommended to leave this field unspecified since the distinction between ""DEVELOPMENT"" and ""PRODUCTION"" instances is going away, and all instances will become ""PRODUCTION"" instances. This means that new and existing ""DEVELOPMENT"" instances will be converted to ""PRODUCTION"" instances. It is recommended for users to use ""PRODUCTION"" instances in any case, since a 1-node ""PRODUCTION"" instance is functionally identical to a ""DEVELOPMENT"" instance, but without the accompanying restrictions. */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The instance type to create. One of `&#34;DEVELOPMENT&#34;` or `&#34;PRODUCTION&#34;`. Defaults to `&#34;PRODUCTION&#34;`.
     * It is recommended to leave this field unspecified since the distinction between `&#34;DEVELOPMENT&#34;` and `&#34;PRODUCTION&#34;` instances is going away,
     * and all instances will become `&#34;PRODUCTION&#34;` instances. This means that new and existing `&#34;DEVELOPMENT&#34;` instances will be converted to
     * `&#34;PRODUCTION&#34;` instances. It is recommended for users to use `&#34;PRODUCTION&#34;` instances in any case, since a 1-node `&#34;PRODUCTION&#34;` instance
     * is functionally identical to a `&#34;DEVELOPMENT&#34;` instance, but without the accompanying restrictions.
     * 
     * @deprecated
     * It is recommended to leave this field unspecified since the distinction between &#34;DEVELOPMENT&#34; and &#34;PRODUCTION&#34; instances is going away, and all instances will become &#34;PRODUCTION&#34; instances. This means that new and existing &#34;DEVELOPMENT&#34; instances will be converted to &#34;PRODUCTION&#34; instances. It is recommended for users to use &#34;PRODUCTION&#34; instances in any case, since a 1-node &#34;PRODUCTION&#34; instance is functionally identical to a &#34;DEVELOPMENT&#34; instance, but without the accompanying restrictions.
     * 
     */
    @Deprecated /* It is recommended to leave this field unspecified since the distinction between ""DEVELOPMENT"" and ""PRODUCTION"" instances is going away, and all instances will become ""PRODUCTION"" instances. This means that new and existing ""DEVELOPMENT"" instances will be converted to ""PRODUCTION"" instances. It is recommended for users to use ""PRODUCTION"" instances in any case, since a 1-node ""PRODUCTION"" instance is functionally identical to a ""DEVELOPMENT"" instance, but without the accompanying restrictions. */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
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

    private InstanceState() {}

    private InstanceState(InstanceState $) {
        this.clusters = $.clusters;
        this.deletionProtection = $.deletionProtection;
        this.displayName = $.displayName;
        this.instanceType = $.instanceType;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceState $;

        public Builder() {
            $ = new InstanceState();
        }

        public Builder(InstanceState defaults) {
            $ = new InstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusters A block of cluster configuration options. This can be specified at least once, and up
         * to as many as possible within 8 cloud regions. See structure below.
         * 
         * @return builder
         * 
         */
        public Builder clusters(@Nullable Output<List<InstanceClusterArgs>> clusters) {
            $.clusters = clusters;
            return this;
        }

        /**
         * @param clusters A block of cluster configuration options. This can be specified at least once, and up
         * to as many as possible within 8 cloud regions. See structure below.
         * 
         * @return builder
         * 
         */
        public Builder clusters(List<InstanceClusterArgs> clusters) {
            return clusters(Output.of(clusters));
        }

        /**
         * @param clusters A block of cluster configuration options. This can be specified at least once, and up
         * to as many as possible within 8 cloud regions. See structure below.
         * 
         * @return builder
         * 
         */
        public Builder clusters(InstanceClusterArgs... clusters) {
            return clusters(List.of(clusters));
        }

        /**
         * @param deletionProtection Whether or not to allow this provider to destroy the instance. Unless this field is set to false
         * in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection Whether or not to allow this provider to destroy the instance. Unless this field is set to false
         * in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param displayName The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param instanceType The instance type to create. One of `&#34;DEVELOPMENT&#34;` or `&#34;PRODUCTION&#34;`. Defaults to `&#34;PRODUCTION&#34;`.
         * It is recommended to leave this field unspecified since the distinction between `&#34;DEVELOPMENT&#34;` and `&#34;PRODUCTION&#34;` instances is going away,
         * and all instances will become `&#34;PRODUCTION&#34;` instances. This means that new and existing `&#34;DEVELOPMENT&#34;` instances will be converted to
         * `&#34;PRODUCTION&#34;` instances. It is recommended for users to use `&#34;PRODUCTION&#34;` instances in any case, since a 1-node `&#34;PRODUCTION&#34;` instance
         * is functionally identical to a `&#34;DEVELOPMENT&#34;` instance, but without the accompanying restrictions.
         * 
         * @return builder
         * 
         * @deprecated
         * It is recommended to leave this field unspecified since the distinction between &#34;DEVELOPMENT&#34; and &#34;PRODUCTION&#34; instances is going away, and all instances will become &#34;PRODUCTION&#34; instances. This means that new and existing &#34;DEVELOPMENT&#34; instances will be converted to &#34;PRODUCTION&#34; instances. It is recommended for users to use &#34;PRODUCTION&#34; instances in any case, since a 1-node &#34;PRODUCTION&#34; instance is functionally identical to a &#34;DEVELOPMENT&#34; instance, but without the accompanying restrictions.
         * 
         */
        @Deprecated /* It is recommended to leave this field unspecified since the distinction between ""DEVELOPMENT"" and ""PRODUCTION"" instances is going away, and all instances will become ""PRODUCTION"" instances. This means that new and existing ""DEVELOPMENT"" instances will be converted to ""PRODUCTION"" instances. It is recommended for users to use ""PRODUCTION"" instances in any case, since a 1-node ""PRODUCTION"" instance is functionally identical to a ""DEVELOPMENT"" instance, but without the accompanying restrictions. */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The instance type to create. One of `&#34;DEVELOPMENT&#34;` or `&#34;PRODUCTION&#34;`. Defaults to `&#34;PRODUCTION&#34;`.
         * It is recommended to leave this field unspecified since the distinction between `&#34;DEVELOPMENT&#34;` and `&#34;PRODUCTION&#34;` instances is going away,
         * and all instances will become `&#34;PRODUCTION&#34;` instances. This means that new and existing `&#34;DEVELOPMENT&#34;` instances will be converted to
         * `&#34;PRODUCTION&#34;` instances. It is recommended for users to use `&#34;PRODUCTION&#34;` instances in any case, since a 1-node `&#34;PRODUCTION&#34;` instance
         * is functionally identical to a `&#34;DEVELOPMENT&#34;` instance, but without the accompanying restrictions.
         * 
         * @return builder
         * 
         * @deprecated
         * It is recommended to leave this field unspecified since the distinction between &#34;DEVELOPMENT&#34; and &#34;PRODUCTION&#34; instances is going away, and all instances will become &#34;PRODUCTION&#34; instances. This means that new and existing &#34;DEVELOPMENT&#34; instances will be converted to &#34;PRODUCTION&#34; instances. It is recommended for users to use &#34;PRODUCTION&#34; instances in any case, since a 1-node &#34;PRODUCTION&#34; instance is functionally identical to a &#34;DEVELOPMENT&#34; instance, but without the accompanying restrictions.
         * 
         */
        @Deprecated /* It is recommended to leave this field unspecified since the distinction between ""DEVELOPMENT"" and ""PRODUCTION"" instances is going away, and all instances will become ""PRODUCTION"" instances. This means that new and existing ""DEVELOPMENT"" instances will be converted to ""PRODUCTION"" instances. It is recommended for users to use ""PRODUCTION"" instances in any case, since a 1-node ""PRODUCTION"" instance is functionally identical to a ""DEVELOPMENT"" instance, but without the accompanying restrictions. */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param labels A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param name The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
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

        public InstanceState build() {
            return $;
        }
    }

}
