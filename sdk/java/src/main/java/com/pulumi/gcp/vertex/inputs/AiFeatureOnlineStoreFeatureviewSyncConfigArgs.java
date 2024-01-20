// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiFeatureOnlineStoreFeatureviewSyncConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final AiFeatureOnlineStoreFeatureviewSyncConfigArgs Empty = new AiFeatureOnlineStoreFeatureviewSyncConfigArgs();

    /**
     * Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.
     * To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: &#34;CRON_TZ=${IANA_TIME_ZONE}&#34; or &#34;TZ=${IANA_TIME_ZONE}&#34;.
     * 
     */
    @Import(name="cron")
    private @Nullable Output<String> cron;

    /**
     * @return Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.
     * To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: &#34;CRON_TZ=${IANA_TIME_ZONE}&#34; or &#34;TZ=${IANA_TIME_ZONE}&#34;.
     * 
     */
    public Optional<Output<String>> cron() {
        return Optional.ofNullable(this.cron);
    }

    private AiFeatureOnlineStoreFeatureviewSyncConfigArgs() {}

    private AiFeatureOnlineStoreFeatureviewSyncConfigArgs(AiFeatureOnlineStoreFeatureviewSyncConfigArgs $) {
        this.cron = $.cron;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiFeatureOnlineStoreFeatureviewSyncConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiFeatureOnlineStoreFeatureviewSyncConfigArgs $;

        public Builder() {
            $ = new AiFeatureOnlineStoreFeatureviewSyncConfigArgs();
        }

        public Builder(AiFeatureOnlineStoreFeatureviewSyncConfigArgs defaults) {
            $ = new AiFeatureOnlineStoreFeatureviewSyncConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cron Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.
         * To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: &#34;CRON_TZ=${IANA_TIME_ZONE}&#34; or &#34;TZ=${IANA_TIME_ZONE}&#34;.
         * 
         * @return builder
         * 
         */
        public Builder cron(@Nullable Output<String> cron) {
            $.cron = cron;
            return this;
        }

        /**
         * @param cron Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.
         * To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: &#34;CRON_TZ=${IANA_TIME_ZONE}&#34; or &#34;TZ=${IANA_TIME_ZONE}&#34;.
         * 
         * @return builder
         * 
         */
        public Builder cron(String cron) {
            return cron(Output.of(cron));
        }

        public AiFeatureOnlineStoreFeatureviewSyncConfigArgs build() {
            return $;
        }
    }

}
