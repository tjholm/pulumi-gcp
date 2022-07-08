// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firestore.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.firestore.inputs.IndexFieldArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IndexState extends com.pulumi.resources.ResourceArgs {

    public static final IndexState Empty = new IndexState();

    /**
     * The collection being indexed.
     * 
     */
    @Import(name="collection")
    private @Nullable Output<String> collection;

    /**
     * @return The collection being indexed.
     * 
     */
    public Optional<Output<String>> collection() {
        return Optional.ofNullable(this.collection);
    }

    /**
     * The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * The fields supported by this index. The last field entry is always for
     * the field path `__name__`. If, on creation, `__name__` was not
     * specified as the last field, it will be added automatically with the
     * same direction as that of the last field defined. If the final field
     * in a composite index is not directional, the `__name__` will be
     * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
     * Structure is documented below.
     * 
     */
    @Import(name="fields")
    private @Nullable Output<List<IndexFieldArgs>> fields;

    /**
     * @return The fields supported by this index. The last field entry is always for
     * the field path `__name__`. If, on creation, `__name__` was not
     * specified as the last field, it will be added automatically with the
     * same direction as that of the last field defined. If the final field
     * in a composite index is not directional, the `__name__` will be
     * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<IndexFieldArgs>>> fields() {
        return Optional.ofNullable(this.fields);
    }

    /**
     * A server defined name for this index. Format:
     * &#39;projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}&#39;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A server defined name for this index. Format:
     * &#39;projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}&#39;
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * The scope at which a query is run.
     * Default value is `COLLECTION`.
     * Possible values are `COLLECTION` and `COLLECTION_GROUP`.
     * 
     */
    @Import(name="queryScope")
    private @Nullable Output<String> queryScope;

    /**
     * @return The scope at which a query is run.
     * Default value is `COLLECTION`.
     * Possible values are `COLLECTION` and `COLLECTION_GROUP`.
     * 
     */
    public Optional<Output<String>> queryScope() {
        return Optional.ofNullable(this.queryScope);
    }

    private IndexState() {}

    private IndexState(IndexState $) {
        this.collection = $.collection;
        this.database = $.database;
        this.fields = $.fields;
        this.name = $.name;
        this.project = $.project;
        this.queryScope = $.queryScope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IndexState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IndexState $;

        public Builder() {
            $ = new IndexState();
        }

        public Builder(IndexState defaults) {
            $ = new IndexState(Objects.requireNonNull(defaults));
        }

        /**
         * @param collection The collection being indexed.
         * 
         * @return builder
         * 
         */
        public Builder collection(@Nullable Output<String> collection) {
            $.collection = collection;
            return this;
        }

        /**
         * @param collection The collection being indexed.
         * 
         * @return builder
         * 
         */
        public Builder collection(String collection) {
            return collection(Output.of(collection));
        }

        /**
         * @param database The Firestore database id. Defaults to `&#34;(default)&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The Firestore database id. Defaults to `&#34;(default)&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param fields The fields supported by this index. The last field entry is always for
         * the field path `__name__`. If, on creation, `__name__` was not
         * specified as the last field, it will be added automatically with the
         * same direction as that of the last field defined. If the final field
         * in a composite index is not directional, the `__name__` will be
         * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder fields(@Nullable Output<List<IndexFieldArgs>> fields) {
            $.fields = fields;
            return this;
        }

        /**
         * @param fields The fields supported by this index. The last field entry is always for
         * the field path `__name__`. If, on creation, `__name__` was not
         * specified as the last field, it will be added automatically with the
         * same direction as that of the last field defined. If the final field
         * in a composite index is not directional, the `__name__` will be
         * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder fields(List<IndexFieldArgs> fields) {
            return fields(Output.of(fields));
        }

        /**
         * @param fields The fields supported by this index. The last field entry is always for
         * the field path `__name__`. If, on creation, `__name__` was not
         * specified as the last field, it will be added automatically with the
         * same direction as that of the last field defined. If the final field
         * in a composite index is not directional, the `__name__` will be
         * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder fields(IndexFieldArgs... fields) {
            return fields(List.of(fields));
        }

        /**
         * @param name A server defined name for this index. Format:
         * &#39;projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}&#39;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A server defined name for this index. Format:
         * &#39;projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}&#39;
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param queryScope The scope at which a query is run.
         * Default value is `COLLECTION`.
         * Possible values are `COLLECTION` and `COLLECTION_GROUP`.
         * 
         * @return builder
         * 
         */
        public Builder queryScope(@Nullable Output<String> queryScope) {
            $.queryScope = queryScope;
            return this;
        }

        /**
         * @param queryScope The scope at which a query is run.
         * Default value is `COLLECTION`.
         * Possible values are `COLLECTION` and `COLLECTION_GROUP`.
         * 
         * @return builder
         * 
         */
        public Builder queryScope(String queryScope) {
            return queryScope(Output.of(queryScope));
        }

        public IndexState build() {
            return $;
        }
    }

}
