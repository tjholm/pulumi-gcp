// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.eventarc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TriggerDestinationGkeArgs extends com.pulumi.resources.ResourceArgs {

    public static final TriggerDestinationGkeArgs Empty = new TriggerDestinationGkeArgs();

    /**
     * Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
     * 
     */
    @Import(name="cluster", required=true)
    private Output<String> cluster;

    /**
     * @return Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
     * 
     */
    public Output<String> cluster() {
        return this.cluster;
    }

    /**
     * Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * Required. The namespace the GKE service is running in.
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return Required. The namespace the GKE service is running in.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }

    /**
     * Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: &#34;/route&#34;, &#34;route&#34;, &#34;route/subroute&#34;.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: &#34;/route&#34;, &#34;route&#34;, &#34;route/subroute&#34;.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Required. Name of the GKE service.
     * 
     */
    @Import(name="service", required=true)
    private Output<String> service;

    /**
     * @return Required. Name of the GKE service.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    private TriggerDestinationGkeArgs() {}

    private TriggerDestinationGkeArgs(TriggerDestinationGkeArgs $) {
        this.cluster = $.cluster;
        this.location = $.location;
        this.namespace = $.namespace;
        this.path = $.path;
        this.service = $.service;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TriggerDestinationGkeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TriggerDestinationGkeArgs $;

        public Builder() {
            $ = new TriggerDestinationGkeArgs();
        }

        public Builder(TriggerDestinationGkeArgs defaults) {
            $ = new TriggerDestinationGkeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cluster Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
         * 
         * @return builder
         * 
         */
        public Builder cluster(Output<String> cluster) {
            $.cluster = cluster;
            return this;
        }

        /**
         * @param cluster Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
         * 
         * @return builder
         * 
         */
        public Builder cluster(String cluster) {
            return cluster(Output.of(cluster));
        }

        /**
         * @param location Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param namespace Required. The namespace the GKE service is running in.
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Required. The namespace the GKE service is running in.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param path Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: &#34;/route&#34;, &#34;route&#34;, &#34;route/subroute&#34;.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: &#34;/route&#34;, &#34;route&#34;, &#34;route/subroute&#34;.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param service Required. Name of the GKE service.
         * 
         * @return builder
         * 
         */
        public Builder service(Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service Required. Name of the GKE service.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        public TriggerDestinationGkeArgs build() {
            $.cluster = Objects.requireNonNull($.cluster, "expected parameter 'cluster' to be non-null");
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            $.service = Objects.requireNonNull($.service, "expected parameter 'service' to be non-null");
            return $;
        }
    }

}
