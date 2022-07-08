// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionCloudResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionCloudResourceArgs Empty = new ConnectionCloudResourceArgs();

    /**
     * - 
     * The account ID of the service created for the purpose of this connection.
     * 
     */
    @Import(name="serviceAccountId")
    private @Nullable Output<String> serviceAccountId;

    /**
     * @return -
     * The account ID of the service created for the purpose of this connection.
     * 
     */
    public Optional<Output<String>> serviceAccountId() {
        return Optional.ofNullable(this.serviceAccountId);
    }

    private ConnectionCloudResourceArgs() {}

    private ConnectionCloudResourceArgs(ConnectionCloudResourceArgs $) {
        this.serviceAccountId = $.serviceAccountId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionCloudResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionCloudResourceArgs $;

        public Builder() {
            $ = new ConnectionCloudResourceArgs();
        }

        public Builder(ConnectionCloudResourceArgs defaults) {
            $ = new ConnectionCloudResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceAccountId -
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
         * @param serviceAccountId -
         * The account ID of the service created for the purpose of this connection.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            return serviceAccountId(Output.of(serviceAccountId));
        }

        public ConnectionCloudResourceArgs build() {
            return $;
        }
    }

}
