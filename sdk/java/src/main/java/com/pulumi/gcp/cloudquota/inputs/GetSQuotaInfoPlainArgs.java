// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudquota.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetSQuotaInfoPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSQuotaInfoPlainArgs Empty = new GetSQuotaInfoPlainArgs();

    /**
     * The parent of the quota info. Allowed parents are &#34;projects/[project-id / number]&#34; or &#34;folders/[folder-id / number]&#34; or &#34;organizations/[org-id / number].
     * 
     */
    @Import(name="parent", required=true)
    private String parent;

    /**
     * @return The parent of the quota info. Allowed parents are &#34;projects/[project-id / number]&#34; or &#34;folders/[folder-id / number]&#34; or &#34;organizations/[org-id / number].
     * 
     */
    public String parent() {
        return this.parent;
    }

    /**
     * The id of the quota, which is unique within the service.
     * 
     */
    @Import(name="quotaId", required=true)
    private String quotaId;

    /**
     * @return The id of the quota, which is unique within the service.
     * 
     */
    public String quotaId() {
        return this.quotaId;
    }

    /**
     * The name of the service in which the quota is defined.
     * 
     */
    @Import(name="service", required=true)
    private String service;

    /**
     * @return The name of the service in which the quota is defined.
     * 
     */
    public String service() {
        return this.service;
    }

    private GetSQuotaInfoPlainArgs() {}

    private GetSQuotaInfoPlainArgs(GetSQuotaInfoPlainArgs $) {
        this.parent = $.parent;
        this.quotaId = $.quotaId;
        this.service = $.service;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSQuotaInfoPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSQuotaInfoPlainArgs $;

        public Builder() {
            $ = new GetSQuotaInfoPlainArgs();
        }

        public Builder(GetSQuotaInfoPlainArgs defaults) {
            $ = new GetSQuotaInfoPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param parent The parent of the quota info. Allowed parents are &#34;projects/[project-id / number]&#34; or &#34;folders/[folder-id / number]&#34; or &#34;organizations/[org-id / number].
         * 
         * @return builder
         * 
         */
        public Builder parent(String parent) {
            $.parent = parent;
            return this;
        }

        /**
         * @param quotaId The id of the quota, which is unique within the service.
         * 
         * @return builder
         * 
         */
        public Builder quotaId(String quotaId) {
            $.quotaId = quotaId;
            return this;
        }

        /**
         * @param service The name of the service in which the quota is defined.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            $.service = service;
            return this;
        }

        public GetSQuotaInfoPlainArgs build() {
            if ($.parent == null) {
                throw new MissingRequiredPropertyException("GetSQuotaInfoPlainArgs", "parent");
            }
            if ($.quotaId == null) {
                throw new MissingRequiredPropertyException("GetSQuotaInfoPlainArgs", "quotaId");
            }
            if ($.service == null) {
                throw new MissingRequiredPropertyException("GetSQuotaInfoPlainArgs", "service");
            }
            return $;
        }
    }

}
