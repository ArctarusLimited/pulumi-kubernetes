// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Scheduling.V1Beta1
{
    /// <summary>
    /// DEPRECATED - This group version of PriorityClass is deprecated by
    /// scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name
    /// to the priority integer value. The value can be any valid integer.
    /// </summary>
    public partial class PriorityClass : Pulumi.CustomResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// description is an arbitrary string that usually provides guidelines on when this
        /// priority class should be used.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// globalDefault specifies whether this PriorityClass should be considered as the default
        /// priority for pods that do not have any priority class. Only one PriorityClass can be
        /// marked as `globalDefault`. However, if more than one PriorityClasses exists with their
        /// `globalDefault` field set to true, the smallest value of such global default
        /// PriorityClasses will be used as the default priority.
        /// </summary>
        [Output("globalDefault")]
        public Output<bool> GlobalDefault { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never,
        /// PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is
        /// alpha-level and is only honored by servers that enable the NonPreemptingPriority
        /// feature.
        /// </summary>
        [Output("preemptionPolicy")]
        public Output<string> PreemptionPolicy { get; private set; } = null!;

        /// <summary>
        /// The value of this priority class. This is the actual priority that pods receive when
        /// they have the name of this class in their pod spec.
        /// </summary>
        [Output("value")]
        public Output<int> Value { get; private set; } = null!;


        /// <summary>
        /// Create a PriorityClass resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PriorityClass(string name, Types.Inputs.Scheduling.V1Beta1.PriorityClassArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:scheduling.k8s.io/v1beta1:PriorityClass", name, SetAPIKindAndVersion(args), MakeResourceOptions(options))
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Scheduling.V1Beta1.PriorityClassArgs? args)
        {
            args ??= new Types.Inputs.Scheduling.V1Beta1.PriorityClassArgs();
            args.ApiVersion = "scheduling.k8s.io/v1beta1";
            args.Kind = "PriorityClass";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            return CustomResourceOptions.Merge(defaultOptions, options);
        }

        /// <summary>
        /// Get an existing PriorityClass resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PriorityClass Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PriorityClass(name, null, CustomResourceOptions.Merge(options, new CustomResourceOptions
            {
                Id = id,
            }));
        }

    }
}
