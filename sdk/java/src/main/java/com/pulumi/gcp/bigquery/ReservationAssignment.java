// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.ReservationAssignmentArgs;
import com.pulumi.gcp.bigquery.inputs.ReservationAssignmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The BigqueryReservation Assignment resource.
 * 
 * To get more information about ReservationAssignment, see:
 * 
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1/projects.locations.reservations.assignments)
 * * How-to Guides
 *     * [Work with reservation assignments](https://cloud.google.com/bigquery/docs/reservations-assignments)
 * 
 * ## Example Usage
 * 
 * ### Bigquery Reservation Assignment Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Reservation;
 * import com.pulumi.gcp.bigquery.ReservationArgs;
 * import com.pulumi.gcp.bigquery.ReservationAssignment;
 * import com.pulumi.gcp.bigquery.ReservationAssignmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var basic = new Reservation("basic", ReservationArgs.builder()
 *             .name("example-reservation")
 *             .project("my-project-name")
 *             .location("us-central1")
 *             .slotCapacity(0)
 *             .ignoreIdleSlots(false)
 *             .build());
 * 
 *         var assignment = new ReservationAssignment("assignment", ReservationAssignmentArgs.builder()
 *             .assignee("projects/my-project-name")
 *             .jobType("PIPELINE")
 *             .reservation(basic.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ReservationAssignment can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{reservation}}/{{name}}`
 * 
 * * `{{location}}/{{reservation}}/{{name}}`
 * 
 * When using the `pulumi import` command, ReservationAssignment can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/reservationAssignment:ReservationAssignment default projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/reservationAssignment:ReservationAssignment default {{project}}/{{location}}/{{reservation}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/reservationAssignment:ReservationAssignment default {{location}}/{{reservation}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigquery/reservationAssignment:ReservationAssignment")
public class ReservationAssignment extends com.pulumi.resources.CustomResource {
    /**
     * The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
     * 
     */
    @Export(name="assignee", refs={String.class}, tree="[0]")
    private Output<String> assignee;

    /**
     * @return The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
     * 
     */
    public Output<String> assignee() {
        return this.assignee;
    }
    /**
     * Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
     * 
     */
    @Export(name="jobType", refs={String.class}, tree="[0]")
    private Output<String> jobType;

    /**
     * @return Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
     * 
     */
    public Output<String> jobType() {
        return this.jobType;
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Output only. The resource name of the assignment.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output only. The resource name of the assignment.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The reservation for the resource
     * 
     * ***
     * 
     */
    @Export(name="reservation", refs={String.class}, tree="[0]")
    private Output<String> reservation;

    /**
     * @return The reservation for the resource
     * 
     * ***
     * 
     */
    public Output<String> reservation() {
        return this.reservation;
    }
    /**
     * Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active.
     * Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active.
     * Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReservationAssignment(java.lang.String name) {
        this(name, ReservationAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReservationAssignment(java.lang.String name, ReservationAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReservationAssignment(java.lang.String name, ReservationAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/reservationAssignment:ReservationAssignment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReservationAssignment(java.lang.String name, Output<java.lang.String> id, @Nullable ReservationAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/reservationAssignment:ReservationAssignment", name, state, makeResourceOptions(options, id), false);
    }

    private static ReservationAssignmentArgs makeArgs(ReservationAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReservationAssignmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ReservationAssignment get(java.lang.String name, Output<java.lang.String> id, @Nullable ReservationAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReservationAssignment(name, id, state, options);
    }
}
