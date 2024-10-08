// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.ConnectionSparkMetastoreServiceConfigArgs;
import com.pulumi.gcp.bigquery.inputs.ConnectionSparkSparkHistoryServerConfigArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionSparkArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionSparkArgs Empty = new ConnectionSparkArgs();

    /**
     * Dataproc Metastore Service configuration for the connection.
     * Structure is documented below.
     * 
     */
    @Import(name="metastoreServiceConfig")
    private @Nullable Output<ConnectionSparkMetastoreServiceConfigArgs> metastoreServiceConfig;

    /**
     * @return Dataproc Metastore Service configuration for the connection.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionSparkMetastoreServiceConfigArgs>> metastoreServiceConfig() {
        return Optional.ofNullable(this.metastoreServiceConfig);
    }

    /**
     * (Output)
     * The account ID of the service created for the purpose of this connection.
     * 
     */
    @Import(name="serviceAccountId")
    private @Nullable Output<String> serviceAccountId;

    /**
     * @return (Output)
     * The account ID of the service created for the purpose of this connection.
     * 
     */
    public Optional<Output<String>> serviceAccountId() {
        return Optional.ofNullable(this.serviceAccountId);
    }

    /**
     * Spark History Server configuration for the connection.
     * Structure is documented below.
     * 
     */
    @Import(name="sparkHistoryServerConfig")
    private @Nullable Output<ConnectionSparkSparkHistoryServerConfigArgs> sparkHistoryServerConfig;

    /**
     * @return Spark History Server configuration for the connection.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionSparkSparkHistoryServerConfigArgs>> sparkHistoryServerConfig() {
        return Optional.ofNullable(this.sparkHistoryServerConfig);
    }

    private ConnectionSparkArgs() {}

    private ConnectionSparkArgs(ConnectionSparkArgs $) {
        this.metastoreServiceConfig = $.metastoreServiceConfig;
        this.serviceAccountId = $.serviceAccountId;
        this.sparkHistoryServerConfig = $.sparkHistoryServerConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionSparkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionSparkArgs $;

        public Builder() {
            $ = new ConnectionSparkArgs();
        }

        public Builder(ConnectionSparkArgs defaults) {
            $ = new ConnectionSparkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metastoreServiceConfig Dataproc Metastore Service configuration for the connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metastoreServiceConfig(@Nullable Output<ConnectionSparkMetastoreServiceConfigArgs> metastoreServiceConfig) {
            $.metastoreServiceConfig = metastoreServiceConfig;
            return this;
        }

        /**
         * @param metastoreServiceConfig Dataproc Metastore Service configuration for the connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metastoreServiceConfig(ConnectionSparkMetastoreServiceConfigArgs metastoreServiceConfig) {
            return metastoreServiceConfig(Output.of(metastoreServiceConfig));
        }

        /**
         * @param serviceAccountId (Output)
         * The account ID of the service created for the purpose of this connection.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(@Nullable Output<String> serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        /**
         * @param serviceAccountId (Output)
         * The account ID of the service created for the purpose of this connection.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            return serviceAccountId(Output.of(serviceAccountId));
        }

        /**
         * @param sparkHistoryServerConfig Spark History Server configuration for the connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder sparkHistoryServerConfig(@Nullable Output<ConnectionSparkSparkHistoryServerConfigArgs> sparkHistoryServerConfig) {
            $.sparkHistoryServerConfig = sparkHistoryServerConfig;
            return this;
        }

        /**
         * @param sparkHistoryServerConfig Spark History Server configuration for the connection.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder sparkHistoryServerConfig(ConnectionSparkSparkHistoryServerConfigArgs sparkHistoryServerConfig) {
            return sparkHistoryServerConfig(Output.of(sparkHistoryServerConfig));
        }

        public ConnectionSparkArgs build() {
            return $;
        }
    }

}
