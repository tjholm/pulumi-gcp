// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ManagedZoneServiceDirectoryConfigNamespaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ManagedZoneServiceDirectoryConfigNamespaceArgs Empty = new ManagedZoneServiceDirectoryConfigNamespaceArgs();

    /**
     * The fully qualified or partial URL of the service directory namespace that should be
     * associated with the zone. This should be formatted like
     * `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
     * or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
     * Ignored for `public` visibility zones.
     * 
     */
    @Import(name="namespaceUrl", required=true)
    private Output<String> namespaceUrl;

    /**
     * @return The fully qualified or partial URL of the service directory namespace that should be
     * associated with the zone. This should be formatted like
     * `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
     * or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
     * Ignored for `public` visibility zones.
     * 
     */
    public Output<String> namespaceUrl() {
        return this.namespaceUrl;
    }

    private ManagedZoneServiceDirectoryConfigNamespaceArgs() {}

    private ManagedZoneServiceDirectoryConfigNamespaceArgs(ManagedZoneServiceDirectoryConfigNamespaceArgs $) {
        this.namespaceUrl = $.namespaceUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ManagedZoneServiceDirectoryConfigNamespaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ManagedZoneServiceDirectoryConfigNamespaceArgs $;

        public Builder() {
            $ = new ManagedZoneServiceDirectoryConfigNamespaceArgs();
        }

        public Builder(ManagedZoneServiceDirectoryConfigNamespaceArgs defaults) {
            $ = new ManagedZoneServiceDirectoryConfigNamespaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespaceUrl The fully qualified or partial URL of the service directory namespace that should be
         * associated with the zone. This should be formatted like
         * `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
         * or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
         * Ignored for `public` visibility zones.
         * 
         * @return builder
         * 
         */
        public Builder namespaceUrl(Output<String> namespaceUrl) {
            $.namespaceUrl = namespaceUrl;
            return this;
        }

        /**
         * @param namespaceUrl The fully qualified or partial URL of the service directory namespace that should be
         * associated with the zone. This should be formatted like
         * `https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}`
         * or simply `projects/{project}/locations/{location}/namespaces/{namespace_id}`
         * Ignored for `public` visibility zones.
         * 
         * @return builder
         * 
         */
        public Builder namespaceUrl(String namespaceUrl) {
            return namespaceUrl(Output.of(namespaceUrl));
        }

        public ManagedZoneServiceDirectoryConfigNamespaceArgs build() {
            $.namespaceUrl = Objects.requireNonNull($.namespaceUrl, "expected parameter 'namespaceUrl' to be non-null");
            return $;
        }
    }

}
