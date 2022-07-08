// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionCloudSpannerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionCloudSpannerArgs Empty = new ConnectionCloudSpannerArgs();

    /**
     * Cloud Spanner database in the form `project/instance/database&#39;
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return Cloud Spanner database in the form `project/instance/database&#39;
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * If parallelism should be used when reading from Cloud Spanner
     * 
     */
    @Import(name="useParallelism")
    private @Nullable Output<Boolean> useParallelism;

    /**
     * @return If parallelism should be used when reading from Cloud Spanner
     * 
     */
    public Optional<Output<Boolean>> useParallelism() {
        return Optional.ofNullable(this.useParallelism);
    }

    private ConnectionCloudSpannerArgs() {}

    private ConnectionCloudSpannerArgs(ConnectionCloudSpannerArgs $) {
        this.database = $.database;
        this.useParallelism = $.useParallelism;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionCloudSpannerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionCloudSpannerArgs $;

        public Builder() {
            $ = new ConnectionCloudSpannerArgs();
        }

        public Builder(ConnectionCloudSpannerArgs defaults) {
            $ = new ConnectionCloudSpannerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database Cloud Spanner database in the form `project/instance/database&#39;
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database Cloud Spanner database in the form `project/instance/database&#39;
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param useParallelism If parallelism should be used when reading from Cloud Spanner
         * 
         * @return builder
         * 
         */
        public Builder useParallelism(@Nullable Output<Boolean> useParallelism) {
            $.useParallelism = useParallelism;
            return this;
        }

        /**
         * @param useParallelism If parallelism should be used when reading from Cloud Spanner
         * 
         * @return builder
         * 
         */
        public Builder useParallelism(Boolean useParallelism) {
            return useParallelism(Output.of(useParallelism));
        }

        public ConnectionCloudSpannerArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            return $;
        }
    }

}
