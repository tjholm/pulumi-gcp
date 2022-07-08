// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DicomStoreNotificationConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final DicomStoreNotificationConfigArgs Empty = new DicomStoreNotificationConfigArgs();

    /**
     * The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
     * PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
     * It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
     * was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
     * project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
     * Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
     * 
     */
    @Import(name="pubsubTopic", required=true)
    private Output<String> pubsubTopic;

    /**
     * @return The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
     * PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
     * It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
     * was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
     * project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
     * Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
     * 
     */
    public Output<String> pubsubTopic() {
        return this.pubsubTopic;
    }

    private DicomStoreNotificationConfigArgs() {}

    private DicomStoreNotificationConfigArgs(DicomStoreNotificationConfigArgs $) {
        this.pubsubTopic = $.pubsubTopic;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DicomStoreNotificationConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DicomStoreNotificationConfigArgs $;

        public Builder() {
            $ = new DicomStoreNotificationConfigArgs();
        }

        public Builder(DicomStoreNotificationConfigArgs defaults) {
            $ = new DicomStoreNotificationConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param pubsubTopic The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
         * PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
         * It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
         * was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
         * project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
         * Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
         * 
         * @return builder
         * 
         */
        public Builder pubsubTopic(Output<String> pubsubTopic) {
            $.pubsubTopic = pubsubTopic;
            return this;
        }

        /**
         * @param pubsubTopic The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
         * PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
         * It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
         * was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
         * project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
         * Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
         * 
         * @return builder
         * 
         */
        public Builder pubsubTopic(String pubsubTopic) {
            return pubsubTopic(Output.of(pubsubTopic));
        }

        public DicomStoreNotificationConfigArgs build() {
            $.pubsubTopic = Objects.requireNonNull($.pubsubTopic, "expected parameter 'pubsubTopic' to be non-null");
            return $;
        }
    }

}
